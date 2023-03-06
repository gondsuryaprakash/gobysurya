package https

import (
	"encoding/json"
	"net/http"
)

func JsonRequestBody() {

	handlerFunc := http.HandlerFunc(JSONHandlerFunc)

	http.Handle("/", handlerFunc)
	http.ListenAndServe(":8080", nil)

}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JSONHandlerFunc(w http.ResponseWriter, r *http.Request) {
	var student *Student
	var unMarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&student)

	if err != nil {
		errorResponse(w, "Bad Request. Wrong Type provided for field "+unMarshalErr.Field, http.StatusBadRequest)
	} else {
		errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
	}
	errorResponse(w, "Success", http.StatusOK)
	return

}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
