# `sync.Mutex` в Go: Синхронизация доступа к общим данным

**`sync.Mutex`** — это примитив синхронизации в Go, который позволяет защитить доступ к общим данным в многопоточной среде. Он обеспечивает взаимное исключение (mutual exclusion), гарантируя, что только одна горутина может получить доступ к защищенным данным в любой момент времени.

---

## Основное использование

`sync.Mutex` имеет два основных метода:
1. **`Lock()`**:
    - Блокирует мьютекс. Если мьютекс уже заблокирован, горутина будет ждать, пока он не освободится.
2. **`Unlock()`**:
    - Разблокирует мьютекс, позволяя другим горутинам получить доступ к данным.

---

### Пример использования

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем мьютекс
	defer c.mu.Unlock() // Разблокируем мьютекс при выходе из функции
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()         // Блокируем мьютекс
	defer c.mu.Unlock() // Разблокируем мьютекс при выходе из функции
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter value:", counter.Value()) // Output: Counter value: 100
}
```

---



## Особенности

1. **Эксклюзивный доступ**:
    - Только одна горутина может владеть мьютексом в любой момент времени.

2. **Блокировка и ожидание**:
    - Если мьютекс заблокирован, другие горутины будут ждать, пока он не освободится.

3. **Рекурсивная блокировка**:
    - В Go мьютекс не является рекурсивным. Если горутина попытается заблокировать мьютекс, который она уже заблокировала, это приведет к deadlock.

4. **Использование с другими примитивами**:
    - `sync.Mutex` часто используется вместе с `sync.WaitGroup` или каналами для более сложных сценариев синхронизации.

---

## Пример: Защита доступа к map

```go
package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	data map[string]int
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.data[key]
	return value, exists
}

func main() {
	sm := SafeMap{data: make(map[string]int)}
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		value, _ := sm.Get(key)
		fmt.Printf("%s: %d\n", key, value)
	}
}
```

#### Вывод:
```
key0: 0
key1: 1
key2: 2
key3: 3
key4: 4
key5: 5
key6: 6
key7: 7
key8: 8
key9: 9
```

---

## Deadlock

Если мьютекс не разблокировать, это приведет к deadlock:

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock() // Deadlock: мьютекс уже заблокирован
	fmt.Println("This will never be printed")
}
```

---

## Заключение

- **`sync.Mutex`** — это простой и эффективный способ защиты доступа к общим данным в многопоточной среде.
- Используйте `Lock()` для блокировки и `Unlock()` для разблокировки мьютекса.
- Всегда используйте `defer` для вызова `Unlock()`, чтобы избежать deadlock.
- `sync.Mutex` идеально подходит для сценариев, где требуется эксклюзивный доступ к данным.
