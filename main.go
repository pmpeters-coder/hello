package main

import (
	"fmt"
	"practice/go-update/depend"
	"time"
)

func main() {
	time.Sleep(2 * time.Second)
	depend.PrintNumber()
	err := doUpdate("https://github.com/pmpeters-coder/hello.git")
	if err != nil {
		fmt.Println(err)
	
}
