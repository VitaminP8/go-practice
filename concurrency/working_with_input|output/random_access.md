# Произвольный доступ к файлам в Go

## Основные интерфейсы для произвольного доступа

### 1. `io.Seeker`
```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```
**Назначение**: Перемещение по файлу без чтения данных.

**Параметры**:
- `offset` - смещение в байтах
- `whence` - точка отсчета:
    - `io.SeekStart` (0) - от начала файла
    - `io.SeekCurrent` (1) - от текущей позиции
    - `io.SeekEnd` (2) - от конца файла

**Пример**:
```go
file, _ := os.Open("data.bin")
file.Seek(100, io.SeekStart) // Переместиться на 100 байт от начала
file.Seek(-50, io.SeekEnd)   // Переместиться на 50 байт от конца
```

### 2. `io.ReaderAt`
```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```
**Особенности**:
- Читает данные с указанной позиции `off`
- Не изменяет текущую позицию в файле
- Всегда читает ровно `len(p)` байт или возвращает ошибку

**Пример**:
```go
buf := make([]byte, 50)
file.ReadAt(buf, 200) // Читает 50 байт начиная с позиции 200
```

### 3. `io.WriterAt`
```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```
**Особенности**:
- Записывает данные в указанную позицию `off`
- Не изменяет текущую позицию в файле
- Всегда записывает ровно `len(p)` байт или возвращает ошибку

**Пример**:
```go
data := []byte("new data")
file.WriteAt(data, 300) // Записывает в позицию 300
```

## Практическое применение

### 1. Чтение заголовка файла
```go
func readHeader(file *os.File) ([]byte, error) {
    header := make([]byte, 128)
    _, err := file.ReadAt(header, 0) // Читаем первые 128 байт
    return header, err
}
```

### 2. Модификация середины файла
```go
func patchFile(file *os.File, patch []byte, offset int64) error {
    _, err := file.WriteAt(patch, offset)
    return err
}
```

### 3. Прямой доступ к структурам данных
```go
type Record struct {
    ID    int32
    Value float64
}

func readRecord(file *os.File, index int) (*Record, error) {
    rec := new(Record)
    offset := int64(index) * int64(unsafe.Sizeof(*rec))
    
    data := make([]byte, unsafe.Sizeof(*rec))
    _, err := file.ReadAt(data, offset)
    if err != nil {
        return nil, err
    }
    
    binary.Read(bytes.NewReader(data), binary.LittleEndian, rec)
    return rec, nil
}
```

## Особенности реализации

1. **Производительность**:
    - Прямой доступ обычно быстрее последовательного чтения
    - Минимизирует количество системных вызовов

2. **Ограничения**:
    - Не все устройства поддерживают произвольный доступ
    - Сетевые соединения и pipes обычно работают только последовательно

3. **Безопасность**:
    - При работе с `ReadAt`/`WriteAt` буфер должен быть инициализирован нужного размера
    - Важно проверять возвращаемые ошибки и количество прочитанных/записанных байт

## Сравнение с последовательным доступом

| Характеристика       | Произвольный доступ       | Последовательный доступ |
|----------------------|--------------------------|-------------------------|
| Позиционирование     | Любая позиция            | Только последовательно  |
| Изменение позиции    | `Seek()`                 | Автоматическое          |
| Буферизация          | Обычно не буферизируется | Часто буферизируется    |
| Применение          | Бинарные форматы, БД    | Текстовые файлы, логи   |

## Оптимизация работы

1. **Использование `io.SectionReader`**:
   ```go
   section := io.NewSectionReader(file, 100, 50)
   // Работает с частью файла (с 100 по 150 байт)
   ```

2. **Комбинирование с буферизацией**:
   ```go
   buffered := bufio.NewReader(file)
   file.Seek(100, io.SeekStart)
   // Дальше работаем с буферизированным чтением
   ```

3. **Работа с памятью**:
   ```go
   data, _ := os.ReadFile("large.bin")
   reader := bytes.NewReader(data) // Реализует все интерфейсы произвольного доступа
   ```