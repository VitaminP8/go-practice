# Работа с ошибками в Go

## 1. Основные принципы работы с ошибками
В Go ошибки представляются через тип `error`, который является встроенным интерфейсом:

```go
package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("что-то пошло не так")
    fmt.Println(err)
}
```

Этот способ создаёт новую ошибку, но чаще используются другие подходы.

## 2. Использование `fmt.Errorf`
Для форматирования ошибок можно использовать `fmt.Errorf`:

```go
package main

import (
    "fmt"
)

func main() {
    name := "Go"
    err := fmt.Errorf("ошибка: неверное значение %s", name)
    fmt.Println(err)
}
```

## 3. Оборачивание ошибок (`errors.Is` и `errors.As`)
В Go 1.13 добавили поддержку оборачивания ошибок с помощью `%w`:

```go
package main

import (
    "errors"
    "fmt"
)

func someFunc() error {
    return errors.New("исходная ошибка")
}

func main() {
    err := fmt.Errorf("высокоуровневая ошибка: %w", someFunc())
    fmt.Println(err)
}
```

Для проверки вложенных ошибок используются `errors.Is` и `errors.As`:

```go
package main

import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("не найдено")

func find() error {
    return fmt.Errorf("ошибка выполнения запроса: %w", ErrNotFound)
}

func main() {
    err := find()
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Элемент не найден")
    }
}
```

## 4. Создание собственных типов ошибок
Можно создать собственный тип ошибки, реализуя интерфейс `error`:

```go
package main

import (
    "fmt"
)

type MyError struct {
    Code int
    Msg  string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("код %d: %s", e.Code, e.Msg)
}

func main() {
    err := &MyError{Code: 404, Msg: "страница не найдена"}
    fmt.Println(err)
}
```

### Итог
- Используйте `fmt.Errorf("%w")` для оборачивания ошибок.
- Применяйте `errors.Is()` и `errors.As()` для проверки типов ошибок.
- Можно создавать собственные типы ошибок для детальной информации.
