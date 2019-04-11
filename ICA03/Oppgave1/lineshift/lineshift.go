package lineshift



import (
"log"
	"io"
	"bytes"
	"fmt"
	"os"
)

func FileToByteslice(file_name string) []byte {

	// Open file https://golang.org/pkg/os/ 
	file, err := os.Open(file_name) // For read access.

		if err != nil {
			log.Fatal(err)
		}

	// check file 	
	file_info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size_of_slice := file_info.Size()

	//create slice for the file 
	byte_Slice := make([]byte, size_of_slice)

	//\x0A is the escaped hexadecimal Line Feed. The equivalent of \n.
	//\x0D is the escaped hexadecimal Carriage Return. The equivalent of \r.
	newLine := ([]byte("\x0A")) 
	newLine2 := ([]byte("\x0D\x0A"))

	_, err = io.ReadFull(file, byte_Slice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	//Check if file consist new lines
	if bytes.Contains(byte_Slice, newLine2) {
		fmt.Println("There is a carriage return /x0D in the file")	
		
	} else if bytes.Contains(byte_Slice, newLine) {
		fmt.Println("There is a line shift /x0A (Line Feed)  in the file ")
	} else {
		fmt.Println("There is no new line in the file.")
	}
	return byte_Slice
}