
# Meter

The metering meter this wallet is paired with, or `null` for a wallet with no meter. Set on create and immutable afterward.

## Structure

`Meter`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `ConversionRate` | `*string` | Optional | How many wallet units one meter unit is worth, as a decimal string. Defaults to `"1.0"`. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    meter := models.Meter{
        Id:                   models.ToPointer(198),
        Name:                 models.ToPointer("name0"),
        Handle:               models.ToPointer("handle6"),
        ConversionRate:       models.ToPointer("conversion_rate2"),
    }

}
```

