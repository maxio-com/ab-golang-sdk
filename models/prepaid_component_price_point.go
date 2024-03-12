package models

import (
	"encoding/json"
)

// PrepaidComponentPricePoint represents a PrepaidComponentPricePoint struct.
type PrepaidComponentPricePoint struct {
	Name   *string `json:"name,omitempty"`
	Handle *string `json:"handle,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme  *PricingScheme  `json:"pricing_scheme,omitempty"`
	Prices         []Price         `json:"prices,omitempty"`
	OveragePricing *OveragePricing `json:"overage_pricing,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidComponentPricePoint.
// It customizes the JSON marshaling process for PrepaidComponentPricePoint objects.
func (p *PrepaidComponentPricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PrepaidComponentPricePoint object to a map representation for JSON marshaling.
func (p *PrepaidComponentPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	if p.Name != nil {
		structMap["name"] = p.Name
	}
	if p.Handle != nil {
		structMap["handle"] = p.Handle
	}
	if p.PricingScheme != nil {
		structMap["pricing_scheme"] = p.PricingScheme
	}
	if p.Prices != nil {
		structMap["prices"] = p.Prices
	}
	if p.OveragePricing != nil {
		structMap["overage_pricing"] = p.OveragePricing.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidComponentPricePoint.
// It customizes the JSON unmarshaling process for PrepaidComponentPricePoint objects.
func (p *PrepaidComponentPricePoint) UnmarshalJSON(input []byte) error {
	var temp prepaidComponentPricePoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	p.Name = temp.Name
	p.Handle = temp.Handle
	p.PricingScheme = temp.PricingScheme
	p.Prices = temp.Prices
	p.OveragePricing = temp.OveragePricing
	return nil
}

// TODO
type prepaidComponentPricePoint struct {
	Name           *string         `json:"name,omitempty"`
	Handle         *string         `json:"handle,omitempty"`
	PricingScheme  *PricingScheme  `json:"pricing_scheme,omitempty"`
	Prices         []Price         `json:"prices,omitempty"`
	OveragePricing *OveragePricing `json:"overage_pricing,omitempty"`
}
