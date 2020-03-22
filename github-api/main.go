package main

import (
	"flag"
	"fmt"
	"log"
)

var loginField = flag.String("login_field", "username", "Username or email address")
var password = flag.String("password", "1234567", "Password")

// go run *.go -login_field= -password=
// go run *.go -login_field=username -password=1234567

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	flag.Parse()

	err := Login(*loginField, *password)
	if err != nil {
		fmt.Println(err)
	}
}
