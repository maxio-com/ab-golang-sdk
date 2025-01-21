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

// ListSubscriptionComponentsResponse represents a ListSubscriptionComponentsResponse struct.
type ListSubscriptionComponentsResponse struct {
    SubscriptionsComponents []SubscriptionComponent `json:"subscriptions_components"`
    AdditionalProperties    map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for ListSubscriptionComponentsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListSubscriptionComponentsResponse) String() string {
    return fmt.Sprintf(
    	"ListSubscriptionComponentsResponse[SubscriptionsComponents=%v, AdditionalProperties=%v]",
    	l.SubscriptionsComponents, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON marshaling process for ListSubscriptionComponentsResponse objects.
func (l ListSubscriptionComponentsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "subscriptions_components"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsResponse object to a map representation for JSON marshaling.
func (l ListSubscriptionComponentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["subscriptions_components"] = l.SubscriptionsComponents
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsResponse objects.
func (l *ListSubscriptionComponentsResponse) UnmarshalJSON(input []byte) error {
    var temp tempListSubscriptionComponentsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscriptions_components")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.SubscriptionsComponents = *temp.SubscriptionsComponents
    return nil
}

// tempListSubscriptionComponentsResponse is a temporary struct used for validating the fields of ListSubscriptionComponentsResponse.
type tempListSubscriptionComponentsResponse  struct {
    SubscriptionsComponents *[]SubscriptionComponent `json:"subscriptions_components"`
}

func (l *tempListSubscriptionComponentsResponse) validate() error {
    var errs []string
    if l.SubscriptionsComponents == nil {
        errs = append(errs, "required field `subscriptions_components` is missing for type `List Subscription Components Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
