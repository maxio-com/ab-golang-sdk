
# Feature Kind

The behavior of a feature:

- `access_right`: a boolean entitlement. A subscriber either has access or does not.
- `usage_limit`: a quantified allowance measured over a recurring period (for example, "10,000 API calls per month").
- `service_right`: a free-form value (text, boolean, or number) that isn't a simple access flag or a metered limit.

## Enumeration

`FeatureKind`

## Fields

| Name |
|  --- |
| `ACCESSRIGHT` |
| `USAGELIMIT` |
| `SERVICERIGHT` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    featureKind := models.FeatureKind_USAGELIMIT

}
```

