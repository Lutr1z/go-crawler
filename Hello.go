package main

import (
	"fmt"
)

func adder() func(int) int {
	num := 0
	return func(v int) int {
		num += v
		return num
	}
}

func main() {
	var do_not_show int
	a := adder()

	for i := 3; i > 0; i-- {
		fmt.Println(a(i))
	}
	b := adder()
	fmt.Println(b(2))

	do_not_show = 4
	fmt.Println(do_not_show)
}
