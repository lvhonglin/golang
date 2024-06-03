package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_client(t *testing.T) {
	resp, err := http.Post("http://localhost:8080/ping", "", nil)
	if err != nil {
		t.Error(err)
		return
	}
	body, _ := ioutil.ReadAll(resp.Body)
	t.Logf("%s", body)
	defer resp.Body.Close()
	t.Error("test")
}
