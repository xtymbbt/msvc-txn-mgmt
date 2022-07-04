package main

import (
	"algorithm-test/test"
	"fmt"
)

func main() {
	//chain(10)
	//tree(1, 3)
	//fourTree(1)
	//test.LoadBalanceTest(5, 100000)
	//test.Slicetest()
	//fmt.Println(common.RandomFixedStr("123456789.", 15))
	x := test.MathTest([]int{40288, 10760, 28092, 15484, 5376})
	fmt.Printf("%#v\n", x)
}
