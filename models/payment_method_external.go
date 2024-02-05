package models

import (
    "encoding/json"
)

// PaymentMethodExternal represents a PaymentMethodExternal struct.
type PaymentMethodExternal struct {
    Details *string                   `json:"details"`
    Kind    string                    `json:"kind"`
    Memo    *string                   `json:"memo"`
    Type    InvoiceEventPaymentMethod `json:"type"`
}

// MarshalJSON implements the json.Marshaler interface for PaymentMethodExternal.
// It customizes the JSON marshaling process for PaymentMethodExternal objects.
func (p *PaymentMethodExternal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PaymentMethodExternal object to a map representation for JSON marshaling.
func (p *PaymentMethodExternal) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["details"] = p.Details
    structMap["kind"] = p.Kind
    structMap["memo"] = p.Memo
    structMap["type"] = p.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentMethodExternal.
// It customizes the JSON unmarshaling process for PaymentMethodExternal objects.
func (p *PaymentMethodExternal) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Details *string                   `json:"details"`
        Kind    string                    `json:"kind"`
        Memo    *string                   `json:"memo"`
        Type    InvoiceEventPaymentMethod `json:"type"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    p.Details = temp.Details
    p.Kind = temp.Kind
    p.Memo = temp.Memo
    p.Type = temp.Type
    return nil
}
