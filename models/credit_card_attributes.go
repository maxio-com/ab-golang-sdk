/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CreditCardAttributes represents a CreditCardAttributes struct.
type CreditCardAttributes struct {
    FullNumber           *string                `json:"full_number,omitempty"`
    ExpirationMonth      *string                `json:"expiration_month,omitempty"`
    ExpirationYear       *string                `json:"expiration_year,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreditCardAttributes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreditCardAttributes) String() string {
    return fmt.Sprintf(
    	"CreditCardAttributes[FullNumber=%v, ExpirationMonth=%v, ExpirationYear=%v, AdditionalProperties=%v]",
    	c.FullNumber, c.ExpirationMonth, c.ExpirationYear, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreditCardAttributes.
// It customizes the JSON marshaling process for CreditCardAttributes objects.
func (c CreditCardAttributes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "full_number", "expiration_month", "expiration_year"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreditCardAttributes object to a map representation for JSON marshaling.
func (c CreditCardAttributes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCreditCardAttributes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "full_number", "expiration_month", "expiration_year")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.FullNumber = temp.FullNumber
    c.ExpirationMonth = temp.ExpirationMonth
    c.ExpirationYear = temp.ExpirationYear
    return nil
}

// tempCreditCardAttributes is a temporary struct used for validating the fields of CreditCardAttributes.
type tempCreditCardAttributes  struct {
    FullNumber      *string `json:"full_number,omitempty"`
    ExpirationMonth *string `json:"expiration_month,omitempty"`
    ExpirationYear  *string `json:"expiration_year,omitempty"`
}
