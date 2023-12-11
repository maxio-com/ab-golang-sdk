
# Refund Prepayment Base Errors Response Exception

Errors returned on creating a refund prepayment when bad request

## Structure

`RefundPrepaymentBaseErrorsResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.RefundPrepaymentBaseRefundError`](refund-prepayment-base-refund-error.md) | Optional | - |

## Example (as JSON)

```json
{
  "errors": {
    "refund": {
      "base": [
        {
          "key1": "val1",
          "key2": "val2"
        }
      ]
    }
  }
}
```

