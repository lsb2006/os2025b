package main

import "fmt"

func main() {
	// var sujects []string
	// sujects = make([]string, 3)

	// subjects := make([]string, 3)

	// sujects[0] = "Go"
	// sujects[2] = "Python"

	subjects := [4]string{"Go", "Javascript", "Python", "Linux"} // slice literal
	subjectsSlice := subjects[:3]                                // slicing
	subjects[0] = "Java"
	subjectsSlice = append(subjectsSlice, "Go")
	for _, subject := range subjects {
		fmt.Println(subject)
	}
	fmt.Println("===================")
	for i := 0; i < len(subjectsSlice); i++ {
		fmt.Println(subjectsSlice[i])
	}
}
