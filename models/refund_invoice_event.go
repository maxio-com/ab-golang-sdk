/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// RefundInvoiceEvent represents a RefundInvoiceEvent struct.
type RefundInvoiceEvent struct {
    Id                   int64                  `json:"id"`
    Timestamp            time.Time              `json:"timestamp"`
    Invoice              Invoice                `json:"invoice"`
    EventType            InvoiceEventType       `json:"event_type"`
    // Example schema for an `refund_invoice` event
    EventData            RefundInvoiceEventData `json:"event_data"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RefundInvoiceEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RefundInvoiceEvent) String() string {
    return fmt.Sprintf(
    	"RefundInvoiceEvent[Id=%v, Timestamp=%v, Invoice=%v, EventType=%v, EventData=%v, AdditionalProperties=%v]",
    	r.Id, r.Timestamp, r.Invoice, r.EventType, r.EventData, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceEvent.
// It customizes the JSON marshaling process for RefundInvoiceEvent objects.
func (r RefundInvoiceEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "timestamp", "invoice", "event_type", "event_data"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceEvent object to a map representation for JSON marshaling.
func (r RefundInvoiceEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["id"] = r.Id
    structMap["timestamp"] = r.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = r.Invoice.toMap()
    structMap["event_type"] = r.EventType
    structMap["event_data"] = r.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceEvent.
// It customizes the JSON unmarshaling process for RefundInvoiceEvent objects.
func (r *RefundInvoiceEvent) UnmarshalJSON(input []byte) error {
    var temp tempRefundInvoiceEvent
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
    r.AdditionalProperties = additionalProperties
    
    r.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    r.Timestamp = TimestampVal
    r.Invoice = *temp.Invoice
    r.EventType = *temp.EventType
    r.EventData = *temp.EventData
    return nil
}

// tempRefundInvoiceEvent is a temporary struct used for validating the fields of RefundInvoiceEvent.
type tempRefundInvoiceEvent  struct {
    Id        *int64                  `json:"id"`
    Timestamp *string                 `json:"timestamp"`
    Invoice   *Invoice                `json:"invoice"`
    EventType *InvoiceEventType       `json:"event_type"`
    EventData *RefundInvoiceEventData `json:"event_data"`
}

func (r *tempRefundInvoiceEvent) validate() error {
    var errs []string
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Refund Invoice Event`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Refund Invoice Event`")
    }
    if r.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Refund Invoice Event`")
    }
    if r.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Refund Invoice Event`")
    }
    if r.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Refund Invoice Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
