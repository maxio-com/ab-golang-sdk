package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListSubscriptionGroupPrepaymentResponse represents a ListSubscriptionGroupPrepaymentResponse struct.
type ListSubscriptionGroupPrepaymentResponse struct {
    Prepayments          []ListSubscriptionGroupPrepayment `json:"prepayments"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON marshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l ListSubscriptionGroupPrepaymentResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionGroupPrepaymentResponse object to a map representation for JSON marshaling.
func (l ListSubscriptionGroupPrepaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["prepayments"] = l.Prepayments
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionGroupPrepaymentResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionGroupPrepaymentResponse objects.
func (l *ListSubscriptionGroupPrepaymentResponse) UnmarshalJSON(input []byte) error {
    var temp listSubscriptionGroupPrepaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "prepayments")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.Prepayments = *temp.Prepayments
    return nil
}

// TODO
type listSubscriptionGroupPrepaymentResponse  struct {
    Prepayments *[]ListSubscriptionGroupPrepayment `json:"prepayments"`
}

func (l *listSubscriptionGroupPrepaymentResponse) validate() error {
    var errs []string
    if l.Prepayments == nil {
        errs = append(errs, "required field `prepayments` is missing for type `List Subscription Group Prepayment Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
