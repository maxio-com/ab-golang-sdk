package models

import (
	"encoding/json"
)

// CouponRestriction represents a CouponRestriction struct.
type CouponRestriction struct {
	Id       *int             `json:"id,omitempty"`
	ItemType *RestrictionType `json:"item_type,omitempty"`
	ItemId   *int             `json:"item_id,omitempty"`
	Name     *string          `json:"name,omitempty"`
	Handle   Optional[string] `json:"handle"`
}

// MarshalJSON implements the json.Marshaler interface for CouponRestriction.
// It customizes the JSON marshaling process for CouponRestriction objects.
func (c *CouponRestriction) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponRestriction object to a map representation for JSON marshaling.
func (c *CouponRestriction) toMap() map[string]any {
	structMap := make(map[string]any)
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
		structMap["handle"] = c.Handle.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponRestriction.
// It customizes the JSON unmarshaling process for CouponRestriction objects.
func (c *CouponRestriction) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id       *int             `json:"id,omitempty"`
		ItemType *RestrictionType `json:"item_type,omitempty"`
		ItemId   *int             `json:"item_id,omitempty"`
		Name     *string          `json:"name,omitempty"`
		Handle   Optional[string] `json:"handle"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Id = temp.Id
	c.ItemType = temp.ItemType
	c.ItemId = temp.ItemId
	c.Name = temp.Name
	c.Handle = temp.Handle
	return nil
}
