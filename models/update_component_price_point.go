package models

import (
    "encoding/json"
)

// UpdateComponentPricePoint represents a UpdateComponentPricePoint struct.
type UpdateComponentPricePoint struct {
    Name         *string       `json:"name,omitempty"`
    // The numerical interval. i.e. an interval of ‘30’ coupled with an interval_unit of day would mean this component price point would renew every 30 days. This property is only available for sites with Multifrequency enabled.
    Interval     *int          `json:"interval,omitempty"`
    // A string representing the interval unit for this component price point, either month or day. This property is only available for sites with Multifrequency enabled.
    IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
    Prices       []UpdatePrice `json:"prices,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateComponentPricePoint.
// It customizes the JSON marshaling process for UpdateComponentPricePoint objects.
func (u *UpdateComponentPricePoint) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateComponentPricePoint object to a map representation for JSON marshaling.
func (u *UpdateComponentPricePoint) toMap() map[string]any {
    structMap := make(map[string]any)
    if u.Name != nil {
        structMap["name"] = u.Name
    }
    if u.Interval != nil {
        structMap["interval"] = u.Interval
    }
    if u.IntervalUnit != nil {
        structMap["interval_unit"] = u.IntervalUnit
    }
    if u.Prices != nil {
        structMap["prices"] = u.Prices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateComponentPricePoint.
// It customizes the JSON unmarshaling process for UpdateComponentPricePoint objects.
func (u *UpdateComponentPricePoint) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Name         *string       `json:"name,omitempty"`
        Interval     *int          `json:"interval,omitempty"`
        IntervalUnit *IntervalUnit `json:"interval_unit,omitempty"`
        Prices       []UpdatePrice `json:"prices,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    u.Name = temp.Name
    u.Interval = temp.Interval
    u.IntervalUnit = temp.IntervalUnit
    u.Prices = temp.Prices
    return nil
}
