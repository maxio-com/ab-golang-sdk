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

// IssueInvoiceEventData represents a IssueInvoiceEventData struct.
// Example schema for an `issue_invoice` event
type IssueInvoiceEventData struct {
    // Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
    // * "none": A normal invoice with no consolidation.
    // * "child": An invoice segment which has been combined into a consolidated invoice.
    // * "parent": A consolidated invoice, whose contents are composed of invoice segments.
    // "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
    // See also the [invoice consolidation documentation](https://maxio.zendesk.com/hc/en-us/articles/24252269909389-Invoice-Consolidation).
    ConsolidationLevel   InvoiceConsolidationLevel `json:"consolidation_level"`
    // The status of the invoice before event occurrence. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more.
    FromStatus           InvoiceStatus             `json:"from_status"`
    // The status of the invoice after event occurrence. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more.
    ToStatus             InvoiceStatus             `json:"to_status"`
    // Amount due on the invoice, which is `total_amount - credit_amount - paid_amount`.
    DueAmount            string                    `json:"due_amount"`
    // The invoice total, which is `subtotal_amount - discount_amount + tax_amount`.'
    TotalAmount          string                    `json:"total_amount"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssueInvoiceEventData.
// It customizes the JSON marshaling process for IssueInvoiceEventData objects.
func (i IssueInvoiceEventData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "consolidation_level", "from_status", "to_status", "due_amount", "total_amount"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssueInvoiceEventData object to a map representation for JSON marshaling.
func (i IssueInvoiceEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    structMap["consolidation_level"] = i.ConsolidationLevel
    structMap["from_status"] = i.FromStatus
    structMap["to_status"] = i.ToStatus
    structMap["due_amount"] = i.DueAmount
    structMap["total_amount"] = i.TotalAmount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssueInvoiceEventData.
// It customizes the JSON unmarshaling process for IssueInvoiceEventData objects.
func (i *IssueInvoiceEventData) UnmarshalJSON(input []byte) error {
    var temp tempIssueInvoiceEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "consolidation_level", "from_status", "to_status", "due_amount", "total_amount")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.ConsolidationLevel = *temp.ConsolidationLevel
    i.FromStatus = *temp.FromStatus
    i.ToStatus = *temp.ToStatus
    i.DueAmount = *temp.DueAmount
    i.TotalAmount = *temp.TotalAmount
    return nil
}

// tempIssueInvoiceEventData is a temporary struct used for validating the fields of IssueInvoiceEventData.
type tempIssueInvoiceEventData  struct {
    ConsolidationLevel *InvoiceConsolidationLevel `json:"consolidation_level"`
    FromStatus         *InvoiceStatus             `json:"from_status"`
    ToStatus           *InvoiceStatus             `json:"to_status"`
    DueAmount          *string                    `json:"due_amount"`
    TotalAmount        *string                    `json:"total_amount"`
}

func (i *tempIssueInvoiceEventData) validate() error {
    var errs []string
    if i.ConsolidationLevel == nil {
        errs = append(errs, "required field `consolidation_level` is missing for type `Issue Invoice Event Data`")
    }
    if i.FromStatus == nil {
        errs = append(errs, "required field `from_status` is missing for type `Issue Invoice Event Data`")
    }
    if i.ToStatus == nil {
        errs = append(errs, "required field `to_status` is missing for type `Issue Invoice Event Data`")
    }
    if i.DueAmount == nil {
        errs = append(errs, "required field `due_amount` is missing for type `Issue Invoice Event Data`")
    }
    if i.TotalAmount == nil {
        errs = append(errs, "required field `total_amount` is missing for type `Issue Invoice Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
