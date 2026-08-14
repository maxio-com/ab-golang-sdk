
# Trial Type

Indicates how a trial is handled when the trial period ends and there is no credit card on file. For `no_obligation`, the subscription transitions to a Trial Ended state. Maxio will not send any emails or statements. For `payment_expected`, the subscription transitions to a Past Due state. Maxio will send normal dunning emails and statements according to your other settings.

## Enumeration

`TrialType`

## Fields

| Name |
|  --- |
| `NOOBLIGATION` |
| `PAYMENTEXPECTED` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    trialType := models.TrialType_NOOBLIGATION

}
```

