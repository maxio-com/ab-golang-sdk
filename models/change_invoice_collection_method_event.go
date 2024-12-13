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

// ChangeInvoiceCollectionMethodEvent represents a ChangeInvoiceCollectionMethodEvent struct.
type ChangeInvoiceCollectionMethodEvent struct {
    Id                   int64                                  `json:"id"`
    Timestamp            time.Time                              `json:"timestamp"`
    Invoice              Invoice                                `json:"invoice"`
    EventType            InvoiceEventType                       `json:"event_type"`
    // Example schema for an `change_invoice_collection_method` event
    EventData            ChangeInvoiceCollectionMethodEventData `json:"event_data"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ChangeInvoiceCollectionMethodEvent.
// It customizes the JSON marshaling process for ChangeInvoiceCollectionMethodEvent objects.
func (c ChangeInvoiceCollectionMethodEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ChangeInvoiceCollectionMethodEvent object to a map representation for JSON marshaling.
func (c ChangeInvoiceCollectionMethodEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["id"] = c.Id
    structMap["timestamp"] = c.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = c.Invoice.toMap()
    structMap["event_type"] = c.EventType
    structMap["event_data"] = c.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChangeInvoiceCollectionMethodEvent.
// It customizes the JSON unmarshaling process for ChangeInvoiceCollectionMethodEvent objects.
func (c *ChangeInvoiceCollectionMethodEvent) UnmarshalJSON(input []byte) error {
    var temp tempChangeInvoiceCollectionMethodEvent
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
    c.AdditionalProperties = additionalProperties
    
    c.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    c.Timestamp = TimestampVal
    c.Invoice = *temp.Invoice
    c.EventType = *temp.EventType
    c.EventData = *temp.EventData
    return nil
}

// tempChangeInvoiceCollectionMethodEvent is a temporary struct used for validating the fields of ChangeInvoiceCollectionMethodEvent.
type tempChangeInvoiceCollectionMethodEvent  struct {
    Id        *int64                                  `json:"id"`
    Timestamp *string                                 `json:"timestamp"`
    Invoice   *Invoice                                `json:"invoice"`
    EventType *InvoiceEventType                       `json:"event_type"`
    EventData *ChangeInvoiceCollectionMethodEventData `json:"event_data"`
}

func (c *tempChangeInvoiceCollectionMethodEvent) validate() error {
    var errs []string
    if c.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Change Invoice Collection Method Event`")
    }
    if c.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Change Invoice Collection Method Event`")
    }
    if c.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Change Invoice Collection Method Event`")
    }
    if c.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Change Invoice Collection Method Event`")
    }
    if c.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Change Invoice Collection Method Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
