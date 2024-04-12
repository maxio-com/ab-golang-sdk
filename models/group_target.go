package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// GroupTarget represents a GroupTarget struct.
// Attributes of the target customer who will be the responsible payer of the created subscription. Required.
type GroupTarget struct {
    // The type of object indicated by the id attribute.
    Type                 GroupTargetType `json:"type"`
    // The id of the target customer or subscription to group the existing subscription with. Ignored and should not be included if type is "self" , "parent", or "eldest"
    Id                   *int            `json:"id,omitempty"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for GroupTarget.
// It customizes the JSON marshaling process for GroupTarget objects.
func (g GroupTarget) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GroupTarget object to a map representation for JSON marshaling.
func (g GroupTarget) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, g.AdditionalProperties)
    structMap["type"] = g.Type
    if g.Id != nil {
        structMap["id"] = g.Id
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupTarget.
// It customizes the JSON unmarshaling process for GroupTarget objects.
func (g *GroupTarget) UnmarshalJSON(input []byte) error {
    var temp groupTarget
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "type", "id")
    if err != nil {
    	return err
    }
    
    g.AdditionalProperties = additionalProperties
    g.Type = *temp.Type
    g.Id = temp.Id
    return nil
}

// TODO
type groupTarget  struct {
    Type *GroupTargetType `json:"type"`
    Id   *int             `json:"id,omitempty"`
}

func (g *groupTarget) validate() error {
    var errs []string
    if g.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Group Target`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
