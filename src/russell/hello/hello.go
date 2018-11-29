/*
官网教程
*/
package main

import (
	"fmt"
	"russell/stringutil"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("\n%d ", i)
		fmt.Print(stringutil.Reverse("Hello, World!\n"))
	}
}
