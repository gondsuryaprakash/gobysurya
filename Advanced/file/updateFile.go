package file

import (
	"errors"
	"log"
	"os"
)

func UpdateFile() {
	_, err := os.Stat("data.txt")

	if errors.Is(err, os.ErrNotExist) {
		log.Fatal("File Doesn't exist: ")
	} else {
		f, err := os.OpenFile("data.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}

	}

}
