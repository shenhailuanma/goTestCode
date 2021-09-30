package main

import (
	"fmt"
	"github.com/leaanthony/go-ansi-parser"
)


func TestAnsiParser(input string) {
	fmt.Println("#TestAnsiParser,input:", input)

	cleaner, _ := ansi.Cleanse(input)

	fmt.Println("#TestAnsiParser,output:", cleaner)
}