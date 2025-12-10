// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ComponentCustomPrice represents a ComponentCustomPrice struct.
// Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
type ComponentCustomPrice struct {
    // Whether or not the price point includes tax
    TaxIncluded              *bool                            `json:"tax_included,omitempty"`
    // Omit for On/Off components
    PricingScheme            *PricingScheme                   `json:"pricing_scheme,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval                 *int                             `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit             Optional[IntervalUnit]           `json:"interval_unit"`
    // On/off components only need one price bracket starting at 1
    Prices                   []Price                          `json:"prices"`
    // Applicable only to prepaid usage components. Controls whether the allocated quantity renews each period.
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    // Applicable only to prepaid usage components. Controls whether remaining units roll over to the next period.
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    // Applicable only when rollover is enabled. Number of `expiration_interval_unit`s after which rollover amounts expire.
    ExpirationInterval       Optional[int]                    `json:"expiration_interval"`
    // Applicable only when rollover is enabled. Interval unit for rollover expiration (month or day).
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
    AdditionalProperties     map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentCustomPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentCustomPrice) String() string {
    return fmt.Sprintf(
    	"ComponentCustomPrice[TaxIncluded=%v, PricingScheme=%v, Interval=%v, IntervalUnit=%v, Prices=%v, RenewPrepaidAllocation=%v, RolloverPrepaidRemainder=%v, ExpirationInterval=%v, ExpirationIntervalUnit=%v, AdditionalProperties=%v]",
    	c.TaxIncluded, c.PricingScheme, c.Interval, c.IntervalUnit, c.Prices, c.RenewPrepaidAllocation, c.RolloverPrepaidRemainder, c.ExpirationInterval, c.ExpirationIntervalUnit, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentCustomPrice.
// It customizes the JSON marshaling process for ComponentCustomPrice objects.
func (c ComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "tax_included", "pricing_scheme", "interval", "interval_unit", "prices", "renew_prepaid_allocation", "rollover_prepaid_remainder", "expiration_interval", "expiration_interval_unit"); err != nil {
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
    if c.RenewPrepaidAllocation != nil {
        structMap["renew_prepaid_allocation"] = c.RenewPrepaidAllocation
    }
    if c.RolloverPrepaidRemainder != nil {
        structMap["rollover_prepaid_remainder"] = c.RolloverPrepaidRemainder
    }
    if c.ExpirationInterval.IsValueSet() {
        if c.ExpirationInterval.Value() != nil {
            structMap["expiration_interval"] = c.ExpirationInterval.Value()
        } else {
            structMap["expiration_interval"] = nil
        }
    }
    if c.ExpirationIntervalUnit.IsValueSet() {
        if c.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = c.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tax_included", "pricing_scheme", "interval", "interval_unit", "prices", "renew_prepaid_allocation", "rollover_prepaid_remainder", "expiration_interval", "expiration_interval_unit")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.TaxIncluded = temp.TaxIncluded
    c.PricingScheme = temp.PricingScheme
    c.Interval = temp.Interval
    c.IntervalUnit = temp.IntervalUnit
    c.Prices = *temp.Prices
    c.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
    c.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
    c.ExpirationInterval = temp.ExpirationInterval
    c.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    return nil
}

// tempComponentCustomPrice is a temporary struct used for validating the fields of ComponentCustomPrice.
type tempComponentCustomPrice  struct {
    TaxIncluded              *bool                            `json:"tax_included,omitempty"`
    PricingScheme            *PricingScheme                   `json:"pricing_scheme,omitempty"`
    Interval                 *int                             `json:"interval,omitempty"`
    IntervalUnit             Optional[IntervalUnit]           `json:"interval_unit"`
    Prices                   *[]Price                         `json:"prices"`
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    ExpirationInterval       Optional[int]                    `json:"expiration_interval"`
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
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
