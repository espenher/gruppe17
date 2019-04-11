package main

import "./lineshift"
import "os"

func main() {
	var file_name = os.Args[1]
	lineshift.FileToByteslice(file_name)
}