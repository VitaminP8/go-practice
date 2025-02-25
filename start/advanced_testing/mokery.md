# Mokery: библиотека для создания моков в Go

Моки (mocks) — это объекты, которые имитируют поведение реальных зависимостей в тестах. Они позволяют изолировать тестируемый код от внешних систем, таких как базы данных, API или другие сервисы. В Go одной из популярных библиотек для создания моков является **Mokery**.

---

## Что такое Mokery?

**Mokery** — это инструмент для генерации моков на основе интерфейсов в Go. Он автоматически создаёт реализации интерфейсов, которые можно использовать в тестах для имитации поведения зависимостей.

---

## Установка

Для установки Mokery используйте команду:

```bash
go install github.com/vektra/mockery/v2@latest
```

После установки вы можете использовать `mockery` в командной строке для генерации моков.

---

## Основные возможности

### 1. **Генерация моков**
Mokery автоматически генерирует моки на основе интерфейсов. Например, если у вас есть интерфейс:

```go
package mypackage

type MyInterface interface {
    DoSomething() error
}
```

Вы можете сгенерировать мок для этого интерфейса с помощью команды:

```bash
mockery --name=MyInterface
```

Эта команда создаст файл `mocks/MyInterface.go` с реализацией мока.

---

### 2. **Использование моков в тестах**

Сгенерированный мок можно использовать в тестах для имитации поведения интерфейса. Например:

```go
package mypackage_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/yourproject/mypackage"
    "github.com/yourproject/mypackage/mocks"
)

func TestSomething(t *testing.T) {
    // Создаём мок
    mockMyInterface := new(mocks.MyInterface)

    // Настраиваем ожидания
    mockMyInterface.On("DoSomething").Return(nil)

    // Используем мок в тесте
    err := mockMyInterface.DoSomething()

    // Проверяем результат
    assert.NoError(t, err)
    mockMyInterface.AssertExpectations(t) // Убедимся, что метод был вызван
}
```

---

### 3. **Настройка поведения моков**

Mokery позволяет настраивать поведение моков с помощью методов `On`, `Return`, `Once`, `Times` и других. Например:

```go
mockMyInterface.On("DoSomething").Return(errors.New("error")).Once()
```

Это означает, что метод `DoSomething` вернёт ошибку только один раз.

---

### 4. **Проверка вызовов**

Mokery автоматически отслеживает вызовы методов моков. Вы можете проверить, что методы были вызваны с правильными аргументами и нужное количество раз:

```go
mockMyInterface.AssertExpectations(t)
```

---

## Преимущества Mokery

- **Автоматическая генерация моков**: Не нужно писать моки вручную.
- **Интеграция с `testify`**: Mokery генерирует моки, совместимые с библиотекой `testify/mock`.
- **Гибкость**: Возможность настраивать поведение моков и проверять вызовы методов.

---

## Пример использования

### Шаг 1: Создайте интерфейс

```go
package mypackage

type MyInterface interface {
    DoSomething() error
}
```

### Шаг 2: Сгенерируйте мок

```bash
mockery --name=MyInterface --output=mocks --dir=.
```

### Шаг 3: Используйте мок в тесте

```go
package mypackage_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/yourproject/mypackage"
    "github.com/yourproject/mypackage/mocks"
)

func TestSomething(t *testing.T) {
    mockMyInterface := new(mocks.MyInterface)
    mockMyInterface.On("DoSomething").Return(nil)

    err := mockMyInterface.DoSomething()

    assert.NoError(t, err)
    mockMyInterface.AssertExpectations(t)
}
```

---

## Заключение

Mokery — это мощный инструмент для создания моков в Go. Он упрощает написание тестов, позволяя изолировать код от внешних зависимостей. Используйте Mokery, чтобы сделать ваши тесты более надёжными и поддерживаемыми.
