package main

import "fmt"

const MAX int = 3

func main() {
	var a int = 20 /* actual variable declaration */
	var ip *int    /* pointer variable declaration */

	ip = &a /* store address of a in a pointer variable */

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in point variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	/* access the value using the pointer */
	fmt.Printf("Value of *ip variable %d\n", *ip)

	hola()
	var m int = 100
	var n int = 200
	fmt.Printf("Before swap, value of m, n: %d, %d\n", m, n)
	swap(&m, &n)

	fmt.Printf("After swap, value of m, n: %d, %d\n", m, n)
}

func hola() {
	a := []int{10, 100, 200}
	var i int
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* assign the address of integer */
		fmt.Printf(" ptr is: %x", ptr[i])
	}
	fmt.Println()

	for i = 0; i < MAX; i++ {
		fmt.Printf("Valor of a[%d] = %d\n", i, *ptr[i])
	}
}

func swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
