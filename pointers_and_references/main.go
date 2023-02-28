package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
	name := "Bill"
	updateValue(name)
	fmt.Println(name)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}

func updateValue(n string) {
	n = "Alex"
}
