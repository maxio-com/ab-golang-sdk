package models

import (
    "encoding/json"
)

// SaleRepItemMrr represents a SaleRepItemMrr struct.
type SaleRepItemMrr struct {
    Mrr                  *string        `json:"mrr,omitempty"`
    Usage                *string        `json:"usage,omitempty"`
    Recurring            *string        `json:"recurring,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SaleRepItemMrr.
// It customizes the JSON marshaling process for SaleRepItemMrr objects.
func (s SaleRepItemMrr) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SaleRepItemMrr object to a map representation for JSON marshaling.
func (s SaleRepItemMrr) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Mrr != nil {
        structMap["mrr"] = s.Mrr
    }
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    if s.Recurring != nil {
        structMap["recurring"] = s.Recurring
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SaleRepItemMrr.
// It customizes the JSON unmarshaling process for SaleRepItemMrr objects.
func (s *SaleRepItemMrr) UnmarshalJSON(input []byte) error {
    var temp saleRepItemMrr
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mrr", "usage", "recurring")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Mrr = temp.Mrr
    s.Usage = temp.Usage
    s.Recurring = temp.Recurring
    return nil
}

// TODO
type saleRepItemMrr  struct {
    Mrr       *string `json:"mrr,omitempty"`
    Usage     *string `json:"usage,omitempty"`
    Recurring *string `json:"recurring,omitempty"`
}
