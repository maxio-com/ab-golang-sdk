
# Chjs Tokenization Success

## Structure

`ChjsTokenizationSuccess`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.TokenizedPaymentProfile`](../../doc/models/tokenized-payment-profile.md) | Required | - |
| `GatewayCustomerId` | `models.Optional[int]` | Optional | - |

## Example (as JSON)

```json
{
  "payment_profile": {
    "id": 44,
    "vault_token": "vault_token6",
    "gateway_handle": "gateway_handle4",
    "customer_vault_token": "customer_vault_token2"
  },
  "gateway_customer_id": 44
}
```

