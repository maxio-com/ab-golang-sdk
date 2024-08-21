/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// BackportInvoiceEvent represents a BackportInvoiceEvent struct.
type BackportInvoiceEvent struct {
    Id                   int64            `json:"id"`
    Timestamp            time.Time        `json:"timestamp"`
    Invoice              Invoice          `json:"invoice"`
    EventType            InvoiceEventType `json:"event_type"`
    // Example schema for an `backport_invoice` event
    EventData            Invoice          `json:"event_data"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BackportInvoiceEvent.
// It customizes the JSON marshaling process for BackportInvoiceEvent objects.
func (b BackportInvoiceEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BackportInvoiceEvent object to a map representation for JSON marshaling.
func (b BackportInvoiceEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    structMap["id"] = b.Id
    structMap["timestamp"] = b.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = b.Invoice.toMap()
    structMap["event_type"] = b.EventType
    structMap["event_data"] = b.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BackportInvoiceEvent.
// It customizes the JSON unmarshaling process for BackportInvoiceEvent objects.
func (b *BackportInvoiceEvent) UnmarshalJSON(input []byte) error {
    var temp tempBackportInvoiceEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "timestamp", "invoice", "event_type", "event_data")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    b.Timestamp = TimestampVal
    b.Invoice = *temp.Invoice
    b.EventType = *temp.EventType
    b.EventData = *temp.EventData
    return nil
}

// tempBackportInvoiceEvent is a temporary struct used for validating the fields of BackportInvoiceEvent.
type tempBackportInvoiceEvent  struct {
    Id        *int64            `json:"id"`
    Timestamp *string           `json:"timestamp"`
    Invoice   *Invoice          `json:"invoice"`
    EventType *InvoiceEventType `json:"event_type"`
    EventData *Invoice          `json:"event_data"`
}

func (b *tempBackportInvoiceEvent) validate() error {
    var errs []string
    if b.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Backport Invoice Event`")
    }
    if b.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Backport Invoice Event`")
    }
    if b.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Backport Invoice Event`")
    }
    if b.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Backport Invoice Event`")
    }
    if b.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Backport Invoice Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
