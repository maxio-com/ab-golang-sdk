package models

import (
    "encoding/json"
)

// ProductPricePointErrors represents a ProductPricePointErrors struct.
type ProductPricePointErrors struct {
    PricePoint           *string        `json:"price_point,omitempty"`
    Interval             []string       `json:"interval,omitempty"`
    IntervalUnit         []string       `json:"interval_unit,omitempty"`
    Name                 []string       `json:"name,omitempty"`
    Price                []string       `json:"price,omitempty"`
    PriceInCents         []string       `json:"price_in_cents,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProductPricePointErrors.
// It customizes the JSON marshaling process for ProductPricePointErrors objects.
func (p ProductPricePointErrors) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProductPricePointErrors object to a map representation for JSON marshaling.
func (p ProductPricePointErrors) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    var temp productPricePointErrors
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_point", "interval", "interval_unit", "name", "price", "price_in_cents")
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

// TODO
type productPricePointErrors  struct {
    PricePoint   *string  `json:"price_point,omitempty"`
    Interval     []string `json:"interval,omitempty"`
    IntervalUnit []string `json:"interval_unit,omitempty"`
    Name         []string `json:"name,omitempty"`
    Price        []string `json:"price,omitempty"`
    PriceInCents []string `json:"price_in_cents,omitempty"`
}
