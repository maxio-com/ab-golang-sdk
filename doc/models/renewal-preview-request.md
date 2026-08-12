
# Renewal Preview Request

## Structure

`RenewalPreviewRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Components` | [`[]models.RenewalPreviewComponent`](../../doc/models/renewal-preview-component.md) | Optional | (Optional) Array of component definitions to preview. Providing any component definitions here will override the actual components on the subscription (and their quantities), and the billing preview will contain only these components (in addition to any product base fees). |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    renewalPreviewRequest := models.RenewalPreviewRequest{
        Components:           []models.RenewalPreviewComponent{
            models.RenewalPreviewComponent{
                ComponentId:          models.ToPointer(models.RenewalPreviewComponentComponentIdContainer.FromString("String5")),
                Quantity:             models.ToPointer(210),
                PricePointId:         models.ToPointer(models.RenewalPreviewComponentPricePointIdContainer.FromString("String3")),
            },
        },
    }

}
```

