package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBook(t *testing.T) {

	expected := `[{"bookId":"101","bookName":"Book Name 1","bookAuthor":"Book Author 1"},{"bookId":"102","bookName":"Book Name 2","bookAuthor":"Book Author 2"}]`

	tests := []struct {
		name           string
		expectedStatus int
		expectedBody   string
	}{

		{
			name:           "GetAllBook",
			expectedStatus: http.StatusOK,
			expectedBody:   expected,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/getBook", nil)
			if err != nil {
				t.Error(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetAllBook)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("got %v want %v ", status, http.StatusOK)
			}

			// Check the content type
			expectedContentType := "application/json"
			if ctype := rr.Header().Get("Content-Type"); ctype != expectedContentType {
				t.Errorf("handler returned wrong content type: got %v want %v",
					ctype, expectedContentType)
			}

			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}

		})
	}

}
