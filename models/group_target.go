package models

import (
	"encoding/json"
)

// GroupTarget represents a GroupTarget struct.
// Attributes of the target customer who will be the responsible payer of the created subscription. Required.
type GroupTarget struct {
	// The type of object indicated by the id attribute.
	Type GroupTargetType `json:"type"`
	// The id of the target customer or subscription to group the existing subscription with. Ignored and should not be included if type is "self" , "parent", or "eldest"
	Id *int `json:"id,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for GroupTarget.
// It customizes the JSON marshaling process for GroupTarget objects.
func (g *GroupTarget) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(g.toMap())
}

// toMap converts the GroupTarget object to a map representation for JSON marshaling.
func (g *GroupTarget) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["type"] = g.Type
	if g.Id != nil {
		structMap["id"] = g.Id
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupTarget.
// It customizes the JSON unmarshaling process for GroupTarget objects.
func (g *GroupTarget) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Type GroupTargetType `json:"type"`
		Id   *int            `json:"id,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	g.Type = temp.Type
	g.Id = temp.Id
	return nil
}
