package models

import (
    "encoding/json"
)

// GroupBilling represents a GroupBilling struct.
// Optional attributes related to billing date and accrual. Note: Only applicable for new subscriptions.
type GroupBilling struct {
    // A flag indicating whether or not to accrue charges on the new subscription.
    Accrue               *bool          `json:"accrue,omitempty"`
    // A flag indicating whether or not to align the billing date of the new subscription with the billing date of the primary subscription of the hierarchy's default subscription group. Required to be true if prorate is also true.
    AlignDate            *bool          `json:"align_date,omitempty"`
    // A flag indicating whether or not to prorate billing of the new subscription for the current period. A value of true is ignored unless align_date is also true.
    Prorate              *bool          `json:"prorate,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GroupBilling.
// It customizes the JSON marshaling process for GroupBilling objects.
func (g GroupBilling) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GroupBilling object to a map representation for JSON marshaling.
func (g GroupBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    if g.Accrue != nil {
        structMap["accrue"] = g.Accrue
    }
    if g.AlignDate != nil {
        structMap["align_date"] = g.AlignDate
    }
    if g.Prorate != nil {
        structMap["prorate"] = g.Prorate
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupBilling.
// It customizes the JSON unmarshaling process for GroupBilling objects.
func (g *GroupBilling) UnmarshalJSON(input []byte) error {
    var temp groupBilling
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accrue", "align_date", "prorate")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Accrue = temp.Accrue
    g.AlignDate = temp.AlignDate
    g.Prorate = temp.Prorate
    return nil
}

// TODO
type groupBilling  struct {
    Accrue    *bool `json:"accrue,omitempty"`
    AlignDate *bool `json:"align_date,omitempty"`
    Prorate   *bool `json:"prorate,omitempty"`
}
