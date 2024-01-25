package models

import (
	"encoding/json"
)

// VoidInvoice represents a VoidInvoice struct.
type VoidInvoice struct {
	Reason string `json:"reason"`
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoice.
// It customizes the JSON marshaling process for VoidInvoice objects.
func (v *VoidInvoice) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoice object to a map representation for JSON marshaling.
func (v *VoidInvoice) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["reason"] = v.Reason
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoice.
// It customizes the JSON unmarshaling process for VoidInvoice objects.
func (v *VoidInvoice) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Reason string `json:"reason"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	v.Reason = temp.Reason
	return nil
}
