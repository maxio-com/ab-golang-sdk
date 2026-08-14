
# Proration

## Structure

`Proration`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PreservePeriod` | `*bool` | Optional | The alternative to sending preserve_period as a direct attribute to migration |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    proration := models.Proration{
        PreservePeriod:       models.ToPointer(false),
    }

}
```

