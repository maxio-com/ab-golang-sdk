package models

import (
    "encoding/json"
)

// ReferralCode represents a ReferralCode struct.
type ReferralCode struct {
    Id                   *int           `json:"id,omitempty"`
    SiteId               *int           `json:"site_id,omitempty"`
    SubscriptionId       *int           `json:"subscription_id,omitempty"`
    Code                 *string        `json:"code,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReferralCode.
// It customizes the JSON marshaling process for ReferralCode objects.
func (r ReferralCode) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ReferralCode object to a map representation for JSON marshaling.
func (r ReferralCode) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp referralCode
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "site_id", "subscription_id", "code")
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

// TODO
type referralCode  struct {
    Id             *int    `json:"id,omitempty"`
    SiteId         *int    `json:"site_id,omitempty"`
    SubscriptionId *int    `json:"subscription_id,omitempty"`
    Code           *string `json:"code,omitempty"`
}
