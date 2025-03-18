package concurrency_patterns

import "fmt"

func generator(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * 2
		}
		close(out)
	}()
	return out
}

func add(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num + 1
		}
		close(out)
	}()
	return out
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("Received:", num)
	}
}

func main() {
	nums := generator(5)
	multiplied := multiply(nums)
	added := add(multiplied)

	consumer(added)
}
