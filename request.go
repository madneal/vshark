package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Request struct {
	Host    string
	Path    string
	Port    int
	Url     string
	Method  string
	Headers map[string]string
	Body    string
}

type Response struct {
	StatusCode int
	Body       string
}

type Headers struct {
	Host        string
	Origin      string
	UserAgent   string
	ContentType string
}

func (request *Request) doRequest() (res *resty.Response) {
	client := resty.New()
	r := client.R()
	if request.Headers != nil {
		r.SetHeaders(request.Headers)
	}
	var err error
	var response *resty.Response
	if request.Method == "GET" {
		response, err = r.Get(request.Url)
	} else if request.Method == "POST" {
		response, err = r.SetBody(request.Body).Post(request.Url)
	}
	if err != nil {
		fmt.Println(err)
	}
	return response
}
