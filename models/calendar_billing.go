/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CalendarBilling represents a CalendarBilling struct.
// (Optional). Cannot be used when also specifying next_billing_at
type CalendarBilling struct {
    // A day of month that subscription will be processed on. Can be 1 up to 28 or 'end'.
    SnapDay                    *CalendarBillingSnapDay `json:"snap_day,omitempty"`
    CalendarBillingFirstCharge *FirstChargeType        `json:"calendar_billing_first_charge,omitempty"`
    AdditionalProperties       map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for CalendarBilling,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CalendarBilling) String() string {
    return fmt.Sprintf(
    	"CalendarBilling[SnapDay=%v, CalendarBillingFirstCharge=%v, AdditionalProperties=%v]",
    	c.SnapDay, c.CalendarBillingFirstCharge, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CalendarBilling.
// It customizes the JSON marshaling process for CalendarBilling objects.
func (c CalendarBilling) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "snap_day", "calendar_billing_first_charge"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CalendarBilling object to a map representation for JSON marshaling.
func (c CalendarBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "snap_day", "calendar_billing_first_charge")
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
