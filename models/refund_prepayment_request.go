package models

import (
    "encoding/json"
)

// RefundPrepaymentRequest represents a RefundPrepaymentRequest struct.
type RefundPrepaymentRequest struct {
    Refund RefundPrepayment `json:"refund"`
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentRequest.
// It customizes the JSON marshaling process for RefundPrepaymentRequest objects.
func (r *RefundPrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentRequest object to a map representation for JSON marshaling.
func (r *RefundPrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["refund"] = r.Refund.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentRequest.
// It customizes the JSON unmarshaling process for RefundPrepaymentRequest objects.
func (r *RefundPrepaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Refund RefundPrepayment `json:"refund"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    r.Refund = temp.Refund
    return nil
}
