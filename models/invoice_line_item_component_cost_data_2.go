/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// InvoiceLineItemComponentCostData2 represents a InvoiceLineItemComponentCostData2 struct.
// This is a container for one-of cases.
type InvoiceLineItemComponentCostData2 struct {
	value                              any
	isInvoiceLineItemComponentCostData bool
}

// String converts the InvoiceLineItemComponentCostData2 object to a string representation.
func (i InvoiceLineItemComponentCostData2) String() string {
	if bytes, err := json.Marshal(i.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for InvoiceLineItemComponentCostData2.
// It customizes the JSON marshaling process for InvoiceLineItemComponentCostData2 objects.
func (i *InvoiceLineItemComponentCostData2) MarshalJSON() (
	[]byte,
	error) {
	if i.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.InvoiceLineItemComponentCostData2Container.From*` functions to initialize the InvoiceLineItemComponentCostData2 object.")
	}
	return json.Marshal(i.toMap())
}

// toMap converts the InvoiceLineItemComponentCostData2 object to a map representation for JSON marshaling.
func (i *InvoiceLineItemComponentCostData2) toMap() any {
	switch obj := i.value.(type) {
	case *InvoiceLineItemComponentCostData:
		return obj.toMap()
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceLineItemComponentCostData2.
// It customizes the JSON unmarshaling process for InvoiceLineItemComponentCostData2 objects.
func (i *InvoiceLineItemComponentCostData2) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&InvoiceLineItemComponentCostData{}, false, &i.isInvoiceLineItemComponentCostData),
	)

	i.value = result
	return err
}

func (i *InvoiceLineItemComponentCostData2) AsInvoiceLineItemComponentCostData() (
	*InvoiceLineItemComponentCostData,
	bool) {
	if !i.isInvoiceLineItemComponentCostData {
		return nil, false
	}
	return i.value.(*InvoiceLineItemComponentCostData), true
}

// internalInvoiceLineItemComponentCostData2 represents a invoiceLineItemComponentCostData2 struct.
// This is a container for one-of cases.
type internalInvoiceLineItemComponentCostData2 struct{}

var InvoiceLineItemComponentCostData2Container internalInvoiceLineItemComponentCostData2

func (i *internalInvoiceLineItemComponentCostData2) FromInvoiceLineItemComponentCostData(val InvoiceLineItemComponentCostData) InvoiceLineItemComponentCostData2 {
	return InvoiceLineItemComponentCostData2{value: &val}
}
