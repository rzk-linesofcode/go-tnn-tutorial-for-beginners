package main

import "fmt"

func main() {

	//var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 15
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slices ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
