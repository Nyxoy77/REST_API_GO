package authservices

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleLogin(t *testing.T) {
	requestBody := `{
		"email":"aaaaa@gmail.com",
		"password":"aaaaa"
	}`
	req, err := http.NewRequest("POST", "/api/v1/login", strings.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	//Creating a response recorer to understand the response
	rr := httptest.NewRecorder()
	HandleLogin(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code , got : %v expected %v ", status, http.StatusOK)
	}
}
