// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AttributeError represents a AttributeError struct.
type AttributeError struct {
    Attribute            []string               `json:"attribute"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AttributeError,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AttributeError) String() string {
    return fmt.Sprintf(
    	"AttributeError[Attribute=%v, AdditionalProperties=%v]",
    	a.Attribute, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AttributeError.
// It customizes the JSON marshaling process for AttributeError objects.
func (a AttributeError) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "attribute"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AttributeError object to a map representation for JSON marshaling.
func (a AttributeError) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["attribute"] = a.Attribute
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AttributeError.
// It customizes the JSON unmarshaling process for AttributeError objects.
func (a *AttributeError) UnmarshalJSON(input []byte) error {
    var temp tempAttributeError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "attribute")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Attribute = *temp.Attribute
    return nil
}

// tempAttributeError is a temporary struct used for validating the fields of AttributeError.
type tempAttributeError  struct {
    Attribute *[]string `json:"attribute"`
}

func (a *tempAttributeError) validate() error {
    var errs []string
    if a.Attribute == nil {
        errs = append(errs, "required field `attribute` is missing for type `Attribute Error`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
