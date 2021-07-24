package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new billing name: ", reader)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price; ", reader)
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount{$): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("you chose s")
	default:
		fmt.Println("that was not a valid option...")
		promptOption(b)
	}

}

func main() {

	mybill := createBill()
	promptOption(mybill)

}
