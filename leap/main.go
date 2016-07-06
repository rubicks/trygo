package main

import (
	"fmt"
	"log"
)

func isleap(n int) bool {
	if 0 == n%400 {
		return true
	}
	if 0 == n%100 {
		return false
	}
	if 0 == n%4 {
		return true
	}
	return false
}

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if nil != err {
		log.Fatal(err)
	}
	if isleap(n) {
		fmt.Printf("%v is a leap year.\n", n)
	} else {
		fmt.Printf("%v is not a leap year.\n", n)
	}
}
