
# Update Invoice

Attributes of a draft ad hoc invoice which can be updated. Only the submitted attributes are changed.

## Structure

`UpdateInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LineItems` | [`[]models.UpdateInvoiceItem`](../../doc/models/update-invoice-item.md) | Optional | Line item changes to apply. Line items without a `uid` are added, line items with a `uid` are updated, and line items with a `uid` and `_destroy` set to `true` are removed. Existing line items not referenced in the array remain unchanged. |
| `IssueDate` | `*time.Time` | Optional | New issue date for the invoice (format YYYY-MM-DD). This date is interpreted and validated in your site's time zone. It must be today or a date in the past — future dates are not accepted. The due date is recalculated from the issue date and net terms. |
| `NetTerms` | `*int` | Optional | Number of days after the issue date on which the invoice is due. The due date is recalculated when net terms or the issue date change. |
| `PaymentInstructions` | `*string` | Optional | Custom payment instructions displayed on the invoice. |
| `Memo` | `*string` | Optional | A custom memo displayed on the invoice. |
| `SellerAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Replaces the seller address on the invoice |
| `BillingAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Replaces the billing address on the invoice |
| `ShippingAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Replaces the shipping address on the invoice |
| `Coupons` | [`[]models.CreateInvoiceCoupon`](../../doc/models/create-invoice-coupon.md) | Optional | When present, replaces all discounts currently applied to the invoice. Send an empty array to remove all discounts. |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    updateInvoice := models.UpdateInvoice{
        LineItems:            []models.UpdateInvoiceItem{
            models.UpdateInvoiceItem{
                Title:                models.ToPointer("title4"),
                Quantity:             models.ToPointer(models.UpdateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                UnitPrice:            models.ToPointer(models.UpdateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                Taxable:              models.ToPointer(false),
                TaxCode:              models.ToPointer("tax_code6"),
            },
        },
        IssueDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        NetTerms:             models.ToPointer(46),
        PaymentInstructions:  models.ToPointer("payment_instructions6"),
        Memo:                 models.ToPointer("memo2"),
    }

}
```

