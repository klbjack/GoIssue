package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("server1.example.com,server2.example.com")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Printf("hello, world\n")

	//test
}
