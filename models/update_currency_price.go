// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateCurrencyPrice represents a UpdateCurrencyPrice struct.
type UpdateCurrencyPrice struct {
    // ID of the currency price record being updated
    Id                   int                    `json:"id"`
    // New price for the given currency
    Price                float64                `json:"price"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateCurrencyPrice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateCurrencyPrice) String() string {
    return fmt.Sprintf(
    	"UpdateCurrencyPrice[Id=%v, Price=%v, AdditionalProperties=%v]",
    	u.Id, u.Price, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateCurrencyPrice.
// It customizes the JSON marshaling process for UpdateCurrencyPrice objects.
func (u UpdateCurrencyPrice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "id", "price"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateCurrencyPrice object to a map representation for JSON marshaling.
func (u UpdateCurrencyPrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "price")
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
    Id    *int     `json:"id"`
    Price *float64 `json:"price"`
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
