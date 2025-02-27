package interfaces

//package main

import "fmt"

type II interface {
	Add() int
}

type T struct {
	a, b int
}

func (t *T) Add() int {
	return t.a + t.b
}

func main() {
	t := T{10, 20}
	i := II(&t)

	fmt.Println(i.Add())
	fmt.Println(t.Add())
	t.a *= 10
	t.b *= 10
	fmt.Println(i.Add())
	fmt.Println(t.Add())
}
