package models

import (
    "encoding/json"
)

// IssueServiceCredit represents a IssueServiceCredit struct.
type IssueServiceCredit struct {
    Amount interface{} `json:"amount"`
    Memo   string      `json:"memo"`
}

// MarshalJSON implements the json.Marshaler interface for IssueServiceCredit.
// It customizes the JSON marshaling process for IssueServiceCredit objects.
func (i *IssueServiceCredit) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the IssueServiceCredit object to a map representation for JSON marshaling.
func (i *IssueServiceCredit) toMap() map[string]any {
    structMap := make(map[string]any)
    structMap["amount"] = i.Amount
    structMap["memo"] = i.Memo
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueServiceCredit.
// It customizes the JSON unmarshaling process for IssueServiceCredit objects.
func (i *IssueServiceCredit) UnmarshalJSON(input []byte) error {
    temp := &struct {
        Amount interface{} `json:"amount"`
        Memo   string      `json:"memo"`
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    i.Amount = temp.Amount
    i.Memo = temp.Memo
    return nil
}
