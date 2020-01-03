package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)
func main() {
	file, err := os.Open("./filename")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res, err := http.Post("http://127.0.0.1:5050/upload", "binary/octet-stream", file)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	message, _ := ioutil.ReadAll(res.Body)
	fmt.Printf(string(message))
}
