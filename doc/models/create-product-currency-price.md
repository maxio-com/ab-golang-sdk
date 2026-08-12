
# Create Product Currency Price

## Structure

`CreateProductCurrencyPrice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Currency` | `string` | Required | ISO code for one of the site level currencies. |
| `Price` | `int` | Required | Price for the given role. |
| `Role` | [`models.CurrencyPriceRole`](../../doc/models/currency-price-role.md) | Required | Role for the price. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    createProductCurrencyPrice := models.CreateProductCurrencyPrice{
        Currency:             "currency2",
        Price:                78,
        Role:                 models.CurrencyPriceRole_BASELINE,
    }

}
```

