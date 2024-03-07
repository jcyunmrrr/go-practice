// variable
// https://ithelp.ithome.com.tw/articles/10293211

// control flow
// https://ithelp.ithome.com.tw/articles/10297366

// data type
// https://ithelp.ithome.com.tw/articles/10293992


package main

import(
	"fmt"
	"strconv"
)

// multi var declaration
var(
	p0 string
	p1 string = "Ben"
)

func main() {

	p0 = "Andy"
	p1 = "Betty"

	// normal declaration
	var p2 string = "Cathy"

	// i is local, can only be used in func
	for i:=0; i<5; i++ {
		switch i {
		case 0:
			fmt.Printf("%s: ", p0)
		case 1:
			fmt.Printf("%s: ", p1)
		case 2:
			fmt.Printf("%s: ", p2)
		default:
			fmt.Printf("Someone: ")
		}

		if i%3==0 {
			fmt.Printf("%d is devided by 3\n", i)
		} else if i%2==0 {
			fmt.Print(strconv.Itoa(i) + " is not devided by 3, but it is even\n")
		} else {
			fmt.Println(i, "is not devided by 3. It is not even neither")
		}
	}

	// const
	fmt.Println()
	const(
		pi = 3.14159
		any_var
		e = 2.718281
		another_var
	)
	fmt.Printf(`\n
pi = %f
any_var = %f
e = %f
another_var = %f` + "\n", pi, any_var, e, another_var)

	var myArr []int
	for i:=0; i<1024; i++ {
		myArr = append(myArr, i)
	}
	fmt.Println(myArr[0:16])
	fmt.Println("len of myArr:", strconv.Itoa(len(myArr)))

	// complex
	fmt.Println()
	c1 := complex(22, 33)
	var c2 = 44 + 55i
	fmt.Println("c1:", c1)
	fmt.Printf("type of c1: %T\n", c1)
	fmt.Println("c2:", c2)
	fmt.Printf("type of c2: %T\n", c2)
}

// 20240307