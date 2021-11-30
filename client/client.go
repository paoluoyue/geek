package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var h HelloService = hello{
		endpoint: "http://localhost:8080/",
	}
	msg, _ := h.SayHello("Go")
	fmt.Print(msg)
}

type HelloService interface {
	SayHello(name string) (string, error)
}

type hello struct {
	endpoint string
}

func (h hello) SayHello(name string) (string, error) {
	client := http.Client{}
	resp, err := client.Get(fmt.Sprintf(h.endpoint + name))
	if err != nil {
		fmt.Printf("%+v", err)
		return "", nil
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return "", nil
	}
	return string(data), nil
}
