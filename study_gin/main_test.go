package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestGetName(t *testing.T) {
	//response, err := http.Get("http://localhost:3000/getName?name=柳岩")
	response, err := http.PostForm("http://localhost:3000/update", url.Values{"name": []string{"柳岩"}})
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
