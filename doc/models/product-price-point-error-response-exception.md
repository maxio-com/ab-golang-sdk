
# Product Price Point Error Response Exception

## Structure

`ProductPricePointErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.ProductPricePointErrors`](../../doc/models/product-price-point-errors.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ProductPricePointErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

