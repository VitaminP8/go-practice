# Работа с файлами в Go (до произвольного доступа)

## Основные операции с файлами

### 1. Открытие файлов
Используются функции из пакета `os`:
```go
// Для чтения
file, err := os.Open("filename.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close() // Важно закрывать файл!

// Для записи (создаст или перезапишет файл)
file, err := os.Create("output.txt")
defer file.Close()

// Расширенное открытие с флагами
file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_APPEND, 0644)
```

### 2. Интерфейсы os.File
Тип `os.File` реализует ключевые интерфейсы:
- **Базовые**: `io.Reader`, `io.Writer`, `io.Closer`
- **Произвольный доступ**: `io.Seeker`, `io.ReaderAt`, `io.WriterAt`
- **Дополнительные**: `io.ReaderFrom`, `io.WriterTo`, `io.StringWriter`


### 3. Чтение файлов

#### Последовательное чтение
```go
buf := make([]byte, 1024) // Буфер для чтения
for {
    n, err := file.Read(buf)
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }
    process(buf[:n]) // Обработка прочитанных данных
}
```

#### Удобные методы чтения
```go
// Чтение всего файла
data, err := io.ReadAll(file)

// Чтение файла по имени (Go 1.16+)
data, err := os.ReadFile("filename.txt")

// Гарантированное заполнение буфера
err := io.ReadFull(file, buf)
```

### 4. Запись в файлы

#### Базовая запись
```go
data := []byte("Hello, World!")
n, err := file.Write(data)
if err != nil || n != len(data) {
    log.Fatal("Ошибка записи")
}
```

#### Удобные методы записи
```go
// Запись всего содержимого
err := os.WriteFile("output.txt", data, 0644)

// Копирование данных между потоками
written, err := io.Copy(file, srcReader)
```

### 5. Буферизованные операции

#### Чтение с буферизацией
```go
reader := bufio.NewReader(file)
line, err := reader.ReadString('\n') // Чтение строки
```

#### Запись с буферизацией
```go
writer := bufio.NewWriter(file)
writer.WriteString("Hello\n")
writer.WriteByte(65)
writer.Flush() // Важно не забывать!
```

### 6. Работа с файловой системой

#### Полезные функции
```go
// Проверка существования
if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
    // Файл не существует
}

// Получение информации о файле
info, err := os.Stat("file.txt")
size := info.Size()
modTime := info.ModTime()

// Работа с путями
import "path/filepath"
absPath := filepath.Abs("relative/path")
dir := filepath.Dir("/path/to/file")
```

## Важные особенности

1. Всегда закрывайте файлы с помощью `defer file.Close()`
2. Обрабатывайте ошибки при операциях ввода-вывода
3. Для текстовых файлов используйте `bufio` для построчного чтения
4. Для бинарных файлов работайте с байтовыми буферами
5. Используйте `os.FileInfo` для получения метаданных о файле
6. Различайте относительные и абсолютные пути (используйте `filepath`)

## Пример: копирование файла
```go
func CopyFile(src, dst string) error {
    in, err := os.Open(src)
    if err != nil {
        return err
    }
    defer in.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, in)
    return err
}
```