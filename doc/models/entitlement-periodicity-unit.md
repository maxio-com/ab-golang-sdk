
# Entitlement Periodicity Unit

The recurring window over which a `usage_limit` feature's allowance resets.

## Enumeration

`EntitlementPeriodicityUnit`

## Fields

| Name |
|  --- |
| `HOUR` |
| `DAY` |
| `WEEK` |
| `MONTH` |
| `YEAR` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    entitlementPeriodicityUnit := models.EntitlementPeriodicityUnit_MONTH

}
```

