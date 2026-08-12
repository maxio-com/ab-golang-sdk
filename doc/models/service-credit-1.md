
# Service Credit 1

## Structure

`ServiceCredit1`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `AmountInCents` | `*int64` | Optional | The amount in cents of the entry |
| `EndingBalanceInCents` | `*int64` | Optional | The new balance for the credit account |
| `EntryType` | [`*models.ServiceCreditType`](../../doc/models/service-credit-type.md) | Optional | The type of entry |
| `Memo` | `*string` | Optional | The memo attached to the entry |
| `InvoiceUid` | `models.Optional[string]` | Optional | The invoice uid associated with the entry. Only present for debit entries. |
| `RemainingBalanceInCents` | `*int64` | Optional | The remaining balance for the entry |
| `CreatedAt` | `*time.Time` | Optional | The date and time the entry was created |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    serviceCredit1 := models.ServiceCredit1{
        Id:                      models.ToPointer(12),
        AmountInCents:           models.ToPointer(int64(158)),
        EndingBalanceInCents:    models.ToPointer(int64(138)),
        EntryType:               models.ToPointer(models.ServiceCreditType_CREDIT),
        Memo:                    models.ToPointer("memo4"),
    }

}
```

