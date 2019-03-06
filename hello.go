package main

import (
	"fmt"

	"github.com/high0815/stringutil"
)

func main() {
	fmt.Println("Hello, World!")

	var x int
	x = 100

	var y = 200

	z := 200

	fmt.Printf("x ist: %d, y ist: %d, z ist: %d\n", x, y, z)

	var ptx *int

	ptx = &x
	fmt.Printf("ptx ist: %p, *ptx ist: %d\n", ptx, *ptx)

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))

	var a [5]int
	a[0] = 1

	fmt.Println("Array ist: ", a[0])

	for i := 0; i < 5; i++ {
		fmt.Printf("Array[%d] = %d\n", i, a[i])
	}

}
