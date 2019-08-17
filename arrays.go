package main

import "fmt"

func main() {

	var a [5] int

	a[3]=99
	a[4] = 100

	fmt.Println("a:", a)
	fmt.Println("a[4]:", a[4])
}
