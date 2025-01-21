/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ListSubscriptionGroupPrepayment represents a ListSubscriptionGroupPrepayment struct.
type ListSubscriptionGroupPrepayment struct {
    Prepayment           ListSubcriptionGroupPrepaymentItem `json:"prepayment"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionGroupPrepayment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionGroupPrepayment) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionGroupPrepayment[Prepayment=%v, AdditionalProperties=%v]",
    	l.Prepayment, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupPrepayment.
// It customizes the JSON marshaling process for ListSubscriptionGroupPrepayment objects.
func (l ListSubscriptionGroupPrepayment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "prepayment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupPrepayment object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupPrepayment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["prepayment"] = l.Prepayment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepayment.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepayment objects.
func (l *ListSubscriptionGroupPrepayment) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionGroupPrepayment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayment")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Prepayment = *temp.Prepayment
    return nil
}

// tempListSubscriptionGroupPrepayment is a temporary struct used for validating the fields of ListSubscriptionGroupPrepayment.
type tempListSubscriptionGroupPrepayment  struct {
    Prepayment *ListSubcriptionGroupPrepaymentItem `json:"prepayment"`
}

func (l *tempListSubscriptionGroupPrepayment) validate() error {
    var errs []string
    if l.Prepayment == nil {
        errs = append(errs, "required field `prepayment` is missing for type `List Subscription Group Prepayment`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
