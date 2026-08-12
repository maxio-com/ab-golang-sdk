
# Component Allocation Error Exception

## Structure

`ComponentAllocationErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`[]models.ComponentAllocationErrorItem`](../../doc/models/component-allocation-error-item.md) | Optional | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.ComponentAllocationErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

