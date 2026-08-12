
# Customer Error Response Exception

## Structure

`CustomerErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.CustomerErrorResponseErrors`](../../doc/models/containers/customer-error-response-errors.md) | Optional | This is a container for one-of cases. |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.CustomerErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

