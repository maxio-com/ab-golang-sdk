
# Component Currency Price

## Structure

`ComponentCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `models.Optional[int]` | Optional | The ID of the stored currency price. This is `null` when the price point uses the site exchange rate (`use_site_exchange_rate: true`), because the price is calculated from the rate and is not stored. |
| `Currency` | `*string` | Optional | - |
| `Price` | `*string` | Optional | - |
| `FormattedPrice` | `*string` | Optional | - |
| `PriceId` | `*int` | Optional | - |
| `PricePointId` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentCurrencyPrice := models.ComponentCurrencyPrice{
        Id:                   models.NewOptional(models.ToPointer(128)),
        Currency:             models.ToPointer("currency2"),
        Price:                models.ToPointer("price4"),
        FormattedPrice:       models.ToPointer("formatted_price6"),
        PriceId:              models.ToPointer(38),
    }

}
```

