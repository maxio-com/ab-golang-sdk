/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// ProductPricePointErrors represents a ProductPricePointErrors struct.
type ProductPricePointErrors struct {
    PricePoint           *string                `json:"price_point,omitempty"`
    Interval             []string               `json:"interval,omitempty"`
    IntervalUnit         []string               `json:"interval_unit,omitempty"`
    Name                 []string               `json:"name,omitempty"`
    Price                []string               `json:"price,omitempty"`
    PriceInCents         []string               `json:"price_in_cents,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ProductPricePointErrors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p ProductPricePointErrors) String() string {
    return fmt.Sprintf(
    	"ProductPricePointErrors[PricePoint=%v, Interval=%v, IntervalUnit=%v, Name=%v, Price=%v, PriceInCents=%v, AdditionalProperties=%v]",
    	p.PricePoint, p.Interval, p.IntervalUnit, p.Name, p.Price, p.PriceInCents, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointErrors.
// It customizes the JSON marshaling process for ProductPricePointErrors objects.
func (p ProductPricePointErrors) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "price_point", "interval", "interval_unit", "name", "price", "price_in_cents"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointErrors object to a map representation for JSON marshaling.
func (p ProductPricePointErrors) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.PricePoint != nil {
        structMap["price_point"] = p.PricePoint
    }
    if p.Interval != nil {
        structMap["interval"] = p.Interval
    }
    if p.IntervalUnit != nil {
        structMap["interval_unit"] = p.IntervalUnit
    }
    if p.Name != nil {
        structMap["name"] = p.Name
    }
    if p.Price != nil {
        structMap["price"] = p.Price
    }
    if p.PriceInCents != nil {
        structMap["price_in_cents"] = p.PriceInCents
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProductPricePointErrors.
// It customizes the JSON unmarshaling process for ProductPricePointErrors objects.
func (p *ProductPricePointErrors) UnmarshalJSON(input []byte) error {
    var temp tempProductPricePointErrors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point", "interval", "interval_unit", "name", "price", "price_in_cents")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.PricePoint = temp.PricePoint
    p.Interval = temp.Interval
    p.IntervalUnit = temp.IntervalUnit
    p.Name = temp.Name
    p.Price = temp.Price
    p.PriceInCents = temp.PriceInCents
    return nil
}

// tempProductPricePointErrors is a temporary struct used for validating the fields of ProductPricePointErrors.
type tempProductPricePointErrors  struct {
    PricePoint   *string  `json:"price_point,omitempty"`
    Interval     []string `json:"interval,omitempty"`
    IntervalUnit []string `json:"interval_unit,omitempty"`
    Name         []string `json:"name,omitempty"`
    Price        []string `json:"price,omitempty"`
    PriceInCents []string `json:"price_in_cents,omitempty"`
}
