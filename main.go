package main

import (
	"fmt"
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
