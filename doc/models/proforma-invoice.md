
# Proforma Invoice

## Structure

`ProformaInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Number` | `*int` | Optional | - |
| `SequenceNumber` | `*int` | Optional | - |
| `CreatedAt` | `*string` | Optional | - |
| `DeliveryDate` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `CollectionMethod` | `*string` | Optional | - |
| `PaymentInstructions` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `ConsolidationLevel` | `*string` | Optional | - |
| `ProductName` | `*string` | Optional | - |
| `ProductFamilyName` | `*string` | Optional | - |
| `Role` | `*string` | Optional | - |
| `Seller` | [`*models.InvoiceSeller`](../../doc/models/invoice-seller.md) | Optional | Information about the seller (merchant) listed on the masthead of the invoice. |
| `Customer` | [`*models.InvoiceCustomer`](../../doc/models/invoice-customer.md) | Optional | Information about the customer who is owner or recipient the invoiced subscription. |
| `Memo` | `*string` | Optional | - |
| `BillingAddress` | [`*models.InvoiceAddress`](../../doc/models/invoice-address.md) | Optional | - |
| `ShippingAddress` | [`*models.InvoiceAddress`](../../doc/models/invoice-address.md) | Optional | - |
| `SubtotalAmount` | `*string` | Optional | - |
| `DiscountAmount` | `*string` | Optional | - |
| `TaxAmount` | `*string` | Optional | - |
| `TotalAmount` | `*string` | Optional | - |
| `CreditAmount` | `*string` | Optional | - |
| `PaidAmount` | `*string` | Optional | - |
| `RefundAmount` | `*string` | Optional | - |
| `DueAmount` | `*string` | Optional | - |
| `LineItems` | [`[]models.InvoiceLineItem`](../../doc/models/invoice-line-item.md) | Optional | - |
| `Discounts` | [`[]models.ProformaInvoiceDiscount`](../../doc/models/proforma-invoice-discount.md) | Optional | - |
| `Taxes` | [`[]models.ProformaInvoiceTax`](../../doc/models/proforma-invoice-tax.md) | Optional | - |
| `Credits` | [`[]models.ProformaInvoiceCredit`](../../doc/models/proforma-invoice-credit.md) | Optional | - |
| `Payments` | [`[]models.ProformaInvoicePayment`](../../doc/models/proforma-invoice-payment.md) | Optional | - |
| `CustomFields` | [`[]models.InvoiceCustomField`](../../doc/models/invoice-custom-field.md) | Optional | - |
| `PublicUrl` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid6",
  "site_id": 196,
  "customer_id": 52,
  "subscription_id": 124,
  "number": 0
}
```

