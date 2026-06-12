
# Tokenized Payment Profile

## Structure

`TokenizedPaymentProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `int` | Required | - |
| `VaultToken` | `*string` | Optional | - |
| `GatewayHandle` | `models.Optional[string]` | Optional | - |
| `CustomerVaultToken` | `models.Optional[string]` | Optional | - |

## Example (as JSON)

```json
{
  "id": 72,
  "vault_token": "vault_token2",
  "gateway_handle": "gateway_handle8",
  "customer_vault_token": "customer_vault_token8"
}
```

