package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		price, _ := getInput("item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOption(b)
			break
		}
		b.addItem(name, p)
		fmt.Println("item added -", name, price)
		promptOption(b)

	case "t":
		tip, _ := getInput("Enter tip amount{$): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOption(b)
			break
		}
		b.updateTip(t)
		fmt.Println("tip added -", tip)
		promptOption(b)

	case "s":
		b.save()
		fmt.Println("you saved the file", b.name)

	default:
		fmt.Println("that was not a valid option...")
		promptOption(b)
	}

}

func main() {

	mybill := createBill()
	promptOption(mybill)

}
