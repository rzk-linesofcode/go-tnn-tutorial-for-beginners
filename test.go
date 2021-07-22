package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffee puding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phoneBook := map[int]string{
		12345678: "mario",
		12345679: "luigi",
		12345670: "peach",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[12345670])

	phoneBook[12345679] = "bowser"
	fmt.Println(phoneBook)

	phoneBook[12345670] = "yoshi"
	fmt.Println(phoneBook)

}
