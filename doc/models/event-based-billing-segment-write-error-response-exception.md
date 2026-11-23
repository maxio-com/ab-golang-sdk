
# Event Based Billing Segment Write Error Response Exception

Errors returned when creating or updating a single segment. Validation problems use the per-field error map; a write to a subscription's custom price point returns a `price_point` error.

## Structure

`EventBasedBillingSegmentWriteErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.EventBasedBillingSegmentWriteErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

