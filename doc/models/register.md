
# Register

## Structure

`Register`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `MaxioId` | `*string` | Optional | - |
| `Name` | `*string` | Optional | - |
| `CurrencyCode` | `*string` | Optional | The ISO 4217 currency code (3 character string) representing the currency of an invoice transaction. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    register := models.Register{
        Id:                   models.ToPointer(54),
        MaxioId:              models.ToPointer("maxio_id4"),
        Name:                 models.ToPointer("name2"),
        CurrencyCode:         models.ToPointer("currency_code2"),
    }

}
```

