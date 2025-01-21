/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "fmt"
)

// CouponRestriction represents a CouponRestriction struct.
type CouponRestriction struct {
    Id                   *int                   `json:"id,omitempty"`
    ItemType             *RestrictionType       `json:"item_type,omitempty"`
    ItemId               *int                   `json:"item_id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Handle               Optional[string]       `json:"handle"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CouponRestriction,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CouponRestriction) String() string {
    return fmt.Sprintf(
    	"CouponRestriction[Id=%v, ItemType=%v, ItemId=%v, Name=%v, Handle=%v, AdditionalProperties=%v]",
    	c.Id, c.ItemType, c.ItemId, c.Name, c.Handle, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CouponRestriction.
// It customizes the JSON marshaling process for CouponRestriction objects.
func (c CouponRestriction) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "id", "item_type", "item_id", "name", "handle"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CouponRestriction object to a map representation for JSON marshaling.
func (c CouponRestriction) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Id != nil {
        structMap["id"] = c.Id
    }
    if c.ItemType != nil {
        structMap["item_type"] = c.ItemType
    }
    if c.ItemId != nil {
        structMap["item_id"] = c.ItemId
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Handle.IsValueSet() {
        if c.Handle.Value() != nil {
            structMap["handle"] = c.Handle.Value()
        } else {
            structMap["handle"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponRestriction.
// It customizes the JSON unmarshaling process for CouponRestriction objects.
func (c *CouponRestriction) UnmarshalJSON(input []byte) error {
    var temp tempCouponRestriction
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "item_type", "item_id", "name", "handle")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Id = temp.Id
    c.ItemType = temp.ItemType
    c.ItemId = temp.ItemId
    c.Name = temp.Name
    c.Handle = temp.Handle
    return nil
}

// tempCouponRestriction is a temporary struct used for validating the fields of CouponRestriction.
type tempCouponRestriction  struct {
    Id       *int             `json:"id,omitempty"`
    ItemType *RestrictionType `json:"item_type,omitempty"`
    ItemId   *int             `json:"item_id,omitempty"`
    Name     *string          `json:"name,omitempty"`
    Handle   Optional[string] `json:"handle"`
}
