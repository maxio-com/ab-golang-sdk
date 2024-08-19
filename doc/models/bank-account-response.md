
# Bank Account Response

## Structure

`BankAccountResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.BankAccountPaymentProfile`](../../doc/models/bank-account-payment-profile.md) | Required | - |

## Example (as JSON)

```json
{
  "payment_profile": {
    "masked_bank_account_number": "masked_bank_account_number6",
    "payment_type": "bank_account",
    "verified": false,
    "id": 44,
    "first_name": "first_name4",
    "last_name": "last_name2",
    "customer_id": 82,
    "current_vault": "authorizenet"
  }
}
```

