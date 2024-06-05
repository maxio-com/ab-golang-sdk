package models

import (
    "encoding/json"
    "errors"
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
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VoidRemainderEvent.
// It customizes the JSON marshaling process for VoidRemainderEvent objects.
func (v VoidRemainderEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VoidRemainderEvent object to a map representation for JSON marshaling.
func (v VoidRemainderEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
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
    var temp voidRemainderEvent
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

// voidRemainderEvent is a temporary struct used for validating the fields of VoidRemainderEvent.
type voidRemainderEvent  struct {
    Id        *int64                  `json:"id"`
    Timestamp *string                 `json:"timestamp"`
    Invoice   *Invoice                `json:"invoice"`
    EventType *InvoiceEventType       `json:"event_type"`
    EventData *VoidRemainderEventData `json:"event_data"`
}

func (v *voidRemainderEvent) validate() error {
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
    return errors.New(strings.Join(errs, "\n"))
}
