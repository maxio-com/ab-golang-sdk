// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AggregatedEntitlementsResponse represents a AggregatedEntitlementsResponse struct.
type AggregatedEntitlementsResponse struct {
    SubscriptionId       int                     `json:"subscription_id"`
    CustomerId           int                     `json:"customer_id"`
    // The subscription's current state, e.g. `active`, `trialing`, `canceled`.
    Status               string                  `json:"status"`
    Entitlements         []AggregatedEntitlement `json:"entitlements"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for AggregatedEntitlementsResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AggregatedEntitlementsResponse) String() string {
    return fmt.Sprintf(
    	"AggregatedEntitlementsResponse[SubscriptionId=%v, CustomerId=%v, Status=%v, Entitlements=%v, AdditionalProperties=%v]",
    	a.SubscriptionId, a.CustomerId, a.Status, a.Entitlements, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AggregatedEntitlementsResponse.
// It customizes the JSON marshaling process for AggregatedEntitlementsResponse objects.
func (a AggregatedEntitlementsResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "subscription_id", "customer_id", "status", "entitlements"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AggregatedEntitlementsResponse object to a map representation for JSON marshaling.
func (a AggregatedEntitlementsResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["subscription_id"] = a.SubscriptionId
    structMap["customer_id"] = a.CustomerId
    structMap["status"] = a.Status
    structMap["entitlements"] = a.Entitlements
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AggregatedEntitlementsResponse.
// It customizes the JSON unmarshaling process for AggregatedEntitlementsResponse objects.
func (a *AggregatedEntitlementsResponse) UnmarshalJSON(input []byte) error {
    var temp tempAggregatedEntitlementsResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "subscription_id", "customer_id", "status", "entitlements")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.SubscriptionId = *temp.SubscriptionId
    a.CustomerId = *temp.CustomerId
    a.Status = *temp.Status
    a.Entitlements = *temp.Entitlements
    return nil
}

// tempAggregatedEntitlementsResponse is a temporary struct used for validating the fields of AggregatedEntitlementsResponse.
type tempAggregatedEntitlementsResponse  struct {
    SubscriptionId *int                     `json:"subscription_id"`
    CustomerId     *int                     `json:"customer_id"`
    Status         *string                  `json:"status"`
    Entitlements   *[]AggregatedEntitlement `json:"entitlements"`
}

func (a *tempAggregatedEntitlementsResponse) validate() error {
    var errs []string
    if a.SubscriptionId == nil {
        errs = append(errs, "required field `subscription_id` is missing for type `Aggregated Entitlements Response`")
    }
    if a.CustomerId == nil {
        errs = append(errs, "required field `customer_id` is missing for type `Aggregated Entitlements Response`")
    }
    if a.Status == nil {
        errs = append(errs, "required field `status` is missing for type `Aggregated Entitlements Response`")
    }
    if a.Entitlements == nil {
        errs = append(errs, "required field `entitlements` is missing for type `Aggregated Entitlements Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
