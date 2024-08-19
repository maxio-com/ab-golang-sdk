/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// RenewalPreviewComponent represents a RenewalPreviewComponent struct.
type RenewalPreviewComponent struct {
    // Either the component's Chargify id or its handle prefixed with `handle:`
    ComponentId          *RenewalPreviewComponentComponentId  `json:"component_id,omitempty"`
    // The quantity for which you wish to preview billing. This is useful if you want to preview a predicted, higher usage value than is currently present on the subscription.
    // This quantity represents:
    // - Whether or not an on/off component is enabled - use 0 for disabled or 1 for enabled
    // - The desired allocated_quantity for a quantity-based component
    // - The desired unit_balance for a metered component
    // - The desired metric quantity for an events-based component
    Quantity             *int                                 `json:"quantity,omitempty"`
    // Either the component price point's Chargify id or its handle prefixed with `handle:`
    PricePointId         *RenewalPreviewComponentPricePointId `json:"price_point_id,omitempty"`
    AdditionalProperties map[string]any                       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RenewalPreviewComponent.
// It customizes the JSON marshaling process for RenewalPreviewComponent objects.
func (r RenewalPreviewComponent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RenewalPreviewComponent object to a map representation for JSON marshaling.
func (r RenewalPreviewComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.ComponentId != nil {
        structMap["component_id"] = r.ComponentId.toMap()
    }
    if r.Quantity != nil {
        structMap["quantity"] = r.Quantity
    }
    if r.PricePointId != nil {
        structMap["price_point_id"] = r.PricePointId.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RenewalPreviewComponent.
// It customizes the JSON unmarshaling process for RenewalPreviewComponent objects.
func (r *RenewalPreviewComponent) UnmarshalJSON(input []byte) error {
    var temp tempRenewalPreviewComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "component_id", "quantity", "price_point_id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ComponentId = temp.ComponentId
    r.Quantity = temp.Quantity
    r.PricePointId = temp.PricePointId
    return nil
}

// tempRenewalPreviewComponent is a temporary struct used for validating the fields of RenewalPreviewComponent.
type tempRenewalPreviewComponent  struct {
    ComponentId  *RenewalPreviewComponentComponentId  `json:"component_id,omitempty"`
    Quantity     *int                                 `json:"quantity,omitempty"`
    PricePointId *RenewalPreviewComponentPricePointId `json:"price_point_id,omitempty"`
}
