
# Subscription Group Signup Error Response Exception

## Structure

`SubscriptionGroupSignupErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.SubscriptionGroupSignupError`](../../doc/models/subscription-group-signup-error.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionGroupSignupErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

