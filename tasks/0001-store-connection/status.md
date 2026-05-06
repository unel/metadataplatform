spec-ft: done
spec-nft: done
spec-review: done
spec: done
acceptance-write: done
acceptance-review: done
test-write: done
test-review: done
implement: done
build: done
test-run: done
code-review: done
tech-doc: done

## test-run fix log

### TestGracefulShutdown_NoNewConnectionsAccepted — panic: close of closed channel

Файл: cmd/store/testhelpers_test.go

Причина: `t.Cleanup` в `newStoreProcess` безусловно вызывал `close(runner.stopCh)`. Если тест уже вызвал `sp.Stop()` — канал закрывался повторно → паника.

Фикс: `t.Cleanup` теперь вызывает `sp.Stop()` вместо прямого `close(runner.stopCh)`. `Stop()` уже идемпотентен через `select { case <-sp.runner.stopCh: ... default: close(...) }` — двойное закрытие исключено.
