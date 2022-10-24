package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type testFile struct {
	path string
	err  error
}

func (f testFile) ReadFile() ([]byte, error)  { return []byte("success"), nil }
func (f testFile) Stat() (os.FileInfo, error) { return nil, f.err }

type testFileInfo struct {
	name string
}

func TestPingSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	pingHandler(resp, req)
	want := http.StatusOK
	got := resp.Code

	if want != resp.Code {
		t.Errorf("\nwant: %d \ngot %d\n", want, got)
	}
}

func TestPingFail(t *testing.T) {
	req, _ := http.NewRequest("GET", "/yolo", nil)
	resp := httptest.NewRecorder()
	pingHandler(resp, req)
	want := http.StatusNotFound
	got := resp.Code

	if want != resp.Code {
		t.Errorf("\nwant: %d \ngot %d\n", want, got)
	}
}

func TestImgSuccess(t *testing.T) {
	tests := []struct {
		name       string
		f          FileReader
		wantData   []byte
		wantStatus int
		wantError  error
	}{
		{
			name:       "happy path",
			f:          &testFile{"happy", nil},
			wantData:   []byte("success"),
			wantStatus: http.StatusOK,
			wantError:  nil,
		},
		{
			name:       "sad path",
			f:          &testFile{"sad", errors.New("file does not exist")},
			wantData:   []byte("503 service unavailable"),
			wantStatus: http.StatusServiceUnavailable,
			wantError:  nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/img", nil)
			w := httptest.NewRecorder()
			imgHandler(w, req, test.f)
			wantStatus := test.wantStatus
			gotStatus := w.Code

			if wantStatus != gotStatus {
				t.Errorf("\nwant status: %d \ngot status %d\n", wantStatus, gotStatus)
			}
			wantData := test.wantData
			gotData := w.Body.Bytes()

			if len(wantData) != len(gotData) {
				t.Errorf("\nwant data: %d \ngot data %d\n", wantData, gotData)
			}

		})
	}
}
