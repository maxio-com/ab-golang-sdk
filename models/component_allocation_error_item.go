package models

import (
    "encoding/json"
)

// ComponentAllocationErrorItem represents a ComponentAllocationErrorItem struct.
type ComponentAllocationErrorItem struct {
    ComponentId          *int           `json:"component_id,omitempty"`
    Message              *string        `json:"message,omitempty"`
    Kind                 *string        `json:"kind,omitempty"`
    On                   *string        `json:"on,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentAllocationErrorItem.
// It customizes the JSON marshaling process for ComponentAllocationErrorItem objects.
func (c ComponentAllocationErrorItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ComponentAllocationErrorItem object to a map representation for JSON marshaling.
func (c ComponentAllocationErrorItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ComponentId != nil {
        structMap["component_id"] = c.ComponentId
    }
    if c.Message != nil {
        structMap["message"] = c.Message
    }
    if c.Kind != nil {
        structMap["kind"] = c.Kind
    }
    if c.On != nil {
        structMap["on"] = c.On
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentAllocationErrorItem.
// It customizes the JSON unmarshaling process for ComponentAllocationErrorItem objects.
func (c *ComponentAllocationErrorItem) UnmarshalJSON(input []byte) error {
    var temp componentAllocationErrorItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "component_id", "message", "kind", "on")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.ComponentId = temp.ComponentId
    c.Message = temp.Message
    c.Kind = temp.Kind
    c.On = temp.On
    return nil
}

// componentAllocationErrorItem is a temporary struct used for validating the fields of ComponentAllocationErrorItem.
type componentAllocationErrorItem  struct {
    ComponentId *int    `json:"component_id,omitempty"`
    Message     *string `json:"message,omitempty"`
    Kind        *string `json:"kind,omitempty"`
    On          *string `json:"on,omitempty"`
}
