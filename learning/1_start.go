package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func Division(a int, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("The value of the denominator can not be 0")
	}
	return float64(a) / float64(b), nil
}

func main() {
	fmt.Print("Hello Go World!!")

	fmt.Println("===== Variable Learning =====")

	name := "Raj"

	var age int = 30
	var score float64 = 1.2
	var active bool = true
	var text string

	text = "Hello Go!!"

	fmt.Println("My name is ", name)
	fmt.Println("I am", age, "years old")
	if active {
		fmt.Println("The score is ", score)
		fmt.Println(text)
	}

	fmt.Println("The add ", add(1, 2))

	result, err := Division(1, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("The division value ", result)
	}

}
