
# Prepaid Product Price Point Filter

## Structure

`PrepaidProductPricePointFilter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProductPricePointId` | `string` | Required, Constant | Passed as a parameter to list methods to return only non null values.<br><br>**Value**: `"not_null"` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    prepaidProductPricePointFilter := models.PrepaidProductPricePointFilter{
        ProductPricePointId:  "not_null",
    }

}
```

