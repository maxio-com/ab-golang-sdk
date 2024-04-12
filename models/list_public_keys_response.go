package models

import (
    "encoding/json"
)

// ListPublicKeysResponse represents a ListPublicKeysResponse struct.
type ListPublicKeysResponse struct {
    ChargifyJsKeys       []PublicKey         `json:"chargify_js_keys,omitempty"`
    Meta                 *ListPublicKeysMeta `json:"meta,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ListPublicKeysResponse.
// It customizes the JSON marshaling process for ListPublicKeysResponse objects.
func (l ListPublicKeysResponse) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(l.toMap())
}

// toMap converts the ListPublicKeysResponse object to a map representation for JSON marshaling.
func (l ListPublicKeysResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, l.AdditionalProperties)
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
    var temp listPublicKeysResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "chargify_js_keys", "meta")
    if err != nil {
    	return err
    }
    
    l.AdditionalProperties = additionalProperties
    l.ChargifyJsKeys = temp.ChargifyJsKeys
    l.Meta = temp.Meta
    return nil
}

// TODO
type listPublicKeysResponse  struct {
    ChargifyJsKeys []PublicKey         `json:"chargify_js_keys,omitempty"`
    Meta           *ListPublicKeysMeta `json:"meta,omitempty"`
}
