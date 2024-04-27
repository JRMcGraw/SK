

package main

import (
	"fmt"
	"math"
)

type arm int

type arms struct {
	arm      int
	n1       float64
	n2       float64
	n3       float64
	arm_calc float64
	isArm    bool
}

func main() {

	arm_test := arms{}
	arm_test.showArm(371)

	fmt.Println("Well, is it?", arm_test)
}

func (a *arms) showArm(n int) {

	n1 := float64(n / 100)
	n2 := float64((n % 100) / 10)
	n3 := float64(n % 10)

	calculated := math.Pow(n1, 3) + math.Pow(n2, 3) + math.Pow(n3, 3)

	a.arm = n
	a.n1 = n1
	a.n2 = n2
	a.n3 = n3
	a.arm_calc = calculated

	a.isArm = float64(n) == calculated

	fmt.Println(a)

	return
}
