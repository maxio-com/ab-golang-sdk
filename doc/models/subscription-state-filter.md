
# Subscription State Filter

Allowed values for filtering by the current state of the subscription.

## Enumeration

`SubscriptionStateFilter`

## Fields

| Name |
|  --- |
| `ACTIVE` |
| `CANCELED` |
| `EXPIRED` |
| `EXPIREDCARDS` |
| `ENUMEXPIREDCARDSLIVESUBSCRIPTIONS` |
| `ENUMEXPIREDCARDSALLSUBSCRIPTIONS` |
| `ONHOLD` |
| `AWAITINGSIGNUP` |
| `AWAITINGSIGNUPDATE` |
| `PASTDUE` |
| `PENDINGCANCELLATION` |
| `PENDINGRENEWAL` |
| `PREPAIDDUNNING` |
| `SUSPENDED` |
| `TRIALENDED` |
| `TRIALING` |
| `UNPAID` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    subscriptionStateFilter := models.SubscriptionStateFilter_AWAITINGSIGNUP

}
```

