/****
test interface
from
http://blog.csdn.net/eclipser1987/article/details/9331881
***/

package main

import "fmt"

type Animal interface {
	move()
}

type Human struct {
	i int
}

type Bird struct {
	i int
}

func (r Human) move() {
	fmt.Println("人类行走")
	r.i++
}

func (r *Bird) move() {
	fmt.Println("鸟类行走")
	r.i++
}

func moveTest1(animal Animal) {
	animal.move()
}

func moveTest2(animal *Animal) {
	(*animal).move()
}

func main() {
	h1 := Human{0}
	moveTest1(h1)
	moveTest1(h1)
	moveTest1(h1)
	fmt.Println(h1.i)
	fmt.Println("------------")

	h2 := &Human{0}
	moveTest1(h2)
	moveTest1(h2)
	moveTest1(h2)
	moveTest1(h2)
	fmt.Println(h2.i)
	fmt.Println("------------")

	// 接口再调用的时候要使用指针
	b2 := &Bird{0}
	moveTest1(b2)
	moveTest1(b2)
	moveTest1(b2)
	moveTest1(b2)
	fmt.Println(b2.i)
}
