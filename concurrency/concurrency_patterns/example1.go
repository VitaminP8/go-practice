package concurrency_patterns

import "fmt"

func main() {
	doWork := func(done <-chan struct{}, strings <-chan string) <-chan struct{} {
		terminated := make(chan struct{})
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(terminated)
			}()

			for {
				select {
				case <-done:
				default:
					select {
					case s := <-strings:
						fmt.Println(s)
					case <-done:
						return
					}
				}
			}
		}()
		return terminated
	}
}
