
# Signup Proforma Preview

## Structure

`SignupProformaPreview`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrentProformaInvoice` | [`*models.ProformaInvoice`](../../doc/models/proforma-invoice.md) | Optional | - |
| `NextProformaInvoice` | [`*models.ProformaInvoice`](../../doc/models/proforma-invoice.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    signupProformaPreview := models.SignupProformaPreview{
        CurrentProformaInvoice: models.ToPointer(models.ProformaInvoice{
            Uid:                  models.ToPointer("uid6"),
            SiteId:               models.ToPointer(72),
            CustomerId:           models.NewOptional(models.ToPointer(184)),
            SubscriptionId:       models.NewOptional(models.ToPointer(0)),
            Number:               models.NewOptional(models.ToPointer(132)),
        }),
        NextProformaInvoice:    models.ToPointer(models.ProformaInvoice{
            Uid:                  models.ToPointer("uid8"),
            SiteId:               models.ToPointer(212),
            CustomerId:           models.NewOptional(models.ToPointer(68)),
            SubscriptionId:       models.NewOptional(models.ToPointer(140)),
            Number:               models.NewOptional(models.ToPointer(16)),
        }),
    }

}
```

