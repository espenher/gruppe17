package main

import (
	"fmt"
	"log"
	"os"
)

func main (){

	file := os.Args[1]

	fi, err := os.Lstat(file)
	if err != nil {
		log.Fatal(err)
	}
	info, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	mode := fi.Mode()
	fmt.Println ("Informasjon om fila",file,":")

	var bytes int64 = info.Size()
	var kb float64 = (float64)(bytes/1024)
	var mb float64 = (kb/1024)
	var gb float64 = (mb/1024)

fmt.Println("Størrelse:",bytes,"i bytes,",kb,"KB,",mb,"MB og",gb,"GB")

	if mode.IsDir () {
		fmt.Println("er directory")
	} else {
		fmt.Println("er ikke directory")
	}

	if mode.IsRegular () {
		fmt.Println("er regular file")
	} else {
		fmt.Println("er ikke regular file")
	}

	var unix = mode & os.ModePerm
	if unix == 0777 {
	fmt.Println("har Unix permission bits: -rwxrwxrwx")
	} else {
	fmt.Println("har ikke Unix permission bits: -rwxrwxrwx")
	}

	var isAppend = mode & os.ModeAppend
	if isAppend != 0 {
		fmt.Println("er append only")
	} else {
		fmt.Println("er ikke append only")
	}

	var device = mode & os.ModeDevice
	if device != 0 {
		fmt.Println("er device file")
	} else {
		fmt.Println("er ikke device file")
	}

	var unixChar = mode & os.ModeCharDevice
	if unixChar != 0 {
		fmt.Println("er Unix character device, when ModeDevice is set")
	} else {
		fmt.Println("er ikke Unix character device, when ModeDevice is set")
	}

	var symbolic = mode & os.ModeSymlink
	if symbolic != 0 {
		fmt.Println("er symbolic link")
	} else {
		fmt.Println("er ikke symbolic link")
	}
}
