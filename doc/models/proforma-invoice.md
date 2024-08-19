
# Proforma Invoice

## Structure

`ProformaInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Uid` | `*string` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `CustomerId` | `models.Optional[int]` | Optional | - |
| `SubscriptionId` | `models.Optional[int]` | Optional | - |
| `Number` | `models.Optional[int]` | Optional | - |
| `SequenceNumber` | `models.Optional[int]` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `DeliveryDate` | `*time.Time` | Optional | - |
| `Status` | [`*models.ProformaInvoiceStatus`](../../doc/models/proforma-invoice-status.md) | Optional | - |
| `CollectionMethod` | [`*models.CollectionMethod`](../../doc/models/collection-method.md) | Optional | The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`. |
| `PaymentInstructions` | `*string` | Optional | - |
| `Currency` | `*string` | Optional | - |
| `ConsolidationLevel` | [`*models.InvoiceConsolidationLevel`](../../doc/models/invoice-consolidation-level.md) | Optional | Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:<br><br>* "none": A normal invoice with no consolidation.<br>* "child": An invoice segment which has been combined into a consolidated invoice.<br>* "parent": A consolidated invoice, whose contents are composed of invoice segments.<br><br>"Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.<br><br>See also the [invoice consolidation documentation](https://maxio.zendesk.com/hc/en-us/articles/24252269909389-Invoice-Consolidation). |
| `ProductName` | `*string` | Optional | - |
| `ProductFamilyName` | `*string` | Optional | - |
| `Role` | [`*models.ProformaInvoiceRole`](../../doc/models/proforma-invoice-role.md) | Optional | 'proforma' value is deprecated in favor of proforma_adhoc and proforma_automatic |
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
| `PublicUrl` | `models.Optional[string]` | Optional | - |

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

