package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoRequest(t *testing.T) {
	url := "https://www.baidu.com"
	method := "GET"
	request := Request{
		Url:    &url,
		Method: &method,
	}
	res := request.doRequest()
	assert.Equal(t, 200, res.StatusCode(), "the status code should be 200")
}
