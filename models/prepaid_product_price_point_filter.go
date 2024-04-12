package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// PrepaidProductPricePointFilter represents a PrepaidProductPricePointFilter struct.
type PrepaidProductPricePointFilter struct {
    // Passed as a parameter to list methods to return only non null values.
    ProductPricePointId  string         `json:"product_price_point_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PrepaidProductPricePointFilter.
// It customizes the JSON marshaling process for PrepaidProductPricePointFilter objects.
func (p PrepaidProductPricePointFilter) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PrepaidProductPricePointFilter object to a map representation for JSON marshaling.
func (p PrepaidProductPricePointFilter) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["product_price_point_id"] = p.ProductPricePointId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaidProductPricePointFilter.
// It customizes the JSON unmarshaling process for PrepaidProductPricePointFilter objects.
func (p *PrepaidProductPricePointFilter) UnmarshalJSON(input []byte) error {
    var temp prepaidProductPricePointFilter
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "product_price_point_id")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.ProductPricePointId = *temp.ProductPricePointId
    return nil
}

// TODO
type prepaidProductPricePointFilter  struct {
    ProductPricePointId *string `json:"product_price_point_id"`
}

func (p *prepaidProductPricePointFilter) validate() error {
    var errs []string
    if p.ProductPricePointId == nil {
        errs = append(errs, "required field `product_price_point_id` is missing for type `Prepaid Product Price Point Filter`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
