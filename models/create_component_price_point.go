/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CreateComponentPricePoint represents a CreateComponentPricePoint struct.
type CreateComponentPricePoint struct {
    Name                 string                 `json:"name"`
    Handle               *string                `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        PricingScheme          `json:"pricing_scheme"`
    Prices               []Price                `json:"prices"`
    // Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
    UseSiteExchangeRate  *bool                  `json:"use_site_exchange_rate,omitempty"`
    // Whether or not the price point includes tax
    TaxIncluded          *bool                  `json:"tax_included,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval             *int                   `json:"interval,omitempty"`
    // A string representing the interval unit for this price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit         Optional[IntervalUnit] `json:"interval_unit"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreateComponentPricePoint.
// It customizes the JSON marshaling process for CreateComponentPricePoint objects.
func (c CreateComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateComponentPricePoint object to a map representation for JSON marshaling.
func (c CreateComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    if c.Handle != nil {
        structMap["handle"] = c.Handle
    }
    structMap["pricing_scheme"] = c.PricingScheme
    structMap["prices"] = c.Prices
    if c.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
    }
    if c.TaxIncluded != nil {
        structMap["tax_included"] = c.TaxIncluded
    }
    if c.Interval != nil {
        structMap["interval"] = c.Interval
    }
    if c.IntervalUnit.IsValueSet() {
        if c.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = c.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateComponentPricePoint.
// It customizes the JSON unmarshaling process for CreateComponentPricePoint objects.
func (c *CreateComponentPricePoint) UnmarshalJSON(input []byte) error {
    var temp tempCreateComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "name", "handle", "pricing_scheme", "prices", "use_site_exchange_rate", "tax_included", "interval", "interval_unit")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Name = *temp.Name
    c.Handle = temp.Handle
    c.PricingScheme = *temp.PricingScheme
    c.Prices = *temp.Prices
    c.UseSiteExchangeRate = temp.UseSiteExchangeRate
    c.TaxIncluded = temp.TaxIncluded
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    return nil
}

// tempCreateComponentPricePoint is a temporary struct used for validating the fields of CreateComponentPricePoint.
type tempCreateComponentPricePoint  struct {
    Name                *string                `json:"name"`
    Handle              *string                `json:"handle,omitempty"`
    PricingScheme       *PricingScheme         `json:"pricing_scheme"`
    Prices              *[]Price               `json:"prices"`
    UseSiteExchangeRate *bool                  `json:"use_site_exchange_rate,omitempty"`
    TaxIncluded         *bool                  `json:"tax_included,omitempty"`
    Interval            *int                   `json:"interval,omitempty"`
    IntervalUnit        Optional[IntervalUnit] `json:"interval_unit"`
}

func (c *tempCreateComponentPricePoint) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Create Component Price Point`")
    }
    if c.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Create Component Price Point`")
    }
    if c.Prices == nil {
        errs = append(errs, "required field `prices` is missing for type `Create Component Price Point`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
