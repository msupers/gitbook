package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieve(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

func main() {

	fmt.Printf("%s\n", retrieve("https://www.imooc.com"))
}
