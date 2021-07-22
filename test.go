package main

import "fmt"

func updateName(x *string) {
	*x = "wedge"
}

func main() {

	name := "tifa"

	//updateName(name)

	//fmt.Println("memory address of name is:", &name)

	m := &name
	// fmt.Println("memory address (value of m):", m)
	// fmt.Println("value of memory address:", *m)

	fmt.Println(name)
	updateName(m)
	fmt.Println(name)

}
