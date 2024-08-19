/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// CalendarBilling represents a CalendarBilling struct.
// (Optional). Cannot be used when also specifying next_billing_at
type CalendarBilling struct {
    // A day of month that subscription will be processed on. Can be 1 up to 28 or 'end'.
    SnapDay                    *CalendarBillingSnapDay `json:"snap_day,omitempty"`
    CalendarBillingFirstCharge *FirstChargeType        `json:"calendar_billing_first_charge,omitempty"`
    AdditionalProperties       map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CalendarBilling.
// It customizes the JSON marshaling process for CalendarBilling objects.
func (c CalendarBilling) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CalendarBilling object to a map representation for JSON marshaling.
func (c CalendarBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.SnapDay != nil {
        structMap["snap_day"] = c.SnapDay.toMap()
    }
    if c.CalendarBillingFirstCharge != nil {
        structMap["calendar_billing_first_charge"] = c.CalendarBillingFirstCharge
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CalendarBilling.
// It customizes the JSON unmarshaling process for CalendarBilling objects.
func (c *CalendarBilling) UnmarshalJSON(input []byte) error {
    var temp tempCalendarBilling
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "snap_day", "calendar_billing_first_charge")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.SnapDay = temp.SnapDay
    c.CalendarBillingFirstCharge = temp.CalendarBillingFirstCharge
    return nil
}

// tempCalendarBilling is a temporary struct used for validating the fields of CalendarBilling.
type tempCalendarBilling  struct {
    SnapDay                    *CalendarBillingSnapDay `json:"snap_day,omitempty"`
    CalendarBillingFirstCharge *FirstChargeType        `json:"calendar_billing_first_charge,omitempty"`
}
