package models

import (
	"encoding/json"
)

// InvoiceEvent represents a InvoiceEvent struct.
type InvoiceEvent struct {
	Id *int `json:"id,omitempty"`
	// Invoice Event Type
	EventType *InvoiceEventTypeEnum `json:"event_type,omitempty"`
	// The event data is the data that, when combined with the command, results in the output invoice found in the invoice field.
	EventData *InvoiceEvent1 `json:"event_data,omitempty"`
	Timestamp *string        `json:"timestamp,omitempty"`
	Invoice   *Invoice       `json:"invoice,omitempty"`
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
		structMap["event_data"] = i.EventData
	}
	if i.Timestamp != nil {
		structMap["timestamp"] = i.Timestamp
	}
	if i.Invoice != nil {
		structMap["invoice"] = i.Invoice
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEvent.
// It customizes the JSON unmarshaling process for InvoiceEvent objects.
func (i *InvoiceEvent) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id        *int                  `json:"id,omitempty"`
		EventType *InvoiceEventTypeEnum `json:"event_type,omitempty"`
		EventData *InvoiceEvent1        `json:"event_data,omitempty"`
		Timestamp *string               `json:"timestamp,omitempty"`
		Invoice   *Invoice              `json:"invoice,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	i.Id = temp.Id
	i.EventType = temp.EventType
	i.EventData = temp.EventData
	i.Timestamp = temp.Timestamp
	i.Invoice = temp.Invoice
	return nil
}
