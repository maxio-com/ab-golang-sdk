
# Component Price Point Error Exception

## Structure

`ComponentPricePointErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`[]models.ComponentPricePointErrorItem`](../../doc/models/component-price-point-error-item.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ComponentPricePointErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

