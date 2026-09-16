package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Printf("H$LLO!")

	resp, err := http.Get("https://methum.me")

	if err == nil {

		print(err)
	}

	fmt.Println(resp.Body)

}
