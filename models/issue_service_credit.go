package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// IssueServiceCredit represents a IssueServiceCredit struct.
type IssueServiceCredit struct {
    Amount               IssueServiceCreditAmount `json:"amount"`
    Memo                 string                   `json:"memo"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCredit.
// It customizes the JSON marshaling process for IssueServiceCredit objects.
func (i IssueServiceCredit) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCredit object to a map representation for JSON marshaling.
func (i IssueServiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["amount"] = i.Amount.toMap()
    structMap["memo"] = i.Memo
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCredit.
// It customizes the JSON unmarshaling process for IssueServiceCredit objects.
func (i *IssueServiceCredit) UnmarshalJSON(input []byte) error {
    var temp issueServiceCredit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amount", "memo")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Amount = *temp.Amount
    i.Memo = *temp.Memo
    return nil
}

// TODO
type issueServiceCredit  struct {
    Amount *IssueServiceCreditAmount `json:"amount"`
    Memo   *string                   `json:"memo"`
}

func (i *issueServiceCredit) validate() error {
    var errs []string
    if i.Amount == nil {
        errs = append(errs, "required field `amount` is missing for type `Issue Service Credit`")
    }
    if i.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Issue Service Credit`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
