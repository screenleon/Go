package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func choose(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2 or 3")
	case 4:
		fallthrough
	case 5:
		fmt.Println("5")
	}
}

func anotherChoose(i int) {
	switch {
	case i > 5:
		fmt.Println(i, "> 5")
	case i < 5:
		fmt.Println(i, "< 5")
	case i == 5:
		fmt.Println(i, "= 5")
	}
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	// while
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// infinite loop
	for {
		sum++
		if sum > 0 {
			break
		}
	}

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println("choose method:")
	choose(1)
	choose(3)
	choose(4)

	fmt.Println("anotherChoose method:")
	anotherChoose(6)

	// goto
	i := 0
HERE:
	fmt.Print(i)
	i++
	if i < 10 {
		goto HERE
	}
}
