package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateCouponCurrency represents a UpdateCouponCurrency struct.
type UpdateCouponCurrency struct {
    // ISO code for the site defined currency.
    Currency             string         `json:"currency"`
    // Price for the given currency.
    Price                int            `json:"price"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCouponCurrency.
// It customizes the JSON marshaling process for UpdateCouponCurrency objects.
func (u UpdateCouponCurrency) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCouponCurrency object to a map representation for JSON marshaling.
func (u UpdateCouponCurrency) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["currency"] = u.Currency
    structMap["price"] = u.Price
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCouponCurrency.
// It customizes the JSON unmarshaling process for UpdateCouponCurrency objects.
func (u *UpdateCouponCurrency) UnmarshalJSON(input []byte) error {
    var temp updateCouponCurrency
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currency", "price")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Currency = *temp.Currency
    u.Price = *temp.Price
    return nil
}

// TODO
type updateCouponCurrency  struct {
    Currency *string `json:"currency"`
    Price    *int    `json:"price"`
}

func (u *updateCouponCurrency) validate() error {
    var errs []string
    if u.Currency == nil {
        errs = append(errs, "required field `currency` is missing for type `Update Coupon Currency`")
    }
    if u.Price == nil {
        errs = append(errs, "required field `price` is missing for type `Update Coupon Currency`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
