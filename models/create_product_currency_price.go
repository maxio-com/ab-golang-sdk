package models

import (
	"encoding/json"
)

// CreateProductCurrencyPrice represents a CreateProductCurrencyPrice struct.
type CreateProductCurrencyPrice struct {
	// ISO code for one of the site level currencies.
	Currency string `json:"currency"`
	// Price for the given role.
	Price int `json:"price"`
	// Role for the price.
	Role CurrencyPriceRole `json:"role"`
}

// MarshalJSON implements the json.Marshaler interface for CreateProductCurrencyPrice.
// It customizes the JSON marshaling process for CreateProductCurrencyPrice objects.
func (c *CreateProductCurrencyPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateProductCurrencyPrice object to a map representation for JSON marshaling.
func (c *CreateProductCurrencyPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["currency"] = c.Currency
	structMap["price"] = c.Price
	structMap["role"] = c.Role
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateProductCurrencyPrice.
// It customizes the JSON unmarshaling process for CreateProductCurrencyPrice objects.
func (c *CreateProductCurrencyPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Currency string            `json:"currency"`
		Price    int               `json:"price"`
		Role     CurrencyPriceRole `json:"role"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Currency = temp.Currency
	c.Price = temp.Price
	c.Role = temp.Role
	return nil
}
