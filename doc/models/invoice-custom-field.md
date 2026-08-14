
# Invoice Custom Field

## Structure

`InvoiceCustomField`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OwnerId` | `*int` | Optional | - |
| `OwnerType` | [`*models.CustomFieldOwner`](../../doc/models/custom-field-owner.md) | Optional | - |
| `Name` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `Value` | `*string` | Optional | **Constraints**: *Minimum Length*: `1` |
| `MetadatumId` | `*int` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    invoiceCustomField := models.InvoiceCustomField{
        OwnerId:              models.ToPointer(14),
        OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
        Name:                 models.ToPointer("name0"),
        Value:                models.ToPointer("value2"),
        MetadatumId:          models.ToPointer(14),
    }

}
```

