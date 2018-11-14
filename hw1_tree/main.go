package main

import (
	"fmt"
	"os"
)

func dirTree(out *os.File, s string, b bool) error {

	//fmt.Println(out)
	//fmt.Println(s)
	//fmt.Println(b)
	return nil
}

func main() {
	//fmt.Println("Message to stdout")
	out := os.Stdout
	//fmt.Println(len(os.Args))
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	fmt.Println(path)
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
