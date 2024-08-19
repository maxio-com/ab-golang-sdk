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

// Event represents a Event struct.
type Event struct {
    Id                   int64                   `json:"id"`
    Key                  string                  `json:"key"`
    Message              string                  `json:"message"`
    SubscriptionId       *int                    `json:"subscription_id"`
    CustomerId           *int                    `json:"customer_id"`
    CreatedAt            time.Time               `json:"created_at"`
    EventSpecificData    *EventEventSpecificData `json:"event_specific_data"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Event.
// It customizes the JSON marshaling process for Event objects.
func (e Event) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the Event object to a map representation for JSON marshaling.
func (e Event) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["id"] = e.Id
    structMap["key"] = e.Key
    structMap["message"] = e.Message
    if e.SubscriptionId != nil {
        structMap["subscription_id"] = e.SubscriptionId
    } else {
        structMap["subscription_id"] = nil
    }
    if e.CustomerId != nil {
        structMap["customer_id"] = e.CustomerId
    } else {
        structMap["customer_id"] = nil
    }
    structMap["created_at"] = e.CreatedAt.Format(time.RFC3339)
    if e.EventSpecificData != nil {
        structMap["event_specific_data"] = e.EventSpecificData.toMap()
    } else {
        structMap["event_specific_data"] = nil
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Event.
// It customizes the JSON unmarshaling process for Event objects.
func (e *Event) UnmarshalJSON(input []byte) error {
    var temp tempEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "key", "message", "subscription_id", "customer_id", "created_at", "event_specific_data")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Id = *temp.Id
    e.Key = *temp.Key
    e.Message = *temp.Message
    e.SubscriptionId = temp.SubscriptionId
    e.CustomerId = temp.CustomerId
    CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
    }
    e.CreatedAt = CreatedAtVal
    e.EventSpecificData = temp.EventSpecificData
    return nil
}

// tempEvent is a temporary struct used for validating the fields of Event.
type tempEvent  struct {
    Id                *int64                  `json:"id"`
    Key               *string                 `json:"key"`
    Message           *string                 `json:"message"`
    SubscriptionId    *int                    `json:"subscription_id"`
    CustomerId        *int                    `json:"customer_id"`
    CreatedAt         *string                 `json:"created_at"`
    EventSpecificData *EventEventSpecificData `json:"event_specific_data"`
}

func (e *tempEvent) validate() error {
    var errs []string
    if e.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Event`")
    }
    if e.Key == nil {
        errs = append(errs, "required field `key` is missing for type `Event`")
    }
    if e.Message == nil {
        errs = append(errs, "required field `message` is missing for type `Event`")
    }
    if e.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
