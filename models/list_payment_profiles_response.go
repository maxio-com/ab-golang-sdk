package models

import (
	"encoding/json"
)

// ListPaymentProfilesResponse represents a ListPaymentProfilesResponse struct.
type ListPaymentProfilesResponse struct {
	PaymentProfile *ListPaymentProfileItem `json:"payment_profile,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ListPaymentProfilesResponse.
// It customizes the JSON marshaling process for ListPaymentProfilesResponse objects.
func (l *ListPaymentProfilesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the ListPaymentProfilesResponse object to a map representation for JSON marshaling.
func (l *ListPaymentProfilesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if l.PaymentProfile != nil {
		structMap["payment_profile"] = l.PaymentProfile
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPaymentProfilesResponse.
// It customizes the JSON unmarshaling process for ListPaymentProfilesResponse objects.
func (l *ListPaymentProfilesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		PaymentProfile *ListPaymentProfileItem `json:"payment_profile,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	l.PaymentProfile = temp.PaymentProfile
	return nil
}
