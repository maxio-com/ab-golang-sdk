
# Single String Error Response Exception

## Structure

`SingleStringErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `*string` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SingleStringErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

