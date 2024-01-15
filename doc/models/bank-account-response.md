
# Bank Account Response

## Structure

`BankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.BankAccount`](../../doc/models/bank-account.md) | Required | - |

## Example (as JSON)

```json
{
  "payment_profile": {
    "verified": false,
    "id": 44,
    "first_name": "first_name4",
    "last_name": "last_name2",
    "customer_id": 82,
    "current_vault": "gocardless"
  }
}
```

