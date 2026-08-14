
# Sale Rep Item Mrr

## Structure

`SaleRepItemMrr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mrr` | `*string` | Optional | - |
| `Usage` | `*string` | Optional | - |
| `Recurring` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    saleRepItemMrr := models.SaleRepItemMrr{
        Mrr:                  models.ToPointer("mrr8"),
        Usage:                models.ToPointer("usage0"),
        Recurring:            models.ToPointer("recurring6"),
    }

}
```

