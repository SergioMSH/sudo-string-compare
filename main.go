package main

import "fmt"

func compareStrings(str1 string, str2 string) bool {
	if str1 == str2 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(compareStrings("hello", "hello")) // true
	fmt.Println(compareStrings("hello", "world")) // false
}
