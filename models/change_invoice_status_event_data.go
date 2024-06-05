package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ChangeInvoiceStatusEventData represents a ChangeInvoiceStatusEventData struct.
// Example schema for an `change_invoice_status` event
type ChangeInvoiceStatusEventData struct {
    // Identifier for the transaction within the payment gateway.
    GatewayTransId       *string                    `json:"gateway_trans_id,omitempty"`
    // The monetary value associated with the linked payment, expressed in dollars.
    Amount               *string                    `json:"amount,omitempty"`
    // The status of the invoice before any changes occurred. See [Invoice Statuses](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405078794253-Introduction-to-Invoices#invoice-statuses) for more.
    FromStatus           InvoiceStatus              `json:"from_status"`
    // The updated status of the invoice after changes have been made. See [Invoice Statuses](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405078794253-Introduction-to-Invoices#invoice-statuses) for more.
    ToStatus             InvoiceStatus              `json:"to_status"`
    ConsolidationLevel   *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ChangeInvoiceStatusEventData.
// It customizes the JSON marshaling process for ChangeInvoiceStatusEventData objects.
func (c ChangeInvoiceStatusEventData) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ChangeInvoiceStatusEventData object to a map representation for JSON marshaling.
func (c ChangeInvoiceStatusEventData) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.GatewayTransId != nil {
        structMap["gateway_trans_id"] = c.GatewayTransId
    }
    if c.Amount != nil {
        structMap["amount"] = c.Amount
    }
    structMap["from_status"] = c.FromStatus
    structMap["to_status"] = c.ToStatus
    if c.ConsolidationLevel != nil {
        structMap["consolidation_level"] = c.ConsolidationLevel
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChangeInvoiceStatusEventData.
// It customizes the JSON unmarshaling process for ChangeInvoiceStatusEventData objects.
func (c *ChangeInvoiceStatusEventData) UnmarshalJSON(input []byte) error {
    var temp changeInvoiceStatusEventData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "gateway_trans_id", "amount", "from_status", "to_status", "consolidation_level")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.GatewayTransId = temp.GatewayTransId
    c.Amount = temp.Amount
    c.FromStatus = *temp.FromStatus
    c.ToStatus = *temp.ToStatus
    c.ConsolidationLevel = temp.ConsolidationLevel
    return nil
}

// changeInvoiceStatusEventData is a temporary struct used for validating the fields of ChangeInvoiceStatusEventData.
type changeInvoiceStatusEventData  struct {
    GatewayTransId     *string                    `json:"gateway_trans_id,omitempty"`
    Amount             *string                    `json:"amount,omitempty"`
    FromStatus         *InvoiceStatus             `json:"from_status"`
    ToStatus           *InvoiceStatus             `json:"to_status"`
    ConsolidationLevel *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
}

func (c *changeInvoiceStatusEventData) validate() error {
    var errs []string
    if c.FromStatus == nil {
        errs = append(errs, "required field `from_status` is missing for type `Change Invoice Status Event Data`")
    }
    if c.ToStatus == nil {
        errs = append(errs, "required field `to_status` is missing for type `Change Invoice Status Event Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
