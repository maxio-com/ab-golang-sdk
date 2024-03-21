package models

import (
	"encoding/json"
	"log"
	"time"
)

// InvoiceEvent represents a InvoiceEvent struct.
type InvoiceEvent struct {
	Id *int `json:"id,omitempty"`
	// Invoice Event Type
	EventType *InvoiceEventType `json:"event_type,omitempty"`
	// The event data is the data that, when combined with the command, results in the output invoice found in the invoice field.
	EventData *InvoiceEventEventData `json:"event_data,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
	Invoice   *Invoice               `json:"invoice,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for InvoiceEvent.
// It customizes the JSON marshaling process for InvoiceEvent objects.
func (i *InvoiceEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceEvent object to a map representation for JSON marshaling.
func (i *InvoiceEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	if i.Id != nil {
		structMap["id"] = i.Id
	}
	if i.EventType != nil {
		structMap["event_type"] = i.EventType
	}
	if i.EventData != nil {
		structMap["event_data"] = i.EventData.toMap()
	}
	if i.Timestamp != nil {
		structMap["timestamp"] = i.Timestamp.Format(time.RFC3339)
	}
	if i.Invoice != nil {
		structMap["invoice"] = i.Invoice.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEvent.
// It customizes the JSON unmarshaling process for InvoiceEvent objects.
func (i *InvoiceEvent) UnmarshalJSON(input []byte) error {
	var temp invoiceEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	i.Id = temp.Id
	i.EventType = temp.EventType
	i.EventData = temp.EventData
	if temp.Timestamp != nil {
		TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
		if err != nil {
			log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
		}
		i.Timestamp = &TimestampVal
	}
	i.Invoice = temp.Invoice
	return nil
}

// TODO
type invoiceEvent struct {
	Id        *int                   `json:"id,omitempty"`
	EventType *InvoiceEventType      `json:"event_type,omitempty"`
	EventData *InvoiceEventEventData `json:"event_data,omitempty"`
	Timestamp *string                `json:"timestamp,omitempty"`
	Invoice   *Invoice               `json:"invoice,omitempty"`
}
