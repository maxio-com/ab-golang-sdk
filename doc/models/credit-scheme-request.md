
# Credit Scheme Request

## Structure

`CreditSchemeRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreditScheme` | [`models.CreditScheme`](../../doc/models/credit-scheme.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    creditSchemeRequest := models.CreditSchemeRequest{
        CreditScheme:         models.CreditScheme_CREDIT,
    }

}
```

