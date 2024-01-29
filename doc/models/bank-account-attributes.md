
# Bank Account Attributes

## Structure

`BankAccountAttributes`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChargifyToken` | `*string` | Optional | - |
| `BankName` | `*string` | Optional | (Required when creating a subscription with ACH or GoCardless) The name of the bank where the customer’s account resides |
| `BankRoutingNumber` | `*string` | Optional | (Required when creating a subscription with ACH. Optional when creating a subscription with GoCardless). The routing number of the bank. It becomes bank_code while passing via GoCardless API |
| `BankAccountNumber` | `*string` | Optional | (Required when creating a subscription with ACH. Required when creating a subscription with GoCardless and bank_iban is blank) The customerʼs bank account number |
| `BankAccountType` | [`*models.BankAccountType`](../../doc/models/bank-account-type.md) | Optional | Defaults to checking<br>**Default**: `"checking"` |
| `BankBranchCode` | `*string` | Optional | (Optional when creating a subscription with GoCardless) Branch code. Alternatively, an IBAN can be provided |
| `BankIban` | `*string` | Optional | (Optional when creating a subscription with GoCardless). International Bank Account Number. Alternatively, local bank details can be provided |
| `BankAccountHolderType` | [`*models.BankAccountHolderType`](../../doc/models/bank-account-holder-type.md) | Optional | Defaults to personal |
| `PaymentType` | [`*models.PaymentType`](../../doc/models/payment-type.md) | Optional | **Default**: `"credit_card"` |
| `CurrentVault` | [`*models.BankAccountVault`](../../doc/models/bank-account-vault.md) | Optional | The vault that stores the payment profile with the provided vault_token. |
| `VaultToken` | `*string` | Optional | - |
| `CustomerVaultToken` | `*string` | Optional | (only for Authorize.Net CIM storage or Square) The customerProfileId for the owner of the customerPaymentProfileId provided as the vault_token |

## Example (as JSON)

```json
{
  "bank_account_type": "checking",
  "payment_type": "credit_card",
  "chargify_token": "chargify_token0",
  "bank_name": "bank_name2",
  "bank_routing_number": "bank_routing_number8",
  "bank_account_number": "bank_account_number4"
}
```

