package main

import (
	"fmt"

	"./fileutils"
)

func main() {
	filename := "./files/lang03.wl"

	var f = fileutils.FileToByteslice(filename)
	fmt.Printf("% X", f)
}
