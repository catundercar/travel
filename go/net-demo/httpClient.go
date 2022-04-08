package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	clt := http.Client{
		Transport: &roundTripper{},
	}
	resp, err := clt.Get("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}

type roundTripper struct {

}

func (*roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Method != "POST" {
		return nil, fmt.Errorf("only support post")
	}

	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}