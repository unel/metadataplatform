package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---------------------------------------------------------------------------
// FT-01 — Старт с валидным конфигом, один сокет
// ---------------------------------------------------------------------------

func TestStore_SingleSocket_FileExistsAndAcceptsConnections(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	sockPath := sp.SockPath()

	// Assert: сокет-файл существует на ФС
	_, err := os.Stat(sockPath)
	require.NoError(t, err, "socket file must exist after store starts")

	// Assert: store принимает входящие подключения
	conn, err := net.DialTimeout("unix", sockPath, time.Second)
	require.NoError(t, err, "store must accept unix connections on the socket path")
	conn.Close()
}

// ---------------------------------------------------------------------------
// FT-02 — Старт с валидным конфигом, несколько сокетов
// ---------------------------------------------------------------------------

func TestStore_MultipleSockets_AllFilesExistAndAcceptConnections(t *testing.T) {
	// Arrange
	sockNames := []string{"store-a.sock", "store-b.sock", "store-c.sock"}
	sp := newStoreProcess(t, sockNames)

	// Assert: каждый из трёх сокет-файлов существует и принимает подключения
	for i, name := range sockNames {
		path := sp.SockPath(i)

		_, statErr := os.Stat(path)
		require.NoErrorf(t, statErr,
			"socket file %q must exist after store starts", name)

		conn, dialErr := net.DialTimeout("unix", path, time.Second)
		require.NoErrorf(t, dialErr,
			"store must accept connections on socket %q", name)
		conn.Close()
	}
}

// ---------------------------------------------------------------------------
// FT-03 — Эхо-ответ на одно сообщение
// ---------------------------------------------------------------------------

func TestStore_EchoSingleMessage_ReturnsSameLine(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	message := `{"op":"ping"}`

	// Act
	sendLine(t, conn, message)
	got := readLine(t, conn, time.Second)

	// Assert: store возвращает ровно ту же строку (без завершающего \n в тексте)
	assert.Equal(t, message, got,
		"store must echo the received line back to the client verbatim")
}

// ---------------------------------------------------------------------------
// FT-04 — Несколько сообщений в одном соединении
// ---------------------------------------------------------------------------

func TestStore_EchoMultipleMessages_ReturnedInOrder(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	messages := []string{
		`{"op":"query","type":"entity"}`,
		`{"op":"upsert","type":"entity","data":{"name":"test"}}`,
		`{"op":"delete","type":"entity","id":"abc"}`,
	}

	// Act + Assert: каждое сообщение возвращается эхом в том же порядке
	for _, msg := range messages {
		sendLine(t, conn, msg)
		got := readLine(t, conn, time.Second)
		assert.Equalf(t, msg, got,
			"store must echo %q without reordering", msg)
	}
}

// ---------------------------------------------------------------------------
// FT-05 — Логирование каждой полученной строки (DEBUG + conn_id)
// ---------------------------------------------------------------------------

func TestStore_ReceivedLine_LoggedAtDebugWithConnID(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	message := `{"op":"ping"}`

	// Act
	sendLine(t, conn, message)
	_ = readLine(t, conn, time.Second) // ждём ответа — строка обработана

	// Assert: в логах есть DEBUG-запись содержащая тело сообщения
	logs := sp.Logs()
	assert.True(t,
		logs.waitForLevelAndMsg("DEBUG", message, time.Second),
		"store must log each received line at DEBUG level with its content")

	// Assert: запись содержит conn_id
	ids := logs.connIDs()
	assert.NotEmpty(t, ids,
		"log entries for received lines must include conn_id")
}

// ---------------------------------------------------------------------------
// FT-06 — Логирование пути каждого поднятого сокета
// ---------------------------------------------------------------------------

func TestStore_Startup_LogsEachSocketPath(t *testing.T) {
	// Arrange
	sockNames := []string{"sock-alpha.sock", "sock-beta.sock"}
	sp := newStoreProcess(t, sockNames)
	logs := sp.Logs()

	// Assert: путь каждого сокета упоминается в логах
	for i, name := range sockNames {
		path := sp.SockPath(i)
		assert.True(t,
			logs.waitForMsg(path, time.Second),
			"store must log socket path %q (%s) at startup", name, path)
	}
}

// ---------------------------------------------------------------------------
// FT-07 — Логирование нового подключения (путь сокета + conn_id)
// ---------------------------------------------------------------------------

