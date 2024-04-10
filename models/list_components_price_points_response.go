package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListComponentsPricePointsResponse represents a ListComponentsPricePointsResponse struct.
type ListComponentsPricePointsResponse struct {
    PricePoints          []ComponentPricePoint `json:"price_points"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON marshaling process for ListComponentsPricePointsResponse objects.
func (l ListComponentsPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListComponentsPricePointsResponse object to a map representation for JSON marshaling.
func (l ListComponentsPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["price_points"] = l.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON unmarshaling process for ListComponentsPricePointsResponse objects.
func (l *ListComponentsPricePointsResponse) UnmarshalJSON(input []byte) error {
    var temp listComponentsPricePointsResponse
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

// TODO
type listComponentsPricePointsResponse  struct {
    PricePoints *[]ComponentPricePoint `json:"price_points"`
}

func (l *listComponentsPricePointsResponse) validate() error {
    var errs []string
    if l.PricePoints == nil {
        errs = append(errs, "required field `price_points` is missing for type `List Components Price Points Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
