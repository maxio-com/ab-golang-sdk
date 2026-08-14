
# Subscription Component Allocation Error Exception

## Structure

`SubscriptionComponentAllocationErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`[]models.SubscriptionComponentAllocationErrorItem`](../../doc/models/subscription-component-allocation-error-item.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionComponentAllocationErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

