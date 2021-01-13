package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "http://0.0.0.0:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	RootHandler(res, req)

	testStatus := res.Code

	if testStatus != 200 {
		t.Errorf("HTTP Reponse Code was %d", testStatus)
	}
}
