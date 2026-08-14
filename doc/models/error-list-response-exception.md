
# Error List Response Exception

Error which contains list of messages.

## Structure

`ErrorListResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `[]string` | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ErrorListResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

