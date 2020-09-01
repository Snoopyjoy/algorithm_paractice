package main

import (
	"encoding/json"
	"fmt"
)

type a struct {
	V int
	P *a
}

func main() {
	b := a{
		V: 1,
	}
	b.P = &b
	r, err := json.Marshal(b)
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("%s", string(r))
}
