package models

import (
    "encoding/json"
)

// OriginInvoice represents a OriginInvoice struct.
type OriginInvoice struct {
    // The UID of the invoice serving as an origin invoice.
    Uid    *string `json:"uid,omitempty"`
    // The number of the invoice serving as an origin invoice.
    Number *string `json:"number,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for OriginInvoice.
// It customizes the JSON marshaling process for OriginInvoice objects.
func (o *OriginInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OriginInvoice object to a map representation for JSON marshaling.
func (o *OriginInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    if o.Uid != nil {
        structMap["uid"] = o.Uid
    }
    if o.Number != nil {
        structMap["number"] = o.Number
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OriginInvoice.
// It customizes the JSON unmarshaling process for OriginInvoice objects.
func (o *OriginInvoice) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Uid    *string `json:"uid,omitempty"`
        Number *string `json:"number,omitempty"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    o.Uid = temp.Uid
    o.Number = temp.Number
    return nil
}
