package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CouponCurrencyRequest represents a CouponCurrencyRequest struct.
type CouponCurrencyRequest struct {
    CurrencyPrices       []UpdateCouponCurrency `json:"currency_prices"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrencyRequest.
// It customizes the JSON marshaling process for CouponCurrencyRequest objects.
func (c CouponCurrencyRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrencyRequest object to a map representation for JSON marshaling.
func (c CouponCurrencyRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["currency_prices"] = c.CurrencyPrices
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrencyRequest.
// It customizes the JSON unmarshaling process for CouponCurrencyRequest objects.
func (c *CouponCurrencyRequest) UnmarshalJSON(input []byte) error {
    var temp couponCurrencyRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currency_prices")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.CurrencyPrices = *temp.CurrencyPrices
    return nil
}

// TODO
type couponCurrencyRequest  struct {
    CurrencyPrices *[]UpdateCouponCurrency `json:"currency_prices"`
}

func (c *couponCurrencyRequest) validate() error {
    var errs []string
    if c.CurrencyPrices == nil {
        errs = append(errs, "required field `currency_prices` is missing for type `Coupon Currency Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
