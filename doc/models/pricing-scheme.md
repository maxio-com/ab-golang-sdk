
# Pricing Scheme

The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.

## Enumeration

`PricingScheme`

## Fields

| Name |
|  --- |
| `STAIRSTEP` |
| `VOLUME` |
| `PERUNIT` |
| `TIERED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    pricingScheme := models.PricingScheme_STAIRSTEP

}
```

