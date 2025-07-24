// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// CreateUsage represents a CreateUsage struct.
type CreateUsage struct {
    // integer by default or decimal number if fractional quantities are enabled for the component
    Quantity             *float64               `json:"quantity,omitempty"`
    PricePointId         *string                `json:"price_point_id,omitempty"`
    Memo                 *string                `json:"memo,omitempty"`
    // This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. Please note this only works for site with Multifrequency enabled
    BillingSchedule      *BillingSchedule       `json:"billing_schedule,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CreateUsage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateUsage) String() string {
    return fmt.Sprintf(
    	"CreateUsage[Quantity=%v, PricePointId=%v, Memo=%v, BillingSchedule=%v, AdditionalProperties=%v]",
    	c.Quantity, c.PricePointId, c.Memo, c.BillingSchedule, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateUsage.
// It customizes the JSON marshaling process for CreateUsage objects.
func (c CreateUsage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "quantity", "price_point_id", "memo", "billing_schedule"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsage object to a map representation for JSON marshaling.
func (c CreateUsage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Quantity != nil {
        structMap["quantity"] = c.Quantity
    }
    if c.PricePointId != nil {
        structMap["price_point_id"] = c.PricePointId
    }
    if c.Memo != nil {
        structMap["memo"] = c.Memo
    }
    if c.BillingSchedule != nil {
        structMap["billing_schedule"] = c.BillingSchedule.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateUsage.
// It customizes the JSON unmarshaling process for CreateUsage objects.
func (c *CreateUsage) UnmarshalJSON(input []byte) error {
    var temp tempCreateUsage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "quantity", "price_point_id", "memo", "billing_schedule")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Quantity = temp.Quantity
    c.PricePointId = temp.PricePointId
    c.Memo = temp.Memo
    c.BillingSchedule = temp.BillingSchedule
    return nil
}

// tempCreateUsage is a temporary struct used for validating the fields of CreateUsage.
type tempCreateUsage  struct {
    Quantity        *float64         `json:"quantity,omitempty"`
    PricePointId    *string          `json:"price_point_id,omitempty"`
    Memo            *string          `json:"memo,omitempty"`
    BillingSchedule *BillingSchedule `json:"billing_schedule,omitempty"`
}
