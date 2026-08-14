
# Line Item Kind

A handle for the line item kind

## Enumeration

`LineItemKind`

## Fields

| Name |
|  --- |
| `BASELINE` |
| `INITIAL` |
| `TRIAL` |
| `QUANTITYBASEDCOMPONENT` |
| `PREPAIDUSAGECOMPONENT` |
| `ONOFFCOMPONENT` |
| `METEREDCOMPONENT` |
| `EVENTBASEDCOMPONENT` |
| `COUPON` |
| `TAX` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    lineItemKind := models.LineItemKind_PREPAIDUSAGECOMPONENT

}
```

