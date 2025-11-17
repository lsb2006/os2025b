package main

import "fmt"

type subcriber struct {
	name  string
	price int
}

func applyPrice(s *subcriber) {
	s.price = 10000
	s.name = "Park Inha"
}

func main() {
	var s1 subcriber
	var p *subcriber = &s1
	// s1.name = "Kim Inha"
	applyPrice(&s1)
	fmt.Println(s1.name, s1.price)
	// fmt.Println(*p.price)
	fmt.Println((*p).price)
	fmt.Println(p.price)
}
