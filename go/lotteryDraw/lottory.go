/****
Control the chance/probability of lottory draw
for example
prize_list = map(
	'0':map('id':1, 'prize'=>'tv','num':1),
	'1':map('id':2, 'prize'=>'phone','num':10),
	'2':map('id':3, 'prize'=>'pad','num':32),
	'3':map('id':4, 'prize'=>'nothing','num':50)
)
prize_list  is read from a conf file
we finally read into a map
@author bravowoo@163.com
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Node struct {
	Mid   int
	Prize string
	v     int
}

func chance(arr map[int]int) (ret int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	total := 0
	for _, v := range arr {
		total += v
	}

	fmt.Println(total)

	for k, v := range arr {

		current := r.Intn(total)

		if current <= v {
			fmt.Println("随机数为:", current)
			ret = k
			break
		} else {
			total -= v
		}
	}

	return
}

func getMap() (ret map[string]*Node) {
	// read from config file
	m := make(map[string]*Node)
	m["0"] = &Node{Mid: 1, Prize: "tv", v: 1}
	m["1"] = &Node{Mid: 2, Prize: "phone", v: 20}
	m["2"] = &Node{Mid: 3, Prize: "pad", v: 32}
	m["3"] = &Node{Mid: 4, Prize: "nothing", v: 50}

	ret = m
	return
}

func main() {
	prize := getMap()
	total_sum := make(map[int]int)

	for _, item := range prize {
		//		fmt.Println(item.Mid, item.Prize, item.v)
		total_sum[item.Mid] = item.v
	}

	index := chance(total_sum)
	idx := strconv.Itoa(index - 1)

	fmt.Println("中了", index, ": ", prize[idx].Prize)
}