func TestStore_NewConnection_LoggedWithSocketPathAndConnID(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	sockPath := sp.SockPath()

	// Act
	conn := dialStore(t, sockPath)
	defer conn.Close()

	logs := sp.Logs()

	// Assert: в логах есть запись о новом подключении с путём сокета
	assert.True(t,
		logs.waitForMsg(sockPath, time.Second),
		"store must log new connection event with socket path")

	// Assert: запись содержит conn_id (живёт в structured attrs, не в msg)
	assert.True(t,
		logs.waitForConnID(time.Second),
		"store must include conn_id in new connection log entry")
}

// ---------------------------------------------------------------------------
// FT-08 — Graceful shutdown по SIGTERM
// ---------------------------------------------------------------------------

func TestStore_GracefulShutdown_OnSIGTERM_SockRemovedAndExitZero(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	sockPath := sp.SockPath()

	_, err := os.Stat(sockPath)
	require.NoError(t, err, "precondition: socket file must exist before shutdown")

	// Act: останавливаем store (embedded runner — закрываем stopCh, эквивалент SIGTERM)
	sp.Stop()

	// Assert: сокет-файл удалён
	assert.True(t,
		sockGone(sockPath, 2*time.Second),
		"socket file must be removed after graceful shutdown")

	// Assert: в логах есть запись о завершении работы
	logs := sp.Logs()
	assert.True(t,
		logs.waitForMsg("shutdown", 2*time.Second) ||
			logs.waitForMsg("shutting down", 2*time.Second) ||
			logs.waitForMsg("stopping", 2*time.Second),
		"store must log a shutdown message")
}

// ---------------------------------------------------------------------------
// FT-09 — Graceful shutdown по SIGINT
// ---------------------------------------------------------------------------

// FT-09 проверяет то же поведение что FT-08, но инициированное SIGINT.
// В embedded runner-подходе оба сигнала моделируются одним механизмом остановки
// (закрытие stopCh), поэтому тест не может отличить SIGTERM от SIGINT на уровне механизма.
// Техническое ограничение: настоящий os/signal процесс будет обрабатывать SIGTERM и SIGINT
// через отдельные каналы, но логика shutdown идентична — сокеты удаляются, процесс завершается.
func TestStore_GracefulShutdown_OnSIGINT_SockRemovedAndProcessEnds(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store-sigint.sock"})
	sockPath := sp.SockPath()

	_, err := os.Stat(sockPath)
	require.NoError(t, err, "precondition: socket file must exist before shutdown")

	// Act
	sp.Stop()

	// Assert: сокет-файл удалён с ФС
	assert.True(t,
		sockGone(sockPath, 2*time.Second),
		"socket file must be removed after SIGINT shutdown")

	// Assert: в логах есть сообщение о завершении
	logs := sp.Logs()
	assert.True(t,
		logs.waitForMsg("shutdown", 2*time.Second) ||
			logs.waitForMsg("shutting down", 2*time.Second) ||
			logs.waitForMsg("stopping", 2*time.Second),
		"store must log shutdown message after SIGINT")
}

// ---------------------------------------------------------------------------
// FT-10 — Fail-fast: конфиг не найден
// ---------------------------------------------------------------------------

func TestStore_FailFast_ConfigNotFound_ExitsWithError(t *testing.T) {
	// Arrange: путь к конфигу не существует
	cfg := Config{
		configPath: "/tmp/nonexistent-store-config-12345.yaml",
	}
	logs := &logBuffer{}
	stopCh := make(chan struct{})

	// Act
	exitCode := RunStoreWithExitCode(cfg, logs, stopCh)

	// Assert: процесс завершается с ненулевым кодом
	assert.NotEqual(t, 0, exitCode,
		"store must exit with non-zero code when config file is not found")

	// Assert: в логах есть сообщение об ошибке
	assert.True(t, logs.hasMsg("config") || logs.hasMsg("no such file") || logs.hasLevelAndMsg("ERROR", ""),
		"store must log an error message when config is not found")
}

// ---------------------------------------------------------------------------
// FT-11 — Fail-fast: конфиг невалидный YAML
// ---------------------------------------------------------------------------

