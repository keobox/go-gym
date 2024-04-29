package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// encoding basic data types to json strings
	boolBinary, _ := json.Marshal(true)
	fmt.Println(string(boolBinary))

	intBinary, _ := json.Marshal(1)
	fmt.Println(string(intBinary))

	floatBinary, _ := json.Marshal(3.14)
	fmt.Println(string(floatBinary))

	sliceData := []string{"apple", "peach", "pear"}
	sliceBinary, _ := json.Marshal(sliceData)
	fmt.Println(string(sliceBinary))

	mapData := map[string]int{"apple": 5, "lettuce": 7}
	mapBinary, _ := json.Marshal(mapData)
	fmt.Println(string(mapBinary))

	// encoding custom data types
	res1Data := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1Binary, _ := json.Marshal(res1Data)
	fmt.Println(string(res1Binary))

	res2Data := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2Binary, _ := json.Marshal(res2Data)
	fmt.Println(string(res2Binary))

}
