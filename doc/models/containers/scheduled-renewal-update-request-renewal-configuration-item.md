
# Scheduled Renewal Update Request Renewal Configuration Item

## Class Name

`ScheduledRenewalUpdateRequestRenewalConfigurationItem`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.ScheduledRenewalItemRequestBodyComponent`](../../../doc/models/scheduled-renewal-item-request-body-component.md) | models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent scheduledRenewalItemRequestBodyComponent) |
| [`models.ScheduledRenewalItemRequestBodyProduct`](../../../doc/models/scheduled-renewal-item-request-body-product.md) | models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyProduct(models.ScheduledRenewalItemRequestBodyProduct scheduledRenewalItemRequestBodyProduct) |

## models.ScheduledRenewalItemRequestBodyComponent

### Initialization Code

#### Example

```go
value := models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyComponent(models.ScheduledRenewalItemRequestBodyComponent{
    ItemType:             "Component",
    ItemId:               108,
})
```

## models.ScheduledRenewalItemRequestBodyProduct

### Initialization Code

#### Example

```go
value := models.ScheduledRenewalUpdateRequestRenewalConfigurationItemContainer.FromScheduledRenewalItemRequestBodyProduct(models.ScheduledRenewalItemRequestBodyProduct{
    ItemType:             "Product",
    ItemId:               32,
})
```

