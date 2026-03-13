// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// ScheduledRenewalLockInRequest represents a ScheduledRenewalLockInRequest struct.
type ScheduledRenewalLockInRequest struct {
    // Date to lock in the renewal.
    LockInAt             time.Time              `json:"lock_in_at"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ScheduledRenewalLockInRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ScheduledRenewalLockInRequest) String() string {
    return fmt.Sprintf(
    	"ScheduledRenewalLockInRequest[LockInAt=%v, AdditionalProperties=%v]",
    	s.LockInAt, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ScheduledRenewalLockInRequest.
// It customizes the JSON marshaling process for ScheduledRenewalLockInRequest objects.
func (s ScheduledRenewalLockInRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "lock_in_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the ScheduledRenewalLockInRequest object to a map representation for JSON marshaling.
func (s ScheduledRenewalLockInRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["lock_in_at"] = s.LockInAt.Format(DEFAULT_DATE)
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ScheduledRenewalLockInRequest.
// It customizes the JSON unmarshaling process for ScheduledRenewalLockInRequest objects.
func (s *ScheduledRenewalLockInRequest) UnmarshalJSON(input []byte) error {
    var temp tempScheduledRenewalLockInRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "lock_in_at")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    LockInAtVal, err := time.Parse(DEFAULT_DATE, *temp.LockInAt)
    if err != nil {
        log.Fatalf("Cannot Parse lock_in_at as % s format.", DEFAULT_DATE)
    }
    s.LockInAt = LockInAtVal
    return nil
}

// tempScheduledRenewalLockInRequest is a temporary struct used for validating the fields of ScheduledRenewalLockInRequest.
type tempScheduledRenewalLockInRequest  struct {
    LockInAt *string `json:"lock_in_at"`
}

func (s *tempScheduledRenewalLockInRequest) validate() error {
    var errs []string
    if s.LockInAt == nil {
        errs = append(errs, "required field `lock_in_at` is missing for type `Scheduled Renewal Lock In Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
