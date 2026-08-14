
# Product Family

## Structure

`ProductFamily`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `Handle` | `*string` | Optional | - |
| `AccountingCode` | `models.Optional[string]` | Optional | - |
| `Description` | `models.Optional[string]` | Optional | - |
| `Surcharging` | `*bool` | Optional | Whether surcharging applies to this product family. Only included on sites where surcharging is enabled. |
| `CreatedAt` | `*time.Time` | Optional | - |
| `UpdatedAt` | `*time.Time` | Optional | - |
| `ArchivedAt` | `models.Optional[time.Time]` | Optional | Timestamp indicating when this product family was archived. `null` if the product family is not archived. |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    productFamily := models.ProductFamily{
        Id:                   models.ToPointer(134),
        Name:                 models.ToPointer("name4"),
        Handle:               models.ToPointer("handle0"),
        AccountingCode:       models.NewOptional(models.ToPointer("accounting_code0")),
        Description:          models.NewOptional(models.ToPointer("description4")),
    }

}
```

