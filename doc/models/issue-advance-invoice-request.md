
# Issue Advance Invoice Request

## Structure

`IssueAdvanceInvoiceRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Force` | `*bool` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    issueAdvanceInvoiceRequest := models.IssueAdvanceInvoiceRequest{
        Force:                models.ToPointer(false),
    }

}
```

