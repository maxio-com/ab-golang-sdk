package models

import (
	"encoding/json"
)

// UpdateCouponCurrency represents a UpdateCouponCurrency struct.
type UpdateCouponCurrency struct {
	// ISO code for the site defined currency.
	Currency string `json:"currency"`
	// Price for the given currency.
	Price int `json:"price"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCouponCurrency.
// It customizes the JSON marshaling process for UpdateCouponCurrency objects.
func (u *UpdateCouponCurrency) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCouponCurrency object to a map representation for JSON marshaling.
func (u *UpdateCouponCurrency) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency"] = u.Currency
	structMap["price"] = u.Price
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCouponCurrency.
// It customizes the JSON unmarshaling process for UpdateCouponCurrency objects.
func (u *UpdateCouponCurrency) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Currency string `json:"currency"`
		Price    int    `json:"price"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Currency = temp.Currency
	u.Price = temp.Price
	return nil
}
