package main

import ("./slice"
		"fmt")

func main () {
	var b int = 1000


	fmt.Println("AllocateVar")
	var bytesliceVar []byte = slice.AllocateVar(b)
	 fmt.Println("&bytesliceVar[0]")
	fmt.Println("AllocateMake")
	 fmt.Println(&bytesliceVar[0])

	var byteslice1 []byte = slice.AllocateMake(b)
	fmt.Println()

	fmt.Println()
	fmt.Println("For Reslice")
	 fmt.Println("&byteslice1[0]")
	 fmt.Println(&byteslice1[0])
	 fmt.Println("After Reslice")
	aslice := slice.Reslice(byteslice1, 50, 100)
	 fmt.Println("&byteslice[50]")
	fmt.Println(&aslice[50])
	fmt.Println("&byteslice1[50]")
	 fmt.Println(&byteslice1[50])


}