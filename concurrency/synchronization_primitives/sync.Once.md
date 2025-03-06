# `sync.Once` в Go: Однократное выполнение

**`sync.Once`** — это примитив синхронизации в Go, который гарантирует, что определенная операция будет выполнена только один раз, даже если она вызывается из нескольких горутин. Это полезно для инициализации ресурсов, которые должны быть созданы только один раз.

---

## Основное использование

`sync.Once` имеет один метод:
- **`Do(f func())`**:
    - Выполняет функцию `f` только один раз, независимо от количества вызовов `Do`.

---

### Пример использования

```go
package main

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	config map[string]string
)

func initializeConfig() {
	fmt.Println("Initializing config...")
	config = map[string]string{
		"host": "localhost",
		"port": "8080",
	}
}

func getConfig() map[string]string {
	once.Do(initializeConfig) // Инициализация выполнится только один раз
	return config
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cfg := getConfig()
			fmt.Println("Config:", cfg)
		}()
	}

	wg.Wait()
}
```

#### Вывод:
```
Initializing config...
Config: map[host:localhost port:8080]
Config: map[host:localhost port:8080]
Config: map[host:localhost port:8080]
```

---


## Особенности

1. **Однократность**:
    - Функция `f` выполняется только один раз, независимо от количества вызовов `Do`.

2. **Потокобезопасность**:
    - `sync.Once` безопасен для использования в многопоточной среде.

3. **Использование с другими примитивами**:
    - `sync.Once` часто используется для ленивой инициализации ресурсов, таких как конфигурации, подключения к базе данных и т.д.

---

## Пример: Ленивая инициализация подключения к базе данных

```go
package main

import (
	"fmt"
	"sync"
)

var (
	once     sync.Once
	dbClient *DBClient
)

type DBClient struct {
	connection string
}

func initializeDB() {
	fmt.Println("Initializing DB client...")
	dbClient = &DBClient{connection: "localhost:5432"}
}

func getDBClient() *DBClient {
	once.Do(initializeDB) // Инициализация выполнится только один раз
	return dbClient
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := getDBClient()
			fmt.Println("DB Client:", client.connection)
		}()
	}

	wg.Wait()
}
```

#### Вывод:
```
Initializing DB client...
DB Client: localhost:5432
DB Client: localhost:5432
DB Client: localhost:5432
```

---

## Заключение

- **`sync.Once`** — это простой и эффективный способ гарантировать однократное выполнение операции.
- Используйте `Do(f)` для выполнения функции `f` только один раз.
- `sync.Once` идеально подходит для ленивой инициализации ресурсов, таких как конфигурации, подключения к базе данных и т.д.
