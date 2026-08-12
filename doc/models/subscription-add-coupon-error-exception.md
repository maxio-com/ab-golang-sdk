
# Subscription Add Coupon Error Exception

## Structure

`SubscriptionAddCouponErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Codes` | `[]string` | Optional | - |
| `CouponCode` | `[]string` | Optional | - |
| `CouponCodes` | `[]string` | Optional | - |
| `Subscription` | `[]string` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SubscriptionAddCouponErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

