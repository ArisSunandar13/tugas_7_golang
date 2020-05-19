package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	number := 13
	word := "hello"

	go printNumber(number)
	printWord(word)

	var input string
	fmt.Scan(&input)
}

func printNumber(x int) {
	reflectNumber := reflect.ValueOf(x)

	fmt.Println(reflectNumber.Type())
}

func printWord(x string) {
	reflectWord := reflect.ValueOf(x)

	fmt.Println(reflectWord.Type())
}
