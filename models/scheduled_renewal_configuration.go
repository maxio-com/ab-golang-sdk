// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ScheduledRenewalConfiguration represents a ScheduledRenewalConfiguration struct.
type ScheduledRenewalConfiguration struct {
    // ID of the renewal.
    Id                                 *int                                `json:"id,omitempty"`
    // ID of the site to which the renewal belongs.
    SiteId                             *int                                `json:"site_id,omitempty"`
    // The id of the subscription.
    SubscriptionId                     *int                                `json:"subscription_id,omitempty"`
    StartsAt                           *time.Time                          `json:"starts_at,omitempty"`
    EndsAt                             *time.Time                          `json:"ends_at,omitempty"`
    LockInAt                           *time.Time                          `json:"lock_in_at,omitempty"`
    CreatedAt                          *time.Time                          `json:"created_at,omitempty"`
    Status                             *string                             `json:"status,omitempty"`
    ScheduledRenewalConfigurationItems []ScheduledRenewalConfigurationItem `json:"scheduled_renewal_configuration_items,omitempty"`
    // Contract linked to the scheduled renewal configuration.
    Contract                           *Contract                           `json:"contract,omitempty"`
    AdditionalProperties               map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfiguration,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfiguration) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfiguration[Id=%v, SiteId=%v, SubscriptionId=%v, StartsAt=%v, EndsAt=%v, LockInAt=%v, CreatedAt=%v, Status=%v, ScheduledRenewalConfigurationItems=%v, Contract=%v, AdditionalProperties=%v]",
    	s.Id, s.SiteId, s.SubscriptionId, s.StartsAt, s.EndsAt, s.LockInAt, s.CreatedAt, s.Status, s.ScheduledRenewalConfigurationItems, s.Contract, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfiguration.
// It customizes the JSON marshaling process for ScheduledRenewalConfiguration objects.
func (s ScheduledRenewalConfiguration) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "site_id", "subscription_id", "starts_at", "ends_at", "lock_in_at", "created_at", "status", "scheduled_renewal_configuration_items", "contract"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfiguration object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfiguration) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.SubscriptionId != nil {
        structMap["subscription_id"] = s.SubscriptionId
    }
    if s.StartsAt != nil {
        structMap["starts_at"] = s.StartsAt.Format(time.RFC3339)
    }
    if s.EndsAt != nil {
        structMap["ends_at"] = s.EndsAt.Format(time.RFC3339)
    }
    if s.LockInAt != nil {
        structMap["lock_in_at"] = s.LockInAt.Format(time.RFC3339)
    }
    if s.CreatedAt != nil {
        structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    if s.ScheduledRenewalConfigurationItems != nil {
        structMap["scheduled_renewal_configuration_items"] = s.ScheduledRenewalConfigurationItems
    }
    if s.Contract != nil {
        structMap["contract"] = s.Contract.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfiguration.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfiguration objects.
func (s *ScheduledRenewalConfiguration) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfiguration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "site_id", "subscription_id", "starts_at", "ends_at", "lock_in_at", "created_at", "status", "scheduled_renewal_configuration_items", "contract")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = temp.Id
    s.SiteId = temp.SiteId
    s.SubscriptionId = temp.SubscriptionId
    if temp.StartsAt != nil {
        StartsAtVal, err := time.Parse(time.RFC3339, *temp.StartsAt)
        if err != nil {
            log.Fatalf("Cannot Parse starts_at as % s format.", time.RFC3339)
        }
        s.StartsAt = &StartsAtVal
    }
    if temp.EndsAt != nil {
        EndsAtVal, err := time.Parse(time.RFC3339, *temp.EndsAt)
        if err != nil {
            log.Fatalf("Cannot Parse ends_at as % s format.", time.RFC3339)
        }
        s.EndsAt = &EndsAtVal
    }
    if temp.LockInAt != nil {
        LockInAtVal, err := time.Parse(time.RFC3339, *temp.LockInAt)
        if err != nil {
            log.Fatalf("Cannot Parse lock_in_at as % s format.", time.RFC3339)
        }
        s.LockInAt = &LockInAtVal
    }
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        s.CreatedAt = &CreatedAtVal
    }
    s.Status = temp.Status
    s.ScheduledRenewalConfigurationItems = temp.ScheduledRenewalConfigurationItems
    s.Contract = temp.Contract
    return nil
}

// tempScheduledRenewalConfiguration is a temporary struct used for validating the fields of ScheduledRenewalConfiguration.
type tempScheduledRenewalConfiguration  struct {
    Id                                 *int                                `json:"id,omitempty"`
    SiteId                             *int                                `json:"site_id,omitempty"`
    SubscriptionId                     *int                                `json:"subscription_id,omitempty"`
    StartsAt                           *string                             `json:"starts_at,omitempty"`
    EndsAt                             *string                             `json:"ends_at,omitempty"`
    LockInAt                           *string                             `json:"lock_in_at,omitempty"`
    CreatedAt                          *string                             `json:"created_at,omitempty"`
    Status                             *string                             `json:"status,omitempty"`
    ScheduledRenewalConfigurationItems []ScheduledRenewalConfigurationItem `json:"scheduled_renewal_configuration_items,omitempty"`
    Contract                           *Contract                           `json:"contract,omitempty"`
}
