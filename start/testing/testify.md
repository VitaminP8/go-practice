# Testify: библиотека для удобного тестирования в Go

`Testify` — это популярная библиотека, которая расширяет возможности встроенного пакета `testing` в Go. Она предоставляет удобные инструменты для написания тестов, такие как утверждения (assertions) и наборы тестов (suites), делая тесты более читаемыми и лаконичными.

## Установка

Для установки библиотеки используйте команду:

```bash
go get github.com/stretchr/testify
```

## Основные компоненты

### 1. **Assertions (Утверждения)**
Пакет `assert` предоставляет множество функций для проверки условий в тестах. Если условие не выполняется, тест завершается с ошибкой.

Пример:

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
    result := 2 + 2
    assert.Equal(t, 4, result, "2 + 2 должно быть равно 4")
}
```

### 2. **Require (Обязательные утверждения)**
Пакет `require` аналогичен `assert`, но если условие не выполняется, тест немедленно завершается. Это полезно, если дальнейшее выполнение теста бессмысленно.

Пример:

```go
import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestDivision(t *testing.T) {
    result := 4 / 2
    require.Equal(t, 2, result, "4 / 2 должно быть равно 2")
}
```

### 3. **Suites (Наборы тестов)**
Пакет `suite` позволяет организовывать тесты в наборы (suites), что упрощает настройку и очистку перед/после выполнения тестов.

Пример:

```go
import (
    "testing"
    "github.com/stretchr/testify/suite"
)

type MySuite struct {
    suite.Suite
    Value int
}

func (s *MySuite) SetupTest() {
    s.Value = 42 // Настройка перед каждым тестом
}

func (s *MySuite) TestExample() {
    s.Equal(42, s.Value, "Значение должно быть 42")
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(MySuite))
}
```

## Преимущества Testify

- **Удобные утверждения**: Упрощает проверку условий с помощью функций вроде `Equal`, `True`, `Nil` и других.
- **Читаемость**: Тесты становятся более читаемыми и лаконичными.
- **Наборы тестов**: Упрощает организацию тестов и управление состоянием.

## Пример использования

```go
package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    assert.Equal(t, 4, Add(2, 2), "2 + 2 должно быть равно 4")
    assert.NotEqual(t, 5, Add(2, 2), "2 + 2 не должно быть равно 5")
}
```

## Заключение

`Testify` — это мощный инструмент для написания тестов в Go. Он делает тесты более читаемыми, удобными и поддерживаемыми. Если вы ещё не используете `testify`, обязательно попробуйте его в своих проектах!
