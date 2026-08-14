
# Deduct Service Credit

## Structure

`DeductServiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | [`models.DeductServiceCreditAmount`](../../doc/models/containers/deduct-service-credit-amount.md) | Required | This is a container for one-of cases. |
| `Memo` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    deductServiceCredit := models.DeductServiceCredit{
        Amount:               models.DeductServiceCreditAmountContainer.FromString("String5"),
        Memo:                 models.ToPointer("memo6"),
    }

}
```

