/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UpdateCurrencyPrice represents a UpdateCurrencyPrice struct.
type UpdateCurrencyPrice struct {
    // ID of the currency price record being updated
    Id                   int            `json:"id"`
    // New price for the given currency
    Price                int            `json:"price"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrencyPrice.
// It customizes the JSON marshaling process for UpdateCurrencyPrice objects.
func (u UpdateCurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrencyPrice object to a map representation for JSON marshaling.
func (u UpdateCurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["id"] = u.Id
    structMap["price"] = u.Price
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateCurrencyPrice.
// It customizes the JSON unmarshaling process for UpdateCurrencyPrice objects.
func (u *UpdateCurrencyPrice) UnmarshalJSON(input []byte) error {
    var temp tempUpdateCurrencyPrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "price")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Id = *temp.Id
    u.Price = *temp.Price
    return nil
}

// tempUpdateCurrencyPrice is a temporary struct used for validating the fields of UpdateCurrencyPrice.
type tempUpdateCurrencyPrice  struct {
    Id    *int `json:"id"`
    Price *int `json:"price"`
}

func (u *tempUpdateCurrencyPrice) validate() error {
    var errs []string
    if u.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Update Currency Price`")
    }
    if u.Price == nil {
        errs = append(errs, "required field `price` is missing for type `Update Currency Price`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
