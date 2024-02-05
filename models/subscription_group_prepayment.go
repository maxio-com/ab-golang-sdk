package models

import (
    "encoding/json"
)

// SubscriptionGroupPrepayment represents a SubscriptionGroupPrepayment struct.
type SubscriptionGroupPrepayment struct {
    Amount  int                               `json:"amount"`
    Details string                            `json:"details"`
    Memo    string                            `json:"memo"`
    Method  SubscriptionGroupPrepaymentMethod `json:"method"`
}

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepayment.
// It customizes the JSON marshaling process for SubscriptionGroupPrepayment objects.
func (s *SubscriptionGroupPrepayment) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SubscriptionGroupPrepayment object to a map representation for JSON marshaling.
func (s *SubscriptionGroupPrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = s.Amount
    structMap["details"] = s.Details
    structMap["memo"] = s.Memo
    structMap["method"] = s.Method
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepayment.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepayment objects.
func (s *SubscriptionGroupPrepayment) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount  int                               `json:"amount"`
        Details string                            `json:"details"`
        Memo    string                            `json:"memo"`
        Method  SubscriptionGroupPrepaymentMethod `json:"method"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    s.Amount = temp.Amount
    s.Details = temp.Details
    s.Memo = temp.Memo
    s.Method = temp.Method
    return nil
}
