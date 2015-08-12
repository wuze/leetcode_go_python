package main

import (
	"fmt"
)

func displayData(line int) {
	data := make([][]int, 7, 7)

	for i := 0; i < line; i++ {
		tmp := make([]int, 7, 7)
		data[i] = tmp
	}

	data[0][0] = 1
	for i := 1; i < line; i++ {
		for j := 0; j <= i; j++ {
			if j == 0 {
				data[i][0] = 1
				continue
			}

			if i == j {
				data[i][j] = 1
			} else {
				data[i][j] = data[i-1][j-1] + data[i-1][j]
			}
		}
	}

	for i := 0; i < line; i++ {
		// 空格
		for j := 0; j < line-i; j++ {
			fmt.Printf("  ")
		}

		// 显示数字
		for _, v := range data[i] {
			if v > 0 {
				fmt.Printf("%2d  ", v)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	line := 7
	fmt.Printf("显示%d行\n", line)

	displayData(line)
}
