/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// SubscriptionGroupComponentCustomPrice represents a SubscriptionGroupComponentCustomPrice struct.
// Used in place of `price_point_id` to define a custom price point unique to the subscription. You still need to provide `component_id`.
type SubscriptionGroupComponentCustomPrice struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices               []Price                `json:"prices,omitempty"`
    OveragePricing       []ComponentCustomPrice `json:"overage_pricing,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionGroupComponentCustomPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionGroupComponentCustomPrice) String() string {
    return fmt.Sprintf(
    	"SubscriptionGroupComponentCustomPrice[PricingScheme=%v, Prices=%v, OveragePricing=%v, AdditionalProperties=%v]",
    	s.PricingScheme, s.Prices, s.OveragePricing, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupComponentCustomPrice.
// It customizes the JSON marshaling process for SubscriptionGroupComponentCustomPrice objects.
func (s SubscriptionGroupComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "pricing_scheme", "prices", "overage_pricing"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupComponentCustomPrice object to a map representation for JSON marshaling.
func (s SubscriptionGroupComponentCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSubscriptionGroupComponentCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "pricing_scheme", "prices", "overage_pricing")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PricingScheme = temp.PricingScheme
    s.Prices = temp.Prices
    s.OveragePricing = temp.OveragePricing
    return nil
}

// tempSubscriptionGroupComponentCustomPrice is a temporary struct used for validating the fields of SubscriptionGroupComponentCustomPrice.
type tempSubscriptionGroupComponentCustomPrice  struct {
    PricingScheme  *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices         []Price                `json:"prices,omitempty"`
    OveragePricing []ComponentCustomPrice `json:"overage_pricing,omitempty"`
}
