package http

import (
	"fmt"
	"net/http"
	"sync"
)

func CallCr


func Crawl(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	resp, err:= http.Get(url)

	if err != nil {
		fmt.Println(err)
	} 


}