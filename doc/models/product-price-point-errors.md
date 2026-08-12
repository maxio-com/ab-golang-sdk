
# Product Price Point Errors

## Structure

`ProductPricePointErrors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PricePoint` | `*string` | Optional | - |
| `Interval` | `[]string` | Optional | - |
| `IntervalUnit` | `[]string` | Optional | - |
| `Name` | `[]string` | Optional | - |
| `Price` | `[]string` | Optional | - |
| `PriceInCents` | `[]string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productPricePointErrors := models.ProductPricePointErrors{
        PricePoint:           models.ToPointer("can't be blank"),
        Interval:             []string{
            "Recurring Interval: cannot be blank.",
            "Recurring Interval: must be greater than or equal to 1.",
        },
        IntervalUnit:         []string{
            "Interval unit: cannot be blank.",
            "Interval unit: must be 'month' or 'day'.",
        },
        Name:                 []string{
            "Name: cannot be blank.",
        },
        Price:                []string{
            "Price: is not a number.",
            "Price: must be greater than or equal to 0.",
        },
        PriceInCents:         []string{
            "Price in cents: cannot be blank.",
        },
    }

}
```

