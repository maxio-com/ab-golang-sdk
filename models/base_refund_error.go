package models

import (
	"encoding/json"
)

// BaseRefundError represents a BaseRefundError struct.
type BaseRefundError struct {
	Base []interface{} `json:"base,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for BaseRefundError.
// It customizes the JSON marshaling process for BaseRefundError objects.
func (b *BaseRefundError) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BaseRefundError object to a map representation for JSON marshaling.
func (b *BaseRefundError) toMap() map[string]any {
	structMap := make(map[string]any)
	if b.Base != nil {
		structMap["base"] = b.Base
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BaseRefundError.
// It customizes the JSON unmarshaling process for BaseRefundError objects.
func (b *BaseRefundError) UnmarshalJSON(input []byte) error {
	var temp baseRefundError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	b.Base = temp.Base
	return nil
}

// TODO
type baseRefundError struct {
	Base []interface{} `json:"base,omitempty"`
}
