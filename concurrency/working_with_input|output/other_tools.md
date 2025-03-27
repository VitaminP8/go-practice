# Другие инструменты ввода-вывода в Go

## Оптимизированные интерфейсы копирования: WriterTo и ReaderFrom

### Интерфейс WriterTo

#### Определение
```go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

#### Назначение
Позволяет объекту оптимизировать процесс записи своих данных в `io.Writer`

#### Как работает
1. Объект сам управляет процессом записи
2. Может использовать внутренние буферы
3. Может выполнять преобразования данных на лету
4. Избегает лишнего копирования через промежуточные буферы

#### Пример реализации
```go
type CustomBuffer struct {
    data [][]byte
}

func (b *CustomBuffer) WriteTo(w io.Writer) (int64, error) {
    var total int64
    for _, chunk := range b.data {
        n, err := w.Write(chunk)
        total += int64(n)
        if err != nil {
            return total, err
        }
    }
    return total, nil
}
```

#### Когда использовать
- Когда ваш тип содержит данные в специальном формате
- Когда можно записать данные эффективнее, чем через стандартный `io.Copy`
- Для больших объектов, где важно минимизировать копирование

### Интерфейс ReaderFrom

#### Определение
```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

#### Назначение
Позволяет объекту оптимизировать процесс чтения данных из `io.Reader`

#### Как работает
1. Объект сам управляет процессом чтения
2. Может использовать внутренние буферы
3. Может выполнять предварительную обработку данных
4. Избегает промежуточных копий данных

#### Пример реализации
```go
type CustomBuffer struct {
    data bytes.Buffer
}

func (b *CustomBuffer) ReadFrom(r io.Reader) (int64, error) {
    // Используем Grow для оптимизации выделения памяти
    b.data.Grow(512)
    return b.data.ReadFrom(r)
}
```

#### Когда использовать
- Когда ваш тип может эффективно накапливать данные
- Когда нужно специальное поведение при чтении
- Для обработки потоковых данных

### Преимущества использования

1. **Производительность**:
    - Уменьшение количества копий данных
    - Возможность использовать специальные оптимизации

2. **Гибкость**:
    - Можно реализовать любую логику преобразования данных
    - Поддержка потоковой обработки

3. **Совместимость**:
    - Прозрачная работа с существующими API
    - Автоматическое использование в `io.Copy`

### Важные нюансы

1. Всегда возвращайте корректное количество записанных/прочитанных байт
2. Обрабатывайте ошибки должным образом
3. Для больших данных используйте чанкированную обработку
4. Документируйте особенности вашей реализации

Эти интерфейсы особенно полезны при работе с:
- Сетевыми протоколами
- Шифрованием/сжатием
- Специализированными форматами данных
- Большими объемами информации

## 2. Комбинированные интерфейсы

### Предопределенные комбинации
```go
type ReadCloser interface { Reader; Closer }
type ReadWriteCloser interface { Reader; Writer; Closer }
type ReadSeeker interface { Reader; Seeker }
// И другие комбинации...
```

**Использование**:
```go
func Process(rwc io.ReadWriteCloser) {
    defer rwc.Close()
    // Работа с данными
}
```

## 3. Объединение потоков

### io.MultiReader
```go
func MultiReader(readers ...Reader) Reader
```
Пример (аналог Unix-команды cat):
```go
r1 := strings.NewReader("first ")
r2 := strings.NewReader("second")
combined := io.MultiReader(r1, r2)
io.Copy(os.Stdout, combined) // "first second"
```

### io.MultiWriter
```go
func MultiWriter(writers ...Writer) Writer
```
Пример (аналог Unix-команды tee):
```go
file, _ := os.Create("log.txt")
multi := io.MultiWriter(os.Stdout, file)
fmt.Fprintf(multi, "Log message\n") // Вывод и в файл, и на консоль
```

## 4. Ограничивающие ридеры

### io.LimitReader
```go
func LimitReader(r Reader, n int64) Reader
```
Пример:
```go
r := strings.NewReader("123456789")
limited := io.LimitReader(r, 5)
io.Copy(os.Stdout, limited) // "12345"
```

