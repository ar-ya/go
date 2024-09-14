package main

// import an package
/* idfs
fsd */

import (
	"fmt"
)

func main() {
	// VARIABLE
	var student1 string = "arya"
	var student2 = "shankar"
	x := 1
	var a, b int = 2, 3
	var c, d = 6, "Hello"
	e, f := 7, "World!"
	var (
		g int
		h int    = 1
		i string = "Hello"
	)
	const PI = 3.14

	var aa float32 = 2.4

	// fmt.Println("Hello World!")
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)

	fmt.Print(d)
	fmt.Print(f)

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(aa)

	fmt.Printf("i has vaule %v and type %T\n", i, i)

	// ARRAY
	arr1 := [5]int{}              //not initialized
	arr2 := [5]int{1, 2}          //partially initialized
	arr3 := [5]int{1, 2, 3, 4, 5} //fully initialized
	arr4 := [...]int{1, 2}
	arr5 := [...]int{0: 1, 3: 40}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

}
