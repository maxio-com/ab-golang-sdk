
# Customer Custom Fields Change

## Structure

`CustomerCustomFieldsChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Before` | [`[]models.InvoiceCustomField`](../../doc/models/invoice-custom-field.md) | Required | - |
| `After` | [`[]models.InvoiceCustomField`](../../doc/models/invoice-custom-field.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerCustomFieldsChange := models.CustomerCustomFieldsChange{
        Before:               []models.InvoiceCustomField{
            models.InvoiceCustomField{
                OwnerId:              models.ToPointer(26),
                OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
                Name:                 models.ToPointer("name0"),
                Value:                models.ToPointer("value2"),
                MetadatumId:          models.ToPointer(26),
            },
        },
        After:                []models.InvoiceCustomField{
            models.InvoiceCustomField{
                OwnerId:              models.ToPointer(130),
                OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
                Name:                 models.ToPointer("name2"),
                Value:                models.ToPointer("value4"),
                MetadatumId:          models.ToPointer(130),
            },
        },
    }

}
```

