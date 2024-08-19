/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GroupSettings represents a GroupSettings struct.
type GroupSettings struct {
    // Attributes of the target customer who will be the responsible payer of the created subscription. Required.
    Target               GroupTarget    `json:"target"`
    // Optional attributes related to billing date and accrual. Note: Only applicable for new subscriptions.
    Billing              *GroupBilling  `json:"billing,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GroupSettings.
// It customizes the JSON marshaling process for GroupSettings objects.
func (g GroupSettings) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GroupSettings object to a map representation for JSON marshaling.
func (g GroupSettings) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["target"] = g.Target.toMap()
    if g.Billing != nil {
        structMap["billing"] = g.Billing.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupSettings.
// It customizes the JSON unmarshaling process for GroupSettings objects.
func (g *GroupSettings) UnmarshalJSON(input []byte) error {
    var temp tempGroupSettings
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "target", "billing")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Target = *temp.Target
    g.Billing = temp.Billing
    return nil
}

// tempGroupSettings is a temporary struct used for validating the fields of GroupSettings.
type tempGroupSettings  struct {
    Target  *GroupTarget  `json:"target"`
    Billing *GroupBilling `json:"billing,omitempty"`
}

func (g *tempGroupSettings) validate() error {
    var errs []string
    if g.Target == nil {
        errs = append(errs, "required field `target` is missing for type `Group Settings`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
