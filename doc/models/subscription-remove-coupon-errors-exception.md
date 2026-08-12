
# Subscription Remove Coupon Errors Exception

## Structure

`SubscriptionRemoveCouponErrorsException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Subscription` | `[]string` | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionRemoveCouponErrorsException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

