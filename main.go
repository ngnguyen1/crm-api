package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

var customers map[string]Customer

func seedCustomers() {
	customers = make(map[string]Customer)

	customer1 := Customer{
		ID:        1,
		Name:      "John Doe",
		Role:      "Admin",
		Email:     "john.doe@example.com",
		Phone:     "123-456-7890",
		Contacted: false,
	}
	customers["1"] = customer1

	customer2 := Customer{
		ID:        2,
		Name:      "Jane Smith",
		Role:      "User",
		Email:     "jane.smith@example.com",
		Phone:     "098-765-4321",
		Contacted: true,
	}
	customers["2"] = customer2

}

func init() {
	seedCustomers()
}

func readCustomer(r *http.Request) (Customer, error) {
	var customer Customer
	err := json.NewDecoder(r.Body).Decode(&customer)
	return customer, err
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer

	customer, err := readCustomer(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid customer data")
		return
	}

	if _, exists := customers[strconv.Itoa(customer.ID)]; exists {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode("Customer already exists")
		return
	}

	customers[strconv.Itoa(customer.ID)] = customer
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["id"]

	if customer, ok := customers[customerID]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with id: " + customerID + " not found.")
	}
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customerID := params["id"]

	id, err := strconv.ParseUint(customerID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid customer ID")
		return
	}

	if _, exists := customers[customerID]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with id: " + customerID + " not found.")
		return
	}

	customer, err := readCustomer(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid customer data")
		return
	}

	if uint64(customer.ID) != id {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Customer ID mismatch")
		return
	}

	customers[customerID] = customer
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	customerID := params["id"]

	if _, exists := customers[customerID]; !exists {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Customer with id: " + customerID + " not found.")
		return
	}

	delete(customers, customerID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)

}

func showIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	seedCustomers()
	router := mux.NewRouter()
	router.HandleFunc("/", showIndexPage).Methods("GET")
	router.HandleFunc("/customers", getAllCustomers).Methods("GET")
	router.HandleFunc("/customers", createCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is running on port 3000...")
	http.ListenAndServe(":3000", router)
}
