// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CustomerCustomFieldsChange represents a CustomerCustomFieldsChange struct.
type CustomerCustomFieldsChange struct {
    Before               []InvoiceCustomField   `json:"before"`
    After                []InvoiceCustomField   `json:"after"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CustomerCustomFieldsChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CustomerCustomFieldsChange) String() string {
    return fmt.Sprintf(
    	"CustomerCustomFieldsChange[Before=%v, After=%v, AdditionalProperties=%v]",
    	c.Before, c.After, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON marshaling process for CustomerCustomFieldsChange objects.
func (c CustomerCustomFieldsChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "before", "after"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CustomerCustomFieldsChange object to a map representation for JSON marshaling.
func (c CustomerCustomFieldsChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["before"] = c.Before
    structMap["after"] = c.After
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON unmarshaling process for CustomerCustomFieldsChange objects.
func (c *CustomerCustomFieldsChange) UnmarshalJSON(input []byte) error {
    var temp tempCustomerCustomFieldsChange
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

// tempCustomerCustomFieldsChange is a temporary struct used for validating the fields of CustomerCustomFieldsChange.
type tempCustomerCustomFieldsChange  struct {
    Before *[]InvoiceCustomField `json:"before"`
    After  *[]InvoiceCustomField `json:"after"`
}

func (c *tempCustomerCustomFieldsChange) validate() error {
    var errs []string
    if c.Before == nil {
        errs = append(errs, "required field `before` is missing for type `Customer Custom Fields Change`")
    }
    if c.After == nil {
        errs = append(errs, "required field `after` is missing for type `Customer Custom Fields Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
