package main

import "fmt"

func main() {
	var a [5] int

	a[4] = 100

	fmt.Println(a)

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(b)

	var c [3][3] int

	for i := 1; i <= 2; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)

}
