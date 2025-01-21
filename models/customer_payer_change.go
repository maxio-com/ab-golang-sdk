/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CustomerPayerChange represents a CustomerPayerChange struct.
type CustomerPayerChange struct {
    Before               InvoicePayerChange     `json:"before"`
    After                InvoicePayerChange     `json:"after"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CustomerPayerChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerPayerChange) String() string {
    return fmt.Sprintf(
    	"CustomerPayerChange[Before=%v, After=%v, AdditionalProperties=%v]",
    	c.Before, c.After, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CustomerPayerChange.
// It customizes the JSON marshaling process for CustomerPayerChange objects.
func (c CustomerPayerChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "before", "after"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerPayerChange object to a map representation for JSON marshaling.
func (c CustomerPayerChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["before"] = c.Before.toMap()
    structMap["after"] = c.After.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerPayerChange.
// It customizes the JSON unmarshaling process for CustomerPayerChange objects.
func (c *CustomerPayerChange) UnmarshalJSON(input []byte) error {
    var temp tempCustomerPayerChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "before", "after")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Before = *temp.Before
    c.After = *temp.After
    return nil
}

// tempCustomerPayerChange is a temporary struct used for validating the fields of CustomerPayerChange.
type tempCustomerPayerChange  struct {
    Before *InvoicePayerChange `json:"before"`
    After  *InvoicePayerChange `json:"after"`
}

func (c *tempCustomerPayerChange) validate() error {
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
    return errors.New(strings.Join (errs, "\n"))
}
