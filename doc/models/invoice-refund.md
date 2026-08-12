
# Invoice Refund

## Structure

`InvoiceRefund`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TransactionId` | `*int` | Optional | - |
| `PaymentId` | `*int` | Optional | - |
| `Memo` | `*string` | Optional | - |
| `OriginalAmount` | `*string` | Optional | - |
| `AppliedAmount` | `*string` | Optional | - |
| `GatewayTransactionId` | `models.Optional[string]` | Optional | The transaction ID for the refund as returned from the payment gateway |
| `GatewayUsed` | `*string` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `AchLateReject` | `models.Optional[bool]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceRefund := models.InvoiceRefund{
        TransactionId:        models.ToPointer(166),
        PaymentId:            models.ToPointer(36),
        Memo:                 models.ToPointer("memo6"),
        OriginalAmount:       models.ToPointer("original_amount6"),
        AppliedAmount:        models.ToPointer("applied_amount6"),
    }

}
```

