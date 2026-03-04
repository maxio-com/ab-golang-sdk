// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalProductPricePoint represents a ScheduledRenewalProductPricePoint struct.
// Custom pricing for a product within a scheduled renewal.
type ScheduledRenewalProductPricePoint struct {
    // (Optional)
    Name                   *string                                       `json:"name,omitempty"`
    // (Optional)
    Handle                 *string                                       `json:"handle,omitempty"`
    // Required if using `custom_price` attribute.
    PriceInCents           ScheduledRenewalProductPricePointPriceInCents `json:"price_in_cents"`
    // Required if using `custom_price` attribute.
    Interval               ScheduledRenewalProductPricePointInterval     `json:"interval"`
    // Required if using `custom_price` attribute.
    IntervalUnit           *IntervalUnit                                 `json:"interval_unit"`
    // (Optional)
    TaxIncluded            *bool                                         `json:"tax_included,omitempty"`
    // The product price point initial charge, in integer cents.
    InitialChargeInCents   *int64                                        `json:"initial_charge_in_cents,omitempty"`
    // The numerical expiration interval. i.e. an expiration_interval of ‘30’ coupled with an expiration_interval_unit of day would mean this product price point would expire after 30 days.
    ExpirationInterval     *int                                          `json:"expiration_interval,omitempty"`
    // A string representing the expiration interval unit for this product price point, either month, day or never
    ExpirationIntervalUnit Optional[ExpirationIntervalUnit]              `json:"expiration_interval_unit"`
    AdditionalProperties   map[string]interface{}                        `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalProductPricePoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalProductPricePoint) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalProductPricePoint[Name=%v, Handle=%v, PriceInCents=%v, Interval=%v, IntervalUnit=%v, TaxIncluded=%v, InitialChargeInCents=%v, ExpirationInterval=%v, ExpirationIntervalUnit=%v, AdditionalProperties=%v]",
    	s.Name, s.Handle, s.PriceInCents, s.Interval, s.IntervalUnit, s.TaxIncluded, s.InitialChargeInCents, s.ExpirationInterval, s.ExpirationIntervalUnit, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalProductPricePoint.
// It customizes the JSON marshaling process for ScheduledRenewalProductPricePoint objects.
func (s ScheduledRenewalProductPricePoint) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "name", "handle", "price_in_cents", "interval", "interval_unit", "tax_included", "initial_charge_in_cents", "expiration_interval", "expiration_interval_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalProductPricePoint object to a map representation for JSON marshaling.
func (s ScheduledRenewalProductPricePoint) toMap() map[string]any {
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
    if s.TaxIncluded != nil {
        structMap["tax_included"] = s.TaxIncluded
    }
    if s.InitialChargeInCents != nil {
        structMap["initial_charge_in_cents"] = s.InitialChargeInCents
    }
    if s.ExpirationInterval != nil {
        structMap["expiration_interval"] = s.ExpirationInterval
    }
    if s.ExpirationIntervalUnit.IsValueSet() {
        if s.ExpirationIntervalUnit.Value() != nil {
            structMap["expiration_interval_unit"] = s.ExpirationIntervalUnit.Value()
        } else {
            structMap["expiration_interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalProductPricePoint.
// It customizes the JSON unmarshaling process for ScheduledRenewalProductPricePoint objects.
func (s *ScheduledRenewalProductPricePoint) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalProductPricePoint
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "name", "handle", "price_in_cents", "interval", "interval_unit", "tax_included", "initial_charge_in_cents", "expiration_interval", "expiration_interval_unit")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Name = temp.Name
    s.Handle = temp.Handle
    s.PriceInCents = *temp.PriceInCents
    s.Interval = *temp.Interval
    s.IntervalUnit = temp.IntervalUnit
    s.TaxIncluded = temp.TaxIncluded
    s.InitialChargeInCents = temp.InitialChargeInCents
    s.ExpirationInterval = temp.ExpirationInterval
    s.ExpirationIntervalUnit = temp.ExpirationIntervalUnit
    return nil
}

// tempScheduledRenewalProductPricePoint is a temporary struct used for validating the fields of ScheduledRenewalProductPricePoint.
type tempScheduledRenewalProductPricePoint  struct {
    Name                   *string                                        `json:"name,omitempty"`
    Handle                 *string                                        `json:"handle,omitempty"`
    PriceInCents           *ScheduledRenewalProductPricePointPriceInCents `json:"price_in_cents"`
    Interval               *ScheduledRenewalProductPricePointInterval     `json:"interval"`
    IntervalUnit           *IntervalUnit                                  `json:"interval_unit"`
    TaxIncluded            *bool                                          `json:"tax_included,omitempty"`
    InitialChargeInCents   *int64                                         `json:"initial_charge_in_cents,omitempty"`
    ExpirationInterval     *int                                           `json:"expiration_interval,omitempty"`
    ExpirationIntervalUnit Optional[ExpirationIntervalUnit]               `json:"expiration_interval_unit"`
}

func (s *tempScheduledRenewalProductPricePoint) validate() error {
    var errs []string
    if s.PriceInCents == nil {
        errs = append(errs, "required field `price_in_cents` is missing for type `Scheduled Renewal Product Price Point`")
    }
    if s.Interval == nil {
        errs = append(errs, "required field `interval` is missing for type `Scheduled Renewal Product Price Point`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
