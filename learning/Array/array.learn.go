package main

import "fmt"

func main() {
	var x [4]float64
	x[0] = 23
	x[1] = 45
	x[2] = 33
	x[3] = 21
	// x := [4]float64{23, 45, 33, 21}
	// x := [4]float64{
	// 	23,
	// 	45,
	// 	33,
	// 	21,
	// }

	var total float64 = 0
	for i := 0; i < 4; i++ {
		total += x[i]
	}
	fmt.Println(total / 4)

	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
	slice3 := append(slice1, 4, 5)
	fmt.Println(slice3)
}
