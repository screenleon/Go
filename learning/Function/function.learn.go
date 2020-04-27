package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func anotheradd(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(anotheradd(42, 58))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
}
