package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type MyJsonStruct struct {
	Test any `json:"test"`
}

func main() {
	testJsonInput := `{"test":{"test1":[1,2,3]}}`
	var jsonParsed MyJsonStruct
	err := json.Unmarshal([]byte(testJsonInput), &jsonParsed)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type of json parsed : %s \n", reflect.TypeOf(jsonParsed.Test))

	switch v := jsonParsed.Test.(type) {
	case map[string]any:
		fmt.Printf("Map Found : %v", v)
		field1, ok := v["test1"]
		if ok {
			fmt.Println("test1 found")
		}
		switch value2 := field1.(type) {
		case []any:
			for _, element := range value2 {
				fmt.Printf("Type of value2 : %v and value :%v \n", reflect.TypeOf(element), element)

			}

		default:
			fmt.Printf("Unexpected output found : %v", reflect.TypeOf(value2))
		}
	default:
		fmt.Printf("Unexpected output found : %v", v)
	}

}
