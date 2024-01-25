package models

import (
	"encoding/json"
)

// UpdateCurrencyPrice represents a UpdateCurrencyPrice struct.
type UpdateCurrencyPrice struct {
	// ID of the currency price record being updated
	Id int `json:"id"`
	// New price for the given currency
	Price int `json:"price"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrencyPrice.
// It customizes the JSON marshaling process for UpdateCurrencyPrice objects.
func (u *UpdateCurrencyPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrencyPrice object to a map representation for JSON marshaling.
func (u *UpdateCurrencyPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = u.Id
	structMap["price"] = u.Price
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrencyPrice.
// It customizes the JSON unmarshaling process for UpdateCurrencyPrice objects.
func (u *UpdateCurrencyPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id    int `json:"id"`
		Price int `json:"price"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.Id = temp.Id
	u.Price = temp.Price
	return nil
}
