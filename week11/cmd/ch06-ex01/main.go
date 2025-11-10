package main

import "fmt"

func main() {
	var sujects []string
	sujects = make([]string, 3)
	sujects[0] = "Go"
	sujects[2] = "Python"

	for _, subject := range sujects {
		fmt.Println(subject)
	}
}
