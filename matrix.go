package main

import "fmt"

func main() {

	var matArry [10][10]int
	a := 100
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			matArry[i][j] = j + a
		}
	}
	printArray(matArry)

}
func printArray(a [10][10]int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(a[i][j], ",")
		}
		fmt.Println()
	}

}
