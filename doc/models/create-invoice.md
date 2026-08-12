
# Create Invoice

## Structure

`CreateInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `LineItems` | [`[]models.CreateInvoiceItem`](../../doc/models/create-invoice-item.md) | Optional | - |
| `IssueDate` | `*time.Time` | Optional | Date on which the invoice will be issued (format YYYY-MM-DD). This date is interpreted and validated in your site's time zone. It must be today or a date in the past — future dates are not accepted. If omitted, defaults to today in your site's time zone. |
| `NetTerms` | `*int` | Optional | By default, invoices will be created with a due date matching the date of invoice creation. If a different due date is desired, the net_terms parameter can be sent indicating the number of days in advance the due date should be. |
| `PaymentInstructions` | `*string` | Optional | - |
| `Memo` | `*string` | Optional | A custom memo can be sent to override the site's default. |
| `SellerAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Overrides the defaults for the site. |
| `BillingAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Overrides the default for the customer. |
| `ShippingAddress` | [`*models.CreateInvoiceAddress`](../../doc/models/create-invoice-address.md) | Optional | Overrides the default for the customer. |
| `Coupons` | [`[]models.CreateInvoiceCoupon`](../../doc/models/create-invoice-coupon.md) | Optional | - |
| `Status` | [`*models.CreateInvoiceStatus`](../../doc/models/create-invoice-status.md) | Optional | **Default**: `"open"` |

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
    createInvoice := models.CreateInvoice{
        LineItems:            []models.CreateInvoiceItem{
            models.CreateInvoiceItem{
                Title:                models.ToPointer("title4"),
                Quantity:             models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                UnitPrice:            models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                Taxable:              models.ToPointer(false),
                TaxCode:              models.ToPointer("tax_code6"),
            },
            models.CreateInvoiceItem{
                Title:                models.ToPointer("title4"),
                Quantity:             models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                UnitPrice:            models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                Taxable:              models.ToPointer(false),
                TaxCode:              models.ToPointer("tax_code6"),
            },
        },
        IssueDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
        NetTerms:             models.ToPointer(202),
        PaymentInstructions:  models.ToPointer("payment_instructions2"),
        Memo:                 models.ToPointer("memo8"),
        Status:               models.ToPointer(models.CreateInvoiceStatus_DRAFT),
    }

}
```

