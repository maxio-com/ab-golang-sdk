// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CreateOfferComponent represents a CreateOfferComponent struct.
type CreateOfferComponent struct {
    ComponentId          *int                   `json:"component_id,omitempty"`
    PricePointId         *int                   `json:"price_point_id,omitempty"`
    StartingQuantity     *int                   `json:"starting_quantity,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateOfferComponent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateOfferComponent) String() string {
    return fmt.Sprintf(
    	"CreateOfferComponent[ComponentId=%v, PricePointId=%v, StartingQuantity=%v, AdditionalProperties=%v]",
    	c.ComponentId, c.PricePointId, c.StartingQuantity, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateOfferComponent.
// It customizes the JSON marshaling process for CreateOfferComponent objects.
func (c CreateOfferComponent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "component_id", "price_point_id", "starting_quantity"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateOfferComponent object to a map representation for JSON marshaling.
func (c CreateOfferComponent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.PricePointId != nil {
        structMap["price_point_id"] = c.PricePointId
    }
    if c.StartingQuantity != nil {
        structMap["starting_quantity"] = c.StartingQuantity
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateOfferComponent.
// It customizes the JSON unmarshaling process for CreateOfferComponent objects.
func (c *CreateOfferComponent) UnmarshalJSON(input []byte) error {
    var temp tempCreateOfferComponent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "component_id", "price_point_id", "starting_quantity")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ComponentId = temp.ComponentId
    c.PricePointId = temp.PricePointId
    c.StartingQuantity = temp.StartingQuantity
    return nil
}

// tempCreateOfferComponent is a temporary struct used for validating the fields of CreateOfferComponent.
type tempCreateOfferComponent  struct {
    ComponentId      *int `json:"component_id,omitempty"`
    PricePointId     *int `json:"price_point_id,omitempty"`
    StartingQuantity *int `json:"starting_quantity,omitempty"`
}
