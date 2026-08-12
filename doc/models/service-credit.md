
# Service Credit

## Structure

`ServiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry |
| `EndingBalanceInCents` | `*int64` | Optional | The new balance for the credit account |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `*string` | Optional | The memo attached to the entry |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    serviceCredit := models.ServiceCredit{
        Id:                   models.ToPointer(132),
        AmountInCents:        models.ToPointer(int64(218)),
        EndingBalanceInCents: models.ToPointer(int64(2)),
        EntryType:            models.ToPointer(models.ServiceCreditType_CREDIT),
        Memo:                 models.ToPointer("memo8"),
    }

}
```