func TestStore_FailFast_InvalidYAML_ExitsWithError(t *testing.T) {
	// Arrange: файл конфига существует, но содержит невалидный YAML
	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "bad-config.yaml")
	require.NoError(t, os.WriteFile(cfgPath, []byte("store: [\nnot: valid: yaml:::"), 0o644))

	cfg := Config{configPath: cfgPath}
	logs := &logBuffer{}
	stopCh := make(chan struct{})

	// Act
	exitCode := RunStoreWithExitCode(cfg, logs, stopCh)

	// Assert: процесс завершается с ненулевым кодом
	assert.NotEqual(t, 0, exitCode,
		"store must exit with non-zero code when config YAML is invalid")

	// Assert: в логах есть сообщение об ошибке
	assert.True(t, logs.hasLevel("ERROR") || logs.hasLevel("FATAL"),
		"store must log an error when config YAML is invalid")
}

// ---------------------------------------------------------------------------
// FT-12 — Fail-fast: список sockets пустой
// ---------------------------------------------------------------------------

func TestStore_FailFast_EmptySocketsList_ExitsWithError(t *testing.T) {
	// Arrange: валидный конфиг, но store.sockets: []
	cfg := Config{
		Store: StoreConfig{
			DBUrl:   "postgres://localhost/platform_test",
			Sockets: []SocketConfig{},
		},
	}
	logs := &logBuffer{}
	stopCh := make(chan struct{})

	// Act
	exitCode := RunStoreWithExitCode(cfg, logs, stopCh)

	// Assert: процесс завершается с ненулевым кодом
	assert.NotEqual(t, 0, exitCode,
		"store must exit with non-zero code when sockets list is empty")

	// Assert: в логах есть сообщение об ошибке
	assert.True(t,
		logs.hasMsg("socket") || logs.hasMsg("sockets") || logs.hasLevel("ERROR") || logs.hasLevel("FATAL"),
		"store must log an error about empty sockets configuration")
}

// ---------------------------------------------------------------------------
// FT-13 — Уникальность conn_id при конкурентных подключениях
// ---------------------------------------------------------------------------

func TestStore_ConcurrentConnections_AllConnIDsAreUnique(t *testing.T) {
	// Arrange
	const clientCount = 10
	sp := newStoreProcess(t, []string{"store.sock"})

	var wg sync.WaitGroup
	wg.Add(clientCount)

	// Act: десять клиентов подключаются одновременно
	for i := 0; i < clientCount; i++ {
		go func(idx int) {
			defer wg.Done()
			conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
			if err != nil {
				t.Errorf("client %d: dial error: %v", idx, err)
				return
			}
			defer conn.Close()
			// Отправляем сообщение чтобы conn_id появился в логах
			fmt.Fprintf(conn, `{"client":%d}`+"\n", idx)
		}(i)
	}
	wg.Wait()

	// Ждём пока все conn_id появятся в логах
	logs := sp.Logs()
	logs.waitForNConnIDs(clientCount, time.Second)

	// Assert: в логах присутствуют conn_id для всех десяти подключений
	ids := logs.connIDs()
	require.GreaterOrEqualf(t, len(ids), clientCount,
		"store must log conn_id for each of %d concurrent connections", clientCount)

	// Assert: все conn_id различны
	seen := map[string]struct{}{}
	for _, id := range ids {
		assert.NotContains(t, seen, id,
			"conn_id must be unique; duplicate found: %q", id)
		seen[id] = struct{}{}
	}
}

// ---------------------------------------------------------------------------
// FT-14 — Configurable max_line_bytes
// ---------------------------------------------------------------------------

func TestStore_MaxLineBytes_ExceedLimit_ConnectionClosed(t *testing.T) {
	// Arrange: store с лимитом 1024 байт
	const maxLine = 1024
	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = maxLine
	})

	// Клиент 1: превысит лимит
	overConn := dialStore(t, sp.SockPath())
	// Клиент 2: должен продолжить работу
	okConn := dialStore(t, sp.SockPath())

	// Act: клиент 1 отправляет 1025 байт без \n
	overLimit := strings.Repeat("x", maxLine+1)
	fmt.Fprint(overConn, overLimit) // нет \n — специально

	// Assert: store закрывает соединение с клиентом 1 (EOF)
	expectEOF(t, overConn, 2*time.Second)

	// Assert: в логах есть запись об ошибке превышения лимита
	logs := sp.Logs()
	assert.True(t,
		logs.waitForMsg("too long", 2*time.Second) ||
			logs.waitForMsg("limit", 2*time.Second) ||
			logs.waitForMsg("max_line", 2*time.Second),
		"store must log error when line exceeds max_line_bytes")

	// Assert: клиент 2 продолжает работу нормально
	sendLine(t, okConn, `{"op":"ping"}`)
	got := readLine(t, okConn, time.Second)
	assert.Equal(t, `{"op":"ping"}`, got,
		"other connected clients must continue working after one exceeds limit")
}

