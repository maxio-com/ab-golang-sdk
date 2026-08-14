
# List Proforma Invoices Response

## Structure

`ListProformaInvoicesResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProformaInvoices` | [`[]models.ProformaInvoice`](../../doc/models/proforma-invoice.md) | Optional | - |
| `Meta` | [`*models.ListProformaInvoicesMeta`](../../doc/models/list-proforma-invoices-meta.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    listProformaInvoicesResponse := models.ListProformaInvoicesResponse{
        ProformaInvoices:     []models.ProformaInvoice{
            models.ProformaInvoice{
                Uid:                  models.ToPointer("uid0"),
                SiteId:               models.ToPointer(140),
                CustomerId:           models.NewOptional(models.ToPointer(252)),
                SubscriptionId:       models.NewOptional(models.ToPointer(68)),
                Number:               models.NewOptional(models.ToPointer(56)),
            },
            models.ProformaInvoice{
                Uid:                  models.ToPointer("uid0"),
                SiteId:               models.ToPointer(140),
                CustomerId:           models.NewOptional(models.ToPointer(252)),
                SubscriptionId:       models.NewOptional(models.ToPointer(68)),
                Number:               models.NewOptional(models.ToPointer(56)),
            },
            models.ProformaInvoice{
                Uid:                  models.ToPointer("uid0"),
                SiteId:               models.ToPointer(140),
                CustomerId:           models.NewOptional(models.ToPointer(252)),
                SubscriptionId:       models.NewOptional(models.ToPointer(68)),
                Number:               models.NewOptional(models.ToPointer(56)),
            },
        },
        Meta:                 models.ToPointer(models.ListProformaInvoicesMeta{
            TotalCount:           models.ToPointer(150),
            CurrentPage:          models.ToPointer(126),
            TotalPages:           models.ToPointer(138),
            StatusCode:           models.ToPointer(168),
        }),
    }

}
```

