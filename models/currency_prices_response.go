package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CurrencyPricesResponse represents a CurrencyPricesResponse struct.
type CurrencyPricesResponse struct {
    CurrencyPrices       []CurrencyPrice `json:"currency_prices"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CurrencyPricesResponse.
// It customizes the JSON marshaling process for CurrencyPricesResponse objects.
func (c CurrencyPricesResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CurrencyPricesResponse object to a map representation for JSON marshaling.
func (c CurrencyPricesResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "currency_prices")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.CurrencyPrices = *temp.CurrencyPrices
    return nil
}

// TODO
type currencyPricesResponse  struct {
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
