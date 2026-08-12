
# Payment Collection Method Changed

## Structure

`PaymentCollectionMethodChanged`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreviousValue` | `string` | Required | - |
| `CurrentValue` | `string` | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentCollectionMethodChanged := models.PaymentCollectionMethodChanged{
        PreviousValue:        "previous_value4",
        CurrentValue:         "current_value2",
    }

}
```

