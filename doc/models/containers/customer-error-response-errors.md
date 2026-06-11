
# Customer Error Response Errors

## Class Name

`CustomerErrorResponseErrors`

## Cases

| Type | Factory Method |
|  --- | --- |
| [`models.CustomerError`](../../../doc/models/customer-error.md) | models.CustomerErrorResponseErrorsContainer.FromCustomerError(models.CustomerError customerError) |
| `[]string` | models.CustomerErrorResponseErrorsContainer.FromArrayOfString([]string arrayOfString) |

## models.CustomerError

### Initialization Code

#### Example

```go
value := models.CustomerErrorResponseErrorsContainer.FromCustomerError(models.CustomerError{
})
```

## []string

### Initialization Code

#### Example

```go
value := models.CustomerErrorResponseErrorsContainer.FromArrayOfString([]string{
    "String1",
})
```

