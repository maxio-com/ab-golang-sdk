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

// ProductPricePointResponse represents a ProductPricePointResponse struct.
type ProductPricePointResponse struct {
    PricePoint           ProductPricePoint      `json:"price_point"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ProductPricePointResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p ProductPricePointResponse) String() string {
    return fmt.Sprintf(
    	"ProductPricePointResponse[PricePoint=%v, AdditionalProperties=%v]",
    	p.PricePoint, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointResponse.
// It customizes the JSON marshaling process for ProductPricePointResponse objects.
func (p ProductPricePointResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointResponse object to a map representation for JSON marshaling.
func (p ProductPricePointResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["price_point"] = p.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointResponse.
// It customizes the JSON unmarshaling process for ProductPricePointResponse objects.
func (p *ProductPricePointResponse) UnmarshalJSON(input []byte) error {
    var temp tempProductPricePointResponse
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
    p.AdditionalProperties = additionalProperties
    
    p.PricePoint = *temp.PricePoint
    return nil
}

// tempProductPricePointResponse is a temporary struct used for validating the fields of ProductPricePointResponse.
type tempProductPricePointResponse  struct {
    PricePoint *ProductPricePoint `json:"price_point"`
}

func (p *tempProductPricePointResponse) validate() error {
    var errs []string
    if p.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Product Price Point Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
