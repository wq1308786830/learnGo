package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	files, _ := ioutil.ReadDir("/Users/russell/GitHub/GoLang/learn/src/")
	for _, f := range files {
		fmt.Println(f.Name())
	}
	fmt.Scanf("%s")
}
