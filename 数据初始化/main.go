package main

import (
	"encoding/json"
	"fmt"
)

type Node struct {
	Name   string
	Number int
	Tags   []string
}

func namingOutput() (output []string) {
	fmt.Println("output:", output)
	jsonString, _ := json.Marshal(output)
	fmt.Println("output json:", string(jsonString))

	return
}

func main() {
	fmt.Println("start")

	var node1 Node
	fmt.Println("node1:", node1)
	fmt.Println("len(node1.Tags):", len(node1.Tags))
	jsonString, _ := json.Marshal(node1)
	fmt.Println("jsonString:", string(jsonString))

	node1.Tags = append(node1.Tags, "test")
	fmt.Println("node1:", node1)
	jsonString, _ = json.Marshal(node1)
	fmt.Println("jsonString:", string(jsonString))

	namingOutput()
}
