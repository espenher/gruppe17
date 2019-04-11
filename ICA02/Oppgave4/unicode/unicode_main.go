package main 

import "./unicode"
import "os"
import "fmt"

func main() {
	var language  = os.Args[1] 
	var text = "nord og sør"
	fmt.Printf("%s", unicode.Translate( text, language))

}