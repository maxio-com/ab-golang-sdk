
# Update Invoice Request

Request payload for updating a draft ad hoc invoice.

## Structure

`UpdateInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoice` | [`models.UpdateInvoice`](../../doc/models/update-invoice.md) | Required | Attributes of a draft ad hoc invoice which can be updated. Only the submitted attributes are changed. |

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
    updateInvoiceRequest := models.UpdateInvoiceRequest{
        Invoice:              models.UpdateInvoice{
            LineItems:            []models.UpdateInvoiceItem{
                models.UpdateInvoiceItem{
                    Title:                models.ToPointer("title4"),
                    Quantity:             models.ToPointer(models.UpdateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                    UnitPrice:            models.ToPointer(models.UpdateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                    Taxable:              models.ToPointer(false),
                    TaxCode:              models.ToPointer("tax_code6"),
                },
                models.UpdateInvoiceItem{
                    Title:                models.ToPointer("title4"),
                    Quantity:             models.ToPointer(models.UpdateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                    UnitPrice:            models.ToPointer(models.UpdateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                    Taxable:              models.ToPointer(false),
                    TaxCode:              models.ToPointer("tax_code6"),
                },
                models.UpdateInvoiceItem{
                    Title:                models.ToPointer("title4"),
                    Quantity:             models.ToPointer(models.UpdateInvoiceItemQuantityContainer.FromPrecision(float64(56.68))),
                    UnitPrice:            models.ToPointer(models.UpdateInvoiceItemUnitPriceContainer.FromPrecision(float64(39.9))),
                    Taxable:              models.ToPointer(false),
                    TaxCode:              models.ToPointer("tax_code6"),
                },
            },
            IssueDate:            models.ToPointer(parseTime(models.DEFAULT_DATE, "2024-01-01", func(err error) { log.Fatalln(err) })),
            NetTerms:             models.ToPointer(144),
            PaymentInstructions:  models.ToPointer("payment_instructions6"),
            Memo:                 models.ToPointer("memo0"),
        },
    }

}
```

