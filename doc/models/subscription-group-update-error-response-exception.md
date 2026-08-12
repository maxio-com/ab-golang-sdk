
# Subscription Group Update Error Response Exception

## Structure

`SubscriptionGroupUpdateErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.SubscriptionGroupUpdateError`](../../doc/models/subscription-group-update-error.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionGroupUpdateErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

