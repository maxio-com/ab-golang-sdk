package models

import (
    "encoding/json"
)

// ListComponentsPricePointsResponse represents a ListComponentsPricePointsResponse struct.
type ListComponentsPricePointsResponse struct {
    PricePoints []ComponentPricePoint `json:"price_points"`
}

// MarshalJSON implements the json.Marshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON marshaling process for ListComponentsPricePointsResponse objects.
func (l *ListComponentsPricePointsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListComponentsPricePointsResponse object to a map representation for JSON marshaling.
func (l *ListComponentsPricePointsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["price_points"] = l.PricePoints
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListComponentsPricePointsResponse.
// It customizes the JSON unmarshaling process for ListComponentsPricePointsResponse objects.
func (l *ListComponentsPricePointsResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricePoints []ComponentPricePoint `json:"price_points"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    l.PricePoints = temp.PricePoints
    return nil
}
