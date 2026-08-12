
# Line Item Transaction Type

A handle for the line item transaction type

## Enumeration

`LineItemTransactionType`

## Fields

| Name |
|  --- |
| `CHARGE` |
| `CREDIT` |
| `ADJUSTMENT` |
| `PAYMENT` |
| `REFUND` |
| `INFOTRANSACTION` |
| `PAYMENTAUTHORIZATION` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    lineItemTransactionType := models.LineItemTransactionType_PAYMENT

}
```

