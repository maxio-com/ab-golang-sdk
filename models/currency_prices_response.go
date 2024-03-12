package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CurrencyPricesResponse represents a CurrencyPricesResponse struct.
type CurrencyPricesResponse struct {
	CurrencyPrices []CurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPricesResponse.
// It customizes the JSON marshaling process for CurrencyPricesResponse objects.
func (c *CurrencyPricesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPricesResponse object to a map representation for JSON marshaling.
func (c *CurrencyPricesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CurrencyPricesResponse.
// It customizes the JSON unmarshaling process for CurrencyPricesResponse objects.
func (c *CurrencyPricesResponse) UnmarshalJSON(input []byte) error {
	var temp currencyPricesResponse
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
type currencyPricesResponse struct {
	CurrencyPrices *[]CurrencyPrice `json:"currency_prices"`
}

func (c *currencyPricesResponse) validate() error {
	var errs []string
	if c.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Currency Prices Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
