package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var 64f float64
	// fmt.Println(64f, reflect.TypeOf(64f))

	totalPrice := 1000
	fmt.Println(totalPrice, reflect.TypeOf(totalPrice))
}
