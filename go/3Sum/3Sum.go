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
 */

package main

import (
	"fmt"
	"sort"
)

func mySum(arr []int) (ret [][3]int) {
	sort.Ints(arr)
	var count = len(arr)
	fmt.Println("After sort: ")
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n---------\n")
	var idx = 0
	ret = make([][3]int, 10)
	for i := 0; i < count-2; i++ {

		var current = i
		var low = i + 1
		var high = count - 1

		for low < high {
			a := arr[low]
			b := arr[high]

			if arr[current]+a+b == 0 {
				ret[idx][0] = arr[current]
				ret[idx][1] = a
				ret[idx][2] = b
				for low < count && arr[low] == arr[low+1] {
					low++
				}

				for high > 0 && arr[high] == arr[high-1] {
					high--
				}

				low++
				high--
				idx++
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