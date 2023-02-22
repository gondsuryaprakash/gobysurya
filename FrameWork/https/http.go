package https

import (
	"fmt"
	"net/http"
	"strings"
)

func CallGet() {
	getProductsHandler := http.HandlerFunc(getProducts)

	http.Handle("/products", getProductsHandler)
	http.ListenAndServe(":8080", nil)
}

func getProducts(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	val, ok := query["filters"]

	if !ok {
		fmt.Println("Filters are not present ")
	}
	w.WriteHeader(200)
	w.Write([]byte(strings.Join(val, ",")))

}
