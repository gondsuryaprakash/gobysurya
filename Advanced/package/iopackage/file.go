package iopackage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Write() {
	file, err := os.Create("surya.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	writer := io.Writer(file)
	for i := 0; i < 10; i++ {
		m := make(map[string]string)
		m["id"] = string(i)
		m["name"] = string(i) + "_surya"
		mByte, err := json.Marshal(m)
		if err != nil {
			fmt.Println(err)
		}
		n, err := writer.Write(mByte)
		if err != nil {
			fmt.Println(n, err)
		}
	}

}

func WriteJson() {

	for i := 0; i < 10; i++ {

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
