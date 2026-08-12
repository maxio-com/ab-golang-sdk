
# Credit Card Attributes

## Structure

`CreditCardAttributes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `FullNumber` | `*string` | Optional | - |
| `ExpirationMonth` | `*string` | Optional | - |
| `ExpirationYear` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    creditCardAttributes := models.CreditCardAttributes{
        FullNumber:           models.ToPointer("full_number8"),
        ExpirationMonth:      models.ToPointer("expiration_month8"),
        ExpirationYear:       models.ToPointer("expiration_year2"),
    }

}
```

