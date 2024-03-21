package models

import (
	"encoding/json"
)

// ListMRRResponseResult represents a ListMRRResponseResult struct.
type ListMRRResponseResult struct {
	Page           *int       `json:"page,omitempty"`
	PerPage        *int       `json:"per_page,omitempty"`
	TotalPages     *int       `json:"total_pages,omitempty"`
	TotalEntries   *int       `json:"total_entries,omitempty"`
	Currency       *string    `json:"currency,omitempty"`
	CurrencySymbol *string    `json:"currency_symbol,omitempty"`
	Movements      []Movement `json:"movements,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListMRRResponseResult.
// It customizes the JSON marshaling process for ListMRRResponseResult objects.
func (l *ListMRRResponseResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListMRRResponseResult object to a map representation for JSON marshaling.
func (l *ListMRRResponseResult) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.Page != nil {
		structMap["page"] = l.Page
	}
	if l.PerPage != nil {
		structMap["per_page"] = l.PerPage
	}
	if l.TotalPages != nil {
		structMap["total_pages"] = l.TotalPages
	}
	if l.TotalEntries != nil {
		structMap["total_entries"] = l.TotalEntries
	}
	if l.Currency != nil {
		structMap["currency"] = l.Currency
	}
	if l.CurrencySymbol != nil {
		structMap["currency_symbol"] = l.CurrencySymbol
	}
	if l.Movements != nil {
		structMap["movements"] = l.Movements
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListMRRResponseResult.
// It customizes the JSON unmarshaling process for ListMRRResponseResult objects.
func (l *ListMRRResponseResult) UnmarshalJSON(input []byte) error {
	var temp listMRRResponseResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	l.Page = temp.Page
	l.PerPage = temp.PerPage
	l.TotalPages = temp.TotalPages
	l.TotalEntries = temp.TotalEntries
	l.Currency = temp.Currency
	l.CurrencySymbol = temp.CurrencySymbol
	l.Movements = temp.Movements
	return nil
}

// TODO
type listMRRResponseResult struct {
	Page           *int       `json:"page,omitempty"`
	PerPage        *int       `json:"per_page,omitempty"`
	TotalPages     *int       `json:"total_pages,omitempty"`
	TotalEntries   *int       `json:"total_entries,omitempty"`
	Currency       *string    `json:"currency,omitempty"`
	CurrencySymbol *string    `json:"currency_symbol,omitempty"`
	Movements      []Movement `json:"movements,omitempty"`
}
