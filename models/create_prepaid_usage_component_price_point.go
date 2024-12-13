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

// CreatePrepaidUsageComponentPricePoint represents a CreatePrepaidUsageComponentPricePoint struct.
type CreatePrepaidUsageComponentPricePoint struct {
    Name                     string                           `json:"name"`
    Handle                   *string                          `json:"handle,omitempty"`
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme            PricingScheme                    `json:"pricing_scheme"`
    Prices                   []Price                          `json:"prices"`
    OveragePricing           OveragePricing                   `json:"overage_pricing"`
    // Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
    UseSiteExchangeRate      *bool                            `json:"use_site_exchange_rate,omitempty"`
    // (only for prepaid usage components) Boolean which controls whether or not remaining units should be rolled over to the next period
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    // (only for prepaid usage components) Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    // (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire
    ExpirationInterval       *float64                         `json:"expiration_interval,omitempty"`
    // (only for prepaid usage components where rollover_prepaid_remainder is true) A string representing the expiration interval unit for this component, either month or day
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
    AdditionalProperties     map[string]interface{}           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CreatePrepaidUsageComponentPricePoint.
// It customizes the JSON marshaling process for CreatePrepaidUsageComponentPricePoint objects.
func (c CreatePrepaidUsageComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "name", "handle", "pricing_scheme", "prices", "overage_pricing", "use_site_exchange_rate", "rollover_prepaid_remainder", "renew_prepaid_allocation", "expiration_interval", "expiration_interval_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreatePrepaidUsageComponentPricePoint object to a map representation for JSON marshaling.
func (c CreatePrepaidUsageComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["name"] = c.Name
    if c.Handle != nil {
        structMap["handle"] = c.Handle
    }
    structMap["pricing_scheme"] = c.PricingScheme
    structMap["prices"] = c.Prices
    structMap["overage_pricing"] = c.OveragePricing.toMap()
    if c.UseSiteExchangeRate != nil {
        structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
    }
    if c.RolloverPrepaidRemainder != nil {
        structMap["rollover_prepaid_remainder"] = c.RolloverPrepaidRemainder
    }
    if c.RenewPrepaidAllocation != nil {
        structMap["renew_prepaid_allocation"] = c.RenewPrepaidAllocation
    }
    if c.ExpirationInterval != nil {
        structMap["expiration_interval"] = c.ExpirationInterval
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

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaidUsageComponentPricePoint.
// It customizes the JSON unmarshaling process for CreatePrepaidUsageComponentPricePoint objects.
func (c *CreatePrepaidUsageComponentPricePoint) UnmarshalJSON(input []byte) error {
    var temp tempCreatePrepaidUsageComponentPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle", "pricing_scheme", "prices", "overage_pricing", "use_site_exchange_rate", "rollover_prepaid_remainder", "renew_prepaid_allocation", "expiration_interval", "expiration_interval_unit")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Name = *temp.Name
    c.Handle = temp.Handle
    c.PricingScheme = *temp.PricingScheme
    c.Prices = *temp.Prices
    c.OveragePricing = *temp.OveragePricing
    c.UseSiteExchangeRate = temp.UseSiteExchangeRate
    c.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
    c.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
    c.ExpirationInterval = temp.ExpirationInterval
    c.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    return nil
}

// tempCreatePrepaidUsageComponentPricePoint is a temporary struct used for validating the fields of CreatePrepaidUsageComponentPricePoint.
type tempCreatePrepaidUsageComponentPricePoint  struct {
    Name                     *string                          `json:"name"`
    Handle                   *string                          `json:"handle,omitempty"`
    PricingScheme            *PricingScheme                   `json:"pricing_scheme"`
    Prices                   *[]Price                         `json:"prices"`
    OveragePricing           *OveragePricing                  `json:"overage_pricing"`
    UseSiteExchangeRate      *bool                            `json:"use_site_exchange_rate,omitempty"`
    RolloverPrepaidRemainder *bool                            `json:"rollover_prepaid_remainder,omitempty"`
    RenewPrepaidAllocation   *bool                            `json:"renew_prepaid_allocation,omitempty"`
    ExpirationInterval       *float64                         `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit   Optional[ExpirationIntervalUnit] `json:"expiration_interval_unit"`
}

func (c *tempCreatePrepaidUsageComponentPricePoint) validate() error {
    var errs []string
    if c.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Create Prepaid Usage Component Price Point`")
    }
    if c.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Create Prepaid Usage Component Price Point`")
    }
    if c.Prices == nil {
        errs = append(errs, "required field `prices` is missing for type `Create Prepaid Usage Component Price Point`")
    }
    if c.OveragePricing == nil {
        errs = append(errs, "required field `overage_pricing` is missing for type `Create Prepaid Usage Component Price Point`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
