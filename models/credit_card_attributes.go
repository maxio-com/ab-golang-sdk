package models

import (
    "encoding/json"
)

// CreditCardAttributes represents a CreditCardAttributes struct.
type CreditCardAttributes struct {
    FullNumber           *string        `json:"full_number,omitempty"`
    ExpirationMonth      *string        `json:"expiration_month,omitempty"`
    ExpirationYear       *string        `json:"expiration_year,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreditCardAttributes.
// It customizes the JSON marshaling process for CreditCardAttributes objects.
func (c CreditCardAttributes) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreditCardAttributes object to a map representation for JSON marshaling.
func (c CreditCardAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.FullNumber != nil {
        structMap["full_number"] = c.FullNumber
    }
    if c.ExpirationMonth != nil {
        structMap["expiration_month"] = c.ExpirationMonth
    }
    if c.ExpirationYear != nil {
        structMap["expiration_year"] = c.ExpirationYear
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditCardAttributes.
// It customizes the JSON unmarshaling process for CreditCardAttributes objects.
func (c *CreditCardAttributes) UnmarshalJSON(input []byte) error {
    var temp creditCardAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "full_number", "expiration_month", "expiration_year")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.FullNumber = temp.FullNumber
    c.ExpirationMonth = temp.ExpirationMonth
    c.ExpirationYear = temp.ExpirationYear
    return nil
}

// TODO
type creditCardAttributes  struct {
    FullNumber      *string `json:"full_number,omitempty"`
    ExpirationMonth *string `json:"expiration_month,omitempty"`
    ExpirationYear  *string `json:"expiration_year,omitempty"`
}
