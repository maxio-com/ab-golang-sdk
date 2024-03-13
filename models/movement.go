package models

import (
	"encoding/json"
	"log"
	"time"
)

// Movement represents a Movement struct.
type Movement struct {
	Timestamp       *time.Time         `json:"timestamp,omitempty"`
	AmountInCents   *int64             `json:"amount_in_cents,omitempty"`
	AmountFormatted *string            `json:"amount_formatted,omitempty"`
	Description     *string            `json:"description,omitempty"`
	Category        *string            `json:"category,omitempty"`
	Breakouts       *Breakouts         `json:"breakouts,omitempty"`
	LineItems       []MovementLineItem `json:"line_items,omitempty"`
	SubscriptionId  *int               `json:"subscription_id,omitempty"`
	SubscriberName  *string            `json:"subscriber_name,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Movement.
// It customizes the JSON marshaling process for Movement objects.
func (m *Movement) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the Movement object to a map representation for JSON marshaling.
func (m *Movement) toMap() map[string]any {
	structMap := make(map[string]any)
	if m.Timestamp != nil {
		structMap["timestamp"] = m.Timestamp.Format(time.RFC3339)
	}
	if m.AmountInCents != nil {
		structMap["amount_in_cents"] = m.AmountInCents
	}
	if m.AmountFormatted != nil {
		structMap["amount_formatted"] = m.AmountFormatted
	}
	if m.Description != nil {
		structMap["description"] = m.Description
	}
	if m.Category != nil {
		structMap["category"] = m.Category
	}
	if m.Breakouts != nil {
		structMap["breakouts"] = m.Breakouts.toMap()
	}
	if m.LineItems != nil {
		structMap["line_items"] = m.LineItems
	}
	if m.SubscriptionId != nil {
		structMap["subscription_id"] = m.SubscriptionId
	}
	if m.SubscriberName != nil {
		structMap["subscriber_name"] = m.SubscriberName
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Movement.
// It customizes the JSON unmarshaling process for Movement objects.
func (m *Movement) UnmarshalJSON(input []byte) error {
	var temp movement
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	if temp.Timestamp != nil {
		TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
		if err != nil {
			log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
		}
		m.Timestamp = &TimestampVal
	}
	m.AmountInCents = temp.AmountInCents
	m.AmountFormatted = temp.AmountFormatted
	m.Description = temp.Description
	m.Category = temp.Category
	m.Breakouts = temp.Breakouts
	m.LineItems = temp.LineItems
	m.SubscriptionId = temp.SubscriptionId
	m.SubscriberName = temp.SubscriberName
	return nil
}

// TODO
type movement struct {
	Timestamp       *string            `json:"timestamp,omitempty"`
	AmountInCents   *int64             `json:"amount_in_cents,omitempty"`
	AmountFormatted *string            `json:"amount_formatted,omitempty"`
	Description     *string            `json:"description,omitempty"`
	Category        *string            `json:"category,omitempty"`
	Breakouts       *Breakouts         `json:"breakouts,omitempty"`
	LineItems       []MovementLineItem `json:"line_items,omitempty"`
	SubscriptionId  *int               `json:"subscription_id,omitempty"`
	SubscriberName  *string            `json:"subscriber_name,omitempty"`
}
