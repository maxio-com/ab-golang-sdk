// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
	"encoding/json"
	"fmt"
)

// ListPublicKeysResponse represents a ListPublicKeysResponse struct.
type ListPublicKeysResponse struct {
	ChargifyJsKeys       []PublicKey            `json:"chargify_js_keys,omitempty"`
	Meta                 *ListPublicKeysMeta    `json:"meta,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ListPublicKeysResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l ListPublicKeysResponse) String() string {
	return fmt.Sprintf(
		"ListPublicKeysResponse[ChargifyJsKeys=%v, Meta=%v, AdditionalProperties=%v]",
		l.ChargifyJsKeys, l.Meta, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ListPublicKeysResponse.
// It customizes the JSON marshaling process for ListPublicKeysResponse objects.
func (l ListPublicKeysResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(l.AdditionalProperties,
		"chargify_js_keys", "meta"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(l.toMap())
}

// toMap converts the ListPublicKeysResponse object to a map representation for JSON marshaling.
func (l ListPublicKeysResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, l.AdditionalProperties)
	if l.ChargifyJsKeys != nil {
		structMap["chargify_js_keys"] = l.ChargifyJsKeys
	}
	if l.Meta != nil {
		structMap["meta"] = l.Meta.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPublicKeysResponse.
// It customizes the JSON unmarshaling process for ListPublicKeysResponse objects.
func (l *ListPublicKeysResponse) UnmarshalJSON(input []byte) error {
	var temp tempListPublicKeysResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "chargify_js_keys", "meta")
	if err != nil {
		return err
	}
	l.AdditionalProperties = additionalProperties

	l.ChargifyJsKeys = temp.ChargifyJsKeys
	l.Meta = temp.Meta
	return nil
}

// tempListPublicKeysResponse is a temporary struct used for validating the fields of ListPublicKeysResponse.
type tempListPublicKeysResponse struct {
	ChargifyJsKeys []PublicKey         `json:"chargify_js_keys,omitempty"`
	Meta           *ListPublicKeysMeta `json:"meta,omitempty"`
}
