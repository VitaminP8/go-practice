# Сортировка слайсов в Go: `slices.Sort` (Go 1.21+)

В версии Go 1.21 добавлен новый удобный метод для сортировки слайсов — `slices.Sort` из пакета `slices`. Этот метод упрощает сортировку по умолчанию (по возрастанию) без необходимости писать кастомные функции сравнения.

## 1. `slices.Sort` для чисел и строк
Теперь сортировка слайсов чисел или строк стала еще проще:

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{5, 3, 8, 1, 9}
    slices.Sort(nums)
    fmt.Println(nums) // [1 3 5 8 9]

    words := []string{"banana", "apple", "cherry"}
    slices.Sort(words)
    fmt.Println(words) // [apple banana cherry]
}
```
`Sort` автоматически использует стандартный порядок сортировки:
- Числа сортируются по возрастанию.
- Строки сортируются в алфавитном порядке.

## 2. `slices.SortFunc` для кастомной сортировки
Если нужно задать пользовательский порядок сортировки, можно использовать `slices.SortFunc`, аналогичный `sort.Slice`:

```go
package main

import (
    "fmt"
    "slices"
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
    slices.SortFunc(people, func(a, b Person) int {
        return a.Age - b.Age // Возвращает отрицательное, если a < b
    })

    fmt.Println(people)
}
```
Вывод:
```
[{Bob 25} {Alice 30} {Charlie 35}]
```
Функция `SortFunc` принимает компаратор, который возвращает:
- Отрицательное число, если `a < b`
- 0, если `a == b`
- Положительное число, если `a > b`

## 3. `slices.IsSorted` — проверка отсортированности
Пакет `slices` также включает `slices.IsSorted`, чтобы проверить, отсортирован ли слайс:

```go
package main

import (
    "fmt"
    "slices"
)

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println(slices.IsSorted(nums)) // true

    nums = []int{5, 3, 1}
    fmt.Println(slices.IsSorted(nums)) // false
}
```

## Итог
- `slices.Sort` (Go 1.21+) — простой способ сортировки слайсов по возрастанию.
- `slices.SortFunc` — для кастомного порядка сортировки.
- `slices.IsSorted` — проверка, отсортирован ли слайс.
- `slices.Sort` работает быстрее и удобнее, чем `sort.Slice`.

### Когда использовать `slices.Sort` вместо `sort.Slice`?
- **Используйте `slices.Sort`**, если нужно просто отсортировать числа или строки.
- **Используйте `slices.SortFunc`**, если требуется кастомная логика сортировки.
- **Используйте `sort.Slice`**, если работаете со старыми версиями Go (<1.21).

Теперь сортировать данные в Go стало еще проще! 🚀