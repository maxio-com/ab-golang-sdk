/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ComponentPricePointResponse represents a ComponentPricePointResponse struct.
type ComponentPricePointResponse struct {
    PricePoint           ComponentPricePoint    `json:"price_point"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointResponse.
// It customizes the JSON marshaling process for ComponentPricePointResponse objects.
func (c ComponentPricePointResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointResponse object to a map representation for JSON marshaling.
func (c ComponentPricePointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointResponse objects.
func (c *ComponentPricePointResponse) UnmarshalJSON(input []byte) error {
    var temp tempComponentPricePointResponse
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

// tempComponentPricePointResponse is a temporary struct used for validating the fields of ComponentPricePointResponse.
type tempComponentPricePointResponse  struct {
    PricePoint *ComponentPricePoint `json:"price_point"`
}

func (c *tempComponentPricePointResponse) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Component Price Point Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
