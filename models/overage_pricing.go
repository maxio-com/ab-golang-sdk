package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OveragePricing represents a OveragePricing struct.
type OveragePricing struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        PricingScheme  `json:"pricing_scheme"`
    Prices               []Price        `json:"prices,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OveragePricing.
// It customizes the JSON marshaling process for OveragePricing objects.
func (o OveragePricing) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OveragePricing object to a map representation for JSON marshaling.
func (o OveragePricing) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["pricing_scheme"] = o.PricingScheme
    if o.Prices != nil {
        structMap["prices"] = o.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OveragePricing.
// It customizes the JSON unmarshaling process for OveragePricing objects.
func (o *OveragePricing) UnmarshalJSON(input []byte) error {
    var temp overagePricing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.PricingScheme = *temp.PricingScheme
    o.Prices = temp.Prices
    return nil
}

// TODO
type overagePricing  struct {
    PricingScheme *PricingScheme `json:"pricing_scheme"`
    Prices        []Price        `json:"prices,omitempty"`
}

func (o *overagePricing) validate() error {
    var errs []string
    if o.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Overage Pricing`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
