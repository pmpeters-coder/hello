package main

import (
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
)

func main() {
	for {
		doUpdate("https://github.com/pmpeters-coder/hello.git")
	}
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
