# Defer в Go

## 1. Что такое `defer`?
В Go ключевое слово `defer` используется для отложенного выполнения функции до выхода из текущей области видимости. Оно полезно для освобождения ресурсов, закрытия файлов и других операций, которые должны выполняться в конце выполнения функции.

Пример:
```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred")
    fmt.Println("End")
}
```
Вывод:
```
Start
End
Deferred
```
Функция, объявленная с `defer`, выполнится после выхода из `main()`, но перед фактическим завершением программы.

## 2. Стек вызовов `defer`
Если несколько `defer` объявлены в одной функции, они выполняются в **обратном порядке** (LIFO — Last In, First Out).

Пример:
```go
package main

import "fmt"

func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    fmt.Println("Main function")
}
```
Вывод:
```
Main function
Third
Second
First
```

## 3. `defer` и работа с файлами
Обычно `defer` используется для закрытия файлов после их открытия, что помогает избежать утечек ресурсов.

Пример:
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    fmt.Println("File opened successfully")
}
```

## 4. `defer` и изменение переменных
Важно помнить, что `defer` захватывает текущее значение аргументов функции в момент объявления.

Пример:
```go
package main

import "fmt"

func main() {
    x := 10
    defer fmt.Println("Deferred x:", x)
    x = 20
    fmt.Println("Updated x:", x)
}
```
Вывод:
```
Updated x: 20
Deferred x: 10
```
Переменная `x` сохраняет своё значение (`10`) в момент вызова `defer`, а не в момент выполнения.

## Итог
- `defer` откладывает выполнение функции до выхода из текущей области видимости.
- Несколько `defer` выполняются в **обратном порядке**.
- Используется для управления ресурсами (например, закрытие файлов).
- Захватывает значения аргументов в момент объявления.