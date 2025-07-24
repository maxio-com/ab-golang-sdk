// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ListComponentsPricePointsResponse represents a ListComponentsPricePointsResponse struct.
type ListComponentsPricePointsResponse struct {
    PricePoints          []ComponentPricePoint  `json:"price_points"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListComponentsPricePointsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListComponentsPricePointsResponse) String() string {
    return fmt.Sprintf(
    	"ListComponentsPricePointsResponse[PricePoints=%v, AdditionalProperties=%v]",
    	l.PricePoints, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON marshaling process for ListComponentsPricePointsResponse objects.
func (l ListComponentsPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "price_points"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListComponentsPricePointsResponse object to a map representation for JSON marshaling.
func (l ListComponentsPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["price_points"] = l.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON unmarshaling process for ListComponentsPricePointsResponse objects.
func (l *ListComponentsPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListComponentsPricePointsResponse
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

// tempListComponentsPricePointsResponse is a temporary struct used for validating the fields of ListComponentsPricePointsResponse.
type tempListComponentsPricePointsResponse  struct {
    PricePoints *[]ComponentPricePoint `json:"price_points"`
}

func (l *tempListComponentsPricePointsResponse) validate() error {
    var errs []string
    if l.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `List Components Price Points Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
