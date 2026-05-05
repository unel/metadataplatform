package adversarial_test

// AC-ROUTER-05—10
// Роутер — граница между "клиент облажался" и "поток испорчен". Разница критична.
// AC-ROUTER-05—08: application-level ошибки, соединение живёт.
// AC-ROUTER-09—10: разница parse vs data error.

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// AC-ROUTER-05: upsert без поля data — INVALID_REQUEST, соединение живёт.
func TestRouter_Upsert_MissingData_ReturnsInvalidRequest_ConnectionAlive(t *testing.T) {
	conn := newTestRouter(t, nil)

	sendLine(t, conn, `{"op":"upsert","type":"entity"}`)
	resp := readResponse(t, conn)

	assert.False(t, resp.OK)
	assert.Equal(t, "INVALID_REQUEST", resp.ErrorCode)
	assert.Contains(t, resp.Error, "data is required")

	// соединение живёт — следующий запрос обрабатывается нормально
	sendLine(t, conn, `{"op":"list","type":"entity"}`)
	resp2 := readResponse(t, conn)
	assert.True(t, resp2.OK, "соединение должно быть живым после application-level ошибки")
}

// AC-ROUTER-06: data=null — INVALID_REQUEST, соединение сохраняется.
func TestRouter_Upsert_DataNull_ReturnsInvalidRequest_ConnectionAlive(t *testing.T) {
	conn := newTestRouter(t, nil)

	sendLine(t, conn, `{"op":"upsert","type":"entity","data":null}`)
	resp := readResponse(t, conn)

	assert.False(t, resp.OK)
	assert.Equal(t, "INVALID_REQUEST", resp.ErrorCode)

	sendLine(t, conn, `{"op":"list","type":"entity"}`)
	resp2 := readResponse(t, conn)
	assert.True(t, resp2.OK, "соединение живёт после data=null")
}

// AC-ROUTER-07: неизвестный op — UNKNOWN_OP, соединение сохраняется.
func TestRouter_UnknownOp_ReturnsUnknownOpError_ConnectionAlive(t *testing.T) {
	conn := newTestRouter(t, nil)

	sendLine(t, conn, `{"op":"patch","type":"entity","id":"x"}`)
	resp := readResponse(t, conn)

	assert.False(t, resp.OK)
	assert.Equal(t, "UNKNOWN_OP", resp.ErrorCode)
	assert.Contains(t, resp.Error, "patch", "ошибка должна содержать неизвестный op")

	sendLine(t, conn, `{"op":"list","type":"entity"}`)
	resp2 := readResponse(t, conn)
	assert.True(t, resp2.OK, "соединение живёт после неизвестного op")
}

// AC-ROUTER-08: неизвестный type — UNKNOWN_TYPE, соединение сохраняется.
func TestRouter_UnknownType_ReturnsUnknownTypeError_ConnectionAlive(t *testing.T) {
	conn := newTestRouter(t, nil)

	sendLine(t, conn, `{"op":"get","type":"table","id":"x"}`)
	resp := readResponse(t, conn)

	assert.False(t, resp.OK)
	assert.Equal(t, "UNKNOWN_TYPE", resp.ErrorCode)
	assert.Contains(t, resp.Error, "table", "ошибка должна содержать неизвестный type")

	sendLine(t, conn, `{"op":"list","type":"entity"}`)
	resp2 := readResponse(t, conn)
	assert.True(t, resp2.OK, "соединение живёт после неизвестного type")
}

// AC-ROUTER-09: невалидный JSON — поток испорчён, соединение закрывается, лог ERROR.
// Фундаментальное отличие от AC-ROUTER-05—08: json.Decoder не восстанавливается.
func TestRouter_InvalidJSON_ClosesConnection_LogsParseError(t *testing.T) {
	var logBuf bytes.Buffer
	conn := newTestRouter(t, &logBuf)

	fmt.Fprintf(conn, "{not-valid-json\n")

	// ожидаем что соединение закрыто — Read вернёт EOF или closed error
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	buf := make([]byte, 256)
	n, err := conn.Read(buf)
	if n > 0 {
		t.Logf("получили неожиданные данные: %s", buf[:n])
	}
	assert.True(t, err == io.EOF || isConnClosed(err),
		"ожидаем EOF или closed соединение после невалидного JSON, получили: %v", err)

	// лог ERROR с обязательными полями op=unknown и parse_error
	logOutput := logBuf.String()
	assert.Contains(t, logOutput, "ERROR", "ожидаем уровень ERROR в логе")
	assert.Contains(t, logOutput, "op=unknown", "ожидаем op=unknown в лог-записи")
	assert.Contains(t, logOutput, "parse_error", "ожидаем поле parse_error в лог-записи")
}

// AC-ROUTER-10: невалидный JSON в data — application level, соединение живёт.
// envelope распарсился нормально, мусор только в data — это другой уровень ошибки.
func TestRouter_InvalidJSONInData_ReturnsInvalidRequest_ConnectionAlive(t *testing.T) {
	conn := newTestRouter(t, nil)

	// data — строка-значение, а не объект; пройдёт первый decode, упадёт на втором
	sendLine(t, conn, `{"op":"upsert","type":"entity","data":"{broken"}`)
	resp := readResponse(t, conn)

	assert.False(t, resp.OK)
	assert.Equal(t, "INVALID_REQUEST", resp.ErrorCode)

	// соединение живёт — следующий запрос обрабатывается
	sendLine(t, conn, `{"op":"list","type":"entity"}`)
	resp2 := readResponse(t, conn)
	assert.True(t, resp2.OK, "соединение живёт после невалидного JSON в data")
}

// AC-ROUTER-16: ошибка операции — уровень ERROR, поля op/type/id/error.
func TestRouter_GetNonexistentEntity_LogsError(t *testing.T) {
	var logBuf bytes.Buffer
	conn := newTestRouter(t, &logBuf)

	sendLine(t, conn, `{"op":"get","type":"entity","id":"ent-999"}`)
	resp := readResponse(t, conn)

	// application-level ошибка — соединение живёт
	assert.False(t, resp.OK)
	assert.Equal(t, "NOT_FOUND", resp.ErrorCode)

	logOutput := logBuf.String()
	assert.Contains(t, logOutput, "ERROR", "ожидаем уровень ERROR в логе")
	assert.Contains(t, logOutput, "op=get", "ожидаем op=get в лог-записи")
	assert.Contains(t, logOutput, "type=entity", "ожидаем type=entity в лог-записи")
	assert.Contains(t, logOutput, "id=ent-999", "ожидаем id в лог-записи")
	assert.Contains(t, logOutput, "error=", "ожидаем поле error в лог-записи")
}
