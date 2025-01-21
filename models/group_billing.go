/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// GroupBilling represents a GroupBilling struct.
// Optional attributes related to billing date and accrual. Note: Only applicable for new subscriptions.
type GroupBilling struct {
    // A flag indicating whether or not to accrue charges on the new subscription.
    Accrue               *bool                  `json:"accrue,omitempty"`
    // A flag indicating whether or not to align the billing date of the new subscription with the billing date of the primary subscription of the hierarchy's default subscription group. Required to be true if prorate is also true.
    AlignDate            *bool                  `json:"align_date,omitempty"`
    // A flag indicating whether or not to prorate billing of the new subscription for the current period. A value of true is ignored unless align_date is also true.
    Prorate              *bool                  `json:"prorate,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for GroupBilling,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (g GroupBilling) String() string {
    return fmt.Sprintf(
    	"GroupBilling[Accrue=%v, AlignDate=%v, Prorate=%v, AdditionalProperties=%v]",
    	g.Accrue, g.AlignDate, g.Prorate, g.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for GroupBilling.
// It customizes the JSON marshaling process for GroupBilling objects.
func (g GroupBilling) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(g.AdditionalProperties,
        "accrue", "align_date", "prorate"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(g.toMap())
}

// toMap converts the GroupBilling object to a map representation for JSON marshaling.
func (g GroupBilling) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, g.AdditionalProperties)
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
    var temp tempGroupBilling
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accrue", "align_date", "prorate")
    if err != nil {
    	return err
    }
    g.AdditionalProperties = additionalProperties
    
    g.Accrue = temp.Accrue
    g.AlignDate = temp.AlignDate
    g.Prorate = temp.Prorate
    return nil
}

// tempGroupBilling is a temporary struct used for validating the fields of GroupBilling.
type tempGroupBilling  struct {
    Accrue    *bool `json:"accrue,omitempty"`
    AlignDate *bool `json:"align_date,omitempty"`
    Prorate   *bool `json:"prorate,omitempty"`
}
