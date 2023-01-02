package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadFile() {

	f, err := os.Open("data.txt")

	if err != nil {
		log.Fatal("err: ", err)
	}

	defer f.Close()

	scan1 := bufio.NewScanner(f)

	for scan1.Scan() {
		fmt.Println(scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		fmt.Println(err)
	}
}
