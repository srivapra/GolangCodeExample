package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		name              string
		url               string
		expected          string
		expectedStausCode int
	}{
		{
			name:              "Positive Test Case",
			url:               "/",
			expected:          "Hello",
			expectedStausCode: http.StatusOK,
		},
		{
			name:              "Negative Test Case",
			url:               "/hello",
			expected:          "404 not found",
			expectedStausCode: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("Get", tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStausCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStausCode)
			}

			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.expected)
			}

		})
	}

}
