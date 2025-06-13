package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	// Avoid parse function for doing parsing
	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}

/*
$ go run 63_flag.go -username=irwan -password="rahasia banget" -host=123.231.23.1 -port=5505

[ Package flag ]
- Package flag berisikan fungsionalitas untuk memparsing command line argument
- https://golang.org/pkg/flag/
*/
