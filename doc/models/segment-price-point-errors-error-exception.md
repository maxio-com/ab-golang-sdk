
# Segment Price Point Errors Error Exception

Returned when a segment write targets a subscription's custom price point. Segments on a custom price point can only be set through `custom_price` on the subscription or component endpoints.

## Structure

`SegmentPricePointErrorsErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.Errors`](../../doc/models/errors.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SegmentPricePointErrorsErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

