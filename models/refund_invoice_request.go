package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RefundInvoiceRequest represents a RefundInvoiceRequest struct.
type RefundInvoiceRequest struct {
    Refund               RefundInvoiceRequestRefund `json:"refund"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RefundInvoiceRequest.
// It customizes the JSON marshaling process for RefundInvoiceRequest objects.
func (r RefundInvoiceRequest) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RefundInvoiceRequest object to a map representation for JSON marshaling.
func (r RefundInvoiceRequest) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["refund"] = r.Refund.toMap()
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RefundInvoiceRequest.
// It customizes the JSON unmarshaling process for RefundInvoiceRequest objects.
func (r *RefundInvoiceRequest) UnmarshalJSON(input []byte) error {
    var temp refundInvoiceRequest
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "refund")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Refund = *temp.Refund
    return nil
}

// TODO
type refundInvoiceRequest  struct {
    Refund *RefundInvoiceRequestRefund `json:"refund"`
}

func (r *refundInvoiceRequest) validate() error {
    var errs []string
    if r.Refund == nil {
        errs = append(errs, "required field `refund` is missing for type `Refund Invoice Request`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
