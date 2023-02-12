package http

import (
	"bufio"
	"fmt"
	"net/http"
)

func CallGet() {
	res, err := http.Get("https://gobyexample.com")

	if err != nil {
		fmt.Println("Error in Fetching Data")
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
}
