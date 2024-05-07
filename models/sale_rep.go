package models

import (
    "encoding/json"
)

// SaleRep represents a SaleRep struct.
type SaleRep struct {
    Id                   *int                  `json:"id,omitempty"`
    FullName             *string               `json:"full_name,omitempty"`
    SubscriptionsCount   *int                  `json:"subscriptions_count,omitempty"`
    TestMode             *bool                 `json:"test_mode,omitempty"`
    Subscriptions        []SaleRepSubscription `json:"subscriptions,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SaleRep.
// It customizes the JSON marshaling process for SaleRep objects.
func (s SaleRep) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRep object to a map representation for JSON marshaling.
func (s SaleRep) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.FullName != nil {
        structMap["full_name"] = s.FullName
    }
    if s.SubscriptionsCount != nil {
        structMap["subscriptions_count"] = s.SubscriptionsCount
    }
    if s.TestMode != nil {
        structMap["test_mode"] = s.TestMode
    }
    if s.Subscriptions != nil {
        structMap["subscriptions"] = s.Subscriptions
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRep.
// It customizes the JSON unmarshaling process for SaleRep objects.
func (s *SaleRep) UnmarshalJSON(input []byte) error {
    var temp saleRep
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "full_name", "subscriptions_count", "test_mode", "subscriptions")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = temp.Id
    s.FullName = temp.FullName
    s.SubscriptionsCount = temp.SubscriptionsCount
    s.TestMode = temp.TestMode
    s.Subscriptions = temp.Subscriptions
    return nil
}

// saleRep is a temporary struct used for validating the fields of SaleRep.
type saleRep  struct {
    Id                 *int                  `json:"id,omitempty"`
    FullName           *string               `json:"full_name,omitempty"`
    SubscriptionsCount *int                  `json:"subscriptions_count,omitempty"`
    TestMode           *bool                 `json:"test_mode,omitempty"`
    Subscriptions      []SaleRepSubscription `json:"subscriptions,omitempty"`
}
