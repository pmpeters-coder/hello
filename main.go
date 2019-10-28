package main

import (
	"fmt"
	"net/http"
	"practice/go-update/depend"
	"time"

	"github.com/inconshreveable/go-update"
)

func main() {
	time.Sleep(2 * time.Second)
	depend.PrintNumber()
	doUpdate("https://github.com/pmpeters-coder/hello.git:go-update")

}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		fmt.Println(err)
	}
	return err
}
