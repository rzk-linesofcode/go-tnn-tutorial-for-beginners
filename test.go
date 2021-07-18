package main

import "fmt"

func main() {

	name := "shaun"
	age := 35

	// print
	// fmt.Print("hello, ")
	// fmt.Print("world! \n")
	// fmt.Print("new line \n")

	// println
	//fmt.Println("Hello Ninjas!")
	//fmt.Println("Goodbye Ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formating string) %_ format specifier
	// fmt.Printf("my age is %v and my name is %v \n", age, name)
	// fmt.Printf("my age is %q and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T \n", age)
	// fmt.Printf("you scored %0.2f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the save string is:", str)

}
