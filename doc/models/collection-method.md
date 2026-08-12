
# Collection Method

The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.

## Enumeration

`CollectionMethod`

## Fields

| Name |
|  --- |
| `AUTOMATIC` |
| `REMITTANCE` |
| `PREPAID` |
| `INVOICE` |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    collectionMethod := models.CollectionMethod_AUTOMATIC

}
```

