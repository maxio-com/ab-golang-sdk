package models

import (
    "encoding/json"
)

// SubscriptionGroupComponentCustomPrice represents a SubscriptionGroupComponentCustomPrice struct.
// Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`.
type SubscriptionGroupComponentCustomPrice struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices               []Price                `json:"prices,omitempty"`
    OveragePricing       []ComponentCustomPrice `json:"overage_pricing,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupComponentCustomPrice.
// It customizes the JSON marshaling process for SubscriptionGroupComponentCustomPrice objects.
func (s SubscriptionGroupComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupComponentCustomPrice object to a map representation for JSON marshaling.
func (s SubscriptionGroupComponentCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.PricingScheme != nil {
        structMap["pricing_scheme"] = s.PricingScheme
    }
    if s.Prices != nil {
        structMap["prices"] = s.Prices
    }
    if s.OveragePricing != nil {
        structMap["overage_pricing"] = s.OveragePricing
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupComponentCustomPrice.
// It customizes the JSON unmarshaling process for SubscriptionGroupComponentCustomPrice objects.
func (s *SubscriptionGroupComponentCustomPrice) UnmarshalJSON(input []byte) error {
    var temp subscriptionGroupComponentCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "pricing_scheme", "prices", "overage_pricing")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.PricingScheme = temp.PricingScheme
    s.Prices = temp.Prices
    s.OveragePricing = temp.OveragePricing
    return nil
}

// subscriptionGroupComponentCustomPrice is a temporary struct used for validating the fields of SubscriptionGroupComponentCustomPrice.
type subscriptionGroupComponentCustomPrice  struct {
    PricingScheme  *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices         []Price                `json:"prices,omitempty"`
    OveragePricing []ComponentCustomPrice `json:"overage_pricing,omitempty"`
}
