
# Debit Note

## Structure

`DebitNote`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | Unique identifier for the debit note. It is generated automatically by Chargify and has the prefix "db_" followed by alphanumeric characters. |
| `SiteId` | `*int` | Optional | ID of the site to which the debit note belongs. |
| `CustomerId` | `*int` | Optional | ID of the customer to which the debit note belongs. |
| `SubscriptionId` | `*int` | Optional | ID of the subscription that generated the debit note. |
| `Number` | `*int` | Optional | A unique, identifier that appears on the debit note and in places it is referenced. |
| `SequenceNumber` | `*int` | Optional | A monotonically increasing number assigned to debit notes as they are created. |
| `OriginCreditNoteUid` | `*string` | Optional | Unique identifier for the connected credit note. It is generated automatically by Chargify and has the prefix "cn_" followed by alphanumeric characters.<br><br>While the UID is long and not appropriate to show to customers, the number is usually shorter and consumable by the customer and the merchant alike. |
| `OriginCreditNoteNumber` | `*string` | Optional | A unique, identifying string of the connected credit note. |
| `IssueDate` | `*time.Time` | Optional | Date the document was issued to the customer. This is the date that the document was made available for payment.<br><br>The format is "YYYY-MM-DD". |
| `AppliedDate` | `*time.Time` | Optional | Debit notes are applied to invoices to offset invoiced amounts - they adjust the amount due. This field is the date the debit note document became fully applied to the invoice.<br><br>The format is "YYYY-MM-DD". |
| `DueDate` | `*time.Time` | Optional | Date the document is due for payment. The format is "YYYY-MM-DD". |
| `Status` | [`*models.DebitNoteStatus`](../../doc/models/debit-note-status.md) | Optional | Current status of the debit note. |
| `Memo` | `*string` | Optional | The memo printed on debit note, which is a description of the reason for the debit. |
| `Role` | [`*models.DebitNoteRole`](../../doc/models/debit-note-role.md) | Optional | The role of the debit note. |
| `Currency` | `*string` | Optional | The ISO 4217 currency code (3 character string) representing the currency of the credit note amount fields. |
| `Seller` | [`*models.InvoiceSeller`](../../doc/models/invoice-seller.md) | Optional | Information about the seller (merchant) listed on the masthead of the debit note. |
| `Customer` | [`*models.InvoiceCustomer`](../../doc/models/invoice-customer.md) | Optional | Information about the customer who is owner or recipient the debited subscription. |
| `BillingAddress` | [`*models.InvoiceAddress`](../../doc/models/invoice-address.md) | Optional | The billing address of the debited subscription. |
| `ShippingAddress` | [`*models.InvoiceAddress`](../../doc/models/invoice-address.md) | Optional | The shipping address of the debited subscription. |
| `LineItems` | [`[]models.CreditNoteLineItem`](../../doc/models/credit-note-line-item.md) | Optional | Line items on the debit note. |
| `Discounts` | [`[]models.InvoiceDiscount`](../../doc/models/invoice-discount.md) | Optional | - |
| `Taxes` | [`[]models.InvoiceTax`](../../doc/models/invoice-tax.md) | Optional | - |
| `Refunds` | [`[]models.InvoiceRefund`](../../doc/models/invoice-refund.md) | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid2",
  "site_id": 112,
  "customer_id": 224,
  "subscription_id": 40,
  "number": 172
}
```

