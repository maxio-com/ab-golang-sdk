package models

import (
    "encoding/json"
)

// CreatePrepaymentResponse represents a CreatePrepaymentResponse struct.
type CreatePrepaymentResponse struct {
    Prepayment CreatedPrepayment `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentResponse.
// It customizes the JSON marshaling process for CreatePrepaymentResponse objects.
func (c *CreatePrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaymentResponse object to a map representation for JSON marshaling.
func (c *CreatePrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepayment"] = c.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentResponse.
// It customizes the JSON unmarshaling process for CreatePrepaymentResponse objects.
func (c *CreatePrepaymentResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Prepayment CreatedPrepayment `json:"prepayment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Prepayment = temp.Prepayment
    return nil
}
