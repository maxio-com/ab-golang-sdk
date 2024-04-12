package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ItemPricePointChanged represents a ItemPricePointChanged struct.
type ItemPricePointChanged struct {
    ItemId               int                `json:"item_id"`
    ItemType             string             `json:"item_type"`
    ItemHandle           string             `json:"item_handle"`
    ItemName             string             `json:"item_name"`
    PreviousPricePoint   ItemPricePointData `json:"previous_price_point"`
    CurrentPricePoint    ItemPricePointData `json:"current_price_point"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ItemPricePointChanged.
// It customizes the JSON marshaling process for ItemPricePointChanged objects.
func (i ItemPricePointChanged) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the ItemPricePointChanged object to a map representation for JSON marshaling.
func (i ItemPricePointChanged) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["item_id"] = i.ItemId
    structMap["item_type"] = i.ItemType
    structMap["item_handle"] = i.ItemHandle
    structMap["item_name"] = i.ItemName
    structMap["previous_price_point"] = i.PreviousPricePoint.toMap()
    structMap["current_price_point"] = i.CurrentPricePoint.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ItemPricePointChanged.
// It customizes the JSON unmarshaling process for ItemPricePointChanged objects.
func (i *ItemPricePointChanged) UnmarshalJSON(input []byte) error {
    var temp itemPricePointChanged
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "item_id", "item_type", "item_handle", "item_name", "previous_price_point", "current_price_point")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.ItemId = *temp.ItemId
    i.ItemType = *temp.ItemType
    i.ItemHandle = *temp.ItemHandle
    i.ItemName = *temp.ItemName
    i.PreviousPricePoint = *temp.PreviousPricePoint
    i.CurrentPricePoint = *temp.CurrentPricePoint
    return nil
}

// TODO
type itemPricePointChanged  struct {
    ItemId             *int                `json:"item_id"`
    ItemType           *string             `json:"item_type"`
    ItemHandle         *string             `json:"item_handle"`
    ItemName           *string             `json:"item_name"`
    PreviousPricePoint *ItemPricePointData `json:"previous_price_point"`
    CurrentPricePoint  *ItemPricePointData `json:"current_price_point"`
}

func (i *itemPricePointChanged) validate() error {
    var errs []string
    if i.ItemId == nil {
        errs = append(errs, "required field `item_id` is missing for type `Item Price Point Changed`")
    }
    if i.ItemType == nil {
        errs = append(errs, "required field `item_type` is missing for type `Item Price Point Changed`")
    }
    if i.ItemHandle == nil {
        errs = append(errs, "required field `item_handle` is missing for type `Item Price Point Changed`")
    }
    if i.ItemName == nil {
        errs = append(errs, "required field `item_name` is missing for type `Item Price Point Changed`")
    }
    if i.PreviousPricePoint == nil {
        errs = append(errs, "required field `previous_price_point` is missing for type `Item Price Point Changed`")
    }
    if i.CurrentPricePoint == nil {
        errs = append(errs, "required field `current_price_point` is missing for type `Item Price Point Changed`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
