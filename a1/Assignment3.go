package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		fmt.Printf("%d ", i)
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
