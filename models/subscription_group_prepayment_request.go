package models

import (
    "encoding/json"
)

// SubscriptionGroupPrepaymentRequest represents a SubscriptionGroupPrepaymentRequest struct.
type SubscriptionGroupPrepaymentRequest struct {
    Prepayment SubscriptionGroupPrepayment `json:"prepayment"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepaymentRequest.
// It customizes the JSON marshaling process for SubscriptionGroupPrepaymentRequest objects.
func (s *SubscriptionGroupPrepaymentRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepaymentRequest object to a map representation for JSON marshaling.
func (s *SubscriptionGroupPrepaymentRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["prepayment"] = s.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepaymentRequest.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepaymentRequest objects.
func (s *SubscriptionGroupPrepaymentRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Prepayment SubscriptionGroupPrepayment `json:"prepayment"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Prepayment = temp.Prepayment
    return nil
}
