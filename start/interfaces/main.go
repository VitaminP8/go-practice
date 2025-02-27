package interfaces

import "fmt"

type Speaker interface {
	SayHello()
}

type Human struct {
	Greeting string
}

func (h Human) SayHello() {
	fmt.Println(h.Greeting)
}

func main() {
	var s Speaker
	h := Human{Greeting: "Hello"}
	s = Speaker(h)
	s.SayHello()
}
