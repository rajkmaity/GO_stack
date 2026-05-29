//go:build !skip
// +build !skip

package main

import "fmt"

func loops() {
	for i := 1; i < 5; i++ {
		fmt.Println("The iteration ", i)
	}
}

func while() {
	i := 0
	for i < 5 {
		fmt.Println("The iteration ", i)
		i++
	}
}

func over_slice() {
	age_list := []int{10, 20, 30}
	for _, age := range age_list {
		fmt.Println(age)
	}

}

func main() {
	loops()
	while()
	over_slice()
}
