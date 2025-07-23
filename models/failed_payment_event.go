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

// FailedPaymentEvent represents a FailedPaymentEvent struct.
type FailedPaymentEvent struct {
	Id        int64            `json:"id"`
	Timestamp time.Time        `json:"timestamp"`
	Invoice   Invoice          `json:"invoice"`
	EventType InvoiceEventType `json:"event_type"`
	// Example schema for an `failed_payment` event
	EventData            FailedPaymentEventData `json:"event_data"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for FailedPaymentEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (f FailedPaymentEvent) String() string {
	return fmt.Sprintf(
		"FailedPaymentEvent[Id=%v, Timestamp=%v, Invoice=%v, EventType=%v, EventData=%v, AdditionalProperties=%v]",
		f.Id, f.Timestamp, f.Invoice, f.EventType, f.EventData, f.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for FailedPaymentEvent.
// It customizes the JSON marshaling process for FailedPaymentEvent objects.
func (f FailedPaymentEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(f.AdditionalProperties,
		"id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(f.toMap())
}

// toMap converts the FailedPaymentEvent object to a map representation for JSON marshaling.
func (f FailedPaymentEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, f.AdditionalProperties)
	structMap["id"] = f.Id
	structMap["timestamp"] = f.Timestamp.Format(time.RFC3339)
	structMap["invoice"] = f.Invoice.toMap()
	structMap["event_type"] = f.EventType
	structMap["event_data"] = f.EventData.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for FailedPaymentEvent.
// It customizes the JSON unmarshaling process for FailedPaymentEvent objects.
func (f *FailedPaymentEvent) UnmarshalJSON(input []byte) error {
	var temp tempFailedPaymentEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "timestamp", "invoice", "event_type", "event_data")
	if err != nil {
		return err
	}
	f.AdditionalProperties = additionalProperties

	f.Id = *temp.Id
	TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
	if err != nil {
		log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
	}
	f.Timestamp = TimestampVal
	f.Invoice = *temp.Invoice
	f.EventType = *temp.EventType
	f.EventData = *temp.EventData
	return nil
}

// tempFailedPaymentEvent is a temporary struct used for validating the fields of FailedPaymentEvent.
type tempFailedPaymentEvent struct {
	Id        *int64                  `json:"id"`
	Timestamp *string                 `json:"timestamp"`
	Invoice   *Invoice                `json:"invoice"`
	EventType *InvoiceEventType       `json:"event_type"`
	EventData *FailedPaymentEventData `json:"event_data"`
}

func (f *tempFailedPaymentEvent) validate() error {
	var errs []string
	if f.Id == nil {
		errs = append(errs, "required field `id` is missing for type `Failed Payment Event`")
	}
	if f.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `Failed Payment Event`")
	}
	if f.Invoice == nil {
		errs = append(errs, "required field `invoice` is missing for type `Failed Payment Event`")
	}
	if f.EventType == nil {
		errs = append(errs, "required field `event_type` is missing for type `Failed Payment Event`")
	}
	if f.EventData == nil {
		errs = append(errs, "required field `event_data` is missing for type `Failed Payment Event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
