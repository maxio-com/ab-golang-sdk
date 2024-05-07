package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListProductPricePointsResponse represents a ListProductPricePointsResponse struct.
type ListProductPricePointsResponse struct {
    PricePoints          []ProductPricePoint `json:"price_points"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListProductPricePointsResponse.
// It customizes the JSON marshaling process for ListProductPricePointsResponse objects.
func (l ListProductPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListProductPricePointsResponse object to a map representation for JSON marshaling.
func (l ListProductPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["price_points"] = l.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductPricePointsResponse.
// It customizes the JSON unmarshaling process for ListProductPricePointsResponse objects.
func (l *ListProductPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp listProductPricePointsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_points")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.PricePoints = *temp.PricePoints
    return nil
}

// listProductPricePointsResponse is a temporary struct used for validating the fields of ListProductPricePointsResponse.
type listProductPricePointsResponse  struct {
    PricePoints *[]ProductPricePoint `json:"price_points"`
}

func (l *listProductPricePointsResponse) validate() error {
    var errs []string
    if l.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `List Product Price Points Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
