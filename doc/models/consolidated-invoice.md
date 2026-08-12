
# Consolidated Invoice

## Structure

`ConsolidatedInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Invoices` | [`[]models.Invoice`](../../doc/models/invoice.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    consolidatedInvoice := models.ConsolidatedInvoice{
        Invoices:             []models.Invoice{
            models.Invoice{
                Id:                         models.ToPointer(int64(196)),
                Uid:                        models.ToPointer("uid6"),
                SiteId:                     models.ToPointer(122),
                CustomerId:                 models.ToPointer(234),
                SubscriptionId:             models.ToPointer(50),
            },
            models.Invoice{
                Id:                         models.ToPointer(int64(196)),
                Uid:                        models.ToPointer("uid6"),
                SiteId:                     models.ToPointer(122),
                CustomerId:                 models.ToPointer(234),
                SubscriptionId:             models.ToPointer(50),
            },
            models.Invoice{
                Id:                         models.ToPointer(int64(196)),
                Uid:                        models.ToPointer("uid6"),
                SiteId:                     models.ToPointer(122),
                CustomerId:                 models.ToPointer(234),
                SubscriptionId:             models.ToPointer(50),
            },
        },
    }

}
```

