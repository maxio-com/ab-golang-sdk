package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CustomerPayerChange represents a CustomerPayerChange struct.
type CustomerPayerChange struct {
    Before               InvoicePayerChange `json:"before"`
    After                InvoicePayerChange `json:"after"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerPayerChange.
// It customizes the JSON marshaling process for CustomerPayerChange objects.
func (c CustomerPayerChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerPayerChange object to a map representation for JSON marshaling.
func (c CustomerPayerChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["before"] = c.Before.toMap()
    structMap["after"] = c.After.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerPayerChange.
// It customizes the JSON unmarshaling process for CustomerPayerChange objects.
func (c *CustomerPayerChange) UnmarshalJSON(input []byte) error {
    var temp customerPayerChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "before", "after")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Before = *temp.Before
    c.After = *temp.After
    return nil
}

// customerPayerChange is a temporary struct used for validating the fields of CustomerPayerChange.
type customerPayerChange  struct {
    Before *InvoicePayerChange `json:"before"`
    After  *InvoicePayerChange `json:"after"`
}

func (c *customerPayerChange) validate() error {
    var errs []string
    if c.Before == nil {
        errs = append(errs, "required field `before` is missing for type `Customer Payer Change`")
    }
    if c.After == nil {
        errs = append(errs, "required field `after` is missing for type `Customer Payer Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
