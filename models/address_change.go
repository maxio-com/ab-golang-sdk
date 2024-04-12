package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AddressChange represents a AddressChange struct.
type AddressChange struct {
    Before               InvoiceAddress `json:"before"`
    After                InvoiceAddress `json:"after"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AddressChange.
// It customizes the JSON marshaling process for AddressChange objects.
func (a AddressChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AddressChange object to a map representation for JSON marshaling.
func (a AddressChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["before"] = a.Before.toMap()
    structMap["after"] = a.After.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddressChange.
// It customizes the JSON unmarshaling process for AddressChange objects.
func (a *AddressChange) UnmarshalJSON(input []byte) error {
    var temp addressChange
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
    
    a.AdditionalProperties = additionalProperties
    a.Before = *temp.Before
    a.After = *temp.After
    return nil
}

// TODO
type addressChange  struct {
    Before *InvoiceAddress `json:"before"`
    After  *InvoiceAddress `json:"after"`
}

func (a *addressChange) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
