package models

import (
	"encoding/json"
)

// ComponentCustomPrice represents a ComponentCustomPrice struct.
// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
type ComponentCustomPrice struct {
	// Omit for On/Off components
	PricingScheme *PricingSchemeEnum `json:"pricing_scheme,omitempty"`
	// On/off components only need one price bracket starting at 1
	Prices []Price `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCustomPrice.
// It customizes the JSON marshaling process for ComponentCustomPrice objects.
func (c *ComponentCustomPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentCustomPrice object to a map representation for JSON marshaling.
func (c *ComponentCustomPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.PricingScheme != nil {
		structMap["pricing_scheme"] = c.PricingScheme
	}
	if c.Prices != nil {
		structMap["prices"] = c.Prices
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCustomPrice.
// It customizes the JSON unmarshaling process for ComponentCustomPrice objects.
func (c *ComponentCustomPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PricingScheme *PricingSchemeEnum `json:"pricing_scheme,omitempty"`
		Prices        []Price            `json:"prices,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.PricingScheme = temp.PricingScheme
	c.Prices = temp.Prices
	return nil
}
