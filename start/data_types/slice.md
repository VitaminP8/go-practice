# Слайсы в Go

## 1. Что такое слайсы?
Слайсы (slices) в Go — это удобная обёртка над массивами, позволяющая работать с динамическими последовательностями данных. В отличие от массивов, размер слайса может изменяться во время выполнения программы.

## 2. Внутреннее устройство слайсов
Слайс в Go представляет собой структуру, которая содержит три элемента:
- **Указатель (Pointer)** — указывает на первый элемент массива в памяти.
- **Длина (Length)** — количество элементов в слайсе.
- **Ёмкость (Capacity)** — максимальное количество элементов, которое можно вместить без перераспределения памяти.

```go
struct SliceHeader {
    Data uintptr   // Указатель на массив
    Len  int       // Текущая длина слайса
    Cap  int       // Ёмкость (capacity)
}
```

## 3. Хранение в памяти
Слайс сам по себе **не содержит** данные, а лишь указывает на массив в памяти.

Пример:
```go
arr := [5]int{1, 2, 3, 4, 5} // Массив в памяти
s := arr[1:4]                // Создаём слайс (указатель на arr[1])
```

Память будет организована так:
```
 Индексы:     0   1   2   3   4  
 Массив:     [1,  2,  3,  4,  5]
 Слайс s:        [2,  3,  4]
```

Структура `s`:
```
Pointer -> arr[1] (указатель на 2)
Length  = 3  (элементы: 2, 3, 4)
Capacity = 4 (максимально возможная длина без перераспределения памяти)
```

### 📌 Картинка хранения в памяти:
```
  +---+---+---+---+---+
  | 1 | 2 | 3 | 4 | 5 |
  +---+---+---+---+---+
      ^   ^   ^
      |   |   |
     s[0] s[1] s[2] (указатель на arr[1], длина = 3)
```

## 4. Изменение слайсов
Так как слайс указывает на массив, изменение элементов внутри слайса **изменяет оригинальный массив**:
```go
s[1] = 99
fmt.Println(arr) // [1 2 99 4 5]
```

Однако если длина слайса превысит `Capacity`, Go создаст новый массив и скопирует в него данные.

```go
s = append(s, 6, 7, 8) // Выход за пределы capacity
fmt.Println(s)         // Новый массив: [2 99 4 6 7 8]
```

Теперь `s` указывает на новый массив, а `arr` остаётся неизменным.

## 5. Итог
- **Слайс — это структура, содержащая указатель, длину и ёмкость.**
- **Слайс не хранит данные, а ссылается на массив.**
- **При добавлении элементов `append()` может перераспределять память.**

Слайсы делают работу с массивами удобной, но важно понимать их устройство, чтобы избежать неожиданных эффектов при изменении данных.

## 6. Основные операции со слайсами

### Создание слайса
```go
s := []int{1, 2, 3} // Создание слайса с элементами 1, 2, 3
```

### Создание слайса с `make`
```go
s := make([]int, 5, 10) // Слайс длиной 5, ёмкостью 10
```

### Добавление элементов (`append`)
```go
s = append(s, 4, 5, 6)
```

### Копирование слайсов (`copy`)
```go
dst := make([]int, len(s))
copy(dst, s) // Копирует элементы s в dst
```

### Обрезка слайса
```go
s = s[1:3] // Оставит элементы с индексами 1 и 2
```

### Итерация по слайсу
```go
for i, v := range s {
    fmt.Println(i, v)
}
```

