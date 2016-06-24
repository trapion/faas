package main

import (
	"fmt"

	"docker.io/chanwit"
)

func main() {
	err := chanwit.Leftpad("foo", 5)
	if err != nil {
		fmt.Println(err)
	}
}
