package models

import (
    "encoding/json"
)

// RecordPaymentRequest represents a RecordPaymentRequest struct.
type RecordPaymentRequest struct {
    Payment CreatePayment `json:"payment"`
}

// MarshalJSON implements the json.Marshaler interface for RecordPaymentRequest.
// It customizes the JSON marshaling process for RecordPaymentRequest objects.
func (r *RecordPaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RecordPaymentRequest object to a map representation for JSON marshaling.
func (r *RecordPaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment"] = r.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecordPaymentRequest.
// It customizes the JSON unmarshaling process for RecordPaymentRequest objects.
func (r *RecordPaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Payment CreatePayment `json:"payment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.Payment = temp.Payment
    return nil
}
