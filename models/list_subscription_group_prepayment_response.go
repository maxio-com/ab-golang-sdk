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

// ListSubscriptionGroupPrepaymentResponse represents a ListSubscriptionGroupPrepaymentResponse struct.
type ListSubscriptionGroupPrepaymentResponse struct {
    Prepayments          []ListSubscriptionGroupPrepayment `json:"prepayments"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionGroupPrepaymentResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionGroupPrepaymentResponse) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionGroupPrepaymentResponse[Prepayments=%v, AdditionalProperties=%v]",
    	l.Prepayments, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON marshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l ListSubscriptionGroupPrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "prepayments"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupPrepaymentResponse object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupPrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["prepayments"] = l.Prepayments
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l *ListSubscriptionGroupPrepaymentResponse) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionGroupPrepaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "prepayments")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Prepayments = *temp.Prepayments
    return nil
}

// tempListSubscriptionGroupPrepaymentResponse is a temporary struct used for validating the fields of ListSubscriptionGroupPrepaymentResponse.
type tempListSubscriptionGroupPrepaymentResponse  struct {
    Prepayments *[]ListSubscriptionGroupPrepayment `json:"prepayments"`
}

func (l *tempListSubscriptionGroupPrepaymentResponse) validate() error {
    var errs []string
    if l.Prepayments == nil {
        errs = append(errs, "required field `prepayments` is missing for type `List Subscription Group Prepayment Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
