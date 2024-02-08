package main

import "fmt"

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func main() {
	a := new(int)
	fmt.Println(&a) // 0xAddressA
	fmt.Println(a)  // 0xPointerAddressA
	fmt.Println(*a) // 0

	*a = 5
	fmt.Println(&a) // 0xAddressA
	fmt.Println(a)  // 0xPointerAddressA
	fmt.Println(*a) // 5

	b := 10
	fmt.Println(&b) // 0xAddressB
	fmt.Println(b)  // 10

	*a = b
	fmt.Println(&a) // 0xAddressA
	fmt.Println(a)  // 0xPointerAddressA
	fmt.Println(*a) // 10

	b = 500
	a = &b
	fmt.Println(&a) // 0xAddressA
	fmt.Println(a)  // 0xAddressB
	fmt.Println(*a) // 500

	x := 5
	y := 10
	fmt.Println(x, y) // 5, 10
	swap(&x, &y)
	fmt.Println(x, y) // 10, 5
}
