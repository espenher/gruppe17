package main

import (
	"fmt"

	"./googlecloud"
	"./watson"
	"./wit"
)

func main() {
	file := "./sound-files/this-is-a-test"

	googleCloudResult := googlecloud.Convert(file)
	fmt.Println("Google Cloud: " + googleCloudResult)

	witResult := wit.Convert(file)
	fmt.Println("Wit: " + witResult)

	watsonResult := watson.Convert(file)
	fmt.Println("Watson: " + watsonResult)
}
