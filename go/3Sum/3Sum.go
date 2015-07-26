/**
* Given an array S of n integers, are there elements a, b, c in S such t
* Find all unique triplets in the array which gives the sum of zero.
*
* Note:
*
* Elements in a triplet (a,b,c) must be in non-descending order. (ie, a
* The solution set must not contain duplicate triplets.
*
*     For example, given array S = {-1 0 1 2 -1 -4},
*
*     A solution set is:
*     (-1, 0, 1)
*     (-1, -1, 2)
@author bravowoo@163.com
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 去重
func checkRet(arr [][3]int) (ret [][3]int) {
	var str []string
	for _, v := range arr {
		if len(v) > 0 {
			str = append(str, strconv.Itoa(v[0])+strconv.Itoa(v[1])+strconv.Itoa(v[2]))
		}
	}

	strLen := len(str)
	if strLen > 2 {
		i := 0
		for i = 0; i < strLen-1; i++ {
			flag := 0
			for j := i + 1; j < strLen; j++ {
				if str[i] == str[j] {
					flag = 1
				}
			}

			if flag == 0 {
				ret = append(ret, arr[i])
			}
		}
		// 最后一个元素
		ret = append(ret, arr[i])
	}

	return
}

func mySum(arr []int) (ret [][3]int) {
	sort.Ints(arr)
	var count = len(arr)
	fmt.Println("After sort: ")
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n---------\n")
	for i := 0; i < count-2; i++ {

		var current = i
		var low = i + 1
		var high = count - 1

		for low < high {
			a := arr[low]
			b := arr[high]

			if arr[current]+a+b == 0 {

				ret = append(ret, [3]int{arr[current], a, b})
				for low < count && arr[low] == arr[low+1] {
					low++
				}

				for high > 0 && arr[high] == arr[high-1] {
					high--
				}

				low++
				high--
			} else if current+a+b > 0 {

				for high > 0 && arr[high] == arr[high-1] {
					high--
				}
				high--
			} else {
				for low < count && arr[low] == arr[low+1] {
					low++
				}

				low++
			}
		}
	}

	ret = checkRet(ret)
	return
}

func main() {
	var arr = []int{-1, 0, 1, 2, -1, -4}
	ret := mySum(arr)

	if len(ret) > 0 {
		fmt.Println("The result is:")
		for i := 0; i < len(ret); i++ {
			fmt.Printf("%d %d %d\n", ret[i][0], ret[i][1], ret[i][2])
		}
	} else {
		fmt.Println("Found Nothing")
	}

}
