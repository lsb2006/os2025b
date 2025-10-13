package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// shadowing
	// var int int = 99
	// var b int = 8
	// var fmt string = "inha"
	// fmt.Println(int, b)
	// fmt.Println()

	fmt.Print("Enter a score: ")
	r := bufio.NewReader(os.Stdin)
	i, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	i = strings.TrimSpace(i)                // 문자열 주위에 붙은 공간 및 탭 키 등 제거
	score, err := strconv.ParseFloat(i, 64) // 정리된 문자열을 실수 타입으로 변환
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if score >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println(score, status)
}
