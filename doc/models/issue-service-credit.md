
# Issue Service Credit

## Structure

`IssueServiceCredit`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Amount` | [`models.IssueServiceCreditAmount`](../../doc/models/containers/issue-service-credit-amount.md) | Required | This is a container for one-of cases. |
| `Memo` | `*string` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    issueServiceCredit := models.IssueServiceCredit{
        Amount:               models.IssueServiceCreditAmountContainer.FromPrecision(float64(44.88)),
        Memo:                 models.ToPointer("memo6"),
    }

}
```

