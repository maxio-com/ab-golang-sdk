package models

import (
	"encoding/json"
)

// SubscriptionCustomPrice represents a SubscriptionCustomPrice struct.
// (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
type SubscriptionCustomPrice struct {
	// (Optional)
	Name *string `json:"name,omitempty"`
	// (Optional)
	Handle *string `json:"handle,omitempty"`
	// Required if using `custom_price` attribute.
	PriceInCents interface{} `json:"price_in_cents"`
	// Required if using `custom_price` attribute.
	Interval interface{} `json:"interval"`
	// Required if using `custom_price` attribute.
	IntervalUnit IntervalUnitEnum `json:"interval_unit"`
	// (Optional)
	TrialPriceInCents *interface{} `json:"trial_price_in_cents,omitempty"`
	// (Optional)
	TrialInterval *interface{} `json:"trial_interval,omitempty"`
	// (Optional)
	TrialIntervalUnit *IntervalUnitEnum `json:"trial_interval_unit,omitempty"`
	// (Optional)
	InitialChargeInCents *interface{} `json:"initial_charge_in_cents,omitempty"`
	// (Optional)
	InitialChargeAfterTrial *bool `json:"initial_charge_after_trial,omitempty"`
	// (Optional)
	ExpirationInterval *interface{} `json:"expiration_interval,omitempty"`
	// (Optional)
	ExpirationIntervalUnit *IntervalUnitEnum `json:"expiration_interval_unit,omitempty"`
	// (Optional)
	TaxIncluded *bool `json:"tax_included,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPrice.
// It customizes the JSON marshaling process for SubscriptionCustomPrice objects.
func (s *SubscriptionCustomPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPrice object to a map representation for JSON marshaling.
func (s *SubscriptionCustomPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Handle != nil {
		structMap["handle"] = s.Handle
	}
	structMap["price_in_cents"] = s.PriceInCents
	structMap["interval"] = s.Interval
	structMap["interval_unit"] = s.IntervalUnit
	if s.TrialPriceInCents != nil {
		structMap["trial_price_in_cents"] = s.TrialPriceInCents
	}
	if s.TrialInterval != nil {
		structMap["trial_interval"] = s.TrialInterval
	}
	if s.TrialIntervalUnit != nil {
		structMap["trial_interval_unit"] = s.TrialIntervalUnit
	}
	if s.InitialChargeInCents != nil {
		structMap["initial_charge_in_cents"] = s.InitialChargeInCents
	}
	if s.InitialChargeAfterTrial != nil {
		structMap["initial_charge_after_trial"] = s.InitialChargeAfterTrial
	}
	if s.ExpirationInterval != nil {
		structMap["expiration_interval"] = s.ExpirationInterval
	}
	if s.ExpirationIntervalUnit != nil {
		structMap["expiration_interval_unit"] = s.ExpirationIntervalUnit
	}
	if s.TaxIncluded != nil {
		structMap["tax_included"] = s.TaxIncluded
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPrice.
// It customizes the JSON unmarshaling process for SubscriptionCustomPrice objects.
func (s *SubscriptionCustomPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Name                    *string           `json:"name,omitempty"`
		Handle                  *string           `json:"handle,omitempty"`
		PriceInCents            interface{}       `json:"price_in_cents"`
		Interval                interface{}       `json:"interval"`
		IntervalUnit            IntervalUnitEnum  `json:"interval_unit"`
		TrialPriceInCents       *interface{}      `json:"trial_price_in_cents,omitempty"`
		TrialInterval           *interface{}      `json:"trial_interval,omitempty"`
		TrialIntervalUnit       *IntervalUnitEnum `json:"trial_interval_unit,omitempty"`
		InitialChargeInCents    *interface{}      `json:"initial_charge_in_cents,omitempty"`
		InitialChargeAfterTrial *bool             `json:"initial_charge_after_trial,omitempty"`
		ExpirationInterval      *interface{}      `json:"expiration_interval,omitempty"`
		ExpirationIntervalUnit  *IntervalUnitEnum `json:"expiration_interval_unit,omitempty"`
		TaxIncluded             *bool             `json:"tax_included,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Name = temp.Name
	s.Handle = temp.Handle
	s.PriceInCents = temp.PriceInCents
	s.Interval = temp.Interval
	s.IntervalUnit = temp.IntervalUnit
	s.TrialPriceInCents = temp.TrialPriceInCents
	s.TrialInterval = temp.TrialInterval
	s.TrialIntervalUnit = temp.TrialIntervalUnit
	s.InitialChargeInCents = temp.InitialChargeInCents
	s.InitialChargeAfterTrial = temp.InitialChargeAfterTrial
	s.ExpirationInterval = temp.ExpirationInterval
	s.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
	s.TaxIncluded = temp.TaxIncluded
	return nil
}
