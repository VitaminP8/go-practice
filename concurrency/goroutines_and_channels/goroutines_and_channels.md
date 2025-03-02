# Concurrency в Go: Goroutines и Channels

Go предоставляет мощные инструменты для работы с конкурентностью: **goroutines** и **channels**. Эти инструменты позволяют эффективно выполнять задачи параллельно и безопасно обмениваться данными между потоками.

---

## Goroutines

**Goroutine** — это легковесный поток, управляемый Go-рантаймом. Goroutines позволяют выполнять функции асинхронно.

### Как создать goroutine?

Для запуска функции в отдельной goroutine используйте ключевое слово `go`:

```go
package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Запуск функции в goroutine
	time.Sleep(2 * time.Second) // Ожидание завершения goroutine
	fmt.Println("Main function")
}
```

### Особенности goroutines:
- Легковесные: тысячи goroutines могут работать одновременно.
- Управляются Go-рантаймом, а не операционной системой.
- Завершаются, когда завершается основная программа (функция `main`).

---

## Channels

**Channel** — это механизм для безопасной передачи данных между goroutines. Каналы позволяют синхронизировать выполнение goroutines.

### Как создать и использовать канал?

Каналы создаются с помощью функции `make` и могут быть буферизированными или небуферизированными.

#### Небуферизированный канал:
```go
package main

import "fmt"

func main() {
	ch := make(chan int) // Создание небуферизированного канала

	go func() {
		ch <- 42 // Отправка данных в канал
	}()

	value := <-ch // Получение данных из канала
	fmt.Println(value) // Output: 42
}
```

#### Буферизированный канал:
```go
package main

import "fmt"

func main() {
	ch := make(chan int, 2) // Создание буферизированного канала с размером буфера 2

	ch <- 1
	ch <- 2

	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
}
```

### Особенности каналов:
- Отправка и получение данных блокируют выполнение, пока данные не будут переданы.
- Закрытие канала: `close(ch)`.
- Проверка на закрытие канала:
  ```go
  value, ok := <-ch
  if !ok {
      fmt.Println("Channel closed")
  }
  ```

---


## Select: Работа с несколькими каналами

**`select`** — это конструкция, которая позволяет goroutine ожидать выполнения одной из нескольких операций с каналами. Она похожа на `switch`, но используется для каналов.

