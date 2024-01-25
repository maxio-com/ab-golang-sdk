package models

import (
	"encoding/json"
)

// ComponentPricePointErrorItem represents a ComponentPricePointErrorItem struct.
type ComponentPricePointErrorItem struct {
	ComponentId *int    `json:"component_id,omitempty"`
	Message     *string `json:"message,omitempty"`
	PricePoint  *int    `json:"price_point,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointErrorItem.
// It customizes the JSON marshaling process for ComponentPricePointErrorItem objects.
func (c *ComponentPricePointErrorItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointErrorItem object to a map representation for JSON marshaling.
func (c *ComponentPricePointErrorItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ComponentId != nil {
		structMap["component_id"] = c.ComponentId
	}
	if c.Message != nil {
		structMap["message"] = c.Message
	}
	if c.PricePoint != nil {
		structMap["price_point"] = c.PricePoint
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointErrorItem.
// It customizes the JSON unmarshaling process for ComponentPricePointErrorItem objects.
func (c *ComponentPricePointErrorItem) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ComponentId *int    `json:"component_id,omitempty"`
		Message     *string `json:"message,omitempty"`
		PricePoint  *int    `json:"price_point,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ComponentId = temp.ComponentId
	c.Message = temp.Message
	c.PricePoint = temp.PricePoint
	return nil
}
