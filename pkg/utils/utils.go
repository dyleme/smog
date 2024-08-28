package utils

import (
	"encoding/json"
	"fmt"
)

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Print(a any) {
	bts, err := json.MarshalIndent(a, "", " ")
	NoErr(err)
	fmt.Println(string(bts))
}
