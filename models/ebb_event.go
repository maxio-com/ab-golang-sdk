package models

import (
    "encoding/json"
)

// EBBEvent represents a EBBEvent struct.
type EBBEvent struct {
    Chargify             *ChargifyEBB   `json:"chargify,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EBBEvent.
// It customizes the JSON marshaling process for EBBEvent objects.
func (e EBBEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EBBEvent object to a map representation for JSON marshaling.
func (e EBBEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Chargify != nil {
        structMap["chargify"] = e.Chargify.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBEvent.
// It customizes the JSON unmarshaling process for EBBEvent objects.
func (e *EBBEvent) UnmarshalJSON(input []byte) error {
    var temp ebbEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chargify")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Chargify = temp.Chargify
    return nil
}

// TODO
type ebbEvent  struct {
    Chargify *ChargifyEBB `json:"chargify,omitempty"`
}
