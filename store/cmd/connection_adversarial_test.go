package main

// connection_adversarial_test.go — adversarial тесты для store/connection.
//
// Acceptance-сценарии (FT/NFT/EDGE из acceptance.md) — в connection_test.go.
// Здесь: граничные условия которые спека упоминает вскользь или молчит,
// race conditions, численные границы, входные данные из ада.
//
// "Это никогда не случится в проде" — любимая цитата перед следующим постмортемом.

import (
	"fmt"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ─── ГРАНИЦА max_line_bytes ───────────────────────────────────────────────────

// Ровно на лимите — должно пройти.
// Спека говорит "больше max_line_bytes → закрыть". Ровно на лимите — это не больше.
func TestMaxLineBytes_ExactlyAtLimit_Accepted(t *testing.T) {
	const limit = 256
	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = limit
	})

	conn := dialStore(t, sp.SockPath())
	line := strings.Repeat("z", limit)
	sendLine(t, conn, line) // ровно limit байт + \n

	reply := readLine(t, conn, time.Second)
	assert.Equal(t, line, reply, "строка ровно в лимит должна пройти и вернуться эхом")
}

// Лимит + 1 байт без \n → должен упасть.
// Это off-by-one который люди регулярно путают.
func TestMaxLineBytes_OneByteBeyondLimit_Rejected(t *testing.T) {
	const limit = 256
	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = limit
	})

	conn := dialStore(t, sp.SockPath())
	payload := strings.Repeat("z", limit+1) // без \n
	_, _ = fmt.Fprint(conn, payload)

	expectEOF(t, conn, 2*time.Second)
}

// Минимально возможный лимит (1 байт): строка из одного символа проходит,
// два символа — нет.
func TestMaxLineBytes_MinimalLimit_OneByte(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = 1
	})

	// один байт — должен пройти
	conn1 := dialStore(t, sp.SockPath())
	sendLine(t, conn1, "x")
	reply := readLine(t, conn1, time.Second)
	assert.Equal(t, "x", reply, "строка из 1 байта при лимите 1 должна пройти")

	// два байта без \n — должен упасть
	conn2 := dialStore(t, sp.SockPath())
	_, _ = fmt.Fprint(conn2, "ab") // 2 байта без \n
	expectEOF(t, conn2, 2*time.Second)
}

// ─── ПРОПУСКНАЯ СПОСОБНОСТЬ ──────────────────────────────────────────────────

// 1000 строк подряд в одном соединении — ни одна не должна потеряться.
// Если store проглатывает строки при высокой нагрузке — это хороший момент это узнать.
//
// Запись и чтение идут параллельно — иначе дедлок:
// unix socket buffer конечен (~208 КБ), store пишет эхо синхронно,
// при синхронной записи 1000 строк буфер забивается и оба блокируются навечно.
func TestHighThroughput_1000Lines_NoneDropped(t *testing.T) {
	if testing.Short() {
		t.Skip("пропускаем в -short режиме")
	}

	const n = 1000
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	writeErr := make(chan error, 1)

	// запись в горутине — не блокирует чтение
	go func() {
		for i := range n {
			_, err := fmt.Fprintf(conn, `{"seq":%d}`+"\n", i)
			if err != nil {
				writeErr <- fmt.Errorf("write seq=%d: %w", i, err)
				return
			}
		}
		writeErr <- nil
	}()

	// читаем ответы параллельно с записью
	received := 0
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	scanner := newTestScanner(conn)
	for scanner.Scan() {
		received++
		if received == n {
			break
		}
	}

	require.NoError(t, <-writeErr, "ошибка записи")
	assert.Equal(t, n, received, "store должен вернуть эхо на каждую из %d строк", n)
}

// ─── СОДЕРЖИМОЕ СТРОК ────────────────────────────────────────────────────────

// Спека явно говорит: "возвращает строку as-is вне зависимости от содержимого".
// Проверяем что store не пытается парсить JSON уже сейчас.
func TestNonJSON_EchoedAsIs(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	garbage := "not json at all ¯\\_(ツ)_/¯"
	sendLine(t, conn, garbage)
	reply := readLine(t, conn, time.Second)
	assert.Equal(t, garbage, reply, "store должен возвращать строку as-is")
}

// Строка только из пробелов — проверяем что нет ответа (тайм-аутом).
// connection_test.go проверяет через счётчик логов; здесь проверяем
// именно что байты в сеть не летят.
func TestWhitespaceOnlyLine_ZeroBytesReceived(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	_, _ = fmt.Fprint(conn, "   \n")

	conn.SetReadDeadline(time.Now().Add(200 * time.Millisecond))
	buf := make([]byte, 64)
	n, _ := conn.Read(buf)
	assert.Equal(t, 0, n, "whitespace-строка не должна вызывать ответ от store")

	_ = sp
}

// Null байты в строке — store не должен крашиться.
// bufio.Scanner читает до \n — null байт это просто данные.
// Но некоторые реализации с этим не справляются.
func TestNullBytes_InLine_NotCrashing(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})
	conn := dialStore(t, sp.SockPath())

	line := "before\x00after"
	_, _ = fmt.Fprintf(conn, "%s\n", line)

	// главное что store не крашится; ответ — как получится
	conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond))
	buf := make([]byte, 256)
	conn.Read(buf)

	// store должен принимать следующее соединение
	conn2 := dialStore(t, sp.SockPath())
	sendLine(t, conn2, `{"alive":true}`)
	reply := readLine(t, conn2, time.Second)
	assert.Equal(t, `{"alive":true}`, reply, "store должен выжить после null-байта в строке")
}

