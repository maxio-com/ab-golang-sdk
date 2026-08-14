
# Event Based Billing Segment Errors Exception

## Structure

`EventBasedBillingSegmentErrorsException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | `map[string]interface{}` | Optional | The key of the object would be a number (an index in the request array) where the error occurred. In the value object, the key represents the field and the value is an array with error messages. In most cases, this object would contain just one key. |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.EventBasedBillingSegmentErrorsException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

