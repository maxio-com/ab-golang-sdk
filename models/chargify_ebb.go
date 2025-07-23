// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ChargifyEBB represents a ChargifyEBB struct.
type ChargifyEBB struct {
    // This timestamp determines what billing period the event will be billed in. If your request payload does not include it, Chargify will add `chargify.timestamp` to the event payload and set the value to `now`.
    Timestamp             *time.Time             `json:"timestamp,omitempty"`
    // A unique ID set by Chargify. Please note that this field is reserved. If `chargify.id` is present in the request payload, it will be overwritten.
    Id                    *string                `json:"id,omitempty"`
    // An ISO-8601 timestamp, set by Chargify at the time each event is recorded. Please note that this field is reserved. If `chargify.created_at` is present in the request payload, it will be overwritten.
    CreatedAt             *time.Time             `json:"created_at,omitempty"`
    // User-defined string scoped per-stream. Duplicate events within a stream will be silently ignored. Tokens expire after 31 days.
    UniquenessToken       *string                `json:"uniqueness_token,omitempty"`
    // Id of Maxio Advanced Billing Subscription which is connected to this event.
    // Provide `subscription_id` if you configured `chargify.subscription_id` as Subscription Identifier in your Event Stream.
    SubscriptionId        *int                   `json:"subscription_id,omitempty"`
    // Reference of Maxio Advanced Billing Subscription which is connected to this event.
    // Provide `subscription_reference` if you configured `chargify.subscription_reference` as Subscription Identifier in your Event Stream.
    SubscriptionReference *string                `json:"subscription_reference,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ChargifyEBB,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ChargifyEBB) String() string {
    return fmt.Sprintf(
    	"ChargifyEBB[Timestamp=%v, Id=%v, CreatedAt=%v, UniquenessToken=%v, SubscriptionId=%v, SubscriptionReference=%v, AdditionalProperties=%v]",
    	c.Timestamp, c.Id, c.CreatedAt, c.UniquenessToken, c.SubscriptionId, c.SubscriptionReference, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ChargifyEBB.
// It customizes the JSON marshaling process for ChargifyEBB objects.
func (c ChargifyEBB) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "timestamp", "id", "created_at", "uniqueness_token", "subscription_id", "subscription_reference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ChargifyEBB object to a map representation for JSON marshaling.
func (c ChargifyEBB) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp.Format(time.RFC3339)
    }
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.CreatedAt != nil {
        structMap["created_at"] = c.CreatedAt.Format(time.RFC3339)
    }
    if c.UniquenessToken != nil {
        structMap["uniqueness_token"] = c.UniquenessToken
    }
    if c.SubscriptionId != nil {
        structMap["subscription_id"] = c.SubscriptionId
    }
    if c.SubscriptionReference != nil {
        structMap["subscription_reference"] = c.SubscriptionReference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChargifyEBB.
// It customizes the JSON unmarshaling process for ChargifyEBB objects.
func (c *ChargifyEBB) UnmarshalJSON(input []byte) error {
    var temp tempChargifyEBB
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "timestamp", "id", "created_at", "uniqueness_token", "subscription_id", "subscription_reference")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    if temp.Timestamp != nil {
        TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
        if err != nil {
            log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
        }
        c.Timestamp = &TimestampVal
    }
    c.Id = temp.Id
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        c.CreatedAt = &CreatedAtVal
    }
    c.UniquenessToken = temp.UniquenessToken
    c.SubscriptionId = temp.SubscriptionId
    c.SubscriptionReference = temp.SubscriptionReference
    return nil
}

// tempChargifyEBB is a temporary struct used for validating the fields of ChargifyEBB.
type tempChargifyEBB  struct {
    Timestamp             *string `json:"timestamp,omitempty"`
    Id                    *string `json:"id,omitempty"`
    CreatedAt             *string `json:"created_at,omitempty"`
    UniquenessToken       *string `json:"uniqueness_token,omitempty"`
    SubscriptionId        *int    `json:"subscription_id,omitempty"`
    SubscriptionReference *string `json:"subscription_reference,omitempty"`
}
