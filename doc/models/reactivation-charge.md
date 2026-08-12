
# Reactivation Charge

You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal.

## Enumeration

`ReactivationCharge`

## Fields

| Name |
|  --- |
| `PRORATED` |
| `IMMEDIATE` |
| `DELAYED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reactivationCharge := models.ReactivationCharge_IMMEDIATE

}
```

