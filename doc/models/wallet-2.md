
# Wallet 2

## Structure

`Wallet2`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `models.Optional[string]` | Optional | A merchant-facing label for the wallet. Surrounding spaces are stripped; a blank or `null` value clears it.<br><br>**Constraints**: *Maximum Length*: `255` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    wallet2 := models.Wallet2{
        Name:                 models.NewOptional(models.ToPointer("name2")),
    }

}
```

