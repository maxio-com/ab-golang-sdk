
# Refund Prepayment Aggregated Errors Response Exception

Errors returned on creating a refund prepayment, grouped by field, as arrays of strings.

## Structure

`RefundPrepaymentAggregatedErrorsResponseException`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Errors` | [`*models.RefundPrepaymentAggregatedError`](../../doc/models/refund-prepayment-aggregated-error.md) | Optional | - |

## Example (as JSON)

```json
{
  "errors": {
    "refund": {
      "amount_in_cents": [
        "amount_in_cents5"
      ],
      "base": [
        "base7"
      ],
      "external": [
        "external0",
        "external1"
      ]
    }
  }
}
```

