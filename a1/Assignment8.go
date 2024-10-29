package main

import (
	"errors"
	"fmt"
)

func divide(lhs, rhs float64) (float64, error) {
	if rhs == 0 {
		return 0, errors.New("division by zero")
	}
	return lhs / rhs, nil
}

func main() {
	ret1, err1 := divide(5, 2)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(ret1)
	}

	ret2, err2 := divide(1, 0)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(ret2)
	}
}
