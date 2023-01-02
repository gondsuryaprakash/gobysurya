package json

import (
	"encoding/json"
	"fmt"
	"log"
)

func ConvertIntoJson() {

	var user = make(map[string]interface{})

	user["name"] = "Surya"
	user["age"] = 28
	user["contact"] = map[string]interface{}{
		"mobile": 7068528089,
		"email":  "suryaprakashgond@gmail.com",
	}

	byteData, err := json.Marshal(&user)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("user data : ", string(byteData))

}

func MarshalSlice() {

	slice := []string{"apple", "banana", "gauava"}

	byteData, _ := json.Marshal(slice)

	fmt.Println(string(byteData))
}

func MarshalMap() {

	mapD := map[string]int{"Surya": 23, "Ajay": 25}

	byteMap, _ := json.Marshal(mapD)
	fmt.Println(string(byteMap))
}

type Response1 struct {
	Page  string
	Fruit []string
}

type Response2 struct {
	Page   string   `json:"page"`
	Fruits []string `json:"fruits"`
}

func MarshalStruct() {

	resp1 := &Response1{
		Page:  "First",
		Fruit: []string{"Banana", "Orange", "Mango"},
	}

	resp2 := &Response2{
		Page:   "Second",
		Fruits: []string{"Banana", "Orange", "Mango"},
	}

	ByteMap1, _ := json.Marshal(resp1)
	ByteMap2, _ := json.Marshal(resp2)

	fmt.Println(string(ByteMap1))
	fmt.Println(string(ByteMap2))

}
