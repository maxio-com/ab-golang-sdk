
# Deduct Service Credit Request

## Structure

`DeductServiceCreditRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Deduction` | [`models.DeductServiceCredit`](../../doc/models/deduct-service-credit.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    deductServiceCreditRequest := models.DeductServiceCreditRequest{
        Deduction:            models.DeductServiceCredit{
            Amount:               models.DeductServiceCreditAmountContainer.FromString("String9"),
            Memo:                 models.ToPointer("memo0"),
        },
    }

}
```

