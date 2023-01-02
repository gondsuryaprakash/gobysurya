package file

import (
	"fmt"
	"log"
	"os"
)

// Use for Creating File
func CreateFile() {

	f, err := os.Create("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	var data = []string{"Surya", "Prakash", "Gond"}

	for _, val := range data {
		_, err := f.WriteString((val + "\n"))

		if err != nil {
			log.Fatal(err)
		}
	}

}
