package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/globalsign/mgo"
)

var db = "go-test"
var globalS *mgo.Session

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost:27017"},
		Timeout:  time.Second * 10,
		Database: db,
	}

	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}

	globalS = s

	http.ListenAndServe(":3000", nil)
	fmt.Println("Server on Port 3000")

}
