/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// PrepaidComponentPricePoint represents a PrepaidComponentPricePoint struct.
type PrepaidComponentPricePoint struct {
    Name                 *string         `json:"name,omitempty"`
    Handle               *string         `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        *PricingScheme  `json:"pricing_scheme,omitempty"`
    Prices               []Price         `json:"prices,omitempty"`
    OveragePricing       *OveragePricing `json:"overage_pricing,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidComponentPricePoint.
// It customizes the JSON marshaling process for PrepaidComponentPricePoint objects.
func (p PrepaidComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidComponentPricePoint object to a map representation for JSON marshaling.
func (p PrepaidComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp tempPrepaidComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "handle", "pricing_scheme", "prices", "overage_pricing")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Name = temp.Name
    p.Handle = temp.Handle
    p.PricingScheme = temp.PricingScheme
    p.Prices = temp.Prices
    p.OveragePricing = temp.OveragePricing
    return nil
}

// tempPrepaidComponentPricePoint is a temporary struct used for validating the fields of PrepaidComponentPricePoint.
type tempPrepaidComponentPricePoint  struct {
    Name           *string         `json:"name,omitempty"`
    Handle         *string         `json:"handle,omitempty"`
    PricingScheme  *PricingScheme  `json:"pricing_scheme,omitempty"`
    Prices         []Price         `json:"prices,omitempty"`
    OveragePricing *OveragePricing `json:"overage_pricing,omitempty"`
}
