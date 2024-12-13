/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
)

// UpdateMetafieldsRequest represents a UpdateMetafieldsRequest struct.
type UpdateMetafieldsRequest struct {
    Metafields           *UpdateMetafieldsRequestMetafields `json:"metafields,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateMetafieldsRequest.
// It customizes the JSON marshaling process for UpdateMetafieldsRequest objects.
func (u UpdateMetafieldsRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "metafields"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateMetafieldsRequest object to a map representation for JSON marshaling.
func (u UpdateMetafieldsRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Metafields != nil {
        structMap["metafields"] = u.Metafields.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateMetafieldsRequest.
// It customizes the JSON unmarshaling process for UpdateMetafieldsRequest objects.
func (u *UpdateMetafieldsRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateMetafieldsRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "metafields")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Metafields = temp.Metafields
    return nil
}

// tempUpdateMetafieldsRequest is a temporary struct used for validating the fields of UpdateMetafieldsRequest.
type tempUpdateMetafieldsRequest  struct {
    Metafields *UpdateMetafieldsRequestMetafields `json:"metafields,omitempty"`
}
