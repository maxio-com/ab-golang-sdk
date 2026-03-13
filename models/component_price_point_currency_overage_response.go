// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ComponentPricePointCurrencyOverageResponse represents a ComponentPricePointCurrencyOverageResponse struct.
type ComponentPricePointCurrencyOverageResponse struct {
    // Extends a component price point with currency overage prices.
    PricePoint           CurrencyOveragePrices  `json:"price_point"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ComponentPricePointCurrencyOverageResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ComponentPricePointCurrencyOverageResponse) String() string {
    return fmt.Sprintf(
    	"ComponentPricePointCurrencyOverageResponse[PricePoint=%v, AdditionalProperties=%v]",
    	c.PricePoint, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointCurrencyOverageResponse.
// It customizes the JSON marshaling process for ComponentPricePointCurrencyOverageResponse objects.
func (c ComponentPricePointCurrencyOverageResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "price_point"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointCurrencyOverageResponse object to a map representation for JSON marshaling.
func (c ComponentPricePointCurrencyOverageResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["price_point"] = c.PricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointCurrencyOverageResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointCurrencyOverageResponse objects.
func (c *ComponentPricePointCurrencyOverageResponse) UnmarshalJSON(input []byte) error {
    var temp tempComponentPricePointCurrencyOverageResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "price_point")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.PricePoint = *temp.PricePoint
    return nil
}

// tempComponentPricePointCurrencyOverageResponse is a temporary struct used for validating the fields of ComponentPricePointCurrencyOverageResponse.
type tempComponentPricePointCurrencyOverageResponse  struct {
    PricePoint *CurrencyOveragePrices `json:"price_point"`
}

func (c *tempComponentPricePointCurrencyOverageResponse) validate() error {
    var errs []string
    if c.PricePoint == nil {
        errs = append(errs, "required field `price_point` is missing for type `Component Price Point Currency Overage Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
