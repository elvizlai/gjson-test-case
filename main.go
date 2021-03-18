package main

import (
	"encoding/json"
	"fmt"

	gjson "github.com/goccy/go-json"
)

type Data struct {
	A []*Id
}

func main() {
	x := Data{
		A: []*Id{{Id: "1"}, {Id: "2"}},
	}

	// this works
	fmt.Println(json.Marshal(x))

	// this panics
	fmt.Println(gjson.Marshal(x))
}
