
# Invoice

## Structure

`Invoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int64` | Optional | - |
| `Uid` | `*string` | Optional | Unique identifier for the invoice. It is generated automatically by Chargify and has the prefix "inv_" followed by alphanumeric characters. |
| `SiteId` | `*int` | Optional | ID of the site to which the invoice belongs. |
| `CustomerId` | `*int` | Optional | ID of the customer to which the invoice belongs. |
| `SubscriptionId` | `*int` | Optional | ID of the subscription that generated the invoice. |
| `Number` | `*string` | Optional | A unique, identifying string that appears on the invoice and in places the invoice is referenced.<br><br>While the UID is long and not appropriate to show to customers, the number is usually shorter and consumable by the customer and the merchant alike. |
| `SequenceNumber` | `*int` | Optional | A monotonically increasing number assigned to invoices as they are created.  This number is unique within a site and can be used to sort and order invoices. |
| `TransactionTime` | `*time.Time` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `IssueDate` | `*string` | Optional | Date the invoice was issued to the customer.  This is the date that the invoice was made available for payment.<br><br>The format is `"YYYY-MM-DD"`. |
| `DueDate` | `*string` | Optional | Date the invoice is due.<br><br>The format is `"YYYY-MM-DD"`. |
| `PaidDate` | `Optional[string]` | Optional | Date the invoice became fully paid.<br><br>If partial payments are applied to the invoice, this date will not be present until payment has been made in full.<br><br>The format is `"YYYY-MM-DD"`. |
| `Status` | [`*models.StatusEnum`](status-enum.md) | Optional | The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more. |
| `Role` | `*string` | Optional | - |
| `ParentInvoiceId` | `Optional[int]` | Optional | - |
| `CollectionMethod` | `*string` | Optional | The collection method of the invoice, which is either "automatic" (tried and retried on an existing payment method by Chargify) or "remittance" (payment must be remitted by the customer or keyed in by the merchant). |
| `PaymentInstructions` | `*string` | Optional | A message that is printed on the invoice when it is marked for remittance collection. It is intended to describe to the customer how they may make payment, and is configured by the merchant. |
| `Currency` | `*string` | Optional | The ISO 4217 currency code (3 character string) representing the currency of invoice transaction. |
| `ConsolidationLevel` | [`*models.InvoiceConsolidationLevelEnum`](invoice-consolidation-level-enum.md) | Optional | Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:<br><br>* "none": A normal invoice with no consolidation.<br>* "child": An invoice segment which has been combined into a consolidated invoice.<br>* "parent": A consolidated invoice, whose contents are composed of invoice segments.<br><br>"Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.<br><br>See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835). |
| `ParentInvoiceUid` | `Optional[string]` | Optional | For invoices with `consolidation_level` of `child`, this specifies the UID of the parent (consolidated) invoice. |
| `SubscriptionGroupId` | `Optional[int]` | Optional | - |
| `ParentInvoiceNumber` | `Optional[int]` | Optional | For invoices with `consolidation_level` of `child`, this specifies the number of the parent (consolidated) invoice. |
| `GroupPrimarySubscriptionId` | `Optional[int]` | Optional | For invoices with `consolidation_level` of `parent`, this specifies the ID of the subscription which was the primary subscription of the subscription group that generated the invoice. |
| `ProductName` | `*string` | Optional | The name of the product subscribed when the invoice was generated. |
| `ProductFamilyName` | `*string` | Optional | The name of the product family subscribed when the invoice was generated. |
| `Seller` | [`*models.InvoiceSeller`](invoice-seller.md) | Optional | Information about the seller (merchant) listed on the masthead of the invoice. |
| `Customer` | [`*models.InvoiceCustomer`](invoice-customer.md) | Optional | Information about the customer who is owner or recipient the invoiced subscription. |
| `Payer` | [`*models.InvoicePayer`](invoice-payer.md) | Optional | - |
| `RecipientEmails` | `[]string` | Optional | **Constraints**: *Maximum Items*: `5` |
| `NetTerms` | `*int` | Optional | - |
| `Memo` | `*string` | Optional | The memo printed on invoices of any collection type.  This message is in control of the merchant. |
| `BillingAddress` | [`*models.InvoiceAddress`](invoice-address.md) | Optional | The invoice billing address. |
| `ShippingAddress` | [`*models.InvoiceAddress`](invoice-address.md) | Optional | The invoice shipping address. |
| `SubtotalAmount` | `*string` | Optional | Subtotal of the invoice, which is the sum of all line items before discounts or taxes. |
| `DiscountAmount` | `*string` | Optional | Total discount applied to the invoice. |
| `TaxAmount` | `*string` | Optional | Total tax on the invoice. |
| `TotalAmount` | `*string` | Optional | The invoice total, which is `subtotal_amount - discount_amount + tax_amount`.' |
| `CreditAmount` | `*string` | Optional | The amount of credit (from credit notes) applied to this invoice.<br><br>Credits offset the amount due from the customer. |
| `RefundAmount` | `*string` | Optional | - |
| `PaidAmount` | `*string` | Optional | The amount paid on the invoice by the customer. |
| `DueAmount` | `*string` | Optional | Amount due on the invoice, which is `total_amount - credit_amount - paid_amount`. |
| `LineItems` | [`[]models.InvoiceLineItem`](invoice-line-item.md) | Optional | Line items on the invoice. |
| `Discounts` | [`[]models.InvoiceDiscount`](invoice-discount.md) | Optional | - |
| `Taxes` | [`[]models.InvoiceTax`](invoice-tax.md) | Optional | - |
| `Credits` | [`[]models.InvoiceCredit`](invoice-credit.md) | Optional | - |
| `Refunds` | [`[]models.InvoiceRefund`](invoice-refund.md) | Optional | - |
| `Payments` | [`[]models.InvoicePayment`](invoice-payment.md) | Optional | - |
| `CustomFields` | [`[]models.InvoiceCustomField`](invoice-custom-field.md) | Optional | - |
| `DisplaySettings` | [`*models.InvoiceDisplaySettings`](invoice-display-settings.md) | Optional | - |
| `PublicUrl` | `*string` | Optional | The public URL of the invoice |
| `PreviousBalanceData` | [`*models.InvoicePreviousBalance`](invoice-previous-balance.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": 252,
  "uid": "uid0",
  "site_id": 178,
  "customer_id": 34,
  "subscription_id": 106
}
```

