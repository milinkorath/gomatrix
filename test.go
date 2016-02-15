package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func is_square_root(number int) bool {
	root := math.Sqrt(float64(number))
	return root == math.Floor(root)
}

func is_square_matrix(matrix [][]string) bool {
	size := len(matrix)
	return is_square_root(size)
}

func transpose_matrix(matrix [][]int64) [][]int64 {
	size := int(len(matrix))

	new_matrix := make([][]int64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			new_matrix[i] = append(new_matrix[i], matrix[j][i])
		}
	}
	return new_matrix
}

func dot_product(array_1, array_2 []int64) int64 {
	size := int(len(array_1))
	total := int64(0)

	for i := 0; i < size; i++ {
		total += array_1[i] * array_2[i]
	}
	return total
}

func multiply_matrices(matrix_1, matrix_2 [][]int64) [][]int64 {
	size := int(len(matrix_1))

	solution_matrix := make([][]int64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			array_1 := make([]int64, 0)
			array_2 := make([]int64, 0)

			for k := 0; k < size; k++ {
				array_1 = append(array_1, matrix_1[i][k])
				array_2 = append(array_2, matrix_2[k][j])
			}

			total := dot_product(array_1, array_2)
			solution_matrix[i] = append(solution_matrix[i], total)
		}
	}
	return solution_matrix
}

func pretty_print(matrix [][]int64) {
	fmt.Println()
	for _, value := range matrix {
		fmt.Println(value)
	}
}

func main() {

	data, _ := ioutil.ReadFile("./input_file10.txt")

	str := string(data)
	str = strings.Trim(str, "\n")

	slices := strings.Split(str, "\n")

	grid := make([][]string, 0)

	for _, value := range slices {
		list := strings.Fields(value)
		grid = append(grid, list)
	}

	if !is_square_matrix(grid) {
		fmt.Println("not a valid matrix")
	}

	size := int(math.Sqrt(float64(len(grid))))

	matrix := make([][]int64, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}

	for _, value := range grid {
		row, _ := strconv.ParseInt(value[0], 10, 64)
		row = row - 1
		column, _ := strconv.ParseInt(value[1], 10, 64)
		column = column - 1
		matrix[row][column], _ = strconv.ParseInt(value[2], 10, 64)
	}

	pretty_print(matrix)
	new_matrix := transpose_matrix(matrix)
	pretty_print(new_matrix)

	solution := multiply_matrices(matrix, new_matrix)

	pretty_print(solution)

}