### Основное использование `select`

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second): // Таймаут
		fmt.Println("Timeout")
	}
}
```

### Объяснение:
1. **`select`** ожидает, пока одна из операций с каналами (`ch1` или `ch2`) не станет доступной.
2. Если доступны несколько операций, `select` выбирает одну из них случайным образом.
3. **`time.After`** используется для добавления таймаута. Если ни один канал не готов в течение указанного времени, выполняется ветка `time.After`.


### Особенности select

- **`select`** позволяет работать с несколькими каналами одновременно.
- Используйте `select` для:
   - Ожидания данных из нескольких каналов.
   - Реализации таймаутов.
   - Неблокирующих операций с каналами.
   - Остановки goroutines.

`select` — это мощный инструмент для управления конкурентностью в Go, который делает код более гибким и эффективным. 

---

# Закрытие каналов в Go: `close`

В Go каналы можно закрывать с помощью функции **`close`**. Закрытие канала используется для сигнализации о завершении работы и предотвращения дальнейшей отправки данных.

---

## Как закрыть канал?

Для закрытия канала используется функция `close`:

```go
ch := make(chan int)
close(ch)
```

---

## Особенности закрытия каналов

1. **Отправка данных в закрытый канал**:
   - При попытке отправить данные в закрытый канал возникает паника (`panic`).

   ```go
   ch := make(chan int)
   close(ch)
   ch <- 42 // panic: send on closed channel
   ```

2. **Чтение из закрытого канала**:
   - Если канал закрыт, чтение из него возвращает нулевое значение и `false` в качестве второго значения.

   ```go
   ch := make(chan int)
   close(ch)
   value, ok := <-ch
   fmt.Println(value, ok) // 0, false
   ```

3. **Итерация по каналу**:
   - Если канал закрыт, цикл `for range` завершается автоматически.

   ```go
   ch := make(chan int, 3)
   ch <- 1
   ch <- 2
   close(ch)

   for value := range ch {
       fmt.Println(value) // 1, 2
   }
   ```

### итоги close

- Используйте `close`, чтобы сигнализировать о завершении работы с каналом.
- Закрытие канала предотвращает дальнейшую отправку данных.
- Чтение из закрытого канала возвращает нулевое значение и `false`.

Закрытие каналов — это важный механизм для управления жизненным циклом goroutines и синхронизации в Go.

---

## GOMAXPROCS: Параллельное выполнение на одном потоке

Go использует **многопоточность** для выполнения goroutines. Количество потоков, которые могут выполняться параллельно, контролируется с помощью `runtime.GOMAXPROCS`.

### Что такое `GOMAXPROCS`?

- `GOMAXPROCS` определяет максимальное количество потоков, которые могут выполняться одновременно.
- По умолчанию `GOMAXPROCS` равно количеству ядер процессора (`runtime.NumCPU()`).
- Вы можете изменить это значение с помощью `runtime.GOMAXPROCS(n)`.

### Пример использования `GOMAXPROCS`:

```go
package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoMaxProc(t *testing.T) {
	runtime.GOMAXPROCS(1) // Устанавливаем максимальное количество потоков
	fmt.Println("CPU cores:", runtime.NumCPU())
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println("start", num)

			// точка возможного переключения горутин
			time.Sleep(1 * time.Millisecond)
			fmt.Println("stop", num)
		}(i)
	}

	time.Sleep(1 * time.Second) // Ожидание завершения goroutines
}
```

### Объяснение:
1. **`runtime.GOMAXPROCS(1)`**:
    - Устанавливает максимальное количество потоков, которые могут выполняться параллельно, равным 1.
    - Это означает, что даже если у вас больше ядер, Go будет использовать только 1 поток.

2. **`runtime.NumCPU()`**:
    - Возвращает количество ядер процессора.

3. **`runtime.GOMAXPROCS(0)`**:
    - Возвращает текущее значение `GOMAXPROCS`.

4. **Параллельное выполнение**:
    - Даже если `GOMAXPROCS` установлено в 1, Go может переключаться между goroutines на одном потоке, создавая иллюзию параллельного выполнения.

5. **`time.Sleep`**:
    - Используется для ожидания завершения goroutines. В реальных приложениях лучше использовать `sync.WaitGroup`.

---

## Пример: Worker Pool

Использование goroutines и каналов для создания пула воркеров:

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Имитация работы
		results <- job * 2
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	// Запуск 3 воркеров
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Отправка задач
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Получение результатов
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
}
```

---

## Синхронизация

Для синхронизации goroutines можно использовать:
- Каналы.
- Примитивы синхронизации из пакета `sync` (например, `sync.WaitGroup`).

### Пример с `sync.WaitGroup`:

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup
	fmt.Printf("Worker %d starting\n", id)
	// Имитация работы
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go worker(i, &wg)
	}

	wg.Wait() // Ожидаем завершения всех goroutines
	fmt.Println("All workers done")
}
```

---

## Заключение

- **Goroutines** позволяют легко запускать асинхронные задачи.
- **Channels** обеспечивают безопасный обмен данными между goroutines.
- Используйте `sync.WaitGroup` для синхронизации выполнения goroutines.
- **`GOMAXPROCS`** контролирует количество потоков, которые могут выполняться параллельно.

Go делает конкурентное программирование простым и эффективным благодаря goroutines, channels и `GOMAXPROCS`. 🚀
```

Теперь заметка включает информацию о `GOMAXPROCS` и примере кода, который демонстрирует, как работает параллельное выполнение на одном потоке.