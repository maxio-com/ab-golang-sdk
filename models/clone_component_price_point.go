// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CloneComponentPricePoint represents a CloneComponentPricePoint struct.
type CloneComponentPricePoint struct {
    Name                 string                 `json:"name"`
    Handle               *string                `json:"handle,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CloneComponentPricePoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CloneComponentPricePoint) String() string {
    return fmt.Sprintf(
    	"CloneComponentPricePoint[Name=%v, Handle=%v, AdditionalProperties=%v]",
    	c.Name, c.Handle, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CloneComponentPricePoint.
// It customizes the JSON marshaling process for CloneComponentPricePoint objects.
func (c CloneComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "name", "handle"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CloneComponentPricePoint object to a map representation for JSON marshaling.
func (c CloneComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    if c.Handle != nil {
        structMap["handle"] = c.Handle
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloneComponentPricePoint.
// It customizes the JSON unmarshaling process for CloneComponentPricePoint objects.
func (c *CloneComponentPricePoint) UnmarshalJSON(input []byte) error {
    var temp tempCloneComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Name = *temp.Name
    c.Handle = temp.Handle
    return nil
}

// tempCloneComponentPricePoint is a temporary struct used for validating the fields of CloneComponentPricePoint.
type tempCloneComponentPricePoint  struct {
    Name   *string `json:"name"`
    Handle *string `json:"handle,omitempty"`
}

func (c *tempCloneComponentPricePoint) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Clone Component Price Point`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
