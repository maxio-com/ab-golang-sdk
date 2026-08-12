
# Too Many Management Link Requests Error Exception

## Structure

`TooManyManagementLinkRequestsErrorException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`models.TooManyManagementLinkRequests`](../../doc/models/too-many-management-link-requests.md) | Required | - |

## Example

```go
if err != nil {
    switch typedErr := err.(type) {
    case *errors.TooManyManagementLinkRequestsErrorException:
        log.Fatalln(typedErr)
    default:
        log.Fatalln(err)
    }
}
```

