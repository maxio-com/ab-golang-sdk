package models

import (
	"encoding/json"
)

// CreateInvoiceCoupon represents a CreateInvoiceCoupon struct.
type CreateInvoiceCoupon struct {
	Code                *string                  `json:"code,omitempty"`
	Percentage          *interface{}             `json:"percentage,omitempty"`
	Amount              *interface{}             `json:"amount,omitempty"`
	Description         *string                  `json:"description,omitempty"`
	ProductFamilyId     *interface{}             `json:"product_family_id,omitempty"`
	CompoundingStrategy *CompoundingStrategyEnum `json:"compounding_strategy,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceCoupon.
// It customizes the JSON marshaling process for CreateInvoiceCoupon objects.
func (c *CreateInvoiceCoupon) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceCoupon object to a map representation for JSON marshaling.
func (c *CreateInvoiceCoupon) toMap() map[string]any {
	structMap := make(map[string]any)
	if c.Code != nil {
		structMap["code"] = c.Code
	}
	if c.Percentage != nil {
		structMap["percentage"] = c.Percentage
	}
	if c.Amount != nil {
		structMap["amount"] = c.Amount
	}
	if c.Description != nil {
		structMap["description"] = c.Description
	}
	if c.ProductFamilyId != nil {
		structMap["product_family_id"] = c.ProductFamilyId
	}
	if c.CompoundingStrategy != nil {
		structMap["compounding_strategy"] = c.CompoundingStrategy
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceCoupon.
// It customizes the JSON unmarshaling process for CreateInvoiceCoupon objects.
func (c *CreateInvoiceCoupon) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Code                *string                  `json:"code,omitempty"`
		Percentage          *interface{}             `json:"percentage,omitempty"`
		Amount              *interface{}             `json:"amount,omitempty"`
		Description         *string                  `json:"description,omitempty"`
		ProductFamilyId     *interface{}             `json:"product_family_id,omitempty"`
		CompoundingStrategy *CompoundingStrategyEnum `json:"compounding_strategy,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	c.Code = temp.Code
	c.Percentage = temp.Percentage
	c.Amount = temp.Amount
	c.Description = temp.Description
	c.ProductFamilyId = temp.ProductFamilyId
	c.CompoundingStrategy = temp.CompoundingStrategy
	return nil
}
