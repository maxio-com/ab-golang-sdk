
# Billing Manifest Line Item Kind

A handle for the billing manifest line item kind

## Enumeration

`BillingManifestLineItemKind`

## Fields

| Name |
|  --- |
| `BASELINE` |
| `INITIAL` |
| `TRIAL` |
| `COUPON` |
| `COMPONENT` |
| `TAX` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    billingManifestLineItemKind := models.BillingManifestLineItemKind_COMPONENT

}
```

