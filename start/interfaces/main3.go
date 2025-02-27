package interfaces

//package main

import "fmt"

type Animal interface {
	Voice()
}

type Cat struct {
	Name  string
	Color string
}

func (c *Cat) Voice() {
	fmt.Println("Мяу...")
}

type Dog struct {
	Name string
	Age  int
}

func (d *Dog) Voice() {
	fmt.Println("Гав...")
}

func main() {
	d := Dog{"Тузик", 5}
	c := Cat{"Богдан", "Черный"}
	Handle(&d)
	Play(&d)
	Play(&c)
}

func Handle(v any) {
	a, ok := v.(Animal)
	if ok {
		a.Voice()
	}
}

func Play(v any) {
	switch pet := v.(type) {
	case *Dog:
		fmt.Println("Чешем собаку: ", pet.Name, " которой ", pet.Age, " лет")
	case *Cat:
		fmt.Println("Гладим кота ", pet.Color, " цвета")
	case Animal:
	default:
		fmt.Println("я не знаю кто это")
	}

}
