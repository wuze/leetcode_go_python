/**********************************************************************************
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * For example,
 * a = "11"
 * b = "1"
 * Return "100".
 *
 *
 **********************************************************************************/

package main

import (
	"fmt"
	"strconv"
)

func addTwoNum(a, b string) (ret string) {

	alen := len(a)
	blen := len(b)
	carray := 0

	for alen > 0 || blen > 0 {
		abit := 0
		bbit := 0
		if alen <= 0 {
			abit = 0
		} else {
			abit = a[alen-1] - '0'
		}
		if blen <= 0 {
			bbit = 0
		} else {
			bbit = b[blen-1] - '0'
		}

		tmp := abit + bbit + carray
		c = strconv.Itoa(tmp&1) + c

	}
	return
}

func main() {
	a := "101"
	b := "1201"
	c := addTwoNum(a, b)
	fmt.Printf("%s + %s = %s\n", a, b, c)
}
