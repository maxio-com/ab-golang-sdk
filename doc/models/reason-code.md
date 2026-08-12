
# Reason Code

## Structure

`ReasonCode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `SiteId` | `*int` | Optional | - |
| `Code` | `*string` | Optional | - |
| `Description` | `*string` | Optional | - |
| `Position` | `*int` | Optional | - |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    reasonCode := models.ReasonCode{
        Id:                   models.ToPointer(174),
        SiteId:               models.ToPointer(100),
        Code:                 models.ToPointer("code4"),
        Description:          models.ToPointer("description6"),
        Position:             models.ToPointer(204),
    }

}
```