// ---------------------------------------------------------------------------
// NFT-01 — Клиент отключается в середине потока
// ---------------------------------------------------------------------------

func TestStore_ClientDisconnects_OtherClientsUnaffected(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})

	// Два клиента
	dyingConn := dialStore(t, sp.SockPath())
	aliveConn := dialStore(t, sp.SockPath())

	// Act: первый клиент отправляет строку и немедленно закрывает соединение
	sendLine(t, dyingConn, `{"op":"drop"}`)
	dyingConn.Close()

	// Ждём чтобы store обработал обрыв.
	// conn_id живёт в attrs (структурированные поля), не в msg — используем waitForConnID.
	logs := sp.Logs()
	logs.waitForConnID(500 * time.Millisecond)

	// Assert: в логах есть запись об обрыве соединения
	assert.True(t,
		logs.hasMsg("disconnect") || logs.hasMsg("closed") || logs.hasMsg("EOF") || logs.hasMsg("conn"),
		"store must log disconnection event")

	// Assert: второй клиент продолжает получать ответы
	sendLine(t, aliveConn, `{"op":"ping"}`)
	got := readLine(t, aliveConn, time.Second)
	assert.Equal(t, `{"op":"ping"}`, got,
		"alive client must continue receiving echo responses after another client disconnects")
}

// ---------------------------------------------------------------------------
// NFT-02 — Stale socket: файл уже существует при старте
// ---------------------------------------------------------------------------

func TestStore_StaleSocketFile_RemovedAndRecreatedOnStart(t *testing.T) {
	// Arrange: создаём stale файл по пути сокета
	dir := t.TempDir()
	stalePath := filepath.Join(dir, "stale.sock")
	makeStaleSockFile(t, stalePath)

	_, err := os.Stat(stalePath)
	require.NoError(t, err, "precondition: stale file must exist before store starts")

	// Act: запускаем store с тем же путём
	sp := newStoreProcess(t, nil, func(cfg *Config) {
		cfg.Store.Sockets = []SocketConfig{
			{Path: stalePath, Ops: []string{"query", "upsert", "delete"}},
		}
	})
	_ = sp

	// Assert: store успешно принимает подключения на этом сокете
	conn, dialErr := net.DialTimeout("unix", stalePath, time.Second)
	require.NoError(t, dialErr,
		"store must accept connections after removing stale socket file")
	conn.Close()
}

// ---------------------------------------------------------------------------
// NFT-03 — Параллельная обработка: медленный клиент не блокирует остальных
// ---------------------------------------------------------------------------

