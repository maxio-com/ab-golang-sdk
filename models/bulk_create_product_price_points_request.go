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

// BulkCreateProductPricePointsRequest represents a BulkCreateProductPricePointsRequest struct.
type BulkCreateProductPricePointsRequest struct {
    PricePoints          []CreateProductPricePoint `json:"price_points"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BulkCreateProductPricePointsRequest.
// It customizes the JSON marshaling process for BulkCreateProductPricePointsRequest objects.
func (b BulkCreateProductPricePointsRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(b.AdditionalProperties,
        "price_points"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(b.toMap())
}

// toMap converts the BulkCreateProductPricePointsRequest object to a map representation for JSON marshaling.
func (b BulkCreateProductPricePointsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["price_points"] = b.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BulkCreateProductPricePointsRequest.
// It customizes the JSON unmarshaling process for BulkCreateProductPricePointsRequest objects.
func (b *BulkCreateProductPricePointsRequest) UnmarshalJSON(input []byte) error {
    var temp tempBulkCreateProductPricePointsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_points")
    if err != nil {
    	return err
    }
    b.AdditionalProperties = additionalProperties
    
    b.PricePoints = *temp.PricePoints
    return nil
}

// tempBulkCreateProductPricePointsRequest is a temporary struct used for validating the fields of BulkCreateProductPricePointsRequest.
type tempBulkCreateProductPricePointsRequest  struct {
    PricePoints *[]CreateProductPricePoint `json:"price_points"`
}

func (b *tempBulkCreateProductPricePointsRequest) validate() error {
    var errs []string
    if b.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `Bulk Create Product Price Points Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
