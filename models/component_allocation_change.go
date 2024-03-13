package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ComponentAllocationChange represents a ComponentAllocationChange struct.
type ComponentAllocationChange struct {
	PreviousAllocation int                                         `json:"previous_allocation"`
	NewAllocation      int                                         `json:"new_allocation"`
	ComponentId        int                                         `json:"component_id"`
	ComponentHandle    string                                      `json:"component_handle"`
	Memo               string                                      `json:"memo"`
	AllocationId       int                                         `json:"allocation_id"`
	AllocatedQuantity  *ComponentAllocationChangeAllocatedQuantity `json:"allocated_quantity,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentAllocationChange.
// It customizes the JSON marshaling process for ComponentAllocationChange objects.
func (c *ComponentAllocationChange) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentAllocationChange object to a map representation for JSON marshaling.
func (c *ComponentAllocationChange) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["previous_allocation"] = c.PreviousAllocation
	structMap["new_allocation"] = c.NewAllocation
	structMap["component_id"] = c.ComponentId
	structMap["component_handle"] = c.ComponentHandle
	structMap["memo"] = c.Memo
	structMap["allocation_id"] = c.AllocationId
	if c.AllocatedQuantity != nil {
		structMap["allocated_quantity"] = c.AllocatedQuantity.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentAllocationChange.
// It customizes the JSON unmarshaling process for ComponentAllocationChange objects.
func (c *ComponentAllocationChange) UnmarshalJSON(input []byte) error {
	var temp componentAllocationChange
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.PreviousAllocation = *temp.PreviousAllocation
	c.NewAllocation = *temp.NewAllocation
	c.ComponentId = *temp.ComponentId
	c.ComponentHandle = *temp.ComponentHandle
	c.Memo = *temp.Memo
	c.AllocationId = *temp.AllocationId
	c.AllocatedQuantity = temp.AllocatedQuantity
	return nil
}

// TODO
type componentAllocationChange struct {
	PreviousAllocation *int                                        `json:"previous_allocation"`
	NewAllocation      *int                                        `json:"new_allocation"`
	ComponentId        *int                                        `json:"component_id"`
	ComponentHandle    *string                                     `json:"component_handle"`
	Memo               *string                                     `json:"memo"`
	AllocationId       *int                                        `json:"allocation_id"`
	AllocatedQuantity  *ComponentAllocationChangeAllocatedQuantity `json:"allocated_quantity,omitempty"`
}

func (c *componentAllocationChange) validate() error {
	var errs []string
	if c.PreviousAllocation == nil {
		errs = append(errs, "required field `previous_allocation` is missing for type `Component Allocation Change`")
	}
	if c.NewAllocation == nil {
		errs = append(errs, "required field `new_allocation` is missing for type `Component Allocation Change`")
	}
	if c.ComponentId == nil {
		errs = append(errs, "required field `component_id` is missing for type `Component Allocation Change`")
	}
	if c.ComponentHandle == nil {
		errs = append(errs, "required field `component_handle` is missing for type `Component Allocation Change`")
	}
	if c.Memo == nil {
		errs = append(errs, "required field `memo` is missing for type `Component Allocation Change`")
	}
	if c.AllocationId == nil {
		errs = append(errs, "required field `allocation_id` is missing for type `Component Allocation Change`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
