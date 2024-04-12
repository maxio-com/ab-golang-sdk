package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// IssueServiceCreditRequest represents a IssueServiceCreditRequest struct.
type IssueServiceCreditRequest struct {
    ServiceCredit        IssueServiceCredit `json:"service_credit"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCreditRequest.
// It customizes the JSON marshaling process for IssueServiceCreditRequest objects.
func (i IssueServiceCreditRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCreditRequest object to a map representation for JSON marshaling.
func (i IssueServiceCreditRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["service_credit"] = i.ServiceCredit.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCreditRequest.
// It customizes the JSON unmarshaling process for IssueServiceCreditRequest objects.
func (i *IssueServiceCreditRequest) UnmarshalJSON(input []byte) error {
    var temp issueServiceCreditRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "service_credit")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.ServiceCredit = *temp.ServiceCredit
    return nil
}

// TODO
type issueServiceCreditRequest  struct {
    ServiceCredit *IssueServiceCredit `json:"service_credit"`
}

func (i *issueServiceCreditRequest) validate() error {
    var errs []string
    if i.ServiceCredit == nil {
        errs = append(errs, "required field `service_credit` is missing for type `Issue Service Credit Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
