/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// OfferItem represents a OfferItem struct.
type OfferItem struct {
    ComponentId          *int                   `json:"component_id,omitempty"`
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    StartingQuantity     *string                `json:"starting_quantity,omitempty"`
    Editable             *bool                  `json:"editable,omitempty"`
    ComponentUnitPrice   *string                `json:"component_unit_price,omitempty"`
    ComponentName        *string                `json:"component_name,omitempty"`
    PricePointName       *string                `json:"price_point_name,omitempty"`
    CurrencyPrices       []CurrencyPrice        `json:"currency_prices,omitempty"`
    // The numerical interval. i.e. an interval of '30' coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval             *int                   `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit         Optional[IntervalUnit] `json:"interval_unit"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OfferItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OfferItem) String() string {
    return fmt.Sprintf(
    	"OfferItem[ComponentId=%v, PricePointId=%v, StartingQuantity=%v, Editable=%v, ComponentUnitPrice=%v, ComponentName=%v, PricePointName=%v, CurrencyPrices=%v, Interval=%v, IntervalUnit=%v, AdditionalProperties=%v]",
    	o.ComponentId, o.PricePointId, o.StartingQuantity, o.Editable, o.ComponentUnitPrice, o.ComponentName, o.PricePointName, o.CurrencyPrices, o.Interval, o.IntervalUnit, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OfferItem.
// It customizes the JSON marshaling process for OfferItem objects.
func (o OfferItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "component_id", "price_point_id", "starting_quantity", "editable", "component_unit_price", "component_name", "price_point_name", "currency_prices", "interval", "interval_unit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OfferItem object to a map representation for JSON marshaling.
func (o OfferItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
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
    if o.IntervalUnit.IsValueSet() {
        if o.IntervalUnit.Value() != nil {
            structMap["interval_unit"] = o.IntervalUnit.Value()
        } else {
            structMap["interval_unit"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OfferItem.
// It customizes the JSON unmarshaling process for OfferItem objects.
func (o *OfferItem) UnmarshalJSON(input []byte) error {
    var temp tempOfferItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "price_point_id", "starting_quantity", "editable", "component_unit_price", "component_name", "price_point_name", "currency_prices", "interval", "interval_unit")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
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

// tempOfferItem is a temporary struct used for validating the fields of OfferItem.
type tempOfferItem  struct {
    ComponentId        *int                   `json:"component_id,omitempty"`
    PricePointId       *int                   `json:"price_point_id,omitempty"`
    StartingQuantity   *string                `json:"starting_quantity,omitempty"`
    Editable           *bool                  `json:"editable,omitempty"`
    ComponentUnitPrice *string                `json:"component_unit_price,omitempty"`
    ComponentName      *string                `json:"component_name,omitempty"`
    PricePointName     *string                `json:"price_point_name,omitempty"`
    CurrencyPrices     []CurrencyPrice        `json:"currency_prices,omitempty"`
    Interval           *int                   `json:"interval,omitempty"`
    IntervalUnit       Optional[IntervalUnit] `json:"interval_unit"`
}
