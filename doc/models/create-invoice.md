
# Create Invoice

## Structure

`CreateInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LineItems` | [`[]models.CreateInvoiceItem`](create-invoice-item.md) | Optional | - |
| `IssueDate` | `*string` | Optional | - |
| `NetTerms` | `*int` | Optional | By default, invoices will be created with a due date matching the date of invoice creation. If a different due date is desired, the net_terms parameter can be sent indicating the number of days in advance the due date should be. |
| `PaymentInstructions` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | A custom memo can be sent to override the site's default. |
| `SellerAddress` | [`*models.CreateInvoiceAddress`](create-invoice-address.md) | Optional | Overrides the defaults for the site |
| `BillingAddress` | [`*models.CreateInvoiceAddress`](create-invoice-address.md) | Optional | Overrides the default for the customer |
| `ShippingAddress` | [`*models.CreateInvoiceAddress`](create-invoice-address.md) | Optional | Overrides the default for the customer |
| `Coupons` | [`[]models.CreateInvoiceCoupon`](create-invoice-coupon.md) | Optional | - |
| `Status` | [`*models.Status1Enum`](status-1-enum.md) | Optional | **Default**: `"open"` |

## Example (as JSON)

```json
{
  "status": "draft",
  "line_items": [
    {
      "title": "title4",
      "quantity": {
        "key1": "val1",
        "key2": "val2"
      },
      "unit_price": {
        "key1": "val1",
        "key2": "val2"
      },
      "taxable": false,
      "tax_code": "tax_code6"
    }
  ],
  "issue_date": "issue_date2",
  "net_terms": 18,
  "payment_instructions": "payment_instructions0",
  "memo": "memo6"
}
```

