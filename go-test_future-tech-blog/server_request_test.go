package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func TestHelloHandler(t *testing.T) {
	r := httptest.NewRequest("GET", "/dummy", nil)
	w := httptest.NewRecorder()

	helloHandler(w, r)

	resp := w.Result()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}

	if string(body) != "hello world!" {
		t.Errorf("got = %s, want = hello world!", body)
	}
}
