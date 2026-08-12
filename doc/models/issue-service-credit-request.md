
# Issue Service Credit Request

## Structure

`IssueServiceCreditRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ServiceCredit` | [`models.IssueServiceCredit`](../../doc/models/issue-service-credit.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    issueServiceCreditRequest := models.IssueServiceCreditRequest{
        ServiceCredit:        models.IssueServiceCredit{
            Amount:               models.IssueServiceCreditAmountContainer.FromPrecision(float64(31.42)),
            Memo:                 models.ToPointer("memo0"),
        },
    }

}
```

