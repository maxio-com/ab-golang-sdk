// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// IssueServiceCreditRequest represents a IssueServiceCreditRequest struct.
type IssueServiceCreditRequest struct {
    ServiceCredit        IssueServiceCredit     `json:"service_credit"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IssueServiceCreditRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssueServiceCreditRequest) String() string {
    return fmt.Sprintf(
    	"IssueServiceCreditRequest[ServiceCredit=%v, AdditionalProperties=%v]",
    	i.ServiceCredit, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCreditRequest.
// It customizes the JSON marshaling process for IssueServiceCreditRequest objects.
func (i IssueServiceCreditRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "service_credit"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCreditRequest object to a map representation for JSON marshaling.
func (i IssueServiceCreditRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["service_credit"] = i.ServiceCredit.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCreditRequest.
// It customizes the JSON unmarshaling process for IssueServiceCreditRequest objects.
func (i *IssueServiceCreditRequest) UnmarshalJSON(input []byte) error {
    var temp tempIssueServiceCreditRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_credit")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.ServiceCredit = *temp.ServiceCredit
    return nil
}

// tempIssueServiceCreditRequest is a temporary struct used for validating the fields of IssueServiceCreditRequest.
type tempIssueServiceCreditRequest  struct {
    ServiceCredit *IssueServiceCredit `json:"service_credit"`
}

func (i *tempIssueServiceCreditRequest) validate() error {
    var errs []string
    if i.ServiceCredit == nil {
        errs = append(errs, "required field `service_credit` is missing for type `Issue Service Credit Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
