// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateSegment represents a UpdateSegment struct.
type UpdateSegment struct {
    // The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
    PricingScheme        PricingScheme                `json:"pricing_scheme"`
    Prices               []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateSegment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateSegment) String() string {
    return fmt.Sprintf(
    	"UpdateSegment[PricingScheme=%v, Prices=%v, AdditionalProperties=%v]",
    	u.PricingScheme, u.Prices, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateSegment.
// It customizes the JSON marshaling process for UpdateSegment objects.
func (u UpdateSegment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "pricing_scheme", "prices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSegment object to a map representation for JSON marshaling.
func (u UpdateSegment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["pricing_scheme"] = u.PricingScheme
    if u.Prices != nil {
        structMap["prices"] = u.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSegment.
// It customizes the JSON unmarshaling process for UpdateSegment objects.
func (u *UpdateSegment) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSegment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "pricing_scheme", "prices")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.PricingScheme = *temp.PricingScheme
    u.Prices = temp.Prices
    return nil
}

// tempUpdateSegment is a temporary struct used for validating the fields of UpdateSegment.
type tempUpdateSegment  struct {
    PricingScheme *PricingScheme               `json:"pricing_scheme"`
    Prices        []CreateOrUpdateSegmentPrice `json:"prices,omitempty"`
}

func (u *tempUpdateSegment) validate() error {
    var errs []string
    if u.PricingScheme == nil {
        errs = append(errs, "required field `pricing_scheme` is missing for type `Update Segment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
