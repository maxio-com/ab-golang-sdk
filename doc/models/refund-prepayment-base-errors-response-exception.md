
# Refund Prepayment Base Errors Response Exception

Errors returned on creating a refund prepayment when bad request

## Structure

`RefundPrepaymentBaseErrorsResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.RefundPrepaymentBaseRefundError`](../../doc/models/refund-prepayment-base-refund-error.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.RefundPrepaymentBaseErrorsResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

