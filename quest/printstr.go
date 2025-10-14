package main

import "fmt"

func PrintStr(s string) {
	for index, s := range s {
		fmt.Printf("the index is: %v and the element is %v\n", index, s)
	}
}

func main() {
	PrintStr("Hello World!")
}
