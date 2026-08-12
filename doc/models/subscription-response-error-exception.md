
# Subscription Response Error Exception

## Structure

`SubscriptionResponseErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | [`*models.Subscription`](../../doc/models/subscription.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionResponseErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

