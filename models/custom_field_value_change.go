package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CustomFieldValueChange represents a CustomFieldValueChange struct.
type CustomFieldValueChange struct {
    EventType            string         `json:"event_type"`
    MetafieldName        string         `json:"metafield_name"`
    MetafieldId          int            `json:"metafield_id"`
    OldValue             *string        `json:"old_value"`
    NewValue             *string        `json:"new_value"`
    ResourceType         string         `json:"resource_type"`
    ResourceId           int            `json:"resource_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CustomFieldValueChange.
// It customizes the JSON marshaling process for CustomFieldValueChange objects.
func (c CustomFieldValueChange) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CustomFieldValueChange object to a map representation for JSON marshaling.
func (c CustomFieldValueChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["event_type"] = c.EventType
    structMap["metafield_name"] = c.MetafieldName
    structMap["metafield_id"] = c.MetafieldId
    if c.OldValue != nil {
        structMap["old_value"] = c.OldValue
    } else {
        structMap["old_value"] = nil
    }
    if c.NewValue != nil {
        structMap["new_value"] = c.NewValue
    } else {
        structMap["new_value"] = nil
    }
    structMap["resource_type"] = c.ResourceType
    structMap["resource_id"] = c.ResourceId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomFieldValueChange.
// It customizes the JSON unmarshaling process for CustomFieldValueChange objects.
func (c *CustomFieldValueChange) UnmarshalJSON(input []byte) error {
    var temp customFieldValueChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "event_type", "metafield_name", "metafield_id", "old_value", "new_value", "resource_type", "resource_id")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.EventType = *temp.EventType
    c.MetafieldName = *temp.MetafieldName
    c.MetafieldId = *temp.MetafieldId
    c.OldValue = temp.OldValue
    c.NewValue = temp.NewValue
    c.ResourceType = *temp.ResourceType
    c.ResourceId = *temp.ResourceId
    return nil
}

// customFieldValueChange is a temporary struct used for validating the fields of CustomFieldValueChange.
type customFieldValueChange  struct {
    EventType     *string `json:"event_type"`
    MetafieldName *string `json:"metafield_name"`
    MetafieldId   *int    `json:"metafield_id"`
    OldValue      *string `json:"old_value"`
    NewValue      *string `json:"new_value"`
    ResourceType  *string `json:"resource_type"`
    ResourceId    *int    `json:"resource_id"`
}

func (c *customFieldValueChange) validate() error {
    var errs []string
    if c.EventType == nil {
        errs = append(errs, "required field `event_type` is missing for type `Custom Field Value Change`")
    }
    if c.MetafieldName == nil {
        errs = append(errs, "required field `metafield_name` is missing for type `Custom Field Value Change`")
    }
    if c.MetafieldId == nil {
        errs = append(errs, "required field `metafield_id` is missing for type `Custom Field Value Change`")
    }
    if c.ResourceType == nil {
        errs = append(errs, "required field `resource_type` is missing for type `Custom Field Value Change`")
    }
    if c.ResourceId == nil {
        errs = append(errs, "required field `resource_id` is missing for type `Custom Field Value Change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
