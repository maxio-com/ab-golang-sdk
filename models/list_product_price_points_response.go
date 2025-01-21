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

// ListProductPricePointsResponse represents a ListProductPricePointsResponse struct.
type ListProductPricePointsResponse struct {
    PricePoints          []ProductPricePoint    `json:"price_points"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListProductPricePointsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListProductPricePointsResponse) String() string {
    return fmt.Sprintf(
    	"ListProductPricePointsResponse[PricePoints=%v, AdditionalProperties=%v]",
    	l.PricePoints, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListProductPricePointsResponse.
// It customizes the JSON marshaling process for ListProductPricePointsResponse objects.
func (l ListProductPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "price_points"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListProductPricePointsResponse object to a map representation for JSON marshaling.
func (l ListProductPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["price_points"] = l.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductPricePointsResponse.
// It customizes the JSON unmarshaling process for ListProductPricePointsResponse objects.
func (l *ListProductPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListProductPricePointsResponse
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
    l.AdditionalProperties = additionalProperties
    
    l.PricePoints = *temp.PricePoints
    return nil
}

// tempListProductPricePointsResponse is a temporary struct used for validating the fields of ListProductPricePointsResponse.
type tempListProductPricePointsResponse  struct {
    PricePoints *[]ProductPricePoint `json:"price_points"`
}

func (l *tempListProductPricePointsResponse) validate() error {
    var errs []string
    if l.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `List Product Price Points Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
