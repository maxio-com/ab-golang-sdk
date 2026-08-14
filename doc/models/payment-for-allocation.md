
# Payment for Allocation

Information for captured payment, if applicable

## Structure

`PaymentForAllocation`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | - |
| `Success` | `*bool` | Optional | - |
| `Memo` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    paymentForAllocation := models.PaymentForAllocation{
        Id:                   models.ToPointer(68),
        AmountInCents:        models.ToPointer(int64(102)),
        Success:              models.ToPointer(false),
        Memo:                 models.ToPointer("memo6"),
    }

}
```

