package models

import (
	"encoding/json"
	"errors"
	"strings"
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
	var temp updateCurrencyPrice
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	u.Id = *temp.Id
	u.Price = *temp.Price
	return nil
}

// TODO
type updateCurrencyPrice struct {
	Id    *int `json:"id"`
	Price *int `json:"price"`
}

func (u *updateCurrencyPrice) validate() error {
	var errs []string
	if u.Id == nil {
		errs = append(errs, "required field `id` is missing for type `Update Currency Price`")
	}
	if u.Price == nil {
		errs = append(errs, "required field `price` is missing for type `Update Currency Price`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
