
# Subscription Group Create Error Response Exception

## Structure

`SubscriptionGroupCreateErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.SubscriptionGroupCreateErrorResponseErrors`](../../doc/models/containers/subscription-group-create-error-response-errors.md) | Required | This is a container for one-of cases. |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionGroupCreateErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

