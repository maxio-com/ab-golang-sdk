package models

import (
	"encoding/json"
)

// OfferItem represents a OfferItem struct.
type OfferItem struct {
	ComponentId        *int            `json:"component_id,omitempty"`
	PricePointId       *int            `json:"price_point_id,omitempty"`
	StartingQuantity   *string         `json:"starting_quantity,omitempty"`
	Editable           *bool           `json:"editable,omitempty"`
	ComponentUnitPrice *string         `json:"component_unit_price,omitempty"`
	ComponentName      *string         `json:"component_name,omitempty"`
	PricePointName     *string         `json:"price_point_name,omitempty"`
	CurrencyPrices     []CurrencyPrice `json:"currency_prices,omitempty"`
	// The numerical interval. i.e. an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
	Interval *int `json:"interval,omitempty"`
	// A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
	IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OfferItem.
// It customizes the JSON marshaling process for OfferItem objects.
func (o *OfferItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OfferItem object to a map representation for JSON marshaling.
func (o *OfferItem) toMap() map[string]any {
	structMap := make(map[string]any)
	if o.ComponentId != nil {
		structMap["component_id"] = o.ComponentId
	}
	if o.PricePointId != nil {
		structMap["price_point_id"] = o.PricePointId
	}
	if o.StartingQuantity != nil {
		structMap["starting_quantity"] = o.StartingQuantity
	}
	if o.Editable != nil {
		structMap["editable"] = o.Editable
	}
	if o.ComponentUnitPrice != nil {
		structMap["component_unit_price"] = o.ComponentUnitPrice
	}
	if o.ComponentName != nil {
		structMap["component_name"] = o.ComponentName
	}
	if o.PricePointName != nil {
		structMap["price_point_name"] = o.PricePointName
	}
	if o.CurrencyPrices != nil {
		structMap["currency_prices"] = o.CurrencyPrices
	}
	if o.Interval != nil {
		structMap["interval"] = o.Interval
	}
	if o.IntervalUnit != nil {
		structMap["interval_unit"] = o.IntervalUnit
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferItem.
// It customizes the JSON unmarshaling process for OfferItem objects.
func (o *OfferItem) UnmarshalJSON(input []byte) error {
	var temp offerItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	o.ComponentId = temp.ComponentId
	o.PricePointId = temp.PricePointId
	o.StartingQuantity = temp.StartingQuantity
	o.Editable = temp.Editable
	o.ComponentUnitPrice = temp.ComponentUnitPrice
	o.ComponentName = temp.ComponentName
	o.PricePointName = temp.PricePointName
	o.CurrencyPrices = temp.CurrencyPrices
	o.Interval = temp.Interval
	o.IntervalUnit = temp.IntervalUnit
	return nil
}

// TODO
type offerItem struct {
	ComponentId        *int            `json:"component_id,omitempty"`
	PricePointId       *int            `json:"price_point_id,omitempty"`
	StartingQuantity   *string         `json:"starting_quantity,omitempty"`
	Editable           *bool           `json:"editable,omitempty"`
	ComponentUnitPrice *string         `json:"component_unit_price,omitempty"`
	ComponentName      *string         `json:"component_name,omitempty"`
	PricePointName     *string         `json:"price_point_name,omitempty"`
	CurrencyPrices     []CurrencyPrice `json:"currency_prices,omitempty"`
	Interval           *int            `json:"interval,omitempty"`
	IntervalUnit       *IntervalUnit   `json:"interval_unit,omitempty"`
}
