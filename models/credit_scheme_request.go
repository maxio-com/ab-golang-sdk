package models

import (
	"encoding/json"
)

// CreditSchemeRequest represents a CreditSchemeRequest struct.
type CreditSchemeRequest struct {
	CreditScheme CreditScheme `json:"credit_scheme"`
}

// MarshalJSON implements the json.Marshaler interface for CreditSchemeRequest.
// It customizes the JSON marshaling process for CreditSchemeRequest objects.
func (c *CreditSchemeRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreditSchemeRequest object to a map representation for JSON marshaling.
func (c *CreditSchemeRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["credit_scheme"] = c.CreditScheme
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditSchemeRequest.
// It customizes the JSON unmarshaling process for CreditSchemeRequest objects.
func (c *CreditSchemeRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CreditScheme CreditScheme `json:"credit_scheme"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CreditScheme = temp.CreditScheme
	return nil
}
