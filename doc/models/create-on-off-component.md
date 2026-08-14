
# Create On Off Component

## Structure

`CreateOnOffComponent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OnOffComponent` | [`models.OnOffComponent`](../../doc/models/on-off-component.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createOnOffComponent := models.CreateOnOffComponent{
        OnOffComponent:       models.OnOffComponent{
            Name:                      "name6",
            Description:               models.ToPointer("description6"),
            Handle:                    models.ToPointer("handle2"),
            Taxable:                   models.ToPointer(false),
            UpgradeCharge:             models.NewOptional(models.ToPointer(models.CreditType_FULL)),
            DowngradeCredit:           models.NewOptional(models.ToPointer(models.CreditType_FULL)),
            UnitPrice:                 models.OnOffComponentUnitPriceContainer.FromString("String5"),
        },
    }

}
```

