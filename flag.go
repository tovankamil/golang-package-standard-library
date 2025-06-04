package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "password", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("Port", 0, "port database")
	flag.Parse()
	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)

}
