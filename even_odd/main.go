package main

import "fmt"

func main() {

	is := []int{}
	for i := 0; i < 11; i++ {
		is = append(is, i)
	}

	for j := range is {
		if j%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}

}
