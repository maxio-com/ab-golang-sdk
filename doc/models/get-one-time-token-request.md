
# Get One Time Token Request

## Structure

`GetOneTimeTokenRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PaymentProfile` | [`models.GetOneTimeTokenPaymentProfile`](../../doc/models/get-one-time-token-payment-profile.md) | Required | - |

## Example (as JSON)

```json
{
  "payment_profile": {
    "id": "id4",
    "first_name": "first_name4",
    "last_name": "last_name2",
    "masked_card_number": "masked_card_number2",
    "card_type": "card_type0",
    "expiration_month": 133.5,
    "expiration_year": 156.84,
    "customer_id": "customer_id2",
    "current_vault": "eway",
    "vault_token": "vault_token6",
    "billing_address": "billing_address4",
    "billing_address_2": "billing_address_26",
    "billing_city": "billing_city8",
    "billing_country": "billing_country2",
    "billing_state": "billing_state2",
    "billing_zip": "billing_zip2",
    "payment_type": "payment_type6",
    "disabled": false,
    "site_gateway_setting_id": 104,
    "customer_vault_token": "customer_vault_token2",
    "gateway_handle": "gateway_handle4"
  }
}
```

