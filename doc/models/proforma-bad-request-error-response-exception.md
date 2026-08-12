
# Proforma Bad Request Error Response Exception

## Structure

`ProformaBadRequestErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.ProformaError`](../../doc/models/proforma-error.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ProformaBadRequestErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

