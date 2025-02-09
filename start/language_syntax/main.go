package language_syntax

import "fmt"

const temp = 451
const (
	zero = iota
	one
	two
	three
	_
	five
)

func div(a, b int) (int, int) {
	return a / b, a % b
}

func main() {
	var temperature = 15
	var message string
	ok := true
	_, _, _ = temperature, message, ok // приемник значений

	div(100, 3)         // игнорируем результат
	d, r := div(100, 3) // сохраняем оба результата
	_, r = div(100, 3)  // сохраняем только второй параметр
	d, _ = div(100, 3)  // сохраняем только первый параметр
	_, _ = div(100, 3)  // ничего не сохраняем
	fmt.Println(d, r)

	if temperature > 20 {
		fmt.Println("Тепло")
	} else if temperature > 10 {
		fmt.Println("Прохладно")
	} else if temperature > 0 {
		fmt.Println("Холодно")
	} else {
		fmt.Println("Мороз")
	}

	pet := "змея"
	var petName string
	switch pet {
	case "кот":
		petName = "Мурзик"
	case "соабка":
		petName = "Тузик"
	case "попугай":
		petName = "Кеша"
	default:
		fmt.Println("такое животное мы заводить не будем")
	}
	fmt.Println(petName)
}
