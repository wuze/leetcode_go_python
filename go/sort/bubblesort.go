package main

import "fmt"

func bubbleSort(arr []int) (ret []int) {
	lenth := len(arr) // {{{

	fmt.Println(lenth)
	for i := 1; i < lenth; i++ {
		for j := 0; j < lenth-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}

	ret = arr
	return // }}}
}

/***
先选出最小的一个与第一个位置的交换，　再在剩下的选取最小的
与第二个数交换
**/

func selectSort(arr []int) (ret []int) {
	lenth := len(arr) // {{{

	for i := 0; i < lenth-1; i++ {
		p := i
		for j := i + 1; j < lenth; j++ {
			if arr[p] > arr[j] {
				p = j
			}
		}

		if p != i {
			tmp := arr[p]
			arr[p] = arr[i]
			arr[i] = tmp
		}
	}

	ret = arr
	return // }}}
}

/**
从第二个数字开始，　如果这个数字比前面的一个小
**/
func insertSort(arr []int) (ret []int) {
	lenth := len(arr) // {{{
	for i := 1; i < lenth; i++ {
		tmp := arr[i]
		for j := i - 1; j > 0; j-- {
			if tmp < arr[j] {
				arr[j+1] = arr[j]
				arr[j] = tmp
			} else {
				break
			}
		}
	}
	ret = arr
	return // }}}
}

/***
 选取一个基准，　基准左右的数据　分别再两个数组中
 递归执行下去
**/
func quckSort(arr []int) (ret []int) {
	lenth := len(arr) // {{{
	base := arr[0]

	var left_arr []int
	var right_arr []int

	for i := 1; i < lenth; i++ {
		if base > arr[i] {
			left_arr = append(left_arr, arr[i])
		} else {
			right_arr = append(right_arr, arr[i])
		}
	}

	if len(left_arr) > 0 {
		left_arr = quckSort(left_arr)
	}

	if len(right_arr) > 0 {
		right_arr = quckSort(right_arr)
	}

	if len(left_arr) > 0 {
		for _, v := range left_arr {
			ret = append(ret, v)
		}
	}

	ret = append(ret, base)

	if len(right_arr) > 0 {
		for _, v := range right_arr {
			ret = append(ret, v)
		}
	}

	return ret // }}}
}

func main() {
	arr := []int{-1, 10, 0, 12, 12, 21, 23, 100, -8} // {{{

	fmt.Println("Origin:")
	fmt.Println(arr)

	fmt.Println("Bubble SORT:")
	ret := bubbleSort(arr)
	fmt.Println(ret)
	fmt.Println("\n-------\n")

	fmt.Println("SELECT Sort:")
	ret = selectSort(arr)
	fmt.Println(ret)
	fmt.Println("\n-------\n")

	fmt.Println("insert Sort:")
	ret = insertSort(arr)
	fmt.Println(ret)
	fmt.Println("\n-------\n")

	fmt.Println("quick Sort:")
	ret = quckSort(arr)
	fmt.Println(ret)
	fmt.Println("\n-------\n")
	// }}}
}
