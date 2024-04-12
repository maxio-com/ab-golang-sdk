package models

import (
    "encoding/json"
)

// RefundPrepaymentAggregatedError represents a RefundPrepaymentAggregatedError struct.
type RefundPrepaymentAggregatedError struct {
    Refund               *PrepaymentAggregatedError `json:"refund,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RefundPrepaymentAggregatedError.
// It customizes the JSON marshaling process for RefundPrepaymentAggregatedError objects.
func (r RefundPrepaymentAggregatedError) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundPrepaymentAggregatedError object to a map representation for JSON marshaling.
func (r RefundPrepaymentAggregatedError) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Refund != nil {
        structMap["refund"] = r.Refund.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundPrepaymentAggregatedError.
// It customizes the JSON unmarshaling process for RefundPrepaymentAggregatedError objects.
func (r *RefundPrepaymentAggregatedError) UnmarshalJSON(input []byte) error {
    var temp refundPrepaymentAggregatedError
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "refund")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Refund = temp.Refund
    return nil
}

// TODO
type refundPrepaymentAggregatedError  struct {
    Refund *PrepaymentAggregatedError `json:"refund,omitempty"`
}
