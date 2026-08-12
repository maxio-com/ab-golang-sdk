
# Create Currency Price

## Structure

`CreateCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currency` | `*string` | Optional | ISO code for a currency defined on the site level |
| `Price` | `*float64` | Optional | Price for the price level in this currency |
| `PriceId` | `*int` | Optional | ID of the price that this corresponds with |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createCurrencyPrice := models.CreateCurrencyPrice{
        Currency:             models.ToPointer("currency2"),
        Price:                models.ToPointer(float64(54.8)),
        PriceId:              models.ToPointer(142),
    }

}
```

