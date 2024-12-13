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

// VoidInvoiceEvent represents a VoidInvoiceEvent struct.
type VoidInvoiceEvent struct {
    Id                   int64                  `json:"id"`
    Timestamp            time.Time              `json:"timestamp"`
    Invoice              Invoice                `json:"invoice"`
    EventType            InvoiceEventType       `json:"event_type"`
    // Example schema for an `void_invoice` event
    EventData            VoidInvoiceEventData   `json:"event_data"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VoidInvoiceEvent.
// It customizes the JSON marshaling process for VoidInvoiceEvent objects.
func (v VoidInvoiceEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VoidInvoiceEvent object to a map representation for JSON marshaling.
func (v VoidInvoiceEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["id"] = v.Id
    structMap["timestamp"] = v.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = v.Invoice.toMap()
    structMap["event_type"] = v.EventType
    structMap["event_data"] = v.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidInvoiceEvent.
// It customizes the JSON unmarshaling process for VoidInvoiceEvent objects.
func (v *VoidInvoiceEvent) UnmarshalJSON(input []byte) error {
    var temp tempVoidInvoiceEvent
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
    v.AdditionalProperties = additionalProperties
    
    v.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    v.Timestamp = TimestampVal
    v.Invoice = *temp.Invoice
    v.EventType = *temp.EventType
    v.EventData = *temp.EventData
    return nil
}

// tempVoidInvoiceEvent is a temporary struct used for validating the fields of VoidInvoiceEvent.
type tempVoidInvoiceEvent  struct {
    Id        *int64                `json:"id"`
    Timestamp *string               `json:"timestamp"`
    Invoice   *Invoice              `json:"invoice"`
    EventType *InvoiceEventType     `json:"event_type"`
    EventData *VoidInvoiceEventData `json:"event_data"`
}

func (v *tempVoidInvoiceEvent) validate() error {
    var errs []string
    if v.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Void Invoice Event`")
    }
    if v.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Void Invoice Event`")
    }
    if v.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Void Invoice Event`")
    }
    if v.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Void Invoice Event`")
    }
    if v.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Void Invoice Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
