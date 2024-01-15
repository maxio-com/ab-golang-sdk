package models

import (
    "encoding/json"
)

// PrepaymentResponse represents a PrepaymentResponse struct.
type PrepaymentResponse struct {
    Prepayment Prepayment `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentResponse.
// It customizes the JSON marshaling process for PrepaymentResponse objects.
func (p *PrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentResponse object to a map representation for JSON marshaling.
func (p *PrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepayment"] = p.Prepayment
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentResponse.
// It customizes the JSON unmarshaling process for PrepaymentResponse objects.
func (p *PrepaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Prepayment Prepayment `json:"prepayment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Prepayment = temp.Prepayment
    return nil
}
