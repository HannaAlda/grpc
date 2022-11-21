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

	bytes, err := json.Marshal(Tesla{
		ProductName:  "CyberTruck",
		ProductPrice: "39900",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
