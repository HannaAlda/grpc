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
	in := `{"productName":"CyberTruck","productPrice": 39900}`

	bytes, err := json.Marshal(in)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
