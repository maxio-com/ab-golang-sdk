
# Customer Changes Preview Response

## Structure

`CustomerChangesPreviewResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Changes` | [`models.CustomerChange`](../../doc/models/customer-change.md) | Required | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    customerChangesPreviewResponse := models.CustomerChangesPreviewResponse{
        Changes:              models.CustomerChange{
            Payer:                models.NewOptional(models.ToPointer(models.CustomerPayerChange{
                Before:               models.InvoicePayerChange{
                    FirstName:            models.ToPointer("first_name0"),
                    LastName:             models.ToPointer("last_name8"),
                    Organization:         models.ToPointer("organization4"),
                    Email:                models.ToPointer("email6"),
                },
                After:                models.InvoicePayerChange{
                    FirstName:            models.ToPointer("first_name2"),
                    LastName:             models.ToPointer("last_name0"),
                    Organization:         models.ToPointer("organization4"),
                    Email:                models.ToPointer("email4"),
                },
            })),
            ShippingAddress:      models.NewOptional(models.ToPointer(models.AddressChange{
                Before:               models.InvoiceAddress{
                    Street:               models.NewOptional(models.ToPointer("street0")),
                    Line2:                models.NewOptional(models.ToPointer("line24")),
                    City:                 models.NewOptional(models.ToPointer("city0")),
                    State:                models.NewOptional(models.ToPointer("state6")),
                    Zip:                  models.NewOptional(models.ToPointer("zip4")),
                },
                After:                models.InvoiceAddress{
                    Street:               models.NewOptional(models.ToPointer("street2")),
                    Line2:                models.NewOptional(models.ToPointer("line26")),
                    City:                 models.NewOptional(models.ToPointer("city8")),
                    State:                models.NewOptional(models.ToPointer("state2")),
                    Zip:                  models.NewOptional(models.ToPointer("zip4")),
                },
            })),
            BillingAddress:       models.NewOptional(models.ToPointer(models.AddressChange{
                Before:               models.InvoiceAddress{
                    Street:               models.NewOptional(models.ToPointer("street0")),
                    Line2:                models.NewOptional(models.ToPointer("line24")),
                    City:                 models.NewOptional(models.ToPointer("city0")),
                    State:                models.NewOptional(models.ToPointer("state6")),
                    Zip:                  models.NewOptional(models.ToPointer("zip4")),
                },
                After:                models.InvoiceAddress{
                    Street:               models.NewOptional(models.ToPointer("street2")),
                    Line2:                models.NewOptional(models.ToPointer("line26")),
                    City:                 models.NewOptional(models.ToPointer("city8")),
                    State:                models.NewOptional(models.ToPointer("state2")),
                    Zip:                  models.NewOptional(models.ToPointer("zip4")),
                },
            })),
            CustomFields:         models.NewOptional(models.ToPointer(models.CustomerCustomFieldsChange{
                Before:               []models.InvoiceCustomField{
                    models.InvoiceCustomField{
                        OwnerId:              models.ToPointer(26),
                        OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
                        Name:                 models.ToPointer("name0"),
                        Value:                models.ToPointer("value2"),
                        MetadatumId:          models.ToPointer(26),
                    },
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
                    models.InvoiceCustomField{
                        OwnerId:              models.ToPointer(130),
                        OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
                        Name:                 models.ToPointer("name2"),
                        Value:                models.ToPointer("value4"),
                        MetadatumId:          models.ToPointer(130),
                    },
                    models.InvoiceCustomField{
                        OwnerId:              models.ToPointer(130),
                        OwnerType:            models.ToPointer(models.CustomFieldOwner_CUSTOMER),
                        Name:                 models.ToPointer("name2"),
                        Value:                models.ToPointer("value4"),
                        MetadatumId:          models.ToPointer(130),
                    },
                },
            })),
        },
    }

}
```

