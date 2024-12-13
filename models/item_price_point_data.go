/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// ItemPricePointData represents a ItemPricePointData struct.
type ItemPricePointData struct {
    Id                   *int                   `json:"id,omitempty"`
    Handle               *string                `json:"handle,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ItemPricePointData.
// It customizes the JSON marshaling process for ItemPricePointData objects.
func (i ItemPricePointData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "id", "handle", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the ItemPricePointData object to a map representation for JSON marshaling.
func (i ItemPricePointData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Id != nil {
        structMap["id"] = i.Id
    }
    if i.Handle != nil {
        structMap["handle"] = i.Handle
    }
    if i.Name != nil {
        structMap["name"] = i.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ItemPricePointData.
// It customizes the JSON unmarshaling process for ItemPricePointData objects.
func (i *ItemPricePointData) UnmarshalJSON(input []byte) error {
    var temp tempItemPricePointData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "handle", "name")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Id = temp.Id
    i.Handle = temp.Handle
    i.Name = temp.Name
    return nil
}

// tempItemPricePointData is a temporary struct used for validating the fields of ItemPricePointData.
type tempItemPricePointData  struct {
    Id     *int    `json:"id,omitempty"`
    Handle *string `json:"handle,omitempty"`
    Name   *string `json:"name,omitempty"`
}
