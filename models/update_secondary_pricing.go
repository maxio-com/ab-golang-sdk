// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// UpdateSecondaryPricing represents a UpdateSecondaryPricing struct.
// Hybrid Pricing only. Changes the secondary pricing model of an existing hybrid price point. Send `pricing_scheme` (or the top-level `secondary_pricing_scheme`) every time you send this block; a block without a scheme returns `422`. Brackets sent without an `id` are added to the existing ones rather than replacing them, so to change a bracket set, update existing brackets by `id` and remove the ones you no longer want with `_destroy: true`. See [Hybrid Pricing](page:introduction/basic-concepts/hybrid-pricing).
type UpdateSecondaryPricing struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    Prices               []UpdatePrice          `json:"prices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateSecondaryPricing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateSecondaryPricing) String() string {
    return fmt.Sprintf(
    	"UpdateSecondaryPricing[PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
    	u.PricingScheme, u.Prices, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateSecondaryPricing.
// It customizes the JSON marshaling process for UpdateSecondaryPricing objects.
func (u UpdateSecondaryPricing) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "pricing_scheme", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSecondaryPricing object to a map representation for JSON marshaling.
func (u UpdateSecondaryPricing) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.PricingScheme != nil {
        structMap["pricing_scheme"] = u.PricingScheme
    }
    if u.Prices != nil {
        structMap["prices"] = u.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSecondaryPricing.
// It customizes the JSON unmarshaling process for UpdateSecondaryPricing objects.
func (u *UpdateSecondaryPricing) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSecondaryPricing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.PricingScheme = temp.PricingScheme
    u.Prices = temp.Prices
    return nil
}

// tempUpdateSecondaryPricing is a temporary struct used for validating the fields of UpdateSecondaryPricing.
type tempUpdateSecondaryPricing  struct {
    PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
    Prices        []UpdatePrice  `json:"prices,omitempty"`
}
