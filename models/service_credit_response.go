package models

import (
	"encoding/json"
)

// ServiceCreditResponse represents a ServiceCreditResponse struct.
type ServiceCreditResponse struct {
	ServiceCredit ServiceCredit `json:"service_credit"`
}

// MarshalJSON implements the json.Marshaler interface for ServiceCreditResponse.
// It customizes the JSON marshaling process for ServiceCreditResponse objects.
func (s *ServiceCreditResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the ServiceCreditResponse object to a map representation for JSON marshaling.
func (s *ServiceCreditResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["service_credit"] = s.ServiceCredit
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCreditResponse.
// It customizes the JSON unmarshaling process for ServiceCreditResponse objects.
func (s *ServiceCreditResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ServiceCredit ServiceCredit `json:"service_credit"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.ServiceCredit = temp.ServiceCredit
	return nil
}
