package connection

import (
	"bufio"
	"encoding/json"
	"fmt"
	"json"
	"net/http"
	"net/url"

	"github.com/TheBromo/goWebService/common"
)

type Connection struct {
	Addr *url.URL
}

func (c *Connection) Available() bool {
	resp, err := http.Get(c.Addr.String() + "/status")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return true
}

func (c *Connection) send(message common.Message) error {
	resp, err := http.Post(c.Addr.String(), "json",)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
