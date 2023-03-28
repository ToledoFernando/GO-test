package main

import (
	"fmt"

	"github.com/ToledoFernando/GO-test/extras"
	"github.com/google/uuid"
)

func main() {
	println(uuid.New().String())
	var xd = extras.Extras()
	fmt.Println(xd)
}