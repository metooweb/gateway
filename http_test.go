package main

import (
	"testing"
	"github.com/go-resty/resty"
	"fmt"
)

func TestHttp(t *testing.T) {

	r, err := resty.R().Get("http://127.0.0.2/test.php")

	fmt.Println(err)
	fmt.Println(r.StatusCode())

}
