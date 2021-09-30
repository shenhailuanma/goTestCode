package main

import (
	"fmt"
	"github.com/robert-nix/ansihtml"
)

func TestAnsi2Html(input string) string {
	fmt.Println("#TestAnsi2Html,input:", input)

	output := ansihtml.ConvertToHTML([]byte(input))

	fmt.Println("#TestAnsi2Html,output:", string(output))

	return string(output)
}