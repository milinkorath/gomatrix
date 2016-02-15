package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	// fd, _ := os.Open("input.txt")
	f, _ := ioutil.ReadFile("input.txt")
	// // fmt.Println(len(f))
	// // a := bytes.LastIndex(f, []byte(" "))
	// // fmt.Println(a)
	// // fmt.Println(fd.Seek(-8, 2))
	// //
	// // fmt.Println(bytes.Count(f, []byte("\n")))
	d := string(f)
	d = strings.Trim(d, "\n")
	size := strings.Count(d, "\n") + 1
	fmt.Println(size)

	s := int(math.Sqrt(float64(size)))
	matrix := make([][]int64, s)
	data := strings.Split(d, "\n")
	for i := 0; i < s; i++ {
		for j := 0; j < s; j++ {
			matrix[i] = append(matrix[i], 0)
		}
		// matrix[i] = append(matrix[i], 0)
	}
	for _, value := range data {
		lines := strings.Fields(value)
		row, _ := strconv.ParseInt(lines[0], 10, 64)
		col, _ := strconv.ParseInt(lines[1], 10, 64)
		val, _ := strconv.ParseInt(lines[2], 10, 64)
		matrix[row][col] = val
	}
	printArray1(matrix, "A =>")
	tran_matrix := transpose(matrix)
	printArray1(tran_matrix, "A` =>")
	result := multiply(matrix, tran_matrix)
	printArray1(result, "A*A` =>")
	// a := strings.Trim("Achtung\n! Achtung\nHello\n", "\n")
	// fmt.Println(strings.Count(a, "\n"))
}

func transpose(matrix [][]int64) [][]int64 {
	transpose := make([][]int64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			transpose[i] = append(transpose[i], matrix[j][i])
		}
	}
	return transpose
}

func multiply(a, b [][]int64) [][]int64 {
	size := len(a)
	multiply := make([][]int64, size)
	// temp := make([]int64, size)
	var sum int64
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				sum = sum + a[i][k]*b[k][j]
			}
			multiply[i] = append(multiply[i], sum)
			sum = 0
		}

	}
	return multiply
}

func printArray1(a [][]int64, b string) {
	fmt.Println(b)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			fmt.Print("\t", a[i][j], " ")
		}
		fmt.Println()
	}

}
