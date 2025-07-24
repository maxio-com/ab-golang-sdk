// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
)

// ReferralCode represents a ReferralCode struct.
type ReferralCode struct {
    Id                   *int                   `json:"id,omitempty"`
    SiteId               *int                   `json:"site_id,omitempty"`
    SubscriptionId       *int                   `json:"subscription_id,omitempty"`
    Code                 *string                `json:"code,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ReferralCode,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReferralCode) String() string {
    return fmt.Sprintf(
    	"ReferralCode[Id=%v, SiteId=%v, SubscriptionId=%v, Code=%v, AdditionalProperties=%v]",
    	r.Id, r.SiteId, r.SubscriptionId, r.Code, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ReferralCode.
// It customizes the JSON marshaling process for ReferralCode objects.
func (r ReferralCode) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "site_id", "subscription_id", "code"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReferralCode object to a map representation for JSON marshaling.
func (r ReferralCode) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.SubscriptionId != nil {
        structMap["subscription_id"] = r.SubscriptionId
    }
    if r.Code != nil {
        structMap["code"] = r.Code
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReferralCode.
// It customizes the JSON unmarshaling process for ReferralCode objects.
func (r *ReferralCode) UnmarshalJSON(input []byte) error {
    var temp tempReferralCode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "site_id", "subscription_id", "code")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = temp.Id
    r.SiteId = temp.SiteId
    r.SubscriptionId = temp.SubscriptionId
    r.Code = temp.Code
    return nil
}

// tempReferralCode is a temporary struct used for validating the fields of ReferralCode.
type tempReferralCode  struct {
    Id             *int    `json:"id,omitempty"`
    SiteId         *int    `json:"site_id,omitempty"`
    SubscriptionId *int    `json:"subscription_id,omitempty"`
    Code           *string `json:"code,omitempty"`
}
