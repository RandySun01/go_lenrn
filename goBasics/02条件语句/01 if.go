package main

import "fmt"

func ifDemo2() {
	score := 10
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	fmt.Println(score)
}
func main() {
	ifDemo2()
}
