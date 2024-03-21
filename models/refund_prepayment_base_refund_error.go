package models

import (
	"encoding/json"
)

// RefundPrepaymentBaseRefundError represents a RefundPrepaymentBaseRefundError struct.
type RefundPrepaymentBaseRefundError struct {
	Refund *BaseRefundError `json:"refund,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentBaseRefundError.
// It customizes the JSON marshaling process for RefundPrepaymentBaseRefundError objects.
func (r *RefundPrepaymentBaseRefundError) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentBaseRefundError object to a map representation for JSON marshaling.
func (r *RefundPrepaymentBaseRefundError) toMap() map[string]any {
	structMap := make(map[string]any)
	if r.Refund != nil {
		structMap["refund"] = r.Refund.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentBaseRefundError.
// It customizes the JSON unmarshaling process for RefundPrepaymentBaseRefundError objects.
func (r *RefundPrepaymentBaseRefundError) UnmarshalJSON(input []byte) error {
	var temp refundPrepaymentBaseRefundError
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	r.Refund = temp.Refund
	return nil
}

// TODO
type refundPrepaymentBaseRefundError struct {
	Refund *BaseRefundError `json:"refund,omitempty"`
}
