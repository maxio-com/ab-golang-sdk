package models

import (
	"encoding/json"
)

// IssueServiceCreditRequest represents a IssueServiceCreditRequest struct.
type IssueServiceCreditRequest struct {
	ServiceCredit IssueServiceCredit `json:"service_credit"`
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCreditRequest.
// It customizes the JSON marshaling process for IssueServiceCreditRequest objects.
func (i *IssueServiceCreditRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCreditRequest object to a map representation for JSON marshaling.
func (i *IssueServiceCreditRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["service_credit"] = i.ServiceCredit
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCreditRequest.
// It customizes the JSON unmarshaling process for IssueServiceCreditRequest objects.
func (i *IssueServiceCreditRequest) UnmarshalJSON(input []byte) error {
	temp := &struct {
		ServiceCredit IssueServiceCredit `json:"service_credit"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.ServiceCredit = temp.ServiceCredit
	return nil
}
