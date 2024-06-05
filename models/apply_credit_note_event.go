package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// ApplyCreditNoteEvent represents a ApplyCreditNoteEvent struct.
type ApplyCreditNoteEvent struct {
    Id                   int64                    `json:"id"`
    Timestamp            time.Time                `json:"timestamp"`
    Invoice              Invoice                  `json:"invoice"`
    EventType            InvoiceEventType         `json:"event_type"`
    // Example schema for an `apply_credit_note` event
    EventData            ApplyCreditNoteEventData `json:"event_data"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApplyCreditNoteEvent.
// It customizes the JSON marshaling process for ApplyCreditNoteEvent objects.
func (a ApplyCreditNoteEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApplyCreditNoteEvent object to a map representation for JSON marshaling.
func (a ApplyCreditNoteEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["id"] = a.Id
    structMap["timestamp"] = a.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = a.Invoice.toMap()
    structMap["event_type"] = a.EventType
    structMap["event_data"] = a.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplyCreditNoteEvent.
// It customizes the JSON unmarshaling process for ApplyCreditNoteEvent objects.
func (a *ApplyCreditNoteEvent) UnmarshalJSON(input []byte) error {
    var temp applyCreditNoteEvent
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
    
    a.AdditionalProperties = additionalProperties
    a.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    a.Timestamp = TimestampVal
    a.Invoice = *temp.Invoice
    a.EventType = *temp.EventType
    a.EventData = *temp.EventData
    return nil
}

// applyCreditNoteEvent is a temporary struct used for validating the fields of ApplyCreditNoteEvent.
type applyCreditNoteEvent  struct {
    Id        *int64                    `json:"id"`
    Timestamp *string                   `json:"timestamp"`
    Invoice   *Invoice                  `json:"invoice"`
    EventType *InvoiceEventType         `json:"event_type"`
    EventData *ApplyCreditNoteEventData `json:"event_data"`
}

func (a *applyCreditNoteEvent) validate() error {
    var errs []string
    if a.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Apply Credit Note Event`")
    }
    if a.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Apply Credit Note Event`")
    }
    if a.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Apply Credit Note Event`")
    }
    if a.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Apply Credit Note Event`")
    }
    if a.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Apply Credit Note Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
