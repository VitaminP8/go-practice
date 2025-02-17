package typical_features

//package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	defer func() {
		fmt.Println("закончилась функция main")
	}()
	defer fmt.Println("почти закончилась функция main")

	//сортировка слайсов
	s := []string{"one", "two", "three"}
	sort.Strings(s)
	fmt.Println("отсортированный слайс строк: ", s)

	n := []int{10, 6, 4}
	sort.Ints(n)
	fmt.Println("отсортированный слайс интов: ", n)

	// новый способ сортировки
	slices.Sort(n)
	fmt.Println("отсортированный слайс интов: ", n)
	slices.Sort(s)
	fmt.Println("отсортированный слайс строк: ", s)

}
