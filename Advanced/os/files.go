package main

import (
	"bufio"
	"fmt"
	"os"
)

// Create the file
func CreateFile() {
	file, err := os.Create("surya.txt")

	if err != nil {
		fmt.Println("Err in Creating file: ", err)
	}

	file.Close()
}

// Open the file
func OpenFile() {
	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("Err in Creating file : ", err)
	}

	file.Close()
}

// Direct readfile
func ReadFile() {

	data, err := os.ReadFile("file.txt")
	if err != nil {
		fmt.Println("Err in data : ", err)
	}

	fmt.Println(string(data))
}

// Read File ChunkWise

func ReadFileChunkWise() {

	chunkSize := 10 // 10 bytes per read time
	b := make([]byte, chunkSize)

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Err in opening files: ", err)
	}

	for {
		bytesRead, _ := file.Read(b)

		if bytesRead == 0 {
			break
		}

		fmt.Println(bytesRead)
		fmt.Println(string(b))

	}

	file.Close()

}

func ReadFileAsyncrosunly() {

	file, err := os.Open("file.txt")

	if err != nil {
		fmt.Println("Err in file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Err while reading : ", err)
	}

}

func main() {
	ReadFile()
	// ReadFileAsyncrosunly()
	ReadFileChunkWise()
}
