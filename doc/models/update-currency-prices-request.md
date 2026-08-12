
# Update Currency Prices Request

## Structure

`UpdateCurrencyPricesRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CurrencyPrices` | [`[]models.UpdateCurrencyPrice`](../../doc/models/update-currency-price.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    updateCurrencyPricesRequest := models.UpdateCurrencyPricesRequest{
        CurrencyPrices:       []models.UpdateCurrencyPrice{
            models.UpdateCurrencyPrice{
                Id:                   50,
                Price:                float64(233.74),
            },
        },
    }

}
```

