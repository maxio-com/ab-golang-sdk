
# MRR Movement

## Structure

`MRRMovement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | `*int` | Optional | - |
| `Category` | `*string` | Optional | - |
| `SubscriberDelta` | `*int` | Optional | - |
| `LeadDelta` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    mrrMovement := models.MRRMovement{
        Amount:               models.ToPointer(10),
        Category:             models.ToPointer("category0"),
        SubscriberDelta:      models.ToPointer(206),
        LeadDelta:            models.ToPointer(234),
    }

}
```

