package main

import (
	"fmt"
	"os"
	"strconv"
)

const size = 9

func isValid(b [][]int, r, c, n int) bool {
	for i := 0; i < size; i++ {
		if b[r][i] == n || b[i][c] == n || b[3*(r/3)+i/3][3*(c/3)+i%3] == n {
			return false
		}
	}
	return true
}

func solveSudoku(b [][]int) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] == 0 {
				for n := 1; n <= size; n++ {
					if isValid(b, i, j, n) {
						b[i][j] = n
						if solveSudoku(b) {
							return true
						}
						b[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != size+1 {
		fmt.Println("Error")
		return
	}

	b := make([][]int, size)
	for i := range b {
		b[i] = make([]int, size)
		for j, c := range os.Args[i+1] {
			if c != '.' {
				b[i][j], _ = strconv.Atoi(string(c))
			}
		}
	}

	if solveSudoku(b) {
		for _, row := range b {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Error")
	}
}
