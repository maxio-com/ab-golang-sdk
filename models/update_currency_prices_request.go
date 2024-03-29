package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// UpdateCurrencyPricesRequest represents a UpdateCurrencyPricesRequest struct.
type UpdateCurrencyPricesRequest struct {
	CurrencyPrices []UpdateCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrencyPricesRequest.
// It customizes the JSON marshaling process for UpdateCurrencyPricesRequest objects.
func (u *UpdateCurrencyPricesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrencyPricesRequest object to a map representation for JSON marshaling.
func (u *UpdateCurrencyPricesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = u.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for UpdateCurrencyPricesRequest objects.
func (u *UpdateCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
	var temp updateCurrencyPricesRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	u.CurrencyPrices = *temp.CurrencyPrices
	return nil
}

// TODO
type updateCurrencyPricesRequest struct {
	CurrencyPrices *[]UpdateCurrencyPrice `json:"currency_prices"`
}

func (u *updateCurrencyPricesRequest) validate() error {
	var errs []string
	if u.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Update Currency Prices Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
