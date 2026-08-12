
# Cancellation Method

The process used to cancel the subscription, if the subscription has been canceled. It is nil if the subscription's state is not canceled.

## Enumeration

`CancellationMethod`

## Fields

| Name |
|  --- |
| `MERCHANTUI` |
| `MERCHANTAPI` |
| `DUNNING` |
| `BILLINGPORTAL` |
| `UNKNOWN` |
| `IMPORTED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    cancellationMethod := models.CancellationMethod_MERCHANTUI

}
```

