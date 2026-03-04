// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CloneComponentPricePointRequest represents a CloneComponentPricePointRequest struct.
type CloneComponentPricePointRequest struct {
    PricePoint           CloneComponentPricePoint `json:"price_point"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for CloneComponentPricePointRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CloneComponentPricePointRequest) String() string {
    return fmt.Sprintf(
    	"CloneComponentPricePointRequest[PricePoint=%v, AdditionalProperties=%v]",
    	c.PricePoint, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CloneComponentPricePointRequest.
// It customizes the JSON marshaling process for CloneComponentPricePointRequest objects.
func (c CloneComponentPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CloneComponentPricePointRequest object to a map representation for JSON marshaling.
func (c CloneComponentPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CloneComponentPricePointRequest.
// It customizes the JSON unmarshaling process for CloneComponentPricePointRequest objects.
func (c *CloneComponentPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp tempCloneComponentPricePointRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PricePoint = *temp.PricePoint
    return nil
}

// tempCloneComponentPricePointRequest is a temporary struct used for validating the fields of CloneComponentPricePointRequest.
type tempCloneComponentPricePointRequest  struct {
    PricePoint *CloneComponentPricePoint `json:"price_point"`
}

func (c *tempCloneComponentPricePointRequest) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Clone Component Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
