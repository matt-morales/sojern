package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	ping(resp, req)
	want := http.StatusOK
	got := resp.Code

	if want != resp.Code {
		t.Errorf("\nwant: %d \ngot %d\n", want, got)
	}
}

func TestPingFail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/yolo", nil)
	resp := httptest.NewRecorder()
	ping(resp, req)
	want := http.StatusNotFound
	got := resp.Code

	if want != resp.Code {
		t.Errorf("\nwant: %d \ngot %d\n", want, got)
	}
}

func TestImgSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/img", nil)
	resp := httptest.NewRecorder()
	img(resp, req)
	want := http.StatusOK
	got := resp.Code

	if want != resp.Code {
		t.Errorf("want: want%d. Got %d\n", want, got)
	}
}
