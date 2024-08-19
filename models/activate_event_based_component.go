/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ActivateEventBasedComponent represents a ActivateEventBasedComponent struct.
type ActivateEventBasedComponent struct {
    // The Chargify id of the price point
    PricePointId         *int                  `json:"price_point_id,omitempty"`
    // This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. Please note this only works for site with Multifrequency enabled
    BillingSchedule      *BillingSchedule      `json:"billing_schedule,omitempty"`
    // Create or update custom pricing unique to the subscription. Used in place of `price_point_id`.
    CustomPrice          *ComponentCustomPrice `json:"custom_price,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ActivateEventBasedComponent.
// It customizes the JSON marshaling process for ActivateEventBasedComponent objects.
func (a ActivateEventBasedComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ActivateEventBasedComponent object to a map representation for JSON marshaling.
func (a ActivateEventBasedComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.PricePointId != nil {
        structMap["price_point_id"] = a.PricePointId
    }
    if a.BillingSchedule != nil {
        structMap["billing_schedule"] = a.BillingSchedule.toMap()
    }
    if a.CustomPrice != nil {
        structMap["custom_price"] = a.CustomPrice.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActivateEventBasedComponent.
// It customizes the JSON unmarshaling process for ActivateEventBasedComponent objects.
func (a *ActivateEventBasedComponent) UnmarshalJSON(input []byte) error {
    var temp tempActivateEventBasedComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "price_point_id", "billing_schedule", "custom_price")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.PricePointId = temp.PricePointId
    a.BillingSchedule = temp.BillingSchedule
    a.CustomPrice = temp.CustomPrice
    return nil
}

// tempActivateEventBasedComponent is a temporary struct used for validating the fields of ActivateEventBasedComponent.
type tempActivateEventBasedComponent  struct {
    PricePointId    *int                  `json:"price_point_id,omitempty"`
    BillingSchedule *BillingSchedule      `json:"billing_schedule,omitempty"`
    CustomPrice     *ComponentCustomPrice `json:"custom_price,omitempty"`
}
