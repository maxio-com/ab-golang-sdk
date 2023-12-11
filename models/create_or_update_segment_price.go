package models

import (
	"encoding/json"
)

// CreateOrUpdateSegmentPrice represents a CreateOrUpdateSegmentPrice struct.
type CreateOrUpdateSegmentPrice struct {
	StartingQuantity *int `json:"starting_quantity,omitempty"`
	EndingQuantity   *int `json:"ending_quantity,omitempty"`
	// The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
	UnitPrice interface{} `json:"unit_price"`
}

// MarshalJSON implements the json.Marshaler interface for CreateOrUpdateSegmentPrice.
// It customizes the JSON marshaling process for CreateOrUpdateSegmentPrice objects.
func (c *CreateOrUpdateSegmentPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateOrUpdateSegmentPrice object to a map representation for JSON marshaling.
func (c *CreateOrUpdateSegmentPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.StartingQuantity != nil {
		structMap["starting_quantity"] = c.StartingQuantity
	}
	if c.EndingQuantity != nil {
		structMap["ending_quantity"] = c.EndingQuantity
	}
	structMap["unit_price"] = c.UnitPrice
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOrUpdateSegmentPrice.
// It customizes the JSON unmarshaling process for CreateOrUpdateSegmentPrice objects.
func (c *CreateOrUpdateSegmentPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		StartingQuantity *int        `json:"starting_quantity,omitempty"`
		EndingQuantity   *int        `json:"ending_quantity,omitempty"`
		UnitPrice        interface{} `json:"unit_price"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.StartingQuantity = temp.StartingQuantity
	c.EndingQuantity = temp.EndingQuantity
	c.UnitPrice = temp.UnitPrice
	return nil
}
