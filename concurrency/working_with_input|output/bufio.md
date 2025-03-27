# Буферизация ввода-вывода в Go

## Пакет `bufio`

### Основные типы

1. **`bufio.Reader`** - буферизированное чтение
2. **`bufio.Writer`** - буферизированная запись
3. **`bufio.Scanner`** - построчное чтение

### Преимущества буферизации
- Уменьшение количества системных вызовов
- Улучшение производительности для мелких операций
- Удобные методы для работы с текстом

## Буферизированное чтение

### Создание Reader
```go
file, _ := os.Open("data.txt")
reader := bufio.NewReader(file)
// С указанием размера буфера (по умолчанию 4096)
bigReader := bufio.NewReaderSize(file, 65536)
```

### Основные методы
```go
// Чтение до разделителя
line, _ := reader.ReadString('\n') 

// Чтение байта
b, _ := reader.ReadByte()

// Чтение руны
r, size, _ := reader.ReadRune()

// Возврат байта в буфер
reader.UnreadByte() 

// Чтение слайса байт
buf := make([]byte, 100)
n, _ := reader.Read(buf)
```

## Буферизированная запись

### Создание Writer
```go
file, _ := os.Create("output.txt")
writer := bufio.NewWriter(file)
// С указанием размера буфера
bigWriter := bufio.NewWriterSize(file, 65536)
```

### Основные методы
```go
// Запись строки
writer.WriteString("text")

// Запись байта
writer.WriteByte('a')

// Запись руны
writer.WriteRune('я')

// Запись слайса байт
writer.Write([]byte{1,2,3})

// Сброс буфера в нижележащий Writer
writer.Flush()
```

### Важно!
Всегда вызывайте `Flush()` после завершения записи:
```go
defer writer.Flush()
```

## Построчное чтение с Scanner

```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    // Обработка строки
}
if err := scanner.Err(); err != nil {
    // Обработка ошибки
}
```

### Кастомизация Scanner
```go
// Чтение по словам
scanner.Split(bufio.ScanWords)

// Чтение по байтам
scanner.Split(bufio.ScanBytes)

// Собственная функция разделения
scanner.Split(mySplitFunction)
```

## Буферизация в реальных примерах

### 1. Обработка большого файла
```go
func CountLines(filename string) (int, error) {
    file, _ := os.Open(filename)
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    count := 0
    for scanner.Scan() {
        count++
    }
    return count, scanner.Err()
}
```

### 2. Эффективная запись логов
```go
func WriteLogs(w io.Writer, logs []string) error {
    bw := bufio.NewWriter(w)
    defer bw.Flush()
    
    for _, log := range logs {
        if _, err := bw.WriteString(log + "\n"); err != nil {
            return err
        }
    }
    return nil
}
```

## Внутренняя работа буферизации

### Для Reader:
1. При создании выделяется буфер
2. При операциях чтения сначала заполняется буфер
3. Последующие чтения идут из буфера
4. При опустошении буфера - новое чтение из источника

### Для Writer:
1. Данные сначала пишутся в буфер
2. При заполнении буфера - сброс в нижележащий Writer
3. Явный `Flush()` принудительно сбрасывает буфер

## Оптимальные размеры буфера

- **По умолчанию**: 4096 байт (4KB)
- **Для HDD**: 32KB-64KB
- **Для SSD**: 4KB-16KB
- **Для сетевых операций**: 8KB-32KB

```go
// Пример настройки размера
reader := bufio.NewReaderSize(conn, 32768) // 32KB
```

## Производительность

### Сравнение с обычным IO
| Операция       | Без буферизации | С буферизацией |
|---------------|----------------|----------------|
| 1000 записей  | 1000 syscalls  | 1-10 syscalls  |
| Чтение 1MB    | 256 syscalls   | 1 syscall      |

## Особенности работы

1. **Не забывайте про Flush()** - иначе данные могут потеряться
2. **Ошибки накапливаются** - проверяйте ошибки после всех операций
3. **Буфер не потокобезопасен** - не используйте один буфер из нескольких горутин
4. **Scanner имеет ограничения** - максимальный размер строки по умолчанию 64KB

