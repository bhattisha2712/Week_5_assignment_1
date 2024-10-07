package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// PhoneCompany struct to define a phone company
type PhoneCompany struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	FoundedYear int    `json:"founded_year"`
	Description string `json:"description"`
}

var companies []PhoneCompany
var idCounter int = 1

// Helper function to get the company ID from the URL
func getIDFromURL(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid URL")
	}
	return strconv.Atoi(parts[2])
}

// Create a new phone company (POST /companies)
func createCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newCompany PhoneCompany
	err := json.NewDecoder(r.Body).Decode(&newCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newCompany.ID = idCounter
	idCounter++
	companies = append(companies, newCompany)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCompany)
}

// Get all phone companies (GET /companies)
func getCompanies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}

// Get a phone company by ID (GET /companies/{id})
func getCompanyByID(w http.ResponseWriter, r *http.Request) {
	companyID, err := getIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	for _, company := range companies {
		if company.ID == companyID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(company)
			return
		}
	}

	http.Error(w, "Company not found", http.StatusNotFound)
}

// Update a phone company (PUT /companies/{id})
func updateCompany(w http.ResponseWriter, r *http.Request) {
	companyID, err := getIDFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid company ID", http.StatusBadRequest)
		return
	}

	var updatedCompany PhoneCompany
	err = json.NewDecoder(r.Body).Decode(&updatedCompany)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, company := range companies {
		if company.ID == companyID {
			companies[i].Name = updatedCompany.Name
			companies[i].Country = updatedCompany.Country
			companies[i].FoundedYear = updatedCompany.FoundedYear
			companies[i].Description = updatedCompany.Description
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(companies[i])
			return
		}
	}

	http.Error(w, "Company not found", http.StatusNotFound)
}