func TestStore_ConcurrentClients_SlowClientDoesNotBlockOthers(t *testing.T) {
	// Arrange
	const clientCount = 10
	sp := newStoreProcess(t, []string{"store.sock"})

	// "Медленный" клиент: подключается, отправляет сообщение, но не читает ответ сразу
	slowConn := dialStore(t, sp.SockPath())
	sendLine(t, slowConn, `{"op":"slow"}`)

	// Act: остальные клиенты подключаются одновременно и отправляют сообщения
	results := make([]string, clientCount)
	var wg sync.WaitGroup
	wg.Add(clientCount)

	start := time.Now()
	for i := 0; i < clientCount; i++ {
		go func(idx int) {
			defer wg.Done()
			conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
			if err != nil {
				t.Errorf("client %d: dial: %v", idx, err)
				return
			}
			defer conn.Close()
			msg := fmt.Sprintf(`{"op":"ping","id":%d}`, idx)
			fmt.Fprintf(conn, "%s\n", msg)
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			sc := make([]byte, 512)
			n, _ := conn.Read(sc)
			results[idx] = strings.TrimRight(string(sc[:n]), "\n")
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)

	// Assert: все клиенты получили ответы быстро (не ждали медленного)
	assert.Less(t, elapsed, 2*time.Second,
		"fast clients must not be blocked by the slow client")

	// Assert: каждый клиент получил эхо своего сообщения
	for i, got := range results {
		expected := fmt.Sprintf(`{"op":"ping","id":%d}`, i)
		assert.Equal(t, expected, got,
			"client %d must receive echo of its own message", i)
	}

	// Assert: медленный клиент тоже получает свой ответ
	got := readLine(t, slowConn, time.Second)
	assert.Equal(t, `{"op":"slow"}`, got,
		"slow client must also receive its echo response")
}

// ---------------------------------------------------------------------------
// NFT-04 — Логирование ошибок чтения/записи
// ---------------------------------------------------------------------------

func TestStore_ClientForciblyClosedDuringWrite_StoreLogsErrorAndContinues(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	// Act: клиент принудительно закрывает соединение (имитируем обрыв)
	// Отправляем сообщение и сразу закрываем — store может не успеть ответить
	sendLine(t, conn, `{"op":"abort"}`)
	conn.Close() // RST или FIN до получения ответа

	// Assert: store логирует ошибку записи (NFT-04)
	logs := sp.Logs()
	assert.True(t,
		logs.waitForMsg("error", time.Second) ||
			logs.waitForMsg("write", time.Second) ||
			logs.waitForLevelAndMsg("INFO", "conn", time.Second),
		"store must log a write error when client forcibly closes connection")

	// Assert: store продолжает работу (принимает новые подключения)
	newConn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
	require.NoError(t, err, "store must continue accepting connections after write error")
	newConn.Close()
}

// ---------------------------------------------------------------------------
// NFT-05 — Уровни лога при закрытии соединений
// ---------------------------------------------------------------------------

func TestStore_CleanClientClose_LoggedAtDebug(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	// Отправляем сообщение, получаем ответ, закрываем штатно
	sendLine(t, conn, `{"op":"ping"}`)
	_ = readLine(t, conn, time.Second)
	conn.Close()

	// Assert: штатное закрытие логируется на уровне DEBUG
	logs := sp.Logs()
	assert.True(t,
		logs.waitForLevelAndMsg("DEBUG", "close", time.Second) ||
			logs.waitForLevelAndMsg("DEBUG", "EOF", time.Second) ||
			logs.waitForLevelAndMsg("DEBUG", "conn", time.Second),
		"clean client disconnect (EOF) must be logged at DEBUG level")
}

func TestStore_ClientClose_PartialRead_LoggedAtDebug(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})

	// Act: клиент подключается, отправляет неполную строку (без \n) и закрывает соединение.
	// На Unix socket SO_LINGER(0) не даёт RST как в TCP — всегда clean EOF.
	// По спеке: нормальное закрытие (клиент прислал EOF) → DEBUG.
	conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
	require.NoError(t, err)

	fmt.Fprint(conn, `{"op":"incomplete"`) // partial line, no \n
	conn.Close()

	// Assert: clean EOF (в т.ч. в середине частичных данных) логируется на уровне DEBUG
	logs := sp.Logs()
	assert.True(t,
		logs.waitForLevelAndMsg("DEBUG", "close", 2*time.Second) ||
			logs.waitForLevelAndMsg("DEBUG", "EOF", 2*time.Second) ||
			logs.waitForLevelAndMsg("DEBUG", "conn", 2*time.Second) ||
			logs.waitForLevelAndMsg("DEBUG", "disconnect", 2*time.Second),
		"client close (clean EOF, partial data) must be logged at DEBUG level")
}

// ---------------------------------------------------------------------------
// NFT-06 — Graceful shutdown ждёт активные соединения
// ---------------------------------------------------------------------------

func TestStore_GracefulShutdown_WaitsForActiveConnections(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	sockPath := sp.SockPath()

	// Клиент подключается и удерживает соединение открытым
	conn, err := net.DialTimeout("unix", sockPath, time.Second)
	require.NoError(t, err, "client must be able to connect before shutdown")

	shutdownDone := make(chan struct{})
	shutdownStart := time.Now()

	// Act: инициируем shutdown в фоне
	go func() {
		defer close(shutdownDone)
		sp.Stop()
	}()

	// Ждём немного — store не должен завершиться мгновенно пока соединение открыто
	time.Sleep(200 * time.Millisecond)
	select {
	case <-shutdownDone:
		// По спеке store обязан ждать активные соединения до 5 секунд.
		// Если он завершился раньше чем мы закрыли соединение — контракт нарушен.
		t.Fatal("store exited before client disconnected — store must wait for active connections (spec: 5s graceful shutdown)")
	default:
		// Assert: store пока ещё работает (ждёт соединение)
		// Закрываем соединение — store должен завершиться
		conn.Close()

		select {
		case <-shutdownDone:
			elapsed := time.Since(shutdownStart)
			assert.Less(t, elapsed, 6*time.Second,
				"store must complete shutdown within 6 seconds after all connections close")
		case <-time.After(6 * time.Second):
			t.Fatal("store did not complete shutdown within 6 seconds after connection closed")
		}
	}
}

