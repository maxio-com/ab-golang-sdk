/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "log"
    "strings"
    "time"
)

// ProformaInvoiceIssued represents a ProformaInvoiceIssued struct.
type ProformaInvoiceIssued struct {
    Uid                  string                     `json:"uid"`
    Number               string                     `json:"number"`
    Role                 string                     `json:"role"`
    DeliveryDate         time.Time                  `json:"delivery_date"`
    CreatedAt            time.Time                  `json:"created_at"`
    DueAmount            string                     `json:"due_amount"`
    PaidAmount           string                     `json:"paid_amount"`
    TaxAmount            string                     `json:"tax_amount"`
    TotalAmount          string                     `json:"total_amount"`
    ProductName          string                     `json:"product_name"`
    LineItems            []InvoiceLineItemEventData `json:"line_items"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceIssued.
// It customizes the JSON marshaling process for ProformaInvoiceIssued objects.
func (p ProformaInvoiceIssued) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoiceIssued object to a map representation for JSON marshaling.
func (p ProformaInvoiceIssued) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    structMap["uid"] = p.Uid
    structMap["number"] = p.Number
    structMap["role"] = p.Role
    structMap["delivery_date"] = p.DeliveryDate.Format(DEFAULT_DATE)
    structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    structMap["due_amount"] = p.DueAmount
    structMap["paid_amount"] = p.PaidAmount
    structMap["tax_amount"] = p.TaxAmount
    structMap["total_amount"] = p.TotalAmount
    structMap["product_name"] = p.ProductName
    structMap["line_items"] = p.LineItems
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceIssued.
// It customizes the JSON unmarshaling process for ProformaInvoiceIssued objects.
func (p *ProformaInvoiceIssued) UnmarshalJSON(input []byte) error {
    var temp tempProformaInvoiceIssued
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "number", "role", "delivery_date", "created_at", "due_amount", "paid_amount", "tax_amount", "total_amount", "product_name", "line_items")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Uid = *temp.Uid
    p.Number = *temp.Number
    p.Role = *temp.Role
    DeliveryDateVal, err := time.Parse(DEFAULT_DATE, *temp.DeliveryDate)
    if err != nil {
        log.Fatalf("Cannot Parse delivery_date as % s format.", DEFAULT_DATE)
    }
    p.DeliveryDate = DeliveryDateVal
    CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
    if err != nil {
        log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
    }
    p.CreatedAt = CreatedAtVal
    p.DueAmount = *temp.DueAmount
    p.PaidAmount = *temp.PaidAmount
    p.TaxAmount = *temp.TaxAmount
    p.TotalAmount = *temp.TotalAmount
    p.ProductName = *temp.ProductName
    p.LineItems = *temp.LineItems
    return nil
}

// tempProformaInvoiceIssued is a temporary struct used for validating the fields of ProformaInvoiceIssued.
type tempProformaInvoiceIssued  struct {
    Uid          *string                     `json:"uid"`
    Number       *string                     `json:"number"`
    Role         *string                     `json:"role"`
    DeliveryDate *string                     `json:"delivery_date"`
    CreatedAt    *string                     `json:"created_at"`
    DueAmount    *string                     `json:"due_amount"`
    PaidAmount   *string                     `json:"paid_amount"`
    TaxAmount    *string                     `json:"tax_amount"`
    TotalAmount  *string                     `json:"total_amount"`
    ProductName  *string                     `json:"product_name"`
    LineItems    *[]InvoiceLineItemEventData `json:"line_items"`
}

func (p *tempProformaInvoiceIssued) validate() error {
    var errs []string
    if p.Uid == nil {
        errs = append(errs, "required field `uid` is missing for type `Proforma Invoice Issued`")
    }
    if p.Number == nil {
        errs = append(errs, "required field `number` is missing for type `Proforma Invoice Issued`")
    }
    if p.Role == nil {
        errs = append(errs, "required field `role` is missing for type `Proforma Invoice Issued`")
    }
    if p.DeliveryDate == nil {
        errs = append(errs, "required field `delivery_date` is missing for type `Proforma Invoice Issued`")
    }
    if p.CreatedAt == nil {
        errs = append(errs, "required field `created_at` is missing for type `Proforma Invoice Issued`")
    }
    if p.DueAmount == nil {
        errs = append(errs, "required field `due_amount` is missing for type `Proforma Invoice Issued`")
    }
    if p.PaidAmount == nil {
        errs = append(errs, "required field `paid_amount` is missing for type `Proforma Invoice Issued`")
    }
    if p.TaxAmount == nil {
        errs = append(errs, "required field `tax_amount` is missing for type `Proforma Invoice Issued`")
    }
    if p.TotalAmount == nil {
        errs = append(errs, "required field `total_amount` is missing for type `Proforma Invoice Issued`")
    }
    if p.ProductName == nil {
        errs = append(errs, "required field `product_name` is missing for type `Proforma Invoice Issued`")
    }
    if p.LineItems == nil {
        errs = append(errs, "required field `line_items` is missing for type `Proforma Invoice Issued`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
