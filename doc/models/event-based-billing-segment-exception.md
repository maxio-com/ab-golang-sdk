
# Event Based Billing Segment Exception

## Structure

`EventBasedBillingSegmentException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.EventBasedBillingSegmentError`](../../doc/models/event-based-billing-segment-error.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.EventBasedBillingSegmentException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

