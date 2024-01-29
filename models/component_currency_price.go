package models

import (
	"encoding/json"
)

// ComponentCurrencyPrice represents a ComponentCurrencyPrice struct.
type ComponentCurrencyPrice struct {
	Id             *int    `json:"id,omitempty"`
	Currency       *string `json:"currency,omitempty"`
	Price          *string `json:"price,omitempty"`
	FormattedPrice *string `json:"formatted_price,omitempty"`
	PriceId        *int    `json:"price_id,omitempty"`
	PricePointId   *int    `json:"price_point_id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentCurrencyPrice.
// It customizes the JSON marshaling process for ComponentCurrencyPrice objects.
func (c *ComponentCurrencyPrice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentCurrencyPrice object to a map representation for JSON marshaling.
func (c *ComponentCurrencyPrice) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Id != nil {
		structMap["id"] = c.Id
	}
	if c.Currency != nil {
		structMap["currency"] = c.Currency
	}
	if c.Price != nil {
		structMap["price"] = c.Price
	}
	if c.FormattedPrice != nil {
		structMap["formatted_price"] = c.FormattedPrice
	}
	if c.PriceId != nil {
		structMap["price_id"] = c.PriceId
	}
	if c.PricePointId != nil {
		structMap["price_point_id"] = c.PricePointId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentCurrencyPrice.
// It customizes the JSON unmarshaling process for ComponentCurrencyPrice objects.
func (c *ComponentCurrencyPrice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id             *int    `json:"id,omitempty"`
		Currency       *string `json:"currency,omitempty"`
		Price          *string `json:"price,omitempty"`
		FormattedPrice *string `json:"formatted_price,omitempty"`
		PriceId        *int    `json:"price_id,omitempty"`
		PricePointId   *int    `json:"price_point_id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.Currency = temp.Currency
	c.Price = temp.Price
	c.FormattedPrice = temp.FormattedPrice
	c.PriceId = temp.PriceId
	c.PricePointId = temp.PricePointId
	return nil
}
