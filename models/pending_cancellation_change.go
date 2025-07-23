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

// PendingCancellationChange represents a PendingCancellationChange struct.
type PendingCancellationChange struct {
	CancellationState    string                 `json:"cancellation_state"`
	CancelsAt            time.Time              `json:"cancels_at"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PendingCancellationChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PendingCancellationChange) String() string {
	return fmt.Sprintf(
		"PendingCancellationChange[CancellationState=%v, CancelsAt=%v, AdditionalProperties=%v]",
		p.CancellationState, p.CancelsAt, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PendingCancellationChange.
// It customizes the JSON marshaling process for PendingCancellationChange objects.
func (p PendingCancellationChange) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"cancellation_state", "cancels_at"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PendingCancellationChange object to a map representation for JSON marshaling.
func (p PendingCancellationChange) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	structMap["cancellation_state"] = p.CancellationState
	structMap["cancels_at"] = p.CancelsAt.Format(time.RFC3339)
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PendingCancellationChange.
// It customizes the JSON unmarshaling process for PendingCancellationChange objects.
func (p *PendingCancellationChange) UnmarshalJSON(input []byte) error {
	var temp tempPendingCancellationChange
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cancellation_state", "cancels_at")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.CancellationState = *temp.CancellationState
	CancelsAtVal, err := time.Parse(time.RFC3339, *temp.CancelsAt)
	if err != nil {
		log.Fatalf("Cannot Parse cancels_at as % s format.", time.RFC3339)
	}
	p.CancelsAt = CancelsAtVal
	return nil
}

// tempPendingCancellationChange is a temporary struct used for validating the fields of PendingCancellationChange.
type tempPendingCancellationChange struct {
	CancellationState *string `json:"cancellation_state"`
	CancelsAt         *string `json:"cancels_at"`
}

func (p *tempPendingCancellationChange) validate() error {
	var errs []string
	if p.CancellationState == nil {
		errs = append(errs, "required field `cancellation_state` is missing for type `Pending Cancellation Change`")
	}
	if p.CancelsAt == nil {
		errs = append(errs, "required field `cancels_at` is missing for type `Pending Cancellation Change`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
