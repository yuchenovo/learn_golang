package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetName(t *testing.T) {
	response, err := http.Get("http://localhost:2345/getName?name=123")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	} else {
		defer response.Body.Close()
		bytes, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		} else {
			fmt.Println(string(bytes))
		}
	}
}
