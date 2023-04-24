package https

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

func CallGetRequest() {

	res, err := http.Get("https://fakestoreapi.com/products")

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	respByte, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(respByte))

}

func CallPostRequest() {

	reqBody, err := json.Marshal(map[string]string{
		"title":       "test product",
		"price":       "13.5",
		"description": "lorem ipsum set",
		"image":       "https://i.pravatar.cc",
		"category":    "electronic",
	})

	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post("https://fakestoreapi.com/products", "application/json", bytes.NewBuffer(reqBody))

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(respByte))
}

func PostingForm() {

	formData := url.Values{
		"name": {"Surya"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	log.Println(res["form"])

}

func CustomRequest() {

	reqBody, err := json.Marshal(map[string]string{
		"name":  "Surya Prakash Gond",
		"email": "suryaprakashgond@gmail.com",
	})
	if err != nil {
		log.Fatalln(err)
	}
	timeOut := time.Duration(time.Second * 5)
	client := http.Client{
		Timeout: timeOut,
	}
	req, err := http.NewRequest("POST", "https://httpbin.org/post", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))

}
