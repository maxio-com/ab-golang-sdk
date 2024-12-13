/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// IssueServiceCredit represents a IssueServiceCredit struct.
type IssueServiceCredit struct {
    Amount               IssueServiceCreditAmount `json:"amount"`
    Memo                 *string                  `json:"memo,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCredit.
// It customizes the JSON marshaling process for IssueServiceCredit objects.
func (i IssueServiceCredit) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "amount", "memo"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCredit object to a map representation for JSON marshaling.
func (i IssueServiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["amount"] = i.Amount.toMap()
    if i.Memo != nil {
        structMap["memo"] = i.Memo
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCredit.
// It customizes the JSON unmarshaling process for IssueServiceCredit objects.
func (i *IssueServiceCredit) UnmarshalJSON(input []byte) error {
    var temp tempIssueServiceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amount", "memo")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Amount = *temp.Amount
    i.Memo = temp.Memo
    return nil
}

// tempIssueServiceCredit is a temporary struct used for validating the fields of IssueServiceCredit.
type tempIssueServiceCredit  struct {
    Amount *IssueServiceCreditAmount `json:"amount"`
    Memo   *string                   `json:"memo,omitempty"`
}

func (i *tempIssueServiceCredit) validate() error {
    var errs []string
    if i.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Issue Service Credit`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
