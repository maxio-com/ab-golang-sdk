package models

import (
    "encoding/json"
)

// PrepaymentsResponse represents a PrepaymentsResponse struct.
type PrepaymentsResponse struct {
    Prepayments []Prepayment `json:"prepayments,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentsResponse.
// It customizes the JSON marshaling process for PrepaymentsResponse objects.
func (p *PrepaymentsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentsResponse object to a map representation for JSON marshaling.
func (p *PrepaymentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    if p.Prepayments != nil {
        structMap["prepayments"] = p.Prepayments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentsResponse.
// It customizes the JSON unmarshaling process for PrepaymentsResponse objects.
func (p *PrepaymentsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Prepayments []Prepayment `json:"prepayments,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Prepayments = temp.Prepayments
    return nil
}
