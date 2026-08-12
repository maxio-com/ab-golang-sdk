
# Component Allocation Error Item

## Structure

`ComponentAllocationErrorItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ComponentId` | `*int` | Optional | - |
| `Message` | `*string` | Optional | - |
| `Kind` | `*string` | Optional | - |
| `On` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    componentAllocationErrorItem := models.ComponentAllocationErrorItem{
        ComponentId:          models.ToPointer(122),
        Message:              models.ToPointer("message8"),
        Kind:                 models.ToPointer("kind6"),
        On:                   models.ToPointer("on8"),
    }

}
```

