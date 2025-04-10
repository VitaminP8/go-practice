# CLI в Go

## Основные цели

- Создавать консольные утилиты на Go
- Работать с аргументами командной строки и флагами
- Обрабатывать переменные окружения
- Запускать внешние программы
- Работать с сигналами
- Создавать временные файлы и управлять файловой системой

---

## 1. Аргументы и флаги

### 1.1 Стандартный пакет `flag`
```go
import "flag"

func main() {
    verbose := flag.Bool("verbose", false, "Verbose mode")
    port := flag.Int("port", 8080, "Port to run server on")

    flag.Parse()

    if *verbose {
        fmt.Println("Verbose mode enabled")
    }
}
```

---

### 1.2 Библиотека `pflag` (совместима с POSIX/GNU стилями)
```go
import "github.com/spf13/pflag"

func main() {
    var cfg string
    verbose := pflag.BoolP("verbose", "v", false, "Verbose output")
    pflag.StringVar(&cfg, "cfg", "config.yaml", "Path to config")
    pflag.Parse()

    if *verbose {
        fmt.Println("Verbose mode:", cfg)
    }
}
```

- Поддерживает `--flag=value` и `-f value`
- `NoOptDefVal` — поведение флага без значения:
```go
pflag.StringVar(&val, "port", "80", "Port")
pflag.Lookup("port").NoOptDefVal = "8080"
```

---

### 1.3 Фреймворки для CLI

#### [`cobra`](https://github.com/spf13/cobra)
- Идеален для CLI с подкомандами (git-like)
- Генератор `cobra-cli`
- Используется во многих проектах (Kubernetes CLI, Hugo)

#### [`urfave/cli`](https://github.com/urfave/cli)
- Прост в освоении
- Подходит для небольших CLI

---

## 2. Переменные окружения

```go
env := os.Environ() // []string{"KEY=VALUE"}
user, ok := os.LookupEnv("USER")     // безопасный доступ
os.Setenv("KEY", "val")              // установка переменной
os.Unsetenv("KEY")                   // удаление
os.ExpandEnv("$USER lives in $CITY") // шаблоны в строке
```

---

## 3. Запуск внешних программ

```go
import "os/exec"

func main() {
    cmd := exec.Command("git", "commit", "-am", "fix")

    // полный вывод (stdout + stderr)
    output, err := cmd.CombinedOutput()
    fmt.Println(string(output), err)

    // отдельно запуск
    cmd.Run()
}
```

- Можно управлять stdin, stdout, stderr через поля структуры `Cmd`
- `Start()` запускает асинхронно, `Wait()` нужно вызывать вручную

---

## 4. Обработка сигналов

Сигналы — механизм ОС для уведомлений процессов (например, Ctrl+C).

```go
import (
    "os"
    "os/signal"
    "syscall"
)

func main() {
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

    for s := range c {
        fmt.Println("Got signal:", s)
    }
}
```

- Некоторые сигналы можно игнорировать (`signal.Ignore`)
- SIGKILL и SIGSTOP не обрабатываются (система завершает/останавливает принудительно)

---

## 5. Временные файлы и файловая система

### Работа с файловой системой (`os`)
```go
os.Mkdir("dir", 0755)
os.Rename("old.txt", "new.txt")
os.Remove("file.txt")
```

### Временные файлы
```go
file, err := os.CreateTemp("", "myapp-*")
defer os.Remove(file.Name())
```

#### Безопасное создание (рекоммендация):
- [`github.com/dchest/safefile`](https://github.com/dchest/safefile)

---

## Ссылки и полезные ресурсы

- 📖 GNU CLI стандарты: https://www.gnu.org/prep/standards/
- 📦 `flag`: https://pkg.go.dev/flag
- 📦 `pflag`: https://github.com/spf13/pflag
- 📦 `cobra`: https://github.com/spf13/cobra
- 🌐 https://clig.dev — советы по CLI-дизайну

