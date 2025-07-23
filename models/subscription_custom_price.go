// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SubscriptionCustomPrice represents a SubscriptionCustomPrice struct.
// (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
type SubscriptionCustomPrice struct {
    // (Optional)
    Name                    *string                                      `json:"name,omitempty"`
    // (Optional)
    Handle                  *string                                      `json:"handle,omitempty"`
    // Required if using `custom_price` attribute.
    PriceInCents            SubscriptionCustomPricePriceInCents          `json:"price_in_cents"`
    // Required if using `custom_price` attribute.
    Interval                SubscriptionCustomPriceInterval              `json:"interval"`
    // Required if using `custom_price` attribute.
    IntervalUnit            *IntervalUnit                                `json:"interval_unit"`
    // (Optional)
    TrialPriceInCents       *SubscriptionCustomPriceTrialPriceInCents    `json:"trial_price_in_cents,omitempty"`
    // (Optional)
    TrialInterval           *SubscriptionCustomPriceTrialInterval        `json:"trial_interval,omitempty"`
    // (Optional)
    TrialIntervalUnit       *IntervalUnit                                `json:"trial_interval_unit,omitempty"`
    // (Optional)
    InitialChargeInCents    *SubscriptionCustomPriceInitialChargeInCents `json:"initial_charge_in_cents,omitempty"`
    // (Optional)
    InitialChargeAfterTrial *bool                                        `json:"initial_charge_after_trial,omitempty"`
    // (Optional)
    ExpirationInterval      *SubscriptionCustomPriceExpirationInterval   `json:"expiration_interval,omitempty"`
    // (Optional)
    ExpirationIntervalUnit  Optional[ExpirationIntervalUnit]             `json:"expiration_interval_unit"`
    // (Optional)
    TaxIncluded             *bool                                        `json:"tax_included,omitempty"`
    AdditionalProperties    map[string]interface{}                       `json:"_"`
}

// String implements the fmt.Stringer interface for SubscriptionCustomPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SubscriptionCustomPrice) String() string {
    return fmt.Sprintf(
    	"SubscriptionCustomPrice[Name=%v, Handle=%v, PriceInCents=%v, Interval=%v, IntervalUnit=%v, TrialPriceInCents=%v, TrialInterval=%v, TrialIntervalUnit=%v, InitialChargeInCents=%v, InitialChargeAfterTrial=%v, ExpirationInterval=%v, ExpirationIntervalUnit=%v, TaxIncluded=%v, AdditionalProperties=%v]",
    	s.Name, s.Handle, s.PriceInCents, s.Interval, s.IntervalUnit, s.TrialPriceInCents, s.TrialInterval, s.TrialIntervalUnit, s.InitialChargeInCents, s.InitialChargeAfterTrial, s.ExpirationInterval, s.ExpirationIntervalUnit, s.TaxIncluded, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionCustomPrice.
// It customizes the JSON marshaling process for SubscriptionCustomPrice objects.
func (s SubscriptionCustomPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "handle", "price_in_cents", "interval", "interval_unit", "trial_price_in_cents", "trial_interval", "trial_interval_unit", "initial_charge_in_cents", "initial_charge_after_trial", "expiration_interval", "expiration_interval_unit", "tax_included"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionCustomPrice object to a map representation for JSON marshaling.
func (s SubscriptionCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Handle != nil {
        structMap["handle"] = s.Handle
    }
    structMap["price_in_cents"] = s.PriceInCents.toMap()
    structMap["interval"] = s.Interval.toMap()
    if s.IntervalUnit != nil {
        structMap["interval_unit"] = s.IntervalUnit
    } else {
        structMap["interval_unit"] = nil
    }
    if s.TrialPriceInCents != nil {
        structMap["trial_price_in_cents"] = s.TrialPriceInCents.toMap()
    }
    if s.TrialInterval != nil {
        structMap["trial_interval"] = s.TrialInterval.toMap()
    }
    if s.TrialIntervalUnit != nil {
        structMap["trial_interval_unit"] = s.TrialIntervalUnit
    }
    if s.InitialChargeInCents != nil {
        structMap["initial_charge_in_cents"] = s.InitialChargeInCents.toMap()
    }
    if s.InitialChargeAfterTrial != nil {
        structMap["initial_charge_after_trial"] = s.InitialChargeAfterTrial
    }
    if s.ExpirationInterval != nil {
        structMap["expiration_interval"] = s.ExpirationInterval.toMap()
    }
    if s.ExpirationIntervalUnit.IsValueSet() {
        if s.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = s.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    if s.TaxIncluded != nil {
        structMap["tax_included"] = s.TaxIncluded
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionCustomPrice.
// It customizes the JSON unmarshaling process for SubscriptionCustomPrice objects.
func (s *SubscriptionCustomPrice) UnmarshalJSON(input []byte) error {
    var temp tempSubscriptionCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle", "price_in_cents", "interval", "interval_unit", "trial_price_in_cents", "trial_interval", "trial_interval_unit", "initial_charge_in_cents", "initial_charge_after_trial", "expiration_interval", "expiration_interval_unit", "tax_included")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Name = temp.Name
    s.Handle = temp.Handle
    s.PriceInCents = *temp.PriceInCents
    s.Interval = *temp.Interval
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

// tempSubscriptionCustomPrice is a temporary struct used for validating the fields of SubscriptionCustomPrice.
type tempSubscriptionCustomPrice  struct {
    Name                    *string                                      `json:"name,omitempty"`
    Handle                  *string                                      `json:"handle,omitempty"`
    PriceInCents            *SubscriptionCustomPricePriceInCents         `json:"price_in_cents"`
    Interval                *SubscriptionCustomPriceInterval             `json:"interval"`
    IntervalUnit            *IntervalUnit                                `json:"interval_unit"`
    TrialPriceInCents       *SubscriptionCustomPriceTrialPriceInCents    `json:"trial_price_in_cents,omitempty"`
    TrialInterval           *SubscriptionCustomPriceTrialInterval        `json:"trial_interval,omitempty"`
    TrialIntervalUnit       *IntervalUnit                                `json:"trial_interval_unit,omitempty"`
    InitialChargeInCents    *SubscriptionCustomPriceInitialChargeInCents `json:"initial_charge_in_cents,omitempty"`
    InitialChargeAfterTrial *bool                                        `json:"initial_charge_after_trial,omitempty"`
    ExpirationInterval      *SubscriptionCustomPriceExpirationInterval   `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit  Optional[ExpirationIntervalUnit]             `json:"expiration_interval_unit"`
    TaxIncluded             *bool                                        `json:"tax_included,omitempty"`
}

func (s *tempSubscriptionCustomPrice) validate() error {
    var errs []string
    if s.PriceInCents == nil {
        errs = append(errs, "required field `price_in_cents` is missing for type `Subscription Custom Price`")
    }
    if s.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `Subscription Custom Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
