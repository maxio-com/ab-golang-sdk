package models

import (
    "encoding/json"
)

// ComponentCustomPrice represents a ComponentCustomPrice struct.
// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
type ComponentCustomPrice struct {
    // Omit for On/Off components
    PricingScheme        *PricingScheme `json:"pricing_scheme,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval             *int           `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit         *IntervalUnit  `json:"interval_unit,omitempty"`
    // On/off components only need one price bracket starting at 1
    Prices               []Price        `json:"prices,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCustomPrice.
// It customizes the JSON marshaling process for ComponentCustomPrice objects.
func (c ComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentCustomPrice object to a map representation for JSON marshaling.
func (c ComponentCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.PricingScheme != nil {
        structMap["pricing_scheme"] = c.PricingScheme
    }
    if c.Interval != nil {
        structMap["interval"] = c.Interval
    }
    if c.IntervalUnit != nil {
        structMap["interval_unit"] = c.IntervalUnit
    }
    if c.Prices != nil {
        structMap["prices"] = c.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCustomPrice.
// It customizes the JSON unmarshaling process for ComponentCustomPrice objects.
func (c *ComponentCustomPrice) UnmarshalJSON(input []byte) error {
    var temp componentCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "pricing_scheme", "interval", "interval_unit", "prices")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.PricingScheme = temp.PricingScheme
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.Prices = temp.Prices
    return nil
}

// componentCustomPrice is a temporary struct used for validating the fields of ComponentCustomPrice.
type componentCustomPrice  struct {
    PricingScheme *PricingScheme `json:"pricing_scheme,omitempty"`
    Interval      *int           `json:"interval,omitempty"`
    IntervalUnit  *IntervalUnit  `json:"interval_unit,omitempty"`
    Prices        []Price        `json:"prices,omitempty"`
}
