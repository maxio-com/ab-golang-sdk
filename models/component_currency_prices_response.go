package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ComponentCurrencyPricesResponse represents a ComponentCurrencyPricesResponse struct.
type ComponentCurrencyPricesResponse struct {
	CurrencyPrices []ComponentCurrencyPrice `json:"currency_prices"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON marshaling process for ComponentCurrencyPricesResponse objects.
func (c *ComponentCurrencyPricesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentCurrencyPricesResponse object to a map representation for JSON marshaling.
func (c *ComponentCurrencyPricesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency_prices"] = c.CurrencyPrices
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCurrencyPricesResponse.
// It customizes the JSON unmarshaling process for ComponentCurrencyPricesResponse objects.
func (c *ComponentCurrencyPricesResponse) UnmarshalJSON(input []byte) error {
	var temp componentCurrencyPricesResponse
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
type componentCurrencyPricesResponse struct {
	CurrencyPrices *[]ComponentCurrencyPrice `json:"currency_prices"`
}

func (c *componentCurrencyPricesResponse) validate() error {
	var errs []string
	if c.CurrencyPrices == nil {
		errs = append(errs, "required field `currency_prices` is missing for type `Component Currency Prices Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
