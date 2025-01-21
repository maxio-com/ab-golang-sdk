/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// VoidRemainderEventData represents a VoidRemainderEventData struct.
// Example schema for an `void_remainder` event
type VoidRemainderEventData struct {
    CreditNoteAttributes CreditNote             `json:"credit_note_attributes"`
    // The memo provided during invoice remainder voiding.
    Memo                 string                 `json:"memo"`
    // The amount of the void.
    AppliedAmount        string                 `json:"applied_amount"`
    // The time the refund was applied, in ISO 8601 format, i.e. "2019-06-07T17:20:06Z"
    TransactionTime      time.Time              `json:"transaction_time"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VoidRemainderEventData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VoidRemainderEventData) String() string {
    return fmt.Sprintf(
    	"VoidRemainderEventData[CreditNoteAttributes=%v, Memo=%v, AppliedAmount=%v, TransactionTime=%v, AdditionalProperties=%v]",
    	v.CreditNoteAttributes, v.Memo, v.AppliedAmount, v.TransactionTime, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VoidRemainderEventData.
// It customizes the JSON marshaling process for VoidRemainderEventData objects.
func (v VoidRemainderEventData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "credit_note_attributes", "memo", "applied_amount", "transaction_time"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VoidRemainderEventData object to a map representation for JSON marshaling.
func (v VoidRemainderEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["credit_note_attributes"] = v.CreditNoteAttributes.toMap()
    structMap["memo"] = v.Memo
    structMap["applied_amount"] = v.AppliedAmount
    structMap["transaction_time"] = v.TransactionTime.Format(time.RFC3339)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidRemainderEventData.
// It customizes the JSON unmarshaling process for VoidRemainderEventData objects.
func (v *VoidRemainderEventData) UnmarshalJSON(input []byte) error {
    var temp tempVoidRemainderEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "credit_note_attributes", "memo", "applied_amount", "transaction_time")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.CreditNoteAttributes = *temp.CreditNoteAttributes
    v.Memo = *temp.Memo
    v.AppliedAmount = *temp.AppliedAmount
    TransactionTimeVal, err := time.Parse(time.RFC3339, *temp.TransactionTime)
    if err != nil {
        log.Fatalf("Cannot Parse transaction_time as % s format.", time.RFC3339)
    }
    v.TransactionTime = TransactionTimeVal
    return nil
}

// tempVoidRemainderEventData is a temporary struct used for validating the fields of VoidRemainderEventData.
type tempVoidRemainderEventData  struct {
    CreditNoteAttributes *CreditNote `json:"credit_note_attributes"`
    Memo                 *string     `json:"memo"`
    AppliedAmount        *string     `json:"applied_amount"`
    TransactionTime      *string     `json:"transaction_time"`
}

func (v *tempVoidRemainderEventData) validate() error {
    var errs []string
    if v.CreditNoteAttributes == nil {
        errs = append(errs, "required field `credit_note_attributes` is missing for type `Void Remainder Event Data`")
    }
    if v.Memo == nil {
        errs = append(errs, "required field `memo` is missing for type `Void Remainder Event Data`")
    }
    if v.AppliedAmount == nil {
        errs = append(errs, "required field `applied_amount` is missing for type `Void Remainder Event Data`")
    }
    if v.TransactionTime == nil {
        errs = append(errs, "required field `transaction_time` is missing for type `Void Remainder Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
