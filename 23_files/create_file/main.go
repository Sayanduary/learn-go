package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Hi Go")
	file.WriteString("Nice Language")

	bytes := []byte("hello")
	file.Write(bytes)

	// Copy Data from one file to another file(read and write streaming fashion)
	source_file, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer source_file.Close()

	dest_file, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}
	defer dest_file.Close()

	// streaming (bufio)

	reader := bufio.NewReader(source_file)

	writer := bufio.NewWriter(dest_file)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}
			break
		}
		err1 := writer.WriteByte(b)
		if err1 != nil {
			panic(err1)
		}
	}
	writer.Flush()
	fmt.Println("Written to new file successfully")

	//delete a file

}
