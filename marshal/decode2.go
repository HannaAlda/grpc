// hello.go

package main

import (
	"encoding/json"
	"fmt"
)

// Tesla struct
type Tesla struct {
	ProductName  string `json:"productName"`
	ProductPrice string `json:"productPrice"`
}

func main() {

	data := `{"productName":"CyberTruck","productPrice":"39900"}`
	bytes := []byte(data)
	var t Tesla
	err := json.Unmarshal(bytes, &t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v \n", t)
	fmt.Println(string(bytes))
}
