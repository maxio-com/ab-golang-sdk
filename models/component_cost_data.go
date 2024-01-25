package models

import (
	"encoding/json"
)

// ComponentCostData represents a ComponentCostData struct.
type ComponentCostData struct {
	ComponentCodeId Optional[int] `json:"component_code_id"`
	PricePointId    *int          `json:"price_point_id,omitempty"`
	ProductId       *int          `json:"product_id,omitempty"`
	Quantity        *string       `json:"quantity,omitempty"`
	Amount          *string       `json:"amount,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme *PricingScheme              `json:"pricing_scheme,omitempty"`
	Tiers         []ComponentCostDataRateTier `json:"tiers,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCostData.
// It customizes the JSON marshaling process for ComponentCostData objects.
func (c *ComponentCostData) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentCostData object to a map representation for JSON marshaling.
func (c *ComponentCostData) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.ComponentCodeId.IsValueSet() {
		structMap["component_code_id"] = c.ComponentCodeId.Value()
	}
	if c.PricePointId != nil {
		structMap["price_point_id"] = c.PricePointId
	}
	if c.ProductId != nil {
		structMap["product_id"] = c.ProductId
	}
	if c.Quantity != nil {
		structMap["quantity"] = c.Quantity
	}
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.PricingScheme != nil {
		structMap["pricing_scheme"] = c.PricingScheme
	}
	if c.Tiers != nil {
		structMap["tiers"] = c.Tiers
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCostData.
// It customizes the JSON unmarshaling process for ComponentCostData objects.
func (c *ComponentCostData) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ComponentCodeId Optional[int]               `json:"component_code_id"`
		PricePointId    *int                        `json:"price_point_id,omitempty"`
		ProductId       *int                        `json:"product_id,omitempty"`
		Quantity        *string                     `json:"quantity,omitempty"`
		Amount          *string                     `json:"amount,omitempty"`
		PricingScheme   *PricingScheme              `json:"pricing_scheme,omitempty"`
		Tiers           []ComponentCostDataRateTier `json:"tiers,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.ComponentCodeId = temp.ComponentCodeId
	c.PricePointId = temp.PricePointId
	c.ProductId = temp.ProductId
	c.Quantity = temp.Quantity
	c.Amount = temp.Amount
	c.PricingScheme = temp.PricingScheme
	c.Tiers = temp.Tiers
	return nil
}
