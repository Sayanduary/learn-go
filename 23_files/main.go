package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("example.txt")
	if err != nil {
		//log the error
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		//log the error
		panic(err)
	}

	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())

	// read file

	file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 12)

	data, err := f.Read(buffer)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buffer); i++ {
		fmt.Println(string(buffer[i]), data)
	}

	// simple way -> for small size file

	new_file, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(new_file))

	// read folders

	new_folder, err := os.Open("../")
	if err != nil {
		panic(err)
	}
	defer new_folder.Close()
	file_info, err := new_folder.ReadDir(-1)

	for _, fi := range file_info {
		fmt.Println(fi.Name(), fi.IsDir())
	}
}
