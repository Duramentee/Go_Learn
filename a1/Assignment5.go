package main

import "fmt"

func add(a int, b int) (sum int) {
	return a + b
}

func main() {
	fmt.Println(add(1, 1))
}
