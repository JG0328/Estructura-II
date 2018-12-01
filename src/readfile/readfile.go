package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("SCC.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bytes read: ", bytesread)
	//fmt.Println("bytestream to string: ", string(buffer))
	//words := strings.Fields(string(buffer))
	//fmt.Println(words[707083])
}
