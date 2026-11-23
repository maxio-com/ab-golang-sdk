
# Event Based Billing List Segments Errors Exception

## Structure

`EventBasedBillingListSegmentsErrorsException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.Errors`](../../doc/models/errors.md) | Optional | Validation error details for listing segments. |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.EventBasedBillingListSegmentsErrorsException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

