package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CustomerCustomFieldsChange represents a CustomerCustomFieldsChange struct.
type CustomerCustomFieldsChange struct {
	Before []InvoiceCustomField `json:"before"`
	After  []InvoiceCustomField `json:"after"`
}

// MarshalJSON implements the json.Marshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON marshaling process for CustomerCustomFieldsChange objects.
func (c *CustomerCustomFieldsChange) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CustomerCustomFieldsChange object to a map representation for JSON marshaling.
func (c *CustomerCustomFieldsChange) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["before"] = c.Before
	structMap["after"] = c.After
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomerCustomFieldsChange.
// It customizes the JSON unmarshaling process for CustomerCustomFieldsChange objects.
func (c *CustomerCustomFieldsChange) UnmarshalJSON(input []byte) error {
	var temp customerCustomFieldsChange
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Before = *temp.Before
	c.After = *temp.After
	return nil
}

// TODO
type customerCustomFieldsChange struct {
	Before *[]InvoiceCustomField `json:"before"`
	After  *[]InvoiceCustomField `json:"after"`
}

func (c *customerCustomFieldsChange) validate() error {
	var errs []string
	if c.Before == nil {
		errs = append(errs, "required field `before` is missing for type `Customer Custom Fields Change`")
	}
	if c.After == nil {
		errs = append(errs, "required field `after` is missing for type `Customer Custom Fields Change`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
