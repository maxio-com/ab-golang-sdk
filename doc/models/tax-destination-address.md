
# Tax Destination Address

## Enumeration

`TaxDestinationAddress`

## Fields

| Name |
|  --- |
| `SHIPPINGTHENBILLING` |
| `BILLINGTHENSHIPPING` |
| `SHIPPINGONLY` |
| `BILLINGONLY` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    taxDestinationAddress := models.TaxDestinationAddress_SHIPPINGONLY

}
```

