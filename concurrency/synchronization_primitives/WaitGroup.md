# WaitGroup в Go: Синхронизация горутин

**`sync.WaitGroup`** — это механизм синхронизации в Go, который позволяет дождаться завершения выполнения группы горутин. Он особенно полезен, когда вам нужно выполнить несколько задач параллельно и дождаться их завершения перед продолжением работы.

---

## Основное использование

`WaitGroup` имеет три основных метода:
1. **`Add(delta int)`**:
    - Увеличивает счетчик на `delta`. Указывает, сколько горутин нужно дождаться.
2. **`Done()`**:
    - Уменьшает счетчик на 1. Вызывается, когда горутина завершает свою работу.
3. **`Wait()`**:
    - Блокирует выполнение, пока счетчик не станет равным 0.

---

### Пример использования

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик при завершении
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second) // Имитация работы
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // Увеличиваем счетчик для каждой горутины
		go worker(i, &wg)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All workers done")
}
```

#### Вывод:
```
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 done
Worker 2 done
Worker 3 done
All workers done
```

---


## Особенности

1. **Счетчик**:
    - Счетчик `WaitGroup` не может быть отрицательным. Если вызвать `Done()` больше раз, чем `Add()`, программа завершится с паникой.

2. **Передача по указателю**:
    - `WaitGroup` всегда передается по указателю (`&wg`), так как он содержит внутреннее состояние, которое должно быть общим для всех горутин.

3. **Использование с другими примитивами**:
    - `WaitGroup` часто используется вместе с каналами или мьютексами для более сложных сценариев синхронизации.

---

## Пример: Параллельная обработка данных

```go
package main

import (
	"fmt"
	"sync"
)

func process(data int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Processing", data)
	// Имитация обработки данных
}

func main() {
	var wg sync.WaitGroup
	data := []int{1, 2, 3, 4, 5}

	for _, value := range data {
		wg.Add(1)
		go process(value, &wg)
	}

	wg.Wait()
	fmt.Println("All data processed")
}
```

#### Вывод:
```
Processing 5
Processing 1
Processing 2
Processing 3
Processing 4
All data processed
```

---

## Пример: Ожидание завершения горутин с результатами

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	result <- id * 2 // Отправляем результат в канал
}

func main() {
	var wg sync.WaitGroup
	result := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg, result)
	}

	// Горутина для закрытия канала после завершения всех задач
	go func() {
		wg.Wait()
		close(result)
	}()

	// Чтение результатов из канала
	for res := range result {
		fmt.Println("Result:", res)
	}
}
```

#### Вывод:
```
Result: 2
Result: 4
Result: 6
```

---

## Заключение

- **`WaitGroup`** — это простой и эффективный способ синхронизации горутин.
- Используйте `Add()` для увеличения счетчика, `Done()` для уменьшения и `Wait()` для ожидания завершения.
- Всегда передавайте `WaitGroup` по указателю.
- `WaitGroup` идеально подходит для сценариев, где нужно дождаться завершения группы горутин.
