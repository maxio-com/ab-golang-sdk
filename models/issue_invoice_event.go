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

// IssueInvoiceEvent represents a IssueInvoiceEvent struct.
type IssueInvoiceEvent struct {
	Id        int64            `json:"id"`
	Timestamp time.Time        `json:"timestamp"`
	Invoice   Invoice          `json:"invoice"`
	EventType InvoiceEventType `json:"event_type"`
	// Example schema for an `issue_invoice` event
	EventData            IssueInvoiceEventData  `json:"event_data"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IssueInvoiceEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssueInvoiceEvent) String() string {
	return fmt.Sprintf(
		"IssueInvoiceEvent[Id=%v, Timestamp=%v, Invoice=%v, EventType=%v, EventData=%v, AdditionalProperties=%v]",
		i.Id, i.Timestamp, i.Invoice, i.EventType, i.EventData, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssueInvoiceEvent.
// It customizes the JSON marshaling process for IssueInvoiceEvent objects.
func (i IssueInvoiceEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the IssueInvoiceEvent object to a map representation for JSON marshaling.
func (i IssueInvoiceEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	structMap["id"] = i.Id
	structMap["timestamp"] = i.Timestamp.Format(time.RFC3339)
	structMap["invoice"] = i.Invoice.toMap()
	structMap["event_type"] = i.EventType
	structMap["event_data"] = i.EventData.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueInvoiceEvent.
// It customizes the JSON unmarshaling process for IssueInvoiceEvent objects.
func (i *IssueInvoiceEvent) UnmarshalJSON(input []byte) error {
	var temp tempIssueInvoiceEvent
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
	i.AdditionalProperties = additionalProperties

	i.Id = *temp.Id
	TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
	if err != nil {
		log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
	}
	i.Timestamp = TimestampVal
	i.Invoice = *temp.Invoice
	i.EventType = *temp.EventType
	i.EventData = *temp.EventData
	return nil
}

// tempIssueInvoiceEvent is a temporary struct used for validating the fields of IssueInvoiceEvent.
type tempIssueInvoiceEvent struct {
	Id        *int64                 `json:"id"`
	Timestamp *string                `json:"timestamp"`
	Invoice   *Invoice               `json:"invoice"`
	EventType *InvoiceEventType      `json:"event_type"`
	EventData *IssueInvoiceEventData `json:"event_data"`
}

func (i *tempIssueInvoiceEvent) validate() error {
	var errs []string
	if i.Id == nil {
		errs = append(errs, "required field `id` is missing for type `Issue Invoice Event`")
	}
	if i.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `Issue Invoice Event`")
	}
	if i.Invoice == nil {
		errs = append(errs, "required field `invoice` is missing for type `Issue Invoice Event`")
	}
	if i.EventType == nil {
		errs = append(errs, "required field `event_type` is missing for type `Issue Invoice Event`")
	}
	if i.EventData == nil {
		errs = append(errs, "required field `event_data` is missing for type `Issue Invoice Event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
