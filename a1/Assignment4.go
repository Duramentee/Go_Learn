package main

import "fmt"

func main() {
	numArr1 := [5]int{1, 2, 3, 4, 5}
	sum1 := 0
	for i := 0; i < len(numArr1); i++ {
		sum1 += numArr1[i]
	}
	fmt.Println("Average(array): ", sum1/len(numArr1))

	numArr2 := []int{1, 2, 3, 4, 5}
	sum2 := 0
	for _, val := range numArr2 {
		sum2 += val
	}
	fmt.Println("Average(slice): ", sum2/len(numArr2))
}
