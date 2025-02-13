package data_types

import (
	"fmt"
)

// тэги элементов структуры
type _ struct {
	id   int    `db:"id"`
	name string `db:"name"`
	age  int    `db:"how_ald"`
}

// замыкание (тк count используется во вложенной функции - удаляться он не будет при завершении функции)
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// в go всегда явное преобразование
	var i int32 = 42
	var k int32 = int32(i)
	var m int64 = int64(i)
	println(k, " ", m)

	// массивы
	var _ [10]int
	var _ [10][10]string
	arr := [10]int{1, 2, 3, 6: 10}
	fmt.Println(arr)

	// слайсы
	var _ []int
	_ = []int{}
	slice := make([]int, 3, 4)
	fmt.Println(slice)

	s := slice[2:4:4]
	s = append(s, 6, 7, 8)
	fmt.Println(s)

	// строки
	str := "Hello world"
	var c byte = str[0]
	l := len(str)
	fmt.Println(str, " ", l)
	fmt.Println(c)

	// итерация по рунам
	for i, v := range str {
		fmt.Println(i, v)
	}
	for i, v := range str {
		fmt.Println(i, string(v))
	}

	// словари
	_ = map[string]int{}
	_ = make(map[string]int)
	market := make(map[string]int, 10)
	market["btc"] = 100
	market["eth"] = 10
	market["eur"] = 50
	market["usd"] = 40
	value, ok := market["usd"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no usd here")
	}
	delete(market, "usd")
	fmt.Println(market)

	// анонимные функции
	func() {
		fmt.Println("hello ")
	}() // () - означает что мы сразу вызвали функцию

	sayWorld := func() {
		fmt.Println("world")
	}
	sayWorld()

	// замыкание
	increment := counter()
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2
	fmt.Println(increment()) // 3
	// создается новый указатель на count
	increment2 := counter()
	fmt.Println(increment2()) // 1
}
