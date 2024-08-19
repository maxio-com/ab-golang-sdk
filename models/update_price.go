/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdatePrice represents a UpdatePrice struct.
type UpdatePrice struct {
    Id                   *int                         `json:"id,omitempty"`
    EndingQuantity       *UpdatePriceEndingQuantity   `json:"ending_quantity,omitempty"`
    // The price can contain up to 8 decimal places. i.e. 1.00 or 0.0012 or 0.00000065
    UnitPrice            *UpdatePriceUnitPrice        `json:"unit_price,omitempty"`
    Destroy              *bool                        `json:"_destroy,omitempty"`
    StartingQuantity     *UpdatePriceStartingQuantity `json:"starting_quantity,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdatePrice.
// It customizes the JSON marshaling process for UpdatePrice objects.
func (u UpdatePrice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdatePrice object to a map representation for JSON marshaling.
func (u UpdatePrice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.EndingQuantity != nil {
        structMap["ending_quantity"] = u.EndingQuantity.toMap()
    }
    if u.UnitPrice != nil {
        structMap["unit_price"] = u.UnitPrice.toMap()
    }
    if u.Destroy != nil {
        structMap["_destroy"] = u.Destroy
    }
    if u.StartingQuantity != nil {
        structMap["starting_quantity"] = u.StartingQuantity.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdatePrice.
// It customizes the JSON unmarshaling process for UpdatePrice objects.
func (u *UpdatePrice) UnmarshalJSON(input []byte) error {
    var temp tempUpdatePrice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "ending_quantity", "unit_price", "_destroy", "starting_quantity")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Id = temp.Id
    u.EndingQuantity = temp.EndingQuantity
    u.UnitPrice = temp.UnitPrice
    u.Destroy = temp.Destroy
    u.StartingQuantity = temp.StartingQuantity
    return nil
}

// tempUpdatePrice is a temporary struct used for validating the fields of UpdatePrice.
type tempUpdatePrice  struct {
    Id               *int                         `json:"id,omitempty"`
    EndingQuantity   *UpdatePriceEndingQuantity   `json:"ending_quantity,omitempty"`
    UnitPrice        *UpdatePriceUnitPrice        `json:"unit_price,omitempty"`
    Destroy          *bool                        `json:"_destroy,omitempty"`
    StartingQuantity *UpdatePriceStartingQuantity `json:"starting_quantity,omitempty"`
}
