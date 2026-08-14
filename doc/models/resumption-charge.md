
# Resumption Charge

(For calendar billing subscriptions only) The way that the resumed subscription's charge should be handled

## Enumeration

`ResumptionCharge`

## Fields

| Name |
|  --- |
| `PRORATED` |
| `IMMEDIATE` |
| `DELAYED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    resumptionCharge := models.ResumptionCharge_IMMEDIATE

}
```

