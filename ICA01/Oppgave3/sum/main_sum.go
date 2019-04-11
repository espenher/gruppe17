package main

import "fmt"
import "os"
import "strconv"
import "bufio"
import "./sum"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input1 string
	var input2 string
	




	fmt.Print("Input number 1: ")
	scanner.Scan()
	input1 = scanner.Text()
	fmt.Print("Input number 2: ")
	scanner.Scan()
	input2 = scanner.Text()
		
	//Parse text to int
	numb1, err := strconv.ParseInt(input1, 10, 64)
	if err != nil {
		fmt.Println("Wrong input,"  + input1 + " is not integer")
		 os.Exit(3)
	} 
	numb2, err := strconv.ParseInt(input2, 10, 64)
	if err != nil {
				fmt.Println("Wrong input,"  + input1 + " is not integer")
				 os.Exit(3)
	} 
	fmt.Print(	"Sum is  "	)
	fmt.Println( sum.SumInt64(numb1, numb2))
			
}
