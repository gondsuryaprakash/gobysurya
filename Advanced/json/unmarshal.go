package json

import (
	"encoding/json"
	"fmt"
	"os"
)

// UnMarshal is used for convert json into different data type

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UnMarshalJson() {
	var person Person
	byteData, err := os.ReadFile("Advanced/json/employee.json")
	if err != nil {
		fmt.Println("Error in reading file :", err)
	}

	err = json.Unmarshal(byteData, &person)
	if err != nil {
		fmt.Println("Error in UnMarshaling :", err)
	}

	fmt.Println("Struct After UnMarshal :", person)
}

type Person1 struct {
	Name string `json:"name"`
	Class string `json:"class"`
}

func CallUnMarshal() {
	
}

