
# Subscriptions Mrr Error Response Exception

## Structure

`SubscriptionsMrrErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.AttributeError`](../../doc/models/attribute-error.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionsMrrErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

