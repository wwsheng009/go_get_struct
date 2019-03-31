package main

import (
	"flag"
	"fmt"
	"github.com/cartmanis/go_get_struct/generator"
	"github.com/cartmanis/go_get_struct/node"
	"os"
)

func main() {
	flag.Parse()
	for _, pathFile := range flag.Args() {
		fmt.Println(pathFile)
		file, err := os.Open(pathFile)
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

}
