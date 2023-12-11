package models

import (
	"encoding/json"
)

// PayerError represents a PayerError struct.
type PayerError struct {
	LastName  []string `json:"last_name,omitempty"`
	FirstName []string `json:"first_name,omitempty"`
	Email     []string `json:"email,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PayerError.
// It customizes the JSON marshaling process for PayerError objects.
func (p *PayerError) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PayerError object to a map representation for JSON marshaling.
func (p *PayerError) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.LastName != nil {
		structMap["last_name"] = p.LastName
	}
	if p.FirstName != nil {
		structMap["first_name"] = p.FirstName
	}
	if p.Email != nil {
		structMap["email"] = p.Email
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PayerError.
// It customizes the JSON unmarshaling process for PayerError objects.
func (p *PayerError) UnmarshalJSON(input []byte) error {
	temp := &struct {
		LastName  []string `json:"last_name,omitempty"`
		FirstName []string `json:"first_name,omitempty"`
		Email     []string `json:"email,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.LastName = temp.LastName
	p.FirstName = temp.FirstName
	p.Email = temp.Email
	return nil
}
