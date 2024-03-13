package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateProductCurrencyPricesRequest represents a CreateProductCurrencyPricesRequest struct.
type CreateProductCurrencyPricesRequest struct {
	CurrencyPrices []CreateProductCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductCurrencyPricesRequest.
// It customizes the JSON marshaling process for CreateProductCurrencyPricesRequest objects.
func (c *CreateProductCurrencyPricesRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductCurrencyPricesRequest object to a map representation for JSON marshaling.
func (c *CreateProductCurrencyPricesRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductCurrencyPricesRequest.
// It customizes the JSON unmarshaling process for CreateProductCurrencyPricesRequest objects.
func (c *CreateProductCurrencyPricesRequest) UnmarshalJSON(input []byte) error {
	var temp createProductCurrencyPricesRequest
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
type createProductCurrencyPricesRequest struct {
	CurrencyPrices *[]CreateProductCurrencyPrice `json:"currency_prices"`
}

func (c *createProductCurrencyPricesRequest) validate() error {
	var errs []string
	if c.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Create Product Currency Prices Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
