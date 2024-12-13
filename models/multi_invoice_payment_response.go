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

// MultiInvoicePaymentResponse represents a MultiInvoicePaymentResponse struct.
type MultiInvoicePaymentResponse struct {
    Payment              MultiInvoicePayment    `json:"payment"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MultiInvoicePaymentResponse.
// It customizes the JSON marshaling process for MultiInvoicePaymentResponse objects.
func (m MultiInvoicePaymentResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "payment"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MultiInvoicePaymentResponse object to a map representation for JSON marshaling.
func (m MultiInvoicePaymentResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["payment"] = m.Payment.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MultiInvoicePaymentResponse.
// It customizes the JSON unmarshaling process for MultiInvoicePaymentResponse objects.
func (m *MultiInvoicePaymentResponse) UnmarshalJSON(input []byte) error {
    var temp tempMultiInvoicePaymentResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "payment")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Payment = *temp.Payment
    return nil
}

// tempMultiInvoicePaymentResponse is a temporary struct used for validating the fields of MultiInvoicePaymentResponse.
type tempMultiInvoicePaymentResponse  struct {
    Payment *MultiInvoicePayment `json:"payment"`
}

func (m *tempMultiInvoicePaymentResponse) validate() error {
    var errs []string
    if m.Payment == nil {
        errs = append(errs, "required field `payment` is missing for type `Multi Invoice Payment Response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
