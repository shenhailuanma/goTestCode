package main

import "fmt"

func namingInt(input int) (output int) {
	output = input
	return
}

func namingArray() (output []string) {
	var output2 []string
	fmt.Println(output2)
	return
}

func main() {
	fmt.Println(namingInt(1))
	fmt.Println(namingArray())
}
