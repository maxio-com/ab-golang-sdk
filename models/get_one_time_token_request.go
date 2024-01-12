package models

import (
    "encoding/json"
)

// GetOneTimeTokenRequest represents a GetOneTimeTokenRequest struct.
type GetOneTimeTokenRequest struct {
    PaymentProfile GetOneTimeTokenPaymentProfile `json:"payment_profile"`
}

// MarshalJSON implements the json.Marshaler interface for GetOneTimeTokenRequest.
// It customizes the JSON marshaling process for GetOneTimeTokenRequest objects.
func (g *GetOneTimeTokenRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(g.toMap())
}

// toMap converts the GetOneTimeTokenRequest object to a map representation for JSON marshaling.
func (g *GetOneTimeTokenRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["payment_profile"] = g.PaymentProfile
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for GetOneTimeTokenRequest.
// It customizes the JSON unmarshaling process for GetOneTimeTokenRequest objects.
func (g *GetOneTimeTokenRequest) UnmarshalJSON(input []byte) error {
    temp := &struct {
        PaymentProfile GetOneTimeTokenPaymentProfile `json:"payment_profile"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    g.PaymentProfile = temp.PaymentProfile
    return nil
}
