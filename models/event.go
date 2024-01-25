package models

import (
	"encoding/json"
)

// Event represents a Event struct.
type Event struct {
	Id                int         `json:"id"`
	Key               string      `json:"key"`
	Message           string      `json:"message"`
	SubscriptionId    *int        `json:"subscription_id"`
	CustomerId        int         `json:"customer_id"`
	CreatedAt         string      `json:"created_at"`
	EventSpecificData interface{} `json:"event_specific_data"`
}

// MarshalJSON implements the json.Marshaler interface for Event.
// It customizes the JSON marshaling process for Event objects.
func (e *Event) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the Event object to a map representation for JSON marshaling.
func (e *Event) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["id"] = e.Id
	structMap["key"] = e.Key
	structMap["message"] = e.Message
	structMap["subscription_id"] = e.SubscriptionId
	structMap["customer_id"] = e.CustomerId
	structMap["created_at"] = e.CreatedAt
	structMap["event_specific_data"] = e.EventSpecificData
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Event.
// It customizes the JSON unmarshaling process for Event objects.
func (e *Event) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                int         `json:"id"`
		Key               string      `json:"key"`
		Message           string      `json:"message"`
		SubscriptionId    *int        `json:"subscription_id"`
		CustomerId        int         `json:"customer_id"`
		CreatedAt         string      `json:"created_at"`
		EventSpecificData interface{} `json:"event_specific_data"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	e.Id = temp.Id
	e.Key = temp.Key
	e.Message = temp.Message
	e.SubscriptionId = temp.SubscriptionId
	e.CustomerId = temp.CustomerId
	e.CreatedAt = temp.CreatedAt
	e.EventSpecificData = temp.EventSpecificData
	return nil
}
