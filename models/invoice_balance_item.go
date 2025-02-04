/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// InvoiceBalanceItem represents a InvoiceBalanceItem struct.
type InvoiceBalanceItem struct {
    Uid                  *string                `json:"uid,omitempty"`
    Number               *string                `json:"number,omitempty"`
    OutstandingAmount    *string                `json:"outstanding_amount,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceBalanceItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceBalanceItem) String() string {
    return fmt.Sprintf(
    	"InvoiceBalanceItem[Uid=%v, Number=%v, OutstandingAmount=%v, AdditionalProperties=%v]",
    	i.Uid, i.Number, i.OutstandingAmount, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceBalanceItem.
// It customizes the JSON marshaling process for InvoiceBalanceItem objects.
func (i InvoiceBalanceItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "number", "outstanding_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceBalanceItem object to a map representation for JSON marshaling.
func (i InvoiceBalanceItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Uid != nil {
        structMap["uid"] = i.Uid
    }
    if i.Number != nil {
        structMap["number"] = i.Number
    }
    if i.OutstandingAmount != nil {
        structMap["outstanding_amount"] = i.OutstandingAmount
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceBalanceItem.
// It customizes the JSON unmarshaling process for InvoiceBalanceItem objects.
func (i *InvoiceBalanceItem) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceBalanceItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "number", "outstanding_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = temp.Uid
    i.Number = temp.Number
    i.OutstandingAmount = temp.OutstandingAmount
    return nil
}

// tempInvoiceBalanceItem is a temporary struct used for validating the fields of InvoiceBalanceItem.
type tempInvoiceBalanceItem  struct {
    Uid               *string `json:"uid,omitempty"`
    Number            *string `json:"number,omitempty"`
    OutstandingAmount *string `json:"outstanding_amount,omitempty"`
}
