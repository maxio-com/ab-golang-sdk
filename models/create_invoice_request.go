package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// CreateInvoiceRequest represents a CreateInvoiceRequest struct.
type CreateInvoiceRequest struct {
	Invoice CreateInvoice `json:"invoice"`
}

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceRequest.
// It customizes the JSON marshaling process for CreateInvoiceRequest objects.
func (c *CreateInvoiceRequest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CreateInvoiceRequest object to a map representation for JSON marshaling.
func (c *CreateInvoiceRequest) toMap() map[string]any {
	structMap := make(map[string]any)
	structMap["invoice"] = c.Invoice.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceRequest.
// It customizes the JSON unmarshaling process for CreateInvoiceRequest objects.
func (c *CreateInvoiceRequest) UnmarshalJSON(input []byte) error {
	var temp createInvoiceRequest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}

	c.Invoice = *temp.Invoice
	return nil
}

// TODO
type createInvoiceRequest struct {
	Invoice *CreateInvoice `json:"invoice"`
}

func (c *createInvoiceRequest) validate() error {
	var errs []string
	if c.Invoice == nil {
		errs = append(errs, "required field `invoice` is missing for type `Create Invoice Request`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
