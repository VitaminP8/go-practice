# Тестирование в Go

Go предоставляет встроенный пакет `testing` для написания тестов. Тесты представляют собой функции с названием, начинающимся на `Test`, и принимают `*testing.T` в качестве аргумента.

## Основы тестирования

В тестах можно использовать методы `t.Errorf` и `t.Fatalf`:

- `t.Errorf(format, args...)` — выводит ошибку, но продолжает выполнение теста.
- `t.Fatalf(format, args...)` — завершает тест сразу после обнаружения ошибки.

Также есть методы:
- `t.Log` — выводит отладочную информацию.
- `t.Skip` — пропускает тест с указанием причины.

## Написание тестов

```go
package main

import (
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}
```

## Запуск тестов

Тесты выполняются с помощью команды:

```sh
go test
```

Для более подробного вывода:

```sh
go test -v
```

Все тесты в пакете и подпакетах:

```sh
go test ./...
go test ./pkg1/...
go test github.com/otus/superapp/...
```

Конкретные тесты по имени:
```sh
go test -run TestFoo
```
По тегам ( //go:build integration ):
```sh
go test -tags=integration
```

## Табличные тесты

Табличные тесты позволяют протестировать несколько случаев в одном тесте:

```go
func TestSumTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b, expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"zero values", 0, 0, 0},
        {"negative numbers", -1, -1, -2},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
        result := Sum(tt.a, tt.b)
        if result != tt.expected {
            t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
        }
        })
    }
}
```

Использование `t.Run` позволяет выполнять каждый тест в отдельной подфункции, что улучшает читаемость вывода тестов и упрощает их отладку.

## Бенчмаркинг

Go позволяет измерять производительность кода с помощью бенчмарков:

```go
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(2, 3)
	}
}
```

Запуск бенчмарков:

```sh
go test -bench .
```

## Заключение

Тестирование в Go — мощный инструмент для обеспечения качества кода. Встроенный пакет `testing`, табличные тесты и бенчмаркинг помогают писать надежные и производительные программы.

