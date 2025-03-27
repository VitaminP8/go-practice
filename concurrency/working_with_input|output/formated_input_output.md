# Форматированный ввод-вывод в Go

## Основные функции вывода

### 1. Стандартный вывод
```go
fmt.Print("Hello")                 // Без перевода строки
fmt.Println("Hello")               // С переводом строки
fmt.Printf("Format: %s", "value")  // Форматированный вывод
```

### 2. Вывод в строку
```go
s := fmt.Sprintf("Formatted: %d", 42)  // Возвращает строку
```

### 3. Вывод в io.Writer
```go
file, _ := os.Create("output.txt")
fmt.Fprintf(file, "Write: %v", data)  // Запись в файл
```

## Форматные спецификаторы

### Базовые спецификаторы
| Спецификатор | Описание                  | Пример         |
|--------------|---------------------------|----------------|
| `%v`         | Значение в default формате | `fmt.Printf("%v", data)` |
| `%+v`        | С именами полей (структуры) | `fmt.Printf("%+v", user)` |
| `%#v`        | Go-синтаксис              | `fmt.Printf("%#v", arr)` |
| `%T`         | Тип переменной            | `fmt.Printf("%T", 42)` → "int" |
| `%%`         | Символ процента           | `fmt.Printf("100%%")` |

### Для разных типов данных
```go
// Целые числа
fmt.Printf("%b", 42)   // 101010 (binary)
fmt.Printf("%d", 42)   // 42 (decimal)
fmt.Printf("%o", 42)   // 52 (octal)
fmt.Printf("%x", 42)   // 2a (hex)

// Строки
fmt.Printf("%s", "text")  // text
fmt.Printf("%q", "text")  // "text" (в кавычках)

// Плавающая точка
fmt.Printf("%f", 3.14)    // 3.140000
fmt.Printf("%.2f", 3.14)  // 3.14
fmt.Printf("%e", 3.14)    // 3.140000e+00

// Указатели
fmt.Printf("%p", &x)      // 0xc000018080
```

## Форматированный ввод

### Основные функции
```go
var name string
var age int

// Чтение с stdin
fmt.Scan(&name)                    // Чтение одного значения
fmt.Scanln(&name, &age)            // Чтение строки
fmt.Scanf("%s %d", &name, &age)    // Форматированное чтение

// Чтение из строки
input := "John 25"
fmt.Sscanf(input, "%s %d", &name, &age)

// Чтение из файла
file, _ := os.Open("data.txt")
fmt.Fscanf(file, "%s %d", &name, &age)
```

### Особенности ввода
1. Разделитель - пробел/перевод строки
2. Для строк - чтение до пробела
3. Ошибки нужно обрабатывать:
```go
n, err := fmt.Scanf("%d", &number)
if err != nil {
    log.Fatal("Ошибка ввода:", err)
}
```

## Пользовательское форматирование

### Интерфейс Stringer
```go
type User struct {
    Name string
    Age  int
}

func (u User) String() string {
    return fmt.Sprintf("%s (%d лет)", u.Name, u.Age)
}

u := User{"John", 25}
fmt.Println(u)  // John (25 лет)
```

### Интерфейс GoStringer (для %#v)
```go
func (u User) GoString() string {
    return fmt.Sprintf("User{Name:%q, Age:%d}", u.Name, u.Age)
}
fmt.Printf("%#v", u)  // User{Name:"John", Age:25}
```

## Ширина и точность

### Управление выводом
```go
fmt.Printf("|%10s|", "text")    // |      text| (ширина 10)
fmt.Printf("|%-10s|", "text")   // |text      | (выравнивание влево)
fmt.Printf("|%10.2f|", 3.14159) // |      3.14| (ширина 10, точность 2)
```

### Динамические параметры
```go
width := 8
precision := 3
fmt.Printf("%*.*f", width, precision, 3.14159)  // "   3.142"
```

## Полезные паттерны

### 1. Построение сложных строк
```go
msg := fmt.Sprintf(
    "[%s] %-15s: %0.2fMB",
    time.Now().Format("2006-01-02"),
    "Memory used",
    123.4567,
)
// [2023-10-25] Memory used   : 123.46MB
```

### 2. Вывод таблиц
```go
fmt.Printf("| %-20s | %10s |\n", "Name", "Age")
fmt.Printf("| %-20s | %10d |\n", "John Smith", 32)
fmt.Printf("| %-20s | %10d |\n", "Alice", 27)
```

### 3. Чтение конфигурации
```go
var config struct {
    Host string
    Port int
}
_, err := fmt.Sscanf("localhost:8080", "%s:%d", &config.Host, &config.Port)
```

## Производительность

Для высоконагруженных сценариев:
- `fmt.Sprintf` медленнее конкатенации
- Для логирования лучше использовать `fmt.Fprint` с `bytes.Buffer`
- В критичных местах - использовать `strconv` для чисел

```go
var buf bytes.Buffer
fmt.Fprintf(&buf, "Value: %d", 42)  // Быстрее чем Sprintf
```