// func
// https://ithelp.ithome.com.tw/articles/10294647

package main

import "fmt"

func foo1(a, b int) int {
	return a*3 + b*4
}

func main() {

	x, y := 1, 2
	fmt.Printf("(x, y) = (%d, %d)\n", x, y)

	fmt.Println("before declaring local foo1:\t3x + 4y =", foo1(x, y))
	foo1 := func(a, b int) int {
		return a*5 + b*7
	}
	fmt.Println("after declaring local foo1:\t5x + 7y =", foo1(x, y))

	// anonymous func
	fmt.Println()
	func() {
		fmt.Println("anonymous func without param")
	}()
	func(a, b int) {
		fmt.Println("anonymous func with param:", a + b + 100)
	}(x, y)

	// func as param
	fmt.Println()
	foo2 := func(myMethod func(int, int) int) int {
		return myMethod(x, y)
	}
	addFunc := func(a, b int) int {
		return a + b
	}
	subFunc := func(a, b int) int {
		return a - b
	}
	fmt.Println("addFunc as param:", foo2(addFunc))
	fmt.Println("subFunc as param:", foo2(subFunc))

	// name the return value
	fmt.Println()
	foo3 := func(a, b int) (myAns int) {
		myAns = a*1000 + b*100
		return
	}
	fmt.Println("name the return value:", foo3(x, y))
}