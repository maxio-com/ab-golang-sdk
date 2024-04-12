package models

import (
    "encoding/json"
)

// PrepaymentsResponse represents a PrepaymentsResponse struct.
type PrepaymentsResponse struct {
    Prepayments          []Prepayment   `json:"prepayments,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaymentsResponse.
// It customizes the JSON marshaling process for PrepaymentsResponse objects.
func (p PrepaymentsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaymentsResponse object to a map representation for JSON marshaling.
func (p PrepaymentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Prepayments != nil {
        structMap["prepayments"] = p.Prepayments
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentsResponse.
// It customizes the JSON unmarshaling process for PrepaymentsResponse objects.
func (p *PrepaymentsResponse) UnmarshalJSON(input []byte) error {
    var temp prepaymentsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepayments")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Prepayments = temp.Prepayments
    return nil
}

// TODO
type prepaymentsResponse  struct {
    Prepayments []Prepayment `json:"prepayments,omitempty"`
}
