package main

import (
	"fmt"
	"github.com/cartmanis/go_get_struct/generator"
	"github.com/cartmanis/go_get_struct/node"
	"os"
)

func main() {
	file, err := os.Open("test.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	n, err := node.Parse(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileGo, err := generator.CreateStruct(n)
	fmt.Println(fileGo)
}
