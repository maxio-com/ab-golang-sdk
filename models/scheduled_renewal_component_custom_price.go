// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ScheduledRenewalComponentCustomPrice represents a ScheduledRenewalComponentCustomPrice struct.
// Custom pricing for a component within a scheduled renewal.
type ScheduledRenewalComponentCustomPrice struct {
    // Whether or not the price point includes tax
    TaxIncluded          *bool                  `json:"tax_included,omitempty"`
    // Omit for On/Off components
    PricingScheme        PricingScheme          `json:"pricing_scheme"`
    // On/off components only need one price bracket starting at 1.
    Prices               []Price                `json:"prices"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalComponentCustomPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalComponentCustomPrice) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalComponentCustomPrice[TaxIncluded=%v, PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
    	s.TaxIncluded, s.PricingScheme, s.Prices, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalComponentCustomPrice.
// It customizes the JSON marshaling process for ScheduledRenewalComponentCustomPrice objects.
func (s ScheduledRenewalComponentCustomPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "tax_included", "pricing_scheme", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalComponentCustomPrice object to a map representation for JSON marshaling.
func (s ScheduledRenewalComponentCustomPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.TaxIncluded != nil {
        structMap["tax_included"] = s.TaxIncluded
    }
    structMap["pricing_scheme"] = s.PricingScheme
    structMap["prices"] = s.Prices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalComponentCustomPrice.
// It customizes the JSON unmarshaling process for ScheduledRenewalComponentCustomPrice objects.
func (s *ScheduledRenewalComponentCustomPrice) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalComponentCustomPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "tax_included", "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.TaxIncluded = temp.TaxIncluded
    s.PricingScheme = *temp.PricingScheme
    s.Prices = *temp.Prices
    return nil
}

// tempScheduledRenewalComponentCustomPrice is a temporary struct used for validating the fields of ScheduledRenewalComponentCustomPrice.
type tempScheduledRenewalComponentCustomPrice  struct {
    TaxIncluded   *bool          `json:"tax_included,omitempty"`
    PricingScheme *PricingScheme `json:"pricing_scheme"`
    Prices        *[]Price       `json:"prices"`
}

func (s *tempScheduledRenewalComponentCustomPrice) validate() error {
    var errs []string
    if s.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Scheduled Renewal Component Custom Price`")
    }
    if s.Prices == nil {
        errs = append(errs, "required field `prices` is missing for type `Scheduled Renewal Component Custom Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
