
# Scheduled Renewal Configuration Item Request Renewal Configuration Item

## Class Name

`ScheduledRenewalConfigurationItemRequestRenewalConfigurationItem`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ScheduledRenewalItemRequestBodyComponent`](../../../doc/models/scheduled-renewal-item-request-body-component.md) | models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent scheduledRenewalItemRequestBodyComponent) |
| [`models.ScheduledRenewalItemRequestBodyProduct`](../../../doc/models/scheduled-renewal-item-request-body-product.md) | models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyProduct(models.ScheduledRenewalItemRequestBodyProduct scheduledRenewalItemRequestBodyProduct) |

## models.ScheduledRenewalItemRequestBodyComponent

### Initialization Code

#### Example

```go
value := models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
    ItemType:             "Component",
    ItemId:               108,
})
```

## models.ScheduledRenewalItemRequestBodyProduct

### Initialization Code

#### Example

```go
value := models.ScheduledRenewalConfigurationItemRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyProduct(models.ScheduledRenewalItemRequestBodyProduct{
    ItemType:             "Product",
    ItemId:               32,
})
```

