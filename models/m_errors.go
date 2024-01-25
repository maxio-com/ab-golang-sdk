package models

import (
	"encoding/json"
)

// Errors represents a Errors struct.
type Errors struct {
	PerPage    []string `json:"per_page,omitempty"`
	PricePoint []string `json:"price_point,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Errors.
// It customizes the JSON marshaling process for Errors objects.
func (e *Errors) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the Errors object to a map representation for JSON marshaling.
func (e *Errors) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.PerPage != nil {
		structMap["per_page"] = e.PerPage
	}
	if e.PricePoint != nil {
		structMap["price_point"] = e.PricePoint
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Errors.
// It customizes the JSON unmarshaling process for Errors objects.
func (e *Errors) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PerPage    []string `json:"per_page,omitempty"`
		PricePoint []string `json:"price_point,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.PerPage = temp.PerPage
	e.PricePoint = temp.PricePoint
	return nil
}
