// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CreateComponentPricePointRequest represents a CreateComponentPricePointRequest struct.
type CreateComponentPricePointRequest struct {
    PricePoint           CreateComponentPricePointRequestPricePoint `json:"price_point"`
    AdditionalProperties map[string]interface{}                     `json:"_"`
}

// String implements the fmt.Stringer interface for CreateComponentPricePointRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateComponentPricePointRequest) String() string {
    return fmt.Sprintf(
    	"CreateComponentPricePointRequest[PricePoint=%v, AdditionalProperties=%v]",
    	c.PricePoint, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePointRequest.
// It customizes the JSON marshaling process for CreateComponentPricePointRequest objects.
func (c CreateComponentPricePointRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePointRequest object to a map representation for JSON marshaling.
func (c CreateComponentPricePointRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePointRequest.
// It customizes the JSON unmarshaling process for CreateComponentPricePointRequest objects.
func (c *CreateComponentPricePointRequest) UnmarshalJSON(input []byte) error {
    var temp tempCreateComponentPricePointRequest
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

// tempCreateComponentPricePointRequest is a temporary struct used for validating the fields of CreateComponentPricePointRequest.
type tempCreateComponentPricePointRequest  struct {
    PricePoint *CreateComponentPricePointRequestPricePoint `json:"price_point"`
}

func (c *tempCreateComponentPricePointRequest) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Create Component Price Point Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
