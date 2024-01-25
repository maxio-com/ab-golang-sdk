package models

import (
	"encoding/json"
)

// EBBEvent represents a EBBEvent struct.
type EBBEvent struct {
	Chargify *ChargifyEBB `json:"chargify,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for EBBEvent.
// It customizes the JSON marshaling process for EBBEvent objects.
func (e *EBBEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EBBEvent object to a map representation for JSON marshaling.
func (e *EBBEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	if e.Chargify != nil {
		structMap["chargify"] = e.Chargify
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EBBEvent.
// It customizes the JSON unmarshaling process for EBBEvent objects.
func (e *EBBEvent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Chargify *ChargifyEBB `json:"chargify,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.Chargify = temp.Chargify
	return nil
}
