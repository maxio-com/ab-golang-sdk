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

// AddressChange represents a AddressChange struct.
type AddressChange struct {
    Before               InvoiceAddress         `json:"before"`
    After                InvoiceAddress         `json:"after"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AddressChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AddressChange) String() string {
    return fmt.Sprintf(
    	"AddressChange[Before=%v, After=%v, AdditionalProperties=%v]",
    	a.Before, a.After, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AddressChange.
// It customizes the JSON marshaling process for AddressChange objects.
func (a AddressChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "before", "after"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AddressChange object to a map representation for JSON marshaling.
func (a AddressChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["before"] = a.Before.toMap()
    structMap["after"] = a.After.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressChange.
// It customizes the JSON unmarshaling process for AddressChange objects.
func (a *AddressChange) UnmarshalJSON(input []byte) error {
    var temp tempAddressChange
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
    a.AdditionalProperties = additionalProperties
    
    a.Before = *temp.Before
    a.After = *temp.After
    return nil
}

// tempAddressChange is a temporary struct used for validating the fields of AddressChange.
type tempAddressChange  struct {
    Before *InvoiceAddress `json:"before"`
    After  *InvoiceAddress `json:"after"`
}

func (a *tempAddressChange) validate() error {
    var errs []string
    if a.Before == nil {
        errs = append(errs, "required field `before` is missing for type `Address Change`")
    }
    if a.After == nil {
        errs = append(errs, "required field `after` is missing for type `Address Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
