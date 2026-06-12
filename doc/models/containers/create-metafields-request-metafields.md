
# Create Metafields Request Metafields

## Class Name

`CreateMetafieldsRequestMetafields`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.CreateMetafield`](../../../doc/models/create-metafield.md) | models.CreateMetafieldsRequestMetafieldsContainer.FromCreateMetafield(models.CreateMetafield createMetafield) |
| [`[]models.CreateMetafield`](../../../doc/models/create-metafield.md) | models.CreateMetafieldsRequestMetafieldsContainer.FromArrayOfCreateMetafield([]models.CreateMetafield arrayOfCreateMetafield) |

## models.CreateMetafield

### Initialization Code

#### Example

```go
value := models.CreateMetafieldsRequestMetafieldsContainer.FromCreateMetafield(models.CreateMetafield{
    Name:                 models.ToPointer("my_field"),
    Scope:                models.ToPointer(models.MetafieldScope{
        Csv:                  models.ToPointer(models.IncludeOption_EXCLUDE),
        Invoices:             models.ToPointer(models.IncludeOption_EXCLUDE),
        Statements:           models.ToPointer(models.IncludeOption_EXCLUDE),
        Portal:               models.ToPointer(models.IncludeOption_EXCLUDE),
        PublicShow:           models.ToPointer(models.IncludeOption_EXCLUDE),
        PublicEdit:           models.ToPointer(models.IncludeOption_EXCLUDE),
    }),
    InputType:            models.ToPointer(models.MetafieldInput_TEXT),
    Enum:                 []string{
        "string",
    },
})
```

## []models.CreateMetafield

### Initialization Code

#### Example

```go
value := models.CreateMetafieldsRequestMetafieldsContainer.FromArrayOfCreateMetafield([]models.CreateMetafield{
    models.CreateMetafield{
    },
})
```

