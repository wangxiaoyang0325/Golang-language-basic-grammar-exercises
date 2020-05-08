package main

import "fmt"

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("未知")
	}

}
