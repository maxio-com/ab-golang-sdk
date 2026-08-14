
# Void Invoice

## Structure

`VoidInvoice`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reason` | `string` | Required | **Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    voidInvoice := models.VoidInvoice{
        Reason:               "reason8",
    }

}
```

