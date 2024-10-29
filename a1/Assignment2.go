package main

import "fmt"

func main() {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("unexpected number.")
		return
	}

	if number%2 == 1 {
		fmt.Println("the number you enter is an odd")
	} else {
		fmt.Println("the number you enter is an even")
	}
}
