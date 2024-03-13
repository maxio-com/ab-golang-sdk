package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ComponentPricePointResponse represents a ComponentPricePointResponse struct.
type ComponentPricePointResponse struct {
	PricePoint ComponentPricePoint `json:"price_point"`
}

// MarshalJSON implements the json.Marshaler interface for ComponentPricePointResponse.
// It customizes the JSON marshaling process for ComponentPricePointResponse objects.
func (c *ComponentPricePointResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ComponentPricePointResponse object to a map representation for JSON marshaling.
func (c *ComponentPricePointResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["price_point"] = c.PricePoint.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentPricePointResponse.
// It customizes the JSON unmarshaling process for ComponentPricePointResponse objects.
func (c *ComponentPricePointResponse) UnmarshalJSON(input []byte) error {
	var temp componentPricePointResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.PricePoint = *temp.PricePoint
	return nil
}

// TODO
type componentPricePointResponse struct {
	PricePoint *ComponentPricePoint `json:"price_point"`
}

func (c *componentPricePointResponse) validate() error {
	var errs []string
	if c.PricePoint == nil {
		errs = append(errs, "required field `price_point` is missing for type `Component Price Point Response`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
