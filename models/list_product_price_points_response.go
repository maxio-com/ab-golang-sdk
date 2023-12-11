package models

import (
	"encoding/json"
)

// ListProductPricePointsResponse represents a ListProductPricePointsResponse struct.
type ListProductPricePointsResponse struct {
	PricePoints []ProductPricePoint `json:"price_points"`
}

// MarshalJSON implements the json.Marshaler interface for ListProductPricePointsResponse.
// It customizes the JSON marshaling process for ListProductPricePointsResponse objects.
func (l *ListProductPricePointsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListProductPricePointsResponse object to a map representation for JSON marshaling.
func (l *ListProductPricePointsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["price_points"] = l.PricePoints
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductPricePointsResponse.
// It customizes the JSON unmarshaling process for ListProductPricePointsResponse objects.
func (l *ListProductPricePointsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PricePoints []ProductPricePoint `json:"price_points"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.PricePoints = temp.PricePoints
	return nil
}
