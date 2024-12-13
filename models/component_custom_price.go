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

// ComponentCustomPrice represents a ComponentCustomPrice struct.
// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
type ComponentCustomPrice struct {
    // Whether or not the price point includes tax
    TaxIncluded          *bool                  `json:"tax_included,omitempty"`
    // Omit for On/Off components
    PricingScheme        *PricingScheme         `json:"pricing_scheme,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval             *int                   `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit         Optional[IntervalUnit] `json:"interval_unit"`
    // On/off components only need one price bracket starting at 1
    Prices               []Price                `json:"prices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCustomPrice.
// It customizes the JSON marshaling process for ComponentCustomPrice objects.
func (c ComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "tax_included", "pricing_scheme", "interval", "interval_unit", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentCustomPrice object to a map representation for JSON marshaling.
func (c ComponentCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.TaxIncluded != nil {
        structMap["tax_included"] = c.TaxIncluded
    }
    if c.PricingScheme != nil {
        structMap["pricing_scheme"] = c.PricingScheme
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
    structMap["prices"] = c.Prices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCustomPrice.
// It customizes the JSON unmarshaling process for ComponentCustomPrice objects.
func (c *ComponentCustomPrice) UnmarshalJSON(input []byte) error {
    var temp tempComponentCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tax_included", "pricing_scheme", "interval", "interval_unit", "prices")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.TaxIncluded = temp.TaxIncluded
    c.PricingScheme = temp.PricingScheme
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.Prices = *temp.Prices
    return nil
}

// tempComponentCustomPrice is a temporary struct used for validating the fields of ComponentCustomPrice.
type tempComponentCustomPrice  struct {
    TaxIncluded   *bool                  `json:"tax_included,omitempty"`
    PricingScheme *PricingScheme         `json:"pricing_scheme,omitempty"`
    Interval      *int                   `json:"interval,omitempty"`
    IntervalUnit  Optional[IntervalUnit] `json:"interval_unit"`
    Prices        *[]Price               `json:"prices"`
}

func (c *tempComponentCustomPrice) validate() error {
    var errs []string
    if c.Prices == nil {
        errs = append(errs, "required field `prices` is missing for type `Component Custom Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
