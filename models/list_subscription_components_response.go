package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ListSubscriptionComponentsResponse represents a ListSubscriptionComponentsResponse struct.
type ListSubscriptionComponentsResponse struct {
    SubscriptionsComponents []SubscriptionComponent `json:"subscriptions_components"`
    AdditionalProperties    map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON marshaling process for ListSubscriptionComponentsResponse objects.
func (l ListSubscriptionComponentsResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListSubscriptionComponentsResponse object to a map representation for JSON marshaling.
func (l ListSubscriptionComponentsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
    structMap["subscriptions_components"] = l.SubscriptionsComponents
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsResponse.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsResponse objects.
func (l *ListSubscriptionComponentsResponse) UnmarshalJSON(input []byte) error {
    var temp listSubscriptionComponentsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "subscriptions_components")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.SubscriptionsComponents = *temp.SubscriptionsComponents
    return nil
}

// TODO
type listSubscriptionComponentsResponse  struct {
    SubscriptionsComponents *[]SubscriptionComponent `json:"subscriptions_components"`
}

func (l *listSubscriptionComponentsResponse) validate() error {
    var errs []string
    if l.SubscriptionsComponents == nil {
        errs = append(errs, "required field `subscriptions_components` is missing for type `List Subscription Components Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
