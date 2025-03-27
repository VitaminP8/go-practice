# Работа с регулярными выражениями в Go

## Пакет `regexp`

### Основные типы
- `regexp.Regexp` - скомпилированное регулярное выражение
- `regexp.Regexp.MatchString` - проверка соответствия строки

### Инициализация
```go
// Компиляция с проверкой ошибок
re, err := regexp.Compile(`pattern`) 

// Компиляция без проверки (паника при ошибке)
re := regexp.MustCompile(`pattern`)
```

## Основные операции

### Проверка соответствия
```go
matched := re.MatchString("sample text") // bool
```

### Поиск совпадений
```go
// Первое совпадение
match := re.FindString("text") // string

// Все совпадения
matches := re.FindAllString("text", -1) // []string

// Индексы совпадений
idx := re.FindStringIndex("text") // [start, end]
```

### Замена
```go
// Простая замена
result := re.ReplaceAllString("text", "replacement")

// Замена с callback
result := re.ReplaceAllStringFunc("text", func(s string) string {
    return strings.ToUpper(s)
})
```

## Синтаксис регулярных выражений

### Основные конструкции
| Паттерн  | Описание                     |
|----------|------------------------------|
| `.`      | Любой символ                 |
| `\d`     | Цифра [0-9]                  |
| `\w`     | Слово [a-zA-Z0-9_]           |
| `\s`     | Пробельный символ            |
| `[abc]`  | Любой из символов a, b, c    |
| `[^abc]` | Любой символ, кроме a, b, c  |
| `a|b`    | a или b                      |
| `^`      | Начало строки                |
| `$`      | Конец строки                 |

### Квантификаторы
| Паттерн | Описание                |
|---------|-------------------------|
| `*`     | 0 или более повторений  |
| `+`     | 1 или более повторений  |
| `?`     | 0 или 1 повторение      |
| `{n}`   | Ровно n повторений      |
| `{n,}`  | n или более повторений  |
| `{n,m}` | От n до m повторений    |

## Группы и подвыражения

### Выделение групп
```go
re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
matches := re.FindStringSubmatch("user@example.com")
// matches: ["user@example.com", "user", "example", "com"]
```

### Именованные группы
```go
re := regexp.MustCompile(`(?P<name>\w+)@(?P<domain>\w+\.\w+)`)
matches := re.FindStringSubmatch("user@example.com")
name := matches[re.SubexpIndex("name")] // "user"
```

## Примеры использования

### Валидация email
```go
func ValidateEmail(email string) bool {
    re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    return re.MatchString(email)
}
```

### Поиск всех ссылок в тексте
```go
func FindLinks(text string) []string {
    re := regexp.MustCompile(`https?://[^\s]+`)
    return re.FindAllString(text, -1)
}
```

### Разбор логов
```go
logEntry := "[ERROR] 2023-10-25: File not found"
re := regexp.MustCompile(`^\[(\w+)\]\s(\d{4}-\d{2}-\d{2}):\s(.+)$`)
parts := re.FindStringSubmatch(logEntry)
// parts: ["[ERROR] 2023-10-25: File not found", "ERROR", "2023-10-25", "File not found"]
```

## Производительность и рекомендации

1. **Компилируйте один раз**:
   ```go
   // Глобальная переменная для часто используемых regexp
   var dateRe = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
   ```

2. **Используйте `regexp.MustCompile`** при инициализации, если уверены в корректности паттерна.

3. **Избегайте сложных выражений** для больших текстов - они могут быть медленными.

4. **Для простых операций** иногда лучше использовать `strings`:
   ```go
   // Вместо regexp.MustCompile(`^prefix`)
   strings.HasPrefix(str, "prefix")
   ```

5. **Ограничения**:
    - Нет поддержки lookahead/lookbehind
    - Максимальная длина совпадения по умолчанию 1MB (можно изменить)

## Отладка регулярных выражений

Используйте онлайн-инструменты:
- [Regex101](https://regex101.com/)
- [Rego](https://regoio.herokuapp.com/)

Пример проверки:
```go
re := regexp.MustCompile(`your pattern`)
fmt.Printf("Regexp: %#v\n", re.String())
```