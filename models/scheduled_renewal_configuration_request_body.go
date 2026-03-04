// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// ScheduledRenewalConfigurationRequestBody represents a ScheduledRenewalConfigurationRequestBody struct.
type ScheduledRenewalConfigurationRequestBody struct {
    // (Optional) Start of the renewal term.
    StartsAt             *time.Time             `json:"starts_at,omitempty"`
    // (Optional) End of the renewal term.
    EndsAt               *time.Time             `json:"ends_at,omitempty"`
    // (Optional) Lock-in date for the renewal.
    LockInAt             *time.Time             `json:"lock_in_at,omitempty"`
    // (Optional) Existing contract to associate with the scheduled renewal. Contracts must be enabled for your site.
    ContractId           *int                   `json:"contract_id,omitempty"`
    // (Optional) Set to true to create a new contract when contracts are enabled. Contracts must be enabled for your site.
    CreateNewContract    *bool                  `json:"create_new_contract,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalConfigurationRequestBody,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalConfigurationRequestBody) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalConfigurationRequestBody[StartsAt=%v, EndsAt=%v, LockInAt=%v, ContractId=%v, CreateNewContract=%v, AdditionalProperties=%v]",
    	s.StartsAt, s.EndsAt, s.LockInAt, s.ContractId, s.CreateNewContract, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalConfigurationRequestBody.
// It customizes the JSON marshaling process for ScheduledRenewalConfigurationRequestBody objects.
func (s ScheduledRenewalConfigurationRequestBody) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "starts_at", "ends_at", "lock_in_at", "contract_id", "create_new_contract"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalConfigurationRequestBody object to a map representation for JSON marshaling.
func (s ScheduledRenewalConfigurationRequestBody) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.StartsAt != nil {
        structMap["starts_at"] = s.StartsAt.Format(time.RFC3339)
    }
    if s.EndsAt != nil {
        structMap["ends_at"] = s.EndsAt.Format(time.RFC3339)
    }
    if s.LockInAt != nil {
        structMap["lock_in_at"] = s.LockInAt.Format(time.RFC3339)
    }
    if s.ContractId != nil {
        structMap["contract_id"] = s.ContractId
    }
    if s.CreateNewContract != nil {
        structMap["create_new_contract"] = s.CreateNewContract
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalConfigurationRequestBody.
// It customizes the JSON unmarshaling process for ScheduledRenewalConfigurationRequestBody objects.
func (s *ScheduledRenewalConfigurationRequestBody) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalConfigurationRequestBody
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "starts_at", "ends_at", "lock_in_at", "contract_id", "create_new_contract")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
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
    s.ContractId = temp.ContractId
    s.CreateNewContract = temp.CreateNewContract
    return nil
}

// tempScheduledRenewalConfigurationRequestBody is a temporary struct used for validating the fields of ScheduledRenewalConfigurationRequestBody.
type tempScheduledRenewalConfigurationRequestBody  struct {
    StartsAt          *string `json:"starts_at,omitempty"`
    EndsAt            *string `json:"ends_at,omitempty"`
    LockInAt          *string `json:"lock_in_at,omitempty"`
    ContractId        *int    `json:"contract_id,omitempty"`
    CreateNewContract *bool   `json:"create_new_contract,omitempty"`
}
