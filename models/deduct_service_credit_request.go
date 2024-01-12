package models

import (
    "encoding/json"
)

// DeductServiceCreditRequest represents a DeductServiceCreditRequest struct.
type DeductServiceCreditRequest struct {
    Deduction DeductServiceCredit `json:"deduction"`
}

// MarshalJSON implements the json.Marshaler interface for DeductServiceCreditRequest.
// It customizes the JSON marshaling process for DeductServiceCreditRequest objects.
func (d *DeductServiceCreditRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DeductServiceCreditRequest object to a map representation for JSON marshaling.
func (d *DeductServiceCreditRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["deduction"] = d.Deduction
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DeductServiceCreditRequest.
// It customizes the JSON unmarshaling process for DeductServiceCreditRequest objects.
func (d *DeductServiceCreditRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Deduction DeductServiceCredit `json:"deduction"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    d.Deduction = temp.Deduction
    return nil
}
