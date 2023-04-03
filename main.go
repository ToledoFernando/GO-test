package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", nil)
	fmt.Println("Server on Port 3000")

}
