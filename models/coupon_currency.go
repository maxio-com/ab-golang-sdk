package models

import (
	"encoding/json"
)

// CouponCurrency represents a CouponCurrency struct.
type CouponCurrency struct {
	Id       *int    `json:"id,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Price    *int    `json:"price,omitempty"`
	CouponId *int    `json:"coupon_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponCurrency.
// It customizes the JSON marshaling process for CouponCurrency objects.
func (c *CouponCurrency) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponCurrency object to a map representation for JSON marshaling.
func (c *CouponCurrency) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.Currency != nil {
		structMap["currency"] = c.Currency
	}
	if c.Price != nil {
		structMap["price"] = c.Price
	}
	if c.CouponId != nil {
		structMap["coupon_id"] = c.CouponId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponCurrency.
// It customizes the JSON unmarshaling process for CouponCurrency objects.
func (c *CouponCurrency) UnmarshalJSON(input []byte) error {
	var temp couponCurrency
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	c.Id = temp.Id
	c.Currency = temp.Currency
	c.Price = temp.Price
	c.CouponId = temp.CouponId
	return nil
}

// TODO
type couponCurrency struct {
	Id       *int    `json:"id,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Price    *int    `json:"price,omitempty"`
	CouponId *int    `json:"coupon_id,omitempty"`
}
