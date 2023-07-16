package main

import (
	"fmt"
	"net/url"

	conn "github.com/TheBromo/goWebService/client/connection"
)

func main() {
	u, err := url.Parse("http://localhost:8090")

	if err != nil {
		fmt.Println("Url could not be parsed")
		panic(err)
	}

	connection := conn.Connection{
		Addr: u,
	}

	for connection.Available() {
		//TODO ask for input
		// display answer
	}

}
