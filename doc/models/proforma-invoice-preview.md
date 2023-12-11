
# Proforma Invoice Preview

## Structure

`ProformaInvoicePreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `SiteId` | `*int` | Optional | - |
| `CustomerId` | `*int` | Optional | - |
| `SubscriptionId` | `*int` | Optional | - |
| `Number` | `*string` | Optional | - |
| `SequenceNumber` | `*int` | Optional | - |
| `CreatedAt` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DeliveryDate` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Status` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `CollectionMethod` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PaymentInstructions` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Currency` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `ConsolidationLevel` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `ProductName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `ProductFamilyName` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Role` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Seller` | [`*models.InvoiceSeller`](invoice-seller.md) | Optional | Information about the seller (merchant) listed on the masthead of the invoice. |
| `Customer` | [`*models.InvoiceCustomer`](invoice-customer.md) | Optional | Information about the customer who is owner or recipient the invoiced subscription. |
| `Memo` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `BillingAddress` | [`*models.InvoiceAddress`](invoice-address.md) | Optional | - |
| `ShippingAddress` | [`*models.InvoiceAddress`](invoice-address.md) | Optional | - |
| `SubtotalAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DiscountAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TaxAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `TotalAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `CreditAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `PaidAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `RefundAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `DueAmount` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `LineItems` | [`[]models.InvoiceLineItem`](invoice-line-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Discounts` | [`[]models.ProformaInvoiceDiscount`](proforma-invoice-discount.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Taxes` | [`[]models.ProformaInvoiceTax`](proforma-invoice-tax.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Credits` | [`[]models.ProformaInvoiceCredit`](proforma-invoice-credit.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Payments` | [`[]models.ProformaInvoicePayment`](proforma-invoice-payment.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `CustomFields` | [`[]models.ProformaCustomField`](proforma-custom-field.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `PublicUrl` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "uid": "uid2",
  "site_id": 214,
  "customer_id": 70,
  "subscription_id": 142,
  "number": "number0"
}
```

