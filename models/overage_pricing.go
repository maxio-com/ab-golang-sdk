package models

import (
    "encoding/json"
)

// OveragePricing represents a OveragePricing struct.
type OveragePricing struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme PricingScheme `json:"pricing_scheme"`
    Prices        []Price       `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OveragePricing.
// It customizes the JSON marshaling process for OveragePricing objects.
func (o *OveragePricing) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OveragePricing object to a map representation for JSON marshaling.
func (o *OveragePricing) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["pricing_scheme"] = o.PricingScheme
    if o.Prices != nil {
        structMap["prices"] = o.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OveragePricing.
// It customizes the JSON unmarshaling process for OveragePricing objects.
func (o *OveragePricing) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PricingScheme PricingScheme `json:"pricing_scheme"`
        Prices        []Price       `json:"prices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    o.PricingScheme = temp.PricingScheme
    o.Prices = temp.Prices
    return nil
}
