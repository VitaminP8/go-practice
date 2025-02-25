# Продвинутое тестирование в Go: Blackbox, Suite, Dependency Injection и Моки

---

## 1. **Blackbox-тестирование**

Blackbox-тестирование — это подход, при котором тестируется только внешнее поведение модуля или пакета без знания его внутренней реализации. В Go это достигается путём создания тестов в отдельном пакете с суффиксом `_test`.

### Пример:

```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

// math_test.go
package math_test

import (
    "testing"
    "github.com/yourproject/math"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    result := math.Add(2, 2)
    assert.Equal(t, 4, result, "2 + 2 должно быть равно 4")
}
```

### Преимущества:
- Тесты не зависят от внутренней реализации.
- Упрощается рефакторинг кода, так как тесты проверяют только внешнее поведение.

---

## 2. **Test Suites (Наборы тестов)**

Test Suites позволяют группировать тесты и разделять общую логику инициализации и завершения. Это особенно полезно для интеграционных тестов или тестов, требующих настройки состояния.

### Пример с использованием `testify/suite`:

```go
package main

import (
    "testing"
    "github.com/stretchr/testify/suite"
)

type MySuite struct {
    suite.Suite
    Value int
}

func (s *MySuite) SetupTest() {
    // Выполняется перед каждым тестом
    s.Value = 42
}

func (s *MySuite) TearDownTest() {
    // Выполняется после каждого теста
    s.Value = 0
}

func (s *MySuite) TestExample() {
    s.Equal(42, s.Value, "Значение должно быть 42")
}

func TestSuite(t *testing.T) {
    suite.Run(t, new(MySuite))
}
```

### Преимущества:
- Группировка тестов.
- Общие настройки и завершение для всех тестов в наборе.

---

## 3. **Dependency Injection (Внедрение зависимостей)**

Dependency Injection (DI) — это подход, при котором зависимости передаются в функцию или структуру извне, что упрощает тестирование и делает код более гибким.

### Пример:

```go
type Database interface {
    GetData() string
}

type Service struct {
    db Database
}

func NewService(db Database) *Service {
    return &Service{db: db}
}

func (s *Service) DoWork() string {
    return s.db.GetData()
}

// Тест с использованием мока
func TestService(t *testing.T) {
    mockDB := new(MockDatabase)
    mockDB.On("GetData").Return("mocked data")

    service := NewService(mockDB)
    result := service.DoWork()

    assert.Equal(t, "mocked data", result)
    mockDB.AssertExpectations(t)
}
```

### Преимущества:
- Упрощает тестирование, так как зависимости можно легко подменить.
- Делает код более модульным и поддерживаемым.

---

## 4. **Моки (Mocking)**

Моки используются для имитации поведения зависимостей в тестах. Это особенно полезно, когда реальные зависимости (например, база данных или внешний API) недоступны или их использование нежелательно.

### Пример с использованием `testify/mock`:

```go
package main

import (
    "testing"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/assert"
)

// Интерфейс для зависимости
type Database interface {
    GetData() string
}

// Мок-реализация
type MockDatabase struct {
    mock.Mock
}

func (m *MockDatabase) GetData() string {
    args := m.Called()
    return args.String(0)
}

func TestServiceWithMock(t *testing.T) {
    // Создаём мок
    mockDB := new(MockDatabase)
    mockDB.On("GetData").Return("mocked data")

    // Используем мок в тесте
    service := NewService(mockDB)
    result := service.DoWork()

    // Проверяем результат
    assert.Equal(t, "mocked data", result)
    mockDB.AssertExpectations(t) // Убедимся, что метод был вызван
}
```

### Преимущества:
- Полный контроль над поведением зависимостей.
- Возможность тестировать сложные сценарии без реальных зависимостей.

---

## 5. **Практические советы**

### a. **Используйте `go test -cover`**
Проверяйте покрытие кода тестами с помощью флага `-cover`:
```bash
go test -cover ./...
```

### b. **Табличные тесты**
Используйте табличные тесты для тестирования множества сценариев:
```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, expected int
    }{
        {1, 1, 2},
        {2, 2, 4},
        {-1, 1, 0},
    }

    for _, tt := range tests {
        result := Add(tt.a, tt.b)
        assert.Equal(t, tt.expected, result)
    }
}
```

---

## Заключение

Продвинутое тестирование в Go включает в себя множество подходов и инструментов:
- **Blackbox-тестирование** для проверки внешнего поведения.
- **Test Suites** для группировки тестов и управления состоянием.
- **Dependency Injection** для упрощения тестирования и повышения гибкости.
- **Моки** для имитации зависимостей.

Используйте эти техники, чтобы писать более надёжные, поддерживаемые и читаемые тесты. Если у вас есть вопросы или нужно больше примеров, дайте знать! 😊