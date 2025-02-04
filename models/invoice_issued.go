/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "log"
    "strings"
    "time"
)

// InvoiceIssued represents a InvoiceIssued struct.
type InvoiceIssued struct {
    Uid                  string                     `json:"uid"`
    Number               string                     `json:"number"`
    Role                 string                     `json:"role"`
    DueDate              *time.Time                 `json:"due_date"`
    // Invoice issue date. Can be an empty string if value is missing.
    IssueDate            string                     `json:"issue_date"`
    // Paid date. Can be an empty string if value is missing.
    PaidDate             string                     `json:"paid_date"`
    DueAmount            string                     `json:"due_amount"`
    PaidAmount           string                     `json:"paid_amount"`
    TaxAmount            string                     `json:"tax_amount"`
    RefundAmount         string                     `json:"refund_amount"`
    TotalAmount          string                     `json:"total_amount"`
    StatusAmount         string                     `json:"status_amount"`
    ProductName          string                     `json:"product_name"`
    ConsolidationLevel   string                     `json:"consolidation_level"`
    LineItems            []InvoiceLineItemEventData `json:"line_items"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for InvoiceIssued,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InvoiceIssued) String() string {
    return fmt.Sprintf(
    	"InvoiceIssued[Uid=%v, Number=%v, Role=%v, DueDate=%v, IssueDate=%v, PaidDate=%v, DueAmount=%v, PaidAmount=%v, TaxAmount=%v, RefundAmount=%v, TotalAmount=%v, StatusAmount=%v, ProductName=%v, ConsolidationLevel=%v, LineItems=%v, AdditionalProperties=%v]",
    	i.Uid, i.Number, i.Role, i.DueDate, i.IssueDate, i.PaidDate, i.DueAmount, i.PaidAmount, i.TaxAmount, i.RefundAmount, i.TotalAmount, i.StatusAmount, i.ProductName, i.ConsolidationLevel, i.LineItems, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InvoiceIssued.
// It customizes the JSON marshaling process for InvoiceIssued objects.
func (i InvoiceIssued) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "uid", "number", "role", "due_date", "issue_date", "paid_date", "due_amount", "paid_amount", "tax_amount", "refund_amount", "total_amount", "status_amount", "product_name", "consolidation_level", "line_items"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InvoiceIssued object to a map representation for JSON marshaling.
func (i InvoiceIssued) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["uid"] = i.Uid
    structMap["number"] = i.Number
    structMap["role"] = i.Role
    if i.DueDate != nil {
        structMap["due_date"] = i.DueDate.Format(DEFAULT_DATE)
    } else {
        structMap["due_date"] = nil
    }
    structMap["issue_date"] = i.IssueDate
    structMap["paid_date"] = i.PaidDate
    structMap["due_amount"] = i.DueAmount
    structMap["paid_amount"] = i.PaidAmount
    structMap["tax_amount"] = i.TaxAmount
    structMap["refund_amount"] = i.RefundAmount
    structMap["total_amount"] = i.TotalAmount
    structMap["status_amount"] = i.StatusAmount
    structMap["product_name"] = i.ProductName
    structMap["consolidation_level"] = i.ConsolidationLevel
    structMap["line_items"] = i.LineItems
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceIssued.
// It customizes the JSON unmarshaling process for InvoiceIssued objects.
func (i *InvoiceIssued) UnmarshalJSON(input []byte) error {
    var temp tempInvoiceIssued
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "uid", "number", "role", "due_date", "issue_date", "paid_date", "due_amount", "paid_amount", "tax_amount", "refund_amount", "total_amount", "status_amount", "product_name", "consolidation_level", "line_items")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Uid = *temp.Uid
    i.Number = *temp.Number
    i.Role = *temp.Role
    if temp.DueDate != nil {
        DueDateVal, err := time.Parse(DEFAULT_DATE, *temp.DueDate)
        if err != nil {
            log.Fatalf("Cannot Parse due_date as % s format.", DEFAULT_DATE)
        }
        i.DueDate = &DueDateVal
    }
    i.IssueDate = *temp.IssueDate
    i.PaidDate = *temp.PaidDate
    i.DueAmount = *temp.DueAmount
    i.PaidAmount = *temp.PaidAmount
    i.TaxAmount = *temp.TaxAmount
    i.RefundAmount = *temp.RefundAmount
    i.TotalAmount = *temp.TotalAmount
    i.StatusAmount = *temp.StatusAmount
    i.ProductName = *temp.ProductName
    i.ConsolidationLevel = *temp.ConsolidationLevel
    i.LineItems = *temp.LineItems
    return nil
}

// tempInvoiceIssued is a temporary struct used for validating the fields of InvoiceIssued.
type tempInvoiceIssued  struct {
    Uid                *string                     `json:"uid"`
    Number             *string                     `json:"number"`
    Role               *string                     `json:"role"`
    DueDate            *string                     `json:"due_date"`
    IssueDate          *string                     `json:"issue_date"`
    PaidDate           *string                     `json:"paid_date"`
    DueAmount          *string                     `json:"due_amount"`
    PaidAmount         *string                     `json:"paid_amount"`
    TaxAmount          *string                     `json:"tax_amount"`
    RefundAmount       *string                     `json:"refund_amount"`
    TotalAmount        *string                     `json:"total_amount"`
    StatusAmount       *string                     `json:"status_amount"`
    ProductName        *string                     `json:"product_name"`
    ConsolidationLevel *string                     `json:"consolidation_level"`
    LineItems          *[]InvoiceLineItemEventData `json:"line_items"`
}

func (i *tempInvoiceIssued) validate() error {
    var errs []string
    if i.Uid == nil {
        errs = append(errs, "required field `uid` is missing for type `Invoice Issued`")
    }
    if i.Number == nil {
        errs = append(errs, "required field `number` is missing for type `Invoice Issued`")
    }
    if i.Role == nil {
        errs = append(errs, "required field `role` is missing for type `Invoice Issued`")
    }
    if i.IssueDate == nil {
        errs = append(errs, "required field `issue_date` is missing for type `Invoice Issued`")
    }
    if i.PaidDate == nil {
        errs = append(errs, "required field `paid_date` is missing for type `Invoice Issued`")
    }
    if i.DueAmount == nil {
        errs = append(errs, "required field `due_amount` is missing for type `Invoice Issued`")
    }
    if i.PaidAmount == nil {
        errs = append(errs, "required field `paid_amount` is missing for type `Invoice Issued`")
    }
    if i.TaxAmount == nil {
        errs = append(errs, "required field `tax_amount` is missing for type `Invoice Issued`")
    }
    if i.RefundAmount == nil {
        errs = append(errs, "required field `refund_amount` is missing for type `Invoice Issued`")
    }
    if i.TotalAmount == nil {
        errs = append(errs, "required field `total_amount` is missing for type `Invoice Issued`")
    }
    if i.StatusAmount == nil {
        errs = append(errs, "required field `status_amount` is missing for type `Invoice Issued`")
    }
    if i.ProductName == nil {
        errs = append(errs, "required field `product_name` is missing for type `Invoice Issued`")
    }
    if i.ConsolidationLevel == nil {
        errs = append(errs, "required field `consolidation_level` is missing for type `Invoice Issued`")
    }
    if i.LineItems == nil {
        errs = append(errs, "required field `line_items` is missing for type `Invoice Issued`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
