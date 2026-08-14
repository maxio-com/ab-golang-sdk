
# Create Invoice Request

## Structure

`CreateInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoice` | [`models.CreateInvoice`](../../doc/models/create-invoice.md) | Required | - |

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
    createInvoiceRequest := models.CreateInvoiceRequest{
        Invoice:              models.CreateInvoice{
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
                models.CreateInvoiceItem{
                    Title:                models.ToPointer("title4"),
                    Quantity:             models.ToPointer(models.CreateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                    UnitPrice:            models.ToPointer(models.CreateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                    Taxable:              models.ToPointer(false),
                    TaxCode:              models.ToPointer("tax_code6"),
                },
            },
            IssueDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
            NetTerms:             models.ToPointer(144),
            PaymentInstructions:  models.ToPointer("payment_instructions6"),
            Memo:                 models.ToPointer("memo0"),
            Status:               models.ToPointer(models.CreateInvoiceStatus_DRAFT),
        },
    }

}
```

