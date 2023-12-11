package models

import (
	"encoding/json"
)

// ComponentPricePointsResponse represents a ComponentPricePointsResponse struct.
type ComponentPricePointsResponse struct {
	PricePoints []ComponentPricePoint `json:"price_points,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointsResponse.
// It customizes the JSON marshaling process for ComponentPricePointsResponse objects.
func (c *ComponentPricePointsResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointsResponse object to a map representation for JSON marshaling.
func (c *ComponentPricePointsResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.PricePoints != nil {
		structMap["price_points"] = c.PricePoints
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointsResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointsResponse objects.
func (c *ComponentPricePointsResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PricePoints []ComponentPricePoint `json:"price_points,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PricePoints = temp.PricePoints
	return nil
}
