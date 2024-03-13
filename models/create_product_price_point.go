package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateProductPricePoint represents a CreateProductPricePoint struct.
type CreateProductPricePoint struct {
	// The product price point name
	Name string `json:"name"`
	// The product price point API handle
	Handle *string `json:"handle,omitempty"`
	// The product price point price, in integer cents
	PriceInCents int64 `json:"price_in_cents"`
	// The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this product price point would renew every 30 days
	Interval int `json:"interval"`
	// A string representing the interval unit for this product price point, either month or day
	IntervalUnit IntervalUnit `json:"interval_unit"`
	// The product price point trial price, in integer cents
	TrialPriceInCents *int64 `json:"trial_price_in_cents,omitempty"`
	// The numerical trial interval. i.e. an interval of ‘30’ coupled with a trial_interval_unit of day would mean this product price point trial would last 30 days.
	TrialInterval *int `json:"trial_interval,omitempty"`
	// A string representing the trial interval unit for this product price point, either month or day
	TrialIntervalUnit *IntervalUnit `json:"trial_interval_unit,omitempty"`
	TrialType         *string       `json:"trial_type,omitempty"`
	// The product price point initial charge, in integer cents
	InitialChargeInCents    *int64 `json:"initial_charge_in_cents,omitempty"`
	InitialChargeAfterTrial *bool  `json:"initial_charge_after_trial,omitempty"`
	// The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days.
	ExpirationInterval *int `json:"expiration_interval,omitempty"`
	// A string representing the expiration interval unit for this product price point, either month or day
	ExpirationIntervalUnit *IntervalUnit `json:"expiration_interval_unit,omitempty"`
	// Whether or not to use the site's exchange rate or define your own pricing when your site has multiple currencies defined.
	UseSiteExchangeRate *bool `json:"use_site_exchange_rate,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductPricePoint.
// It customizes the JSON marshaling process for CreateProductPricePoint objects.
func (c *CreateProductPricePoint) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductPricePoint object to a map representation for JSON marshaling.
func (c *CreateProductPricePoint) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["name"] = c.Name
	if c.Handle != nil {
		structMap["handle"] = c.Handle
	}
	structMap["price_in_cents"] = c.PriceInCents
	structMap["interval"] = c.Interval
	structMap["interval_unit"] = c.IntervalUnit
	if c.TrialPriceInCents != nil {
		structMap["trial_price_in_cents"] = c.TrialPriceInCents
	}
	if c.TrialInterval != nil {
		structMap["trial_interval"] = c.TrialInterval
	}
	if c.TrialIntervalUnit != nil {
		structMap["trial_interval_unit"] = c.TrialIntervalUnit
	}
	if c.TrialType != nil {
		structMap["trial_type"] = c.TrialType
	}
	if c.InitialChargeInCents != nil {
		structMap["initial_charge_in_cents"] = c.InitialChargeInCents
	}
	if c.InitialChargeAfterTrial != nil {
		structMap["initial_charge_after_trial"] = c.InitialChargeAfterTrial
	}
	if c.ExpirationInterval != nil {
		structMap["expiration_interval"] = c.ExpirationInterval
	}
	if c.ExpirationIntervalUnit != nil {
		structMap["expiration_interval_unit"] = c.ExpirationIntervalUnit
	}
	if c.UseSiteExchangeRate != nil {
		structMap["use_site_exchange_rate"] = c.UseSiteExchangeRate
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductPricePoint.
// It customizes the JSON unmarshaling process for CreateProductPricePoint objects.
func (c *CreateProductPricePoint) UnmarshalJSON(input []byte) error {
	var temp createProductPricePoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Name = *temp.Name
	c.Handle = temp.Handle
	c.PriceInCents = *temp.PriceInCents
	c.Interval = *temp.Interval
	c.IntervalUnit = *temp.IntervalUnit
	c.TrialPriceInCents = temp.TrialPriceInCents
	c.TrialInterval = temp.TrialInterval
	c.TrialIntervalUnit = temp.TrialIntervalUnit
	c.TrialType = temp.TrialType
	c.InitialChargeInCents = temp.InitialChargeInCents
	c.InitialChargeAfterTrial = temp.InitialChargeAfterTrial
	c.ExpirationInterval = temp.ExpirationInterval
	c.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
	c.UseSiteExchangeRate = temp.UseSiteExchangeRate
	return nil
}

// TODO
type createProductPricePoint struct {
	Name                    *string       `json:"name"`
	Handle                  *string       `json:"handle,omitempty"`
	PriceInCents            *int64        `json:"price_in_cents"`
	Interval                *int          `json:"interval"`
	IntervalUnit            *IntervalUnit `json:"interval_unit"`
	TrialPriceInCents       *int64        `json:"trial_price_in_cents,omitempty"`
	TrialInterval           *int          `json:"trial_interval,omitempty"`
	TrialIntervalUnit       *IntervalUnit `json:"trial_interval_unit,omitempty"`
	TrialType               *string       `json:"trial_type,omitempty"`
	InitialChargeInCents    *int64        `json:"initial_charge_in_cents,omitempty"`
	InitialChargeAfterTrial *bool         `json:"initial_charge_after_trial,omitempty"`
	ExpirationInterval      *int          `json:"expiration_interval,omitempty"`
	ExpirationIntervalUnit  *IntervalUnit `json:"expiration_interval_unit,omitempty"`
	UseSiteExchangeRate     *bool         `json:"use_site_exchange_rate,omitempty"`
}

func (c *createProductPricePoint) validate() error {
	var errs []string
	if c.Name == nil {
		errs = append(errs, "required field `name` is missing for type `Create Product Price Point`")
	}
	if c.PriceInCents == nil {
		errs = append(errs, "required field `price_in_cents` is missing for type `Create Product Price Point`")
	}
	if c.Interval == nil {
		errs = append(errs, "required field `interval` is missing for type `Create Product Price Point`")
	}
	if c.IntervalUnit == nil {
		errs = append(errs, "required field `interval_unit` is missing for type `Create Product Price Point`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