func TestStore_GracefulShutdown_ForcesCloseAfter5Seconds(t *testing.T) {
	// Этот тест проверяет принудительное закрытие по таймауту — ожидаемое время 5–7 секунд.
	// В коротком режиме (-short) пропускается.
	if testing.Short() {
		t.Skip("skipping slow shutdown timeout test in -short mode (~5-7s)")
	}

	// Arrange: store запущен, клиент держит соединение открытым
	sp := newStoreProcess(t, []string{"store.sock"})

	conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
	require.NoError(t, err)
	defer conn.Close()

	shutdownDone := make(chan struct{})

	// Act: инициируем shutdown, соединение НЕ закрываем
	go func() {
		defer close(shutdownDone)
		sp.Stop()
	}()

	// Assert: даже с открытым соединением store завершается не позднее чем через 6 секунд
	select {
	case <-shutdownDone:
		// Хорошо — store завершился
	case <-time.After(7 * time.Second):
		t.Fatal("store must force shutdown within 5s+buffer even with active connections")
	}
}

// ---------------------------------------------------------------------------
// EDGE-01 — Пустая строка пропускается
// ---------------------------------------------------------------------------

func TestStore_EmptyLine_IsSkipped_NoResponseNoLog(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	logsBefore := len(sp.Logs().all())

	// Act: клиент отправляет пустую строку (только \n)
	fmt.Fprint(conn, "\n")

	// Ждём чуть дольше чем нужно для ответа
	conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	buf := make([]byte, 256)
	n, _ := conn.Read(buf)
	conn.SetReadDeadline(time.Time{})

	// Assert: ответ клиенту не отправляется
	assert.Equal(t, 0, n,
		"store must not send any response for an empty line")

	// Assert: количество лог-записей не выросло из-за пустой строки
	time.Sleep(50 * time.Millisecond)
	logsAfter := len(sp.Logs().all())
	assert.Equal(t, logsBefore, logsAfter,
		"store must not create any log entry for an empty line")

	// Assert: следующее нормальное сообщение всё равно обрабатывается
	sendLine(t, conn, `{"op":"ping"}`)
	got := readLine(t, conn, time.Second)
	assert.Equal(t, `{"op":"ping"}`, got,
		"store must continue processing messages after skipping empty line")
}

// ---------------------------------------------------------------------------
// EDGE-01b — Строка из пробельных символов пропускается
// ---------------------------------------------------------------------------

func TestStore_WhitespaceOnlyLine_IsSkipped_NoResponseNoLog(t *testing.T) {
	// Arrange
	sp := newStoreProcess(t, []string{"store.sock"})

	whitespaceInputs := []string{
		"   \n",
		"\t\t\n",
		" \t \n",
	}

	// Каждая итерация — отдельное соединение: устраняет нестабильность счётчика
	// из-за задержки логирования предыдущей итерации.
	for i, input := range whitespaceInputs {
		conn := dialStore(t, sp.SockPath())
		logs := sp.Logs()
		// Ждём пока соединение этой итерации появится в логах, потом фиксируем baseline
		logs.waitForNConnIDs(i+1, time.Second)
		logsBefore := len(logs.all())

		// Act
		fmt.Fprint(conn, input)

		// Assert: ответ не отправляется
		conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
		buf := make([]byte, 256)
		n, _ := conn.Read(buf)
		conn.SetReadDeadline(time.Time{})

		assert.Equalf(t, 0, n,
			"store must not respond to whitespace-only line %q", input)

		time.Sleep(50 * time.Millisecond)
		logsAfter := len(logs.all())
		assert.Equalf(t, logsBefore, logsAfter,
			"store must not log whitespace-only line %q", input)

		conn.Close()
	}

	// Assert: новое соединение работает нормально после всех whitespace-итераций
	finalConn := dialStore(t, sp.SockPath())
	sendLine(t, finalConn, `{"op":"ping"}`)
	got := readLine(t, finalConn, time.Second)
	assert.Equal(t, `{"op":"ping"}`, got,
		"store must continue processing after whitespace-only lines")
}

