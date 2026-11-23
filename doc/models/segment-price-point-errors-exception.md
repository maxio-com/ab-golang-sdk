
# Segment Price Point Errors Exception

Returned when a segment write targets a subscription's custom price point. Segments on a custom price point can only be set through `custom_price` on the subscription or component endpoints.

## Structure

`SegmentPricePointErrorsException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.SegmentPricePointErrorDetails`](../../doc/models/segment-price-point-error-details.md) | Required | Price-point validation error details for segment writes. |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SegmentPricePointErrorsException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

