
# Single Error Response Exception

## Structure

`SingleErrorResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Error` | `string` | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.SingleErrorResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

