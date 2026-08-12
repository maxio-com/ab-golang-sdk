
# Error Array Map Response Exception

## Structure

`ErrorArrayMapResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `map[string]interface{}` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ErrorArrayMapResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