// Очень длинная строка ровно у лимита — и сразу за ней нормальная строка.
// Проверяем что reader не "съедает" следующую строку после переполнения.
func TestMaxLineBytes_AfterOverflow_NextConnectionWorks(t *testing.T) {
	const limit = 128
	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = limit
	})

	// первое соединение: overflow → EOF
	badConn := dialStore(t, sp.SockPath())
	_, _ = fmt.Fprint(badConn, strings.Repeat("x", limit+1))
	expectEOF(t, badConn, 2*time.Second)

	// второе соединение — должно работать нормально
	goodConn := dialStore(t, sp.SockPath())
	sendLine(t, goodConn, `{"seq":"after_overflow"}`)
	reply := readLine(t, goodConn, time.Second)
	assert.Equal(t, `{"seq":"after_overflow"}`, reply,
		"новое соединение после overflow должно работать нормально",
	)
}

// ─── CONN_ID АТОМАРНОСТЬ ─────────────────────────────────────────────────────

// conn_id монотонно возрастает — проверяем что при последовательных подключениях
// id не повторяются и в логах они различимы.
func TestConnID_Sequential_Unique(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})

	for i := range 5 {
		conn := dialStore(t, sp.SockPath())
		sendLine(t, conn, fmt.Sprintf(`{"seq":%d}`, i))
		readLine(t, conn, time.Second)
		conn.Close()
		time.Sleep(10 * time.Millisecond) // даём время залогировать
	}

	ids := sp.Logs().connIDs()
	require.GreaterOrEqual(t, len(ids), 5,
		"должны найти conn_id для 5 последовательных подключений",
	)

	seen := map[string]struct{}{}
	for _, id := range ids {
		_, dup := seen[id]
		assert.False(t, dup, "conn_id %q встречается дважды — нарушена уникальность", id)
		seen[id] = struct{}{}
	}
}

// ─── SHUTDOWN ГРАНИЧНЫЕ СЛУЧАИ ────────────────────────────────────────────────

// После shutdown новые подключения не принимаются.
// connection_test.go проверяет что сокет удалён; здесь проверяем именно
// что dial возвращает ошибку — не просто файла нет.
func TestGracefulShutdown_NoNewConnectionsAccepted(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})
	sockPath := sp.SockPath()

	sp.Stop()

	// ждём пока сокет удалится
	if !sockGone(sockPath, 2*time.Second) {
		t.Log("сокет-файл не удалился за 2с, но проверяем подключение")
	}

	conn, err := net.DialTimeout("unix", sockPath, 200*time.Millisecond)
	if err == nil {
		conn.Close()
		t.Fatal("не должны были подключиться после shutdown")
	}
	// ожидаем ошибку подключения — это правильное поведение
}

// Shutdown с нулём активных соединений — должен завершиться быстро.
// Не ждать 5 секунд без причины.
func TestGracefulShutdown_NoActiveConns_CompletesQuickly(t *testing.T) {
	sp := newStoreProcess(t, []string{"store.sock"})

	start := time.Now()
	sp.Stop()
	elapsed := time.Since(start)

	assert.Less(t, elapsed, 2*time.Second,
		"shutdown без активных соединений должен завершиться за <2с, прошло: %v", elapsed,
	)
}

// ─── RACE CONDITIONS ─────────────────────────────────────────────────────────

// Одновременный overflow от нескольких клиентов — store не должен крашиться.
// Любимый сценарий для data race в обработке ошибок.
func TestConcurrentOverflows_NoCrash(t *testing.T) {
	const limit = 64
	const nClients = 5

	sp := newStoreProcess(t, []string{"store.sock"}, func(cfg *Config) {
		cfg.Store.MaxLineBytes = limit
	})

	done := make(chan struct{}, nClients)
	for range nClients {
		go func() {
			defer func() { done <- struct{}{} }()
			conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
			if err != nil {
				return
			}
			defer conn.Close()
			_, _ = fmt.Fprint(conn, strings.Repeat("X", limit+1))
			// ждём EOF или ошибку
			conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			buf := make([]byte, 1)
			conn.Read(buf)
		}()
	}

	// ждём завершения всех горутин
	for range nClients {
		select {
		case <-done:
		case <-time.After(5 * time.Second):
			t.Fatal("клиент не завершился за 5с")
		}
	}

	// store должен выжить
	conn := dialStore(t, sp.SockPath())
	sendLine(t, conn, `{"survived":true}`)
	reply := readLine(t, conn, time.Second)
	assert.Equal(t, `{"survived":true}`, reply,
		"store должен выжить после конкурентных overflow",
	)
}

// Подключение и отключение в быстром цикле — проверяем что нет утечек горутин
// или дедлоков при высокой частоте connect/disconnect.
func TestRapidConnectDisconnect_NoCrash(t *testing.T) {
	if testing.Short() {
		t.Skip("пропускаем в -short режиме")
	}

	const iterations = 50
	sp := newStoreProcess(t, []string{"store.sock"})

	for i := range iterations {
		conn, err := net.DialTimeout("unix", sp.SockPath(), time.Second)
		if err != nil {
			t.Fatalf("iteration %d: dial failed: %v", i, err)
		}
		conn.Close() // сразу закрываем
	}

	// после 50 быстрых connect/disconnect store должен работать
	conn := dialStore(t, sp.SockPath())
	sendLine(t, conn, `{"final":"check"}`)
	reply := readLine(t, conn, time.Second)
	assert.Equal(t, `{"final":"check"}`, reply,
		"store должен работать после %d быстрых connect/disconnect", iterations,
	)
}
