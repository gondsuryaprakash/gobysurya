package https

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

func HeaderFunc() {
	handler := http.HandlerFunc(handlerFunc)

	http.Handle("/book", handler)
	http.ListenAndServe(":8080", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {

	// Set the value in header
	w.Header().Set("content-type", "application/json")
	w.Header().Add("foo", "bar1")
	w.Header().Add("foo", "bar2")

	// r.Header.Values -- > Give the all values assosiated with key
	// func (h Header) Values(key string) []stringe
	fooValues := r.Header.Values("foo")
	fmt.Println("fooValues :", fooValues)

	// r.Header.Get(key) ---> Give the value corresponding to the Key
	// func (h Header) Get(key string) string
	contentType := r.Header.Get("content-type")
	fmt.Println("ContentType : ", contentType)

	resp := make(map[string]string)
	resp["message"] = "success"

	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)

}
