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

// ChangeChargebackStatusEvent represents a ChangeChargebackStatusEvent struct.
type ChangeChargebackStatusEvent struct {
    Id                   int64                           `json:"id"`
    Timestamp            time.Time                       `json:"timestamp"`
    Invoice              Invoice                         `json:"invoice"`
    EventType            InvoiceEventType                `json:"event_type"`
    // Example schema for an `change_chargeback_status` event
    EventData            ChangeChargebackStatusEventData `json:"event_data"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ChangeChargebackStatusEvent.
// It customizes the JSON marshaling process for ChangeChargebackStatusEvent objects.
func (c ChangeChargebackStatusEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ChangeChargebackStatusEvent object to a map representation for JSON marshaling.
func (c ChangeChargebackStatusEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["id"] = c.Id
    structMap["timestamp"] = c.Timestamp.Format(time.RFC3339)
    structMap["invoice"] = c.Invoice.toMap()
    structMap["event_type"] = c.EventType
    structMap["event_data"] = c.EventData.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChangeChargebackStatusEvent.
// It customizes the JSON unmarshaling process for ChangeChargebackStatusEvent objects.
func (c *ChangeChargebackStatusEvent) UnmarshalJSON(input []byte) error {
    var temp tempChangeChargebackStatusEvent
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
    
    c.AdditionalProperties = additionalProperties
    c.Id = *temp.Id
    TimestampVal, err := time.Parse(time.RFC3339, *temp.Timestamp)
    if err != nil {
        log.Fatalf("Cannot Parse timestamp as % s format.", time.RFC3339)
    }
    c.Timestamp = TimestampVal
    c.Invoice = *temp.Invoice
    c.EventType = *temp.EventType
    c.EventData = *temp.EventData
    return nil
}

// tempChangeChargebackStatusEvent is a temporary struct used for validating the fields of ChangeChargebackStatusEvent.
type tempChangeChargebackStatusEvent  struct {
    Id        *int64                           `json:"id"`
    Timestamp *string                          `json:"timestamp"`
    Invoice   *Invoice                         `json:"invoice"`
    EventType *InvoiceEventType                `json:"event_type"`
    EventData *ChangeChargebackStatusEventData `json:"event_data"`
}

func (c *tempChangeChargebackStatusEvent) validate() error {
    var errs []string
    if c.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Change Chargeback Status Event`")
    }
    if c.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Change Chargeback Status Event`")
    }
    if c.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Change Chargeback Status Event`")
    }
    if c.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Change Chargeback Status Event`")
    }
    if c.EventData == nil {
        errs = append(errs, "required field `event_data` is missing for type `Change Chargeback Status Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
