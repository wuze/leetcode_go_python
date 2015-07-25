/**********************************************************************************
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * For example,
 * a = "11"
 * b = "1"
 * Return "100".
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
	cbit := 0
	var carry bool

	for alen > 0 || blen > 0 {
		abit := 0
		bbit := 0

		if alen <= 0 {
			abit = 0
		} else {
			abit = int(a[alen-1]) - 48
		}

		if blen <= 0 {
			bbit = 0
		} else {
			bbit = int(b[blen-1]) - 48
		}

		ret = strconv.Itoa((abit+bbit+cbit)&1) + ret

		carry = (abit + bbit + cbit) > 1

		if carry {
			cbit = 1
		} else {
			cbit = 0
		}

		fmt.Println("--", abit, bbit, cbit, "--")

		alen--
		blen--
	}

	if carry {
		cbit = 1
	} else {
		cbit = 0
	}

	if cbit == 1 {
		ret = "1" + ret
	}
	return
}

func main() {
	a := "101"
	b := "1101"
	c := addTwoNum(a, b)
	fmt.Printf("%s + %s = %s\n", a, b, c)
}
