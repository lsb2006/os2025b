package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n') // ignore error
	//fmt.Println(err)
	if err != nil {
		log.Fatal(err) // report the error, exit the program
	}
	fmt.Println(i)
}
