
# Contract

Contract linked to the scheduled renewal configuration.

## Structure

`Contract`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `MaxioId` | `*string` | Optional | - |
| `Number` | `models.Optional[string]` | Optional | - |
| `Register` | [`*models.Register`](../../doc/models/register.md) | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    contract := models.Contract{
        Id:                   models.ToPointer(112),
        MaxioId:              models.ToPointer("maxio_id6"),
        Number:               models.NewOptional(models.ToPointer("number2")),
        Register:             models.ToPointer(models.Register{
            Id:                   models.ToPointer(54),
            MaxioId:              models.ToPointer("maxio_id4"),
            Name:                 models.ToPointer("name2"),
            CurrencyCode:         models.ToPointer("currency_code2"),
        }),
    }

}
```

