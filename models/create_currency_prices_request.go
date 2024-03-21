package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateCurrencyPricesRequest represents a CreateCurrencyPricesRequest struct.
type CreateCurrencyPricesRequest struct {
	CurrencyPrices []CreateCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON marshaling process for CreateCurrencyPricesRequest objects.
func (c *CreateCurrencyPricesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateCurrencyPricesRequest object to a map representation for JSON marshaling.
func (c *CreateCurrencyPricesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for CreateCurrencyPricesRequest objects.
func (c *CreateCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
	var temp createCurrencyPricesRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.CurrencyPrices = *temp.CurrencyPrices
	return nil
}

// TODO
type createCurrencyPricesRequest struct {
	CurrencyPrices *[]CreateCurrencyPrice `json:"currency_prices"`
}

func (c *createCurrencyPricesRequest) validate() error {
	var errs []string
	if c.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Create Currency Prices Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