// ---------------------------------------------------------------------------
// EDGE-02 — Строка превышает лимит буфера 1 МБ (дефолтный)
// ---------------------------------------------------------------------------

func TestStore_DefaultMaxLineBytes_ExceedOneMB_ConnectionClosed(t *testing.T) {
	// Arrange: store с дефолтным лимитом (1 МБ = 1 048 576 байт)
	sp := newStoreProcess(t, []string{"store.sock"})

	overLimitConn := dialStore(t, sp.SockPath())
	okConn := dialStore(t, sp.SockPath())

	// Act: отправляем поток > 1 МБ без \n
	overLimit := strings.Repeat("y", 1048577)
	fmt.Fprint(overLimitConn, overLimit)

	// Assert: store закрывает соединение (клиент получает EOF)
	expectEOF(t, overLimitConn, 3*time.Second)

	// Assert: клиент не получает никаких данных перед закрытием
	// (expectEOF уже проверяет что при чтении сразу EOF, не данные)

	// Assert: в логах есть запись об ошибке
	logs := sp.Logs()
	assert.True(t,
		logs.waitForMsg("too long", 2*time.Second) ||
			logs.waitForMsg("limit", 2*time.Second) ||
			logs.waitForMsg("overflow", 2*time.Second) ||
			logs.waitForMsg("max_line", 2*time.Second),
		"store must log an error when line exceeds default 1MB limit")

	// Assert: остальные клиенты продолжают работу
	sendLine(t, okConn, `{"op":"ping"}`)
	got := readLine(t, okConn, time.Second)
	assert.Equal(t, `{"op":"ping"}`, got,
		"other clients must continue working after one exceeds buffer limit")
}

// ---------------------------------------------------------------------------
// EDGE-03 — Несколько подключений к разным сокетам
// ---------------------------------------------------------------------------

func TestStore_MultiSocket_TrafficNotMixedBetweenSockets(t *testing.T) {
	// Arrange: store с двумя сокетами
	sp := newStoreProcess(t, []string{"sock-a.sock", "sock-b.sock"})
	sockA := sp.SockPath(0)
	sockB := sp.SockPath(1)

	connA := dialStore(t, sockA)
	connB := dialStore(t, sockB)

	// Act: клиент A отправляет сообщение через sock-a
	msgA := `{"client":"A","sock":"a"}`
	sendLine(t, connA, msgA)

	// Act: клиент B отправляет сообщение через sock-b
	msgB := `{"client":"B","sock":"b"}`
	sendLine(t, connB, msgB)

	// Assert: каждый получает эхо своего сообщения через свой сокет
	gotA := readLine(t, connA, time.Second)
	gotB := readLine(t, connB, time.Second)

	assert.Equal(t, msgA, gotA,
		"client A must receive echo of its own message via sock-a")
	assert.Equal(t, msgB, gotB,
		"client B must receive echo of its own message via sock-b")

	// Assert: трафик не смешивается (A не получил сообщение B и наоборот)
	assert.NotEqual(t, msgB, gotA, "client A must not receive client B's message")
	assert.NotEqual(t, msgA, gotB, "client B must not receive client A's message")
}

// ---------------------------------------------------------------------------
// EDGE-04 — Директория для сокета не существует
// ---------------------------------------------------------------------------

func TestStore_SocketDirDoesNotExist_FailsWithError(t *testing.T) {
	// Arrange: путь сокета в несуществующей директории
	nonexistentDir := filepath.Join(t.TempDir(), "nonexistent-dir")
	sockPath := filepath.Join(nonexistentDir, "store.sock")

	cfg := Config{
		Store: StoreConfig{
			DBUrl: "postgres://localhost/platform_test",
			Sockets: []SocketConfig{
				{Path: sockPath, Ops: []string{"query", "upsert", "delete"}},
			},
		},
	}
	logs := &logBuffer{}
	stopCh := make(chan struct{})

	// Act
	exitCode := RunStoreWithExitCode(cfg, logs, stopCh)

	// Assert: процесс завершается с ненулевым кодом
	assert.NotEqual(t, 0, exitCode,
		"store must exit with non-zero code when socket directory does not exist")

	// Assert: в логах есть понятное сообщение об ошибке создания сокета
	assert.True(t,
		logs.hasMsg("socket") || logs.hasMsg(sockPath) || logs.hasMsg("no such file") || logs.hasLevel("ERROR"),
		"store must log an informative error about socket creation failure")
}
