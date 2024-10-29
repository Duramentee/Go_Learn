package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args {
		if idx == 0 {
			fmt.Println("instruction name: " + arg)
		}
		fmt.Println(idx, "=", arg)
	}
}
