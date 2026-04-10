# Таблично-управляемые тесты

Стандартный паттерн Go для тестирования нескольких сценариев одной функции.

```go
func TestFindEntity(t *testing.T) {
    tests := []struct {
        name    string
        id      string
        want    *Entity
        wantErr error
    }{
        {
            name: "найдена существующая entity",
            id:   "uuid-123",
            want: &Entity{ID: "uuid-123", Type: "file"},
        },
        {
            name:    "не найдена — возвращает ErrNotFound",
            id:      "uuid-999",
            wantErr: ErrNotFound,
        },
        {
            name:    "пустой id — возвращает ошибку валидации",
            id:      "",
            wantErr: ErrInvalidID,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := store.FindEntity(context.Background(), tt.id)
            if tt.wantErr != nil {
                if !errors.Is(err, tt.wantErr) {
                    t.Errorf("got error %v, want %v", err, tt.wantErr)
                }
                return
            }
            if err != nil {
                t.Fatalf("unexpected error: %v", err)
            }
            if got.ID != tt.want.ID {
                t.Errorf("got ID %q, want %q", got.ID, tt.want.ID)
            }
        })
    }
}
```

## Субтесты и параллельность

```go
t.Run(tt.name, func(t *testing.T) {
    t.Parallel() // запускать параллельно если тесты независимы
    ...
})
```
