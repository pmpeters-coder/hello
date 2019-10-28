package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/inconshreveable/go-update"
)

func main() {
	for {
		time.Sleep(2 * time.Second)
		fmt.Println(2)
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
