package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllCustomers(t *testing.T) {
	req, err := http.NewRequest("GET", "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAllCustomers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var got = customers
	// Check if the response body can be decoded into a slice of Customer
	err = json.NewDecoder(rr.Body).Decode(&got)
	if err != nil {
		t.Fatal("Failed to decode response body")
	}
}
func TestReadCustomer(t *testing.T) {
	customerJSON := `{"id": 3, "firstName": "Alice", "lastName": "Johnson", "email": "alice.johnson@example.com"}`
	req, err := http.NewRequest("POST", "/customers", strings.NewReader(customerJSON))
	if err != nil {
		t.Fatal(err)
	}

	customer, err := readCustomer(req)
	if err != nil {
		t.Fatal(err)
	}

	expectedCustomer := Customer{
		ID:        3,
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alice.johnson@example.com",
	}

	if customer != expectedCustomer {
		t.Errorf("readCustomer returned unexpected customer: got %v, want %v", customer, expectedCustomer)
	}
}
