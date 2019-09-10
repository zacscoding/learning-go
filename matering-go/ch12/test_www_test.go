package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// $ go test test_www.go test_www_test.go -v
func TestCheckStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/CheckStatusOK", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CheckStatusOK)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusOK {
		t.Errorf("handler returned %v", status)
	}

	expect := `Fine!`
	if rr.Body.String() != expect {
		t.Errorf("handler returned %v", rr.Body.String())
	}
}

func TestStatusNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/StatusNotFound", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusNotFound)
	handler.ServeHTTP(rr, req)

	status := rr.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned %v", status)
	}
}
