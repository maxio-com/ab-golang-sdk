// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// SecondaryPricing represents a SecondaryPricing struct.
// Hybrid Pricing only. The secondary pricing model that takes over once usage passes the primary model's included threshold. Applies to Quantity-Based and Metered Components on Sites with Hybrid Pricing enabled; ignored otherwise. Its `prices` brackets start at `1`, not at the primary model's highest `ending_quantity`. See [Hybrid Pricing](page:introduction/basic-concepts/hybrid-pricing).
type SecondaryPricing struct {
    // Required, unless you send the top-level `secondary_pricing_scheme` instead (Create and Update Component Price Point only). A block with neither returns `422` (`Secondary pricing scheme must be defined`).
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices               []Price                `json:"prices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SecondaryPricing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SecondaryPricing) String() string {
    return fmt.Sprintf(
    	"SecondaryPricing[PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
    	s.PricingScheme, s.Prices, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SecondaryPricing.
// It customizes the JSON marshaling process for SecondaryPricing objects.
func (s SecondaryPricing) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "pricing_scheme", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SecondaryPricing object to a map representation for JSON marshaling.
func (s SecondaryPricing) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.PricingScheme != nil {
        structMap["pricing_scheme"] = s.PricingScheme
    }
    if s.Prices != nil {
        structMap["prices"] = s.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SecondaryPricing.
// It customizes the JSON unmarshaling process for SecondaryPricing objects.
func (s *SecondaryPricing) UnmarshalJSON(input []byte) error {
    var temp tempSecondaryPricing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.PricingScheme = temp.PricingScheme
    s.Prices = temp.Prices
    return nil
}

// tempSecondaryPricing is a temporary struct used for validating the fields of SecondaryPricing.
type tempSecondaryPricing  struct {
    PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
    Prices        []Price        `json:"prices,omitempty"`
}
