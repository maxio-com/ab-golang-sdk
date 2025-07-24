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

// VoidRemainderEvent represents a VoidRemainderEvent struct.
type VoidRemainderEvent struct {
    Id                   int64                  `json:"id"`
    Timestamp            time.Time              `json:"timestamp"`
    Invoice              Invoice                `json:"invoice"`
    EventType            InvoiceEventType       `json:"event_type"`
    // Example schema for an `void_remainder` event
    EventData            VoidRemainderEventData `json:"event_data"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for VoidRemainderEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VoidRemainderEvent) String() string {
    return fmt.Sprintf(
    	"VoidRemainderEvent[Id=%v, Timestamp=%v, Invoice=%v, EventType=%v, EventData=%v, AdditionalProperties=%v]",
    	v.Id, v.Timestamp, v.Invoice, v.EventType, v.EventData, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VoidRemainderEvent.
// It customizes the JSON marshaling process for VoidRemainderEvent objects.
func (v VoidRemainderEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VoidRemainderEvent object to a map representation for JSON marshaling.
func (v VoidRemainderEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["id"] = v.Id
    structMap["timestamp"] = v.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = v.Invoice.toMap()
    structMap["event_type"] = v.EventType
    structMap["event_data"] = v.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VoidRemainderEvent.
// It customizes the JSON unmarshaling process for VoidRemainderEvent objects.
func (v *VoidRemainderEvent) UnmarshalJSON(input []byte) error {
    var temp tempVoidRemainderEvent
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

// tempVoidRemainderEvent is a temporary struct used for validating the fields of VoidRemainderEvent.
type tempVoidRemainderEvent  struct {
    Id        *int64                  `json:"id"`
    Timestamp *string                 `json:"timestamp"`
    Invoice   *Invoice                `json:"invoice"`
    EventType *InvoiceEventType       `json:"event_type"`
    EventData *VoidRemainderEventData `json:"event_data"`
}

func (v *tempVoidRemainderEvent) validate() error {
    var errs []string
    if v.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Void Remainder Event`")
    }
    if v.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Void Remainder Event`")
    }
    if v.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Void Remainder Event`")
    }
    if v.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Void Remainder Event`")
    }
    if v.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Void Remainder Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
