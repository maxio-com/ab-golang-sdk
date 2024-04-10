package models

import (
    "encoding/json"
    "log"
    "time"
)

// ProformaInvoice represents a ProformaInvoice struct.
type ProformaInvoice struct {
    Uid                  *string                    `json:"uid,omitempty"`
    SiteId               *int                       `json:"site_id,omitempty"`
    CustomerId           Optional[int]              `json:"customer_id"`
    SubscriptionId       Optional[int]              `json:"subscription_id"`
    Number               Optional[int]              `json:"number"`
    SequenceNumber       Optional[int]              `json:"sequence_number"`
    CreatedAt            *time.Time                 `json:"created_at,omitempty"`
    DeliveryDate         *time.Time                 `json:"delivery_date,omitempty"`
    Status               *ProformaInvoiceStatus     `json:"status,omitempty"`
    // The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
    CollectionMethod     *CollectionMethod          `json:"collection_method,omitempty"`
    PaymentInstructions  *string                    `json:"payment_instructions,omitempty"`
    Currency             *string                    `json:"currency,omitempty"`
    // Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
    // * "none": A normal invoice with no consolidation.
    // * "child": An invoice segment which has been combined into a consolidated invoice.
    // * "parent": A consolidated invoice, whose contents are composed of invoice segments.
    // "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
    // See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).
    ConsolidationLevel   *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
    ProductName          *string                    `json:"product_name,omitempty"`
    ProductFamilyName    *string                    `json:"product_family_name,omitempty"`
    // 'proforma' value is deprecated in favor of proforma_adhoc and proforma_automatic
    Role                 *ProformaInvoiceRole       `json:"role,omitempty"`
    // Information about the seller (merchant) listed on the masthead of the invoice.
    Seller               *InvoiceSeller             `json:"seller,omitempty"`
    // Information about the customer who is owner or recipient the invoiced subscription.
    Customer             *InvoiceCustomer           `json:"customer,omitempty"`
    Memo                 *string                    `json:"memo,omitempty"`
    BillingAddress       *InvoiceAddress            `json:"billing_address,omitempty"`
    ShippingAddress      *InvoiceAddress            `json:"shipping_address,omitempty"`
    SubtotalAmount       *string                    `json:"subtotal_amount,omitempty"`
    DiscountAmount       *string                    `json:"discount_amount,omitempty"`
    TaxAmount            *string                    `json:"tax_amount,omitempty"`
    TotalAmount          *string                    `json:"total_amount,omitempty"`
    CreditAmount         *string                    `json:"credit_amount,omitempty"`
    PaidAmount           *string                    `json:"paid_amount,omitempty"`
    RefundAmount         *string                    `json:"refund_amount,omitempty"`
    DueAmount            *string                    `json:"due_amount,omitempty"`
    LineItems            []InvoiceLineItem          `json:"line_items,omitempty"`
    Discounts            []ProformaInvoiceDiscount  `json:"discounts,omitempty"`
    Taxes                []ProformaInvoiceTax       `json:"taxes,omitempty"`
    Credits              []ProformaInvoiceCredit    `json:"credits,omitempty"`
    Payments             []ProformaInvoicePayment   `json:"payments,omitempty"`
    CustomFields         []InvoiceCustomField       `json:"custom_fields,omitempty"`
    PublicUrl            Optional[string]           `json:"public_url"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ProformaInvoice.
// It customizes the JSON marshaling process for ProformaInvoice objects.
func (p ProformaInvoice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the ProformaInvoice object to a map representation for JSON marshaling.
func (p ProformaInvoice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Uid != nil {
        structMap["uid"] = p.Uid
    }
    if p.SiteId != nil {
        structMap["site_id"] = p.SiteId
    }
    if p.CustomerId.IsValueSet() {
        if p.CustomerId.Value() != nil {
            structMap["customer_id"] = p.CustomerId.Value()
        } else {
            structMap["customer_id"] = nil
        }
    }
    if p.SubscriptionId.IsValueSet() {
        if p.SubscriptionId.Value() != nil {
            structMap["subscription_id"] = p.SubscriptionId.Value()
        } else {
            structMap["subscription_id"] = nil
        }
    }
    if p.Number.IsValueSet() {
        if p.Number.Value() != nil {
            structMap["number"] = p.Number.Value()
        } else {
            structMap["number"] = nil
        }
    }
    if p.SequenceNumber.IsValueSet() {
        if p.SequenceNumber.Value() != nil {
            structMap["sequence_number"] = p.SequenceNumber.Value()
        } else {
            structMap["sequence_number"] = nil
        }
    }
    if p.CreatedAt != nil {
        structMap["created_at"] = p.CreatedAt.Format(time.RFC3339)
    }
    if p.DeliveryDate != nil {
        structMap["delivery_date"] = p.DeliveryDate.Format(DEFAULT_DATE)
    }
    if p.Status != nil {
        structMap["status"] = p.Status
    }
    if p.CollectionMethod != nil {
        structMap["collection_method"] = p.CollectionMethod
    }
    if p.PaymentInstructions != nil {
        structMap["payment_instructions"] = p.PaymentInstructions
    }
    if p.Currency != nil {
        structMap["currency"] = p.Currency
    }
    if p.ConsolidationLevel != nil {
        structMap["consolidation_level"] = p.ConsolidationLevel
    }
    if p.ProductName != nil {
        structMap["product_name"] = p.ProductName
    }
    if p.ProductFamilyName != nil {
        structMap["product_family_name"] = p.ProductFamilyName
    }
    if p.Role != nil {
        structMap["role"] = p.Role
    }
    if p.Seller != nil {
        structMap["seller"] = p.Seller.toMap()
    }
    if p.Customer != nil {
        structMap["customer"] = p.Customer.toMap()
    }
    if p.Memo != nil {
        structMap["memo"] = p.Memo
    }
    if p.BillingAddress != nil {
        structMap["billing_address"] = p.BillingAddress.toMap()
    }
    if p.ShippingAddress != nil {
        structMap["shipping_address"] = p.ShippingAddress.toMap()
    }
    if p.SubtotalAmount != nil {
        structMap["subtotal_amount"] = p.SubtotalAmount
    }
    if p.DiscountAmount != nil {
        structMap["discount_amount"] = p.DiscountAmount
    }
    if p.TaxAmount != nil {
        structMap["tax_amount"] = p.TaxAmount
    }
    if p.TotalAmount != nil {
        structMap["total_amount"] = p.TotalAmount
    }
    if p.CreditAmount != nil {
        structMap["credit_amount"] = p.CreditAmount
    }
    if p.PaidAmount != nil {
        structMap["paid_amount"] = p.PaidAmount
    }
    if p.RefundAmount != nil {
        structMap["refund_amount"] = p.RefundAmount
    }
    if p.DueAmount != nil {
        structMap["due_amount"] = p.DueAmount
    }
    if p.LineItems != nil {
        structMap["line_items"] = p.LineItems
    }
    if p.Discounts != nil {
        structMap["discounts"] = p.Discounts
    }
    if p.Taxes != nil {
        structMap["taxes"] = p.Taxes
    }
    if p.Credits != nil {
        structMap["credits"] = p.Credits
    }
    if p.Payments != nil {
        structMap["payments"] = p.Payments
    }
    if p.CustomFields != nil {
        structMap["custom_fields"] = p.CustomFields
    }
    if p.PublicUrl.IsValueSet() {
        if p.PublicUrl.Value() != nil {
            structMap["public_url"] = p.PublicUrl.Value()
        } else {
            structMap["public_url"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoice.
// It customizes the JSON unmarshaling process for ProformaInvoice objects.
func (p *ProformaInvoice) UnmarshalJSON(input []byte) error {
    var temp proformaInvoice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "uid", "site_id", "customer_id", "subscription_id", "number", "sequence_number", "created_at", "delivery_date", "status", "collection_method", "payment_instructions", "currency", "consolidation_level", "product_name", "product_family_name", "role", "seller", "customer", "memo", "billing_address", "shipping_address", "subtotal_amount", "discount_amount", "tax_amount", "total_amount", "credit_amount", "paid_amount", "refund_amount", "due_amount", "line_items", "discounts", "taxes", "credits", "payments", "custom_fields", "public_url")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Uid = temp.Uid
    p.SiteId = temp.SiteId
    p.CustomerId = temp.CustomerId
    p.SubscriptionId = temp.SubscriptionId
    p.Number = temp.Number
    p.SequenceNumber = temp.SequenceNumber
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        p.CreatedAt = &CreatedAtVal
    }
    if temp.DeliveryDate != nil {
        DeliveryDateVal, err := time.Parse(DEFAULT_DATE, *temp.DeliveryDate)
        if err != nil {
            log.Fatalf("Cannot Parse delivery_date as % s format.", DEFAULT_DATE)
        }
        p.DeliveryDate = &DeliveryDateVal
    }
    p.Status = temp.Status
    p.CollectionMethod = temp.CollectionMethod
    p.PaymentInstructions = temp.PaymentInstructions
    p.Currency = temp.Currency
    p.ConsolidationLevel = temp.ConsolidationLevel
    p.ProductName = temp.ProductName
    p.ProductFamilyName = temp.ProductFamilyName
    p.Role = temp.Role
    p.Seller = temp.Seller
    p.Customer = temp.Customer
    p.Memo = temp.Memo
    p.BillingAddress = temp.BillingAddress
    p.ShippingAddress = temp.ShippingAddress
    p.SubtotalAmount = temp.SubtotalAmount
    p.DiscountAmount = temp.DiscountAmount
    p.TaxAmount = temp.TaxAmount
    p.TotalAmount = temp.TotalAmount
    p.CreditAmount = temp.CreditAmount
    p.PaidAmount = temp.PaidAmount
    p.RefundAmount = temp.RefundAmount
    p.DueAmount = temp.DueAmount
    p.LineItems = temp.LineItems
    p.Discounts = temp.Discounts
    p.Taxes = temp.Taxes
    p.Credits = temp.Credits
    p.Payments = temp.Payments
    p.CustomFields = temp.CustomFields
    p.PublicUrl = temp.PublicUrl
    return nil
}

// TODO
type proformaInvoice  struct {
    Uid                 *string                    `json:"uid,omitempty"`
    SiteId              *int                       `json:"site_id,omitempty"`
    CustomerId          Optional[int]              `json:"customer_id"`
    SubscriptionId      Optional[int]              `json:"subscription_id"`
    Number              Optional[int]              `json:"number"`
    SequenceNumber      Optional[int]              `json:"sequence_number"`
    CreatedAt           *string                    `json:"created_at,omitempty"`
    DeliveryDate        *string                    `json:"delivery_date,omitempty"`
    Status              *ProformaInvoiceStatus     `json:"status,omitempty"`
    CollectionMethod    *CollectionMethod          `json:"collection_method,omitempty"`
    PaymentInstructions *string                    `json:"payment_instructions,omitempty"`
    Currency            *string                    `json:"currency,omitempty"`
    ConsolidationLevel  *InvoiceConsolidationLevel `json:"consolidation_level,omitempty"`
    ProductName         *string                    `json:"product_name,omitempty"`
    ProductFamilyName   *string                    `json:"product_family_name,omitempty"`
    Role                *ProformaInvoiceRole       `json:"role,omitempty"`
    Seller              *InvoiceSeller             `json:"seller,omitempty"`
    Customer            *InvoiceCustomer           `json:"customer,omitempty"`
    Memo                *string                    `json:"memo,omitempty"`
    BillingAddress      *InvoiceAddress            `json:"billing_address,omitempty"`
    ShippingAddress     *InvoiceAddress            `json:"shipping_address,omitempty"`
    SubtotalAmount      *string                    `json:"subtotal_amount,omitempty"`
    DiscountAmount      *string                    `json:"discount_amount,omitempty"`
    TaxAmount           *string                    `json:"tax_amount,omitempty"`
    TotalAmount         *string                    `json:"total_amount,omitempty"`
    CreditAmount        *string                    `json:"credit_amount,omitempty"`
    PaidAmount          *string                    `json:"paid_amount,omitempty"`
    RefundAmount        *string                    `json:"refund_amount,omitempty"`
    DueAmount           *string                    `json:"due_amount,omitempty"`
    LineItems           []InvoiceLineItem          `json:"line_items,omitempty"`
    Discounts           []ProformaInvoiceDiscount  `json:"discounts,omitempty"`
    Taxes               []ProformaInvoiceTax       `json:"taxes,omitempty"`
    Credits             []ProformaInvoiceCredit    `json:"credits,omitempty"`
    Payments            []ProformaInvoicePayment   `json:"payments,omitempty"`
    CustomFields        []InvoiceCustomField       `json:"custom_fields,omitempty"`
    PublicUrl           Optional[string]           `json:"public_url"`
}
