
# Feature Owner Price Point Type

Identifies which kind of price point a feature catalog item override applies to: `ProductPricePoint` for a product price point, `PricePoint` for a component price point. Only relevant when `price_point_id` is set.

## Enumeration

`FeatureOwnerPricePointType`

## Fields

| Name |
|  --- |
| `PRODUCTPRICEPOINT` |
| `PRICEPOINT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureOwnerPricePointType := models.FeatureOwnerPricePointType_PRODUCTPRICEPOINT

}
```

