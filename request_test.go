package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoRequest(t *testing.T) {
	request := Request{
		Url:    "https://www.baidu.com",
		Method: "GET",
	}
	res := request.doRequest()
	assert.Equal(t, 200, res.StatusCode(), "the status code should be 200")
}
