package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println(1, "+", 1, "=", 1+1)

	var floatValue float64
	floatValue = 7.0
	var floatValue2 = 3.0
	fmt.Println("7.0/3.0 =", floatValue/floatValue2)

	var test float64
	var test2 float32
	test = 1.1
	test2 = 2.2
	// float32 cant calculate with float64, need to Forced transformation
	fmt.Println("test + test2 =", float32(test)+test2)

	var complexValue complex64
	complexValue = 1.2 + 12i
	complexValue2 := 1.2 + 12i
	complexValue3 := complex(3.2, 12)

	fmt.Println("complexValue =", complexValue)
	fmt.Println("complexValue2 =", complexValue2)
	fmt.Println("complexValue3 =", complexValue3)

	fmt.Println("complexValue3 實數 =", real(complexValue3))
	fmt.Println("complexValue3 虛數 =", imag(complexValue3))

	fmt.Println("1" + "1")
	fmt.Println(len("Hello World"))
	// print ascii code
	fmt.Println("Hello World"[0])
	fmt.Println("Hello" + "World")

	a := "Hello World"
	fmt.Printf("%c\n", a[1])

	var c bool
	c = true
	fmt.Println("c =", c)

	b := false
	fmt.Println("b =", b)

	// fmt.Println(true && true) run with warning redundant
	fmt.Println(true && false)
	// fmt.Println(true || true) run with warning redundant
	fmt.Println(true || false)
	fmt.Println(!true)
}
