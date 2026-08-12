// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UpdateInvoiceRequest represents a UpdateInvoiceRequest struct.
// Request payload for updating a draft ad hoc invoice.
type UpdateInvoiceRequest struct {
    // Attributes of a draft ad hoc invoice which can be updated. Only the submitted attributes are changed.
    Invoice              UpdateInvoice          `json:"invoice"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateInvoiceRequest,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateInvoiceRequest) String() string {
    return fmt.Sprintf(
    	"UpdateInvoiceRequest[Invoice=%v, AdditionalProperties=%v]",
    	u.Invoice, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateInvoiceRequest.
// It customizes the JSON marshaling process for UpdateInvoiceRequest objects.
func (u UpdateInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "invoice"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateInvoiceRequest object to a map representation for JSON marshaling.
func (u UpdateInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["invoice"] = u.Invoice.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateInvoiceRequest.
// It customizes the JSON unmarshaling process for UpdateInvoiceRequest objects.
func (u *UpdateInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp tempUpdateInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "invoice")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Invoice = *temp.Invoice
    return nil
}

// tempUpdateInvoiceRequest is a temporary struct used for validating the fields of UpdateInvoiceRequest.
type tempUpdateInvoiceRequest  struct {
    Invoice *UpdateInvoice `json:"invoice"`
}

func (u *tempUpdateInvoiceRequest) validate() error {
    var errs []string
    if u.Invoice == nil {
        errs = append(errs, "required field `invoice` is missing for type `Update Invoice Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
