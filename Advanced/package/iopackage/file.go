package iopackage

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cast"
)

func Write() {
	file, err := os.Create("file.text")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := io.Writer(file)
	n, err := writer.Write([]byte("Hello Surya"))

	if err != nil {
		fmt.Println(n, err)
	}

	n, err = io.WriteString(writer, "!")
	fmt.Println(n, err)

	for i := 0; i < 10; i++ {
		n, err := io.WriteString(writer, cast.ToString(i)+"\n")
		if err != nil {
			fmt.Println(n, err)
		}
	}

}

func Read() {

	file, _ := os.Open("file.text")

	defer file.Close()

	reader := io.Reader(file)
	buffer := make([]byte, 100)
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Println(n, err)
	}

	newBuffer, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ReadAll :", string(newBuffer))

	fmt.Printf("Read n={%v}, err = {%v} buffer = {%v}", n, err, string(buffer))
}
