# Методы в Go

Методы в Go — это функции, привязанные к конкретному типу. Они позволяют добавлять поведение к структурам или другим пользовательским типам. Методы похожи на функции, но имеют получателя (receiver), который указывает, к какому типу они принадлежат.

## 1. Определение методов

Метод определяется с использованием получателя (receiver) перед именем функции. Получатель может быть как значением, так и указателем.

Пример метода для структуры:
```go
package main

import "fmt"

// Определяем структуру
type Rectangle struct {
    width, height float64
}

// Метод для структуры Rectangle
func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    fmt.Println("Area:", rect.Area()) // Area: 50
}
```

Здесь метод `Area()` привязан к типу `Rectangle` и может быть вызван на экземпляре этой структуры.

## 2. Методы с указателями

Если метод изменяет состояние структуры, получатель должен быть указателем. Это позволяет изменять оригинальные данные, а не их копию.

Пример:
```go
package main

import "fmt"

type Rectangle struct {
    width, height float64
}

// Метод с получателем-указателем
func (r *Rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}

func main() {
    rect := Rectangle{width: 10, height: 5}
    rect.Scale(2)
    fmt.Println("Scaled Rectangle:", rect) // Scaled Rectangle: {20 10}
}
```

В этом примере метод `Scale()` изменяет значения полей структуры `Rectangle`.

## 3. Методы для пользовательских типов

Методы можно определять не только для структур, но и для любых пользовательских типов.

Пример:
```go
package main

import "fmt"

// Определяем пользовательский тип
type MyInt int

// Метод для пользовательского типа
func (m MyInt) IsEven() bool {
    return m%2 == 0
}

func main() {
    num := MyInt(4)
    fmt.Println("Is even?", num.IsEven()) // Is even? true
}
```

Здесь метод `IsEven()` привязан к типу `MyInt`.

## 4. Выбор между значением и указателем в получателе

- **Получатель-значение**: Используется, если метод не изменяет состояние структуры. Работает с копией данных.
- **Получатель-указатель**: Используется, если метод изменяет состояние структуры. Работает с оригинальными данными.

Пример:
```go
package main

import "fmt"

type Circle struct {
    radius float64
}

// Метод с получателем-значением
func (c Circle) Area() float64 {
    return 3.14 * c.radius * c.radius
}

// Метод с получателем-указателем
func (c *Circle) Scale(factor float64) {
    c.radius *= factor
}

func main() {
    circle := Circle{radius: 5}
    fmt.Println("Area:", circle.Area()) // Area: 78.5

    circle.Scale(2)
    fmt.Println("Scaled Area:", circle.Area()) // Scaled Area: 314
}
```

## 5. Встраивание методов (Embedding)

В Go можно встраивать структуры, что позволяет использовать методы встроенной структуры в родительской.

Пример:
```go
package main

import "fmt"

type Person struct {
    Name string
}

func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

type Employee struct {
    Person
    JobTitle string
}

func main() {
    emp := Employee{
        Person:   Person{Name: "John"},
        JobTitle: "Developer",
    }
    emp.Greet() // Hello, my name is John
}
```

Здесь структура `Employee` встраивает структуру `Person`, и метод `Greet()` становится доступным для `Employee`.

### Итог
- **Методы** позволяют добавлять поведение к типам.
- Получатель может быть **значением** или **указателем**.
- Методы можно определять для **структур** и **пользовательских типов**.
- **Встраивание** позволяет использовать методы вложенных структур.