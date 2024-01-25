package models

import (
	"encoding/json"
)

// CouponSubcodesResponse represents a CouponSubcodesResponse struct.
type CouponSubcodesResponse struct {
	CreatedCodes   []string `json:"created_codes,omitempty"`
	DuplicateCodes []string `json:"duplicate_codes,omitempty"`
	InvalidCodes   []string `json:"invalid_codes,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CouponSubcodesResponse.
// It customizes the JSON marshaling process for CouponSubcodesResponse objects.
func (c *CouponSubcodesResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CouponSubcodesResponse object to a map representation for JSON marshaling.
func (c *CouponSubcodesResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.CreatedCodes != nil {
		structMap["created_codes"] = c.CreatedCodes
	}
	if c.DuplicateCodes != nil {
		structMap["duplicate_codes"] = c.DuplicateCodes
	}
	if c.InvalidCodes != nil {
		structMap["invalid_codes"] = c.InvalidCodes
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CouponSubcodesResponse.
// It customizes the JSON unmarshaling process for CouponSubcodesResponse objects.
func (c *CouponSubcodesResponse) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CreatedCodes   []string `json:"created_codes,omitempty"`
		DuplicateCodes []string `json:"duplicate_codes,omitempty"`
		InvalidCodes   []string `json:"invalid_codes,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.CreatedCodes = temp.CreatedCodes
	c.DuplicateCodes = temp.DuplicateCodes
	c.InvalidCodes = temp.InvalidCodes
	return nil
}
