package main

// PhoneCompany struct to define a phone company
type PhoneCompany struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Country     string `json:"country"`
	FoundedYear int    `json:"founded_year"`
	Description string `json:"description"`
}
