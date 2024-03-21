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
}

// MarshalJSON implements the json.Marshaler interface for CalendarBilling.
// It customizes the JSON marshaling process for CalendarBilling objects.
func (c *CalendarBilling) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CalendarBilling object to a map representation for JSON marshaling.
func (c *CalendarBilling) toMap() map[string]any {
	structMap := make(map[string]any)
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
	var temp calendarBilling
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	c.SnapDay = temp.SnapDay
	c.CalendarBillingFirstCharge = temp.CalendarBillingFirstCharge
	return nil
}

// TODO
type calendarBilling struct {
	SnapDay                    *CalendarBillingSnapDay `json:"snap_day,omitempty"`
	CalendarBillingFirstCharge *FirstChargeType        `json:"calendar_billing_first_charge,omitempty"`
}
