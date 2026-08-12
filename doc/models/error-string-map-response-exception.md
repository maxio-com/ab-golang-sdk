
# Error String Map Response Exception

## Structure

`ErrorStringMapResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `map[string]string` | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ErrorStringMapResponseException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

