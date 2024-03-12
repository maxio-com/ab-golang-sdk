package models

import (
	"encoding/json"
)

// AddSubscriptionToAGroup represents a AddSubscriptionToAGroup struct.
type AddSubscriptionToAGroup struct {
	Group *AddSubscriptionToAGroupGroup `json:"group,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for AddSubscriptionToAGroup.
// It customizes the JSON marshaling process for AddSubscriptionToAGroup objects.
func (a *AddSubscriptionToAGroup) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AddSubscriptionToAGroup object to a map representation for JSON marshaling.
func (a *AddSubscriptionToAGroup) toMap() map[string]any {
	structMap := make(map[string]any)
	if a.Group != nil {
		structMap["group"] = a.Group.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AddSubscriptionToAGroup.
// It customizes the JSON unmarshaling process for AddSubscriptionToAGroup objects.
func (a *AddSubscriptionToAGroup) UnmarshalJSON(input []byte) error {
	var temp addSubscriptionToAGroup
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	a.Group = temp.Group
	return nil
}

// TODO
type addSubscriptionToAGroup struct {
	Group *AddSubscriptionToAGroupGroup `json:"group,omitempty"`
}
