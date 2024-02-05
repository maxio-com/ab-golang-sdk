package models

import (
    "encoding/json"
)

// CreateUsage represents a CreateUsage struct.
type CreateUsage struct {
    // integer by default or decimal number if fractional quantities are enabled for the component
    Quantity        *float64         `json:"quantity,omitempty"`
    PricePointId    *string          `json:"price_point_id,omitempty"`
    Memo            *string          `json:"memo,omitempty"`
    // This attribute is particularly useful when you need to align billing events for different components on distinct schedules within a subscription. Please note this only works for site with Multifrequency enabled
    BillingSchedule *BillingSchedule `json:"billing_schedule,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateUsage.
// It customizes the JSON marshaling process for CreateUsage objects.
func (c *CreateUsage) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CreateUsage object to a map representation for JSON marshaling.
func (c *CreateUsage) toMap() map[string]any {
    structMap := make(map[string]any)
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
    temp := &struct {
        Quantity        *float64         `json:"quantity,omitempty"`
        PricePointId    *string          `json:"price_point_id,omitempty"`
        Memo            *string          `json:"memo,omitempty"`
        BillingSchedule *BillingSchedule `json:"billing_schedule,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    c.Quantity = temp.Quantity
    c.PricePointId = temp.PricePointId
    c.Memo = temp.Memo
    c.BillingSchedule = temp.BillingSchedule
    return nil
}
