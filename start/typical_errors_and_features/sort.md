# Сортировка слайсов в Go

## 1. Использование `sort.Slice`
В Go стандартная библиотека `sort` предоставляет функцию `sort.Slice`, которая позволяет сортировать слайсы по заданному критерию.

Пример сортировки чисел:
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    nums := []int{5, 3, 8, 1, 9}
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j] // сортировка по возрастанию
    })
    fmt.Println(nums) // [1 3 5 8 9]
}
```

Пример сортировки строк:
```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    words := []string{"banana", "apple", "cherry"}
    sort.Slice(words, func(i, j int) bool {
        return words[i] < words[j] // сортировка в алфавитном порядке
    })
    fmt.Println(words) // [apple banana cherry]
}
```

## 2. Сортировка слайсов структур
Если у вас есть слайс структур, можно сортировать его по определённому полю.

Пример:
```go
package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }

    // Сортировка по возрасту
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age < people[j].Age
    })

    fmt.Println(people)
}
```
Вывод:
```
[{Bob 25} {Alice 30} {Charlie 35}]
```

## 3. Использование `sort.Ints`, `sort.Strings`, `sort.Float64s`
Если требуется сортировка чисел, строк или чисел с плавающей точкой, можно использовать готовые функции:

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    ints := []int{5, 3, 8, 1}
    sort.Ints(ints) // Сортировка по возрастанию
    fmt.Println(ints) // [1 3 5 8]

    strs := []string{"banana", "apple", "cherry"}
    sort.Strings(strs) // Алфавитный порядок
    fmt.Println(strs) // [apple banana cherry]
}
```

## Итог
- `sort.Slice` позволяет задавать кастомный порядок сортировки.
- `sort.Ints`, `sort.Strings`, `sort.Float64s` удобны для сортировки стандартных типов.
- Можно сортировать структуры, передавая кастомную функцию сравнения.

Эти методы позволяют гибко работать с сортировкой в Go. 🚀