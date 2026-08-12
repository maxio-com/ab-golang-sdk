
# Component Currency Price

## Structure

`ComponentCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
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
        Id:                   models.ToPointer(128),
        Currency:             models.ToPointer("currency2"),
        Price:                models.ToPointer("price4"),
        FormattedPrice:       models.ToPointer("formatted_price6"),
        PriceId:              models.ToPointer(38),
    }

}
```

