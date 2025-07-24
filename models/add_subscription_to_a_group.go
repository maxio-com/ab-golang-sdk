// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// AddSubscriptionToAGroup represents a AddSubscriptionToAGroup struct.
type AddSubscriptionToAGroup struct {
    Group                *GroupSettings         `json:"group,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AddSubscriptionToAGroup,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AddSubscriptionToAGroup) String() string {
    return fmt.Sprintf(
    	"AddSubscriptionToAGroup[Group=%v, AdditionalProperties=%v]",
    	a.Group, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AddSubscriptionToAGroup.
// It customizes the JSON marshaling process for AddSubscriptionToAGroup objects.
func (a AddSubscriptionToAGroup) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "group"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AddSubscriptionToAGroup object to a map representation for JSON marshaling.
func (a AddSubscriptionToAGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Group != nil {
        structMap["group"] = a.Group.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddSubscriptionToAGroup.
// It customizes the JSON unmarshaling process for AddSubscriptionToAGroup objects.
func (a *AddSubscriptionToAGroup) UnmarshalJSON(input []byte) error {
    var temp tempAddSubscriptionToAGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "group")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Group = temp.Group
    return nil
}

// tempAddSubscriptionToAGroup is a temporary struct used for validating the fields of AddSubscriptionToAGroup.
type tempAddSubscriptionToAGroup  struct {
    Group *GroupSettings `json:"group,omitempty"`
}
