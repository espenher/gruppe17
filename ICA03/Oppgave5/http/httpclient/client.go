package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(getURL("http://localhost:5010"))
	fmt.Println(getURL("http://localhost:5010/user"))

	// Get everyone's response
	//fmt.Println(getURL("http://192.168.2.2:5010/user"))
	//fmt.Println(getURL("http://192.168.2.3:5010/user"))
	//fmt.Println(getURL("http://192.168.2.4:5010/user"))
	//fmt.Println(getURL("http://192.168.2.5:5010/user"))
}

func getURL(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	return string(body)
}
