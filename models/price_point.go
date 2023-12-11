package models

import (
	"encoding/json"
)

// PricePoint represents a PricePoint struct.
type PricePoint struct {
	Name   *string `json:"name,omitempty"`
	Handle *string `json:"handle,omitempty"`
	// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
	PricingScheme *PricingSchemeEnum `json:"pricing_scheme,omitempty"`
	Prices        []Price            `json:"prices,omitempty"`
	// Whether to use the site level exchange rate or define your own prices for each currency if you have multiple currencies defined on the site.
	UseSiteExchangeRate *bool           `json:"use_site_exchange_rate,omitempty"`
	OveragePricing      *OveragePricing `json:"overage_pricing,omitempty"`
	// Boolean which controls whether or not remaining units should be rolled over to the next period
	RolloverPrepaidRemainder *bool `json:"rollover_prepaid_remainder,omitempty"`
	// Boolean which controls whether or not the allocated quantity should be renewed at the beginning of each period
	RenewPrepaidAllocation *bool `json:"renew_prepaid_allocation,omitempty"`
	// (only for prepaid usage components where rollover_prepaid_remainder is true) The number of `expiration_interval_unit`s after which rollover amounts should expire
	ExpirationInterval     *float64          `json:"expiration_interval,omitempty"`
	ExpirationIntervalUnit *IntervalUnitEnum `json:"expiration_interval_unit,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for PricePoint.
// It customizes the JSON marshaling process for PricePoint objects.
func (p *PricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PricePoint object to a map representation for JSON marshaling.
func (p *PricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
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
	if p.UseSiteExchangeRate != nil {
		structMap["use_site_exchange_rate"] = p.UseSiteExchangeRate
	}
	if p.OveragePricing != nil {
		structMap["overage_pricing"] = p.OveragePricing
	}
	if p.RolloverPrepaidRemainder != nil {
		structMap["rollover_prepaid_remainder"] = p.RolloverPrepaidRemainder
	}
	if p.RenewPrepaidAllocation != nil {
		structMap["renew_prepaid_allocation"] = p.RenewPrepaidAllocation
	}
	if p.ExpirationInterval != nil {
		structMap["expiration_interval"] = p.ExpirationInterval
	}
	if p.ExpirationIntervalUnit != nil {
		structMap["expiration_interval_unit"] = p.ExpirationIntervalUnit
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PricePoint.
// It customizes the JSON unmarshaling process for PricePoint objects.
func (p *PricePoint) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                     *string            `json:"name,omitempty"`
		Handle                   *string            `json:"handle,omitempty"`
		PricingScheme            *PricingSchemeEnum `json:"pricing_scheme,omitempty"`
		Prices                   []Price            `json:"prices,omitempty"`
		UseSiteExchangeRate      *bool              `json:"use_site_exchange_rate,omitempty"`
		OveragePricing           *OveragePricing    `json:"overage_pricing,omitempty"`
		RolloverPrepaidRemainder *bool              `json:"rollover_prepaid_remainder,omitempty"`
		RenewPrepaidAllocation   *bool              `json:"renew_prepaid_allocation,omitempty"`
		ExpirationInterval       *float64           `json:"expiration_interval,omitempty"`
		ExpirationIntervalUnit   *IntervalUnitEnum  `json:"expiration_interval_unit,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	p.Name = temp.Name
	p.Handle = temp.Handle
	p.PricingScheme = temp.PricingScheme
	p.Prices = temp.Prices
	p.UseSiteExchangeRate = temp.UseSiteExchangeRate
	p.OveragePricing = temp.OveragePricing
	p.RolloverPrepaidRemainder = temp.RolloverPrepaidRemainder
	p.RenewPrepaidAllocation = temp.RenewPrepaidAllocation
	p.ExpirationInterval = temp.ExpirationInterval
	p.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
	return nil
}
