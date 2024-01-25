
# Agreement Acceptance

Required when creating a subscription with Maxio Payments.

## Structure

`AgreementAcceptance`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `IpAddress` | `*string` | Optional | Required when providing agreement acceptance params. |
| `TermsUrl` | `*string` | Optional | Required when creating a subscription with Maxio Payments. Either terms_url or provacy_policy_url required when providing agreement_acceptance params. |
| `PrivacyPolicyUrl` | `*string` | Optional | - |
| `ReturnRefundPolicyUrl` | `*string` | Optional | - |
| `DeliveryPolicyUrl` | `*string` | Optional | - |
| `SecureCheckoutPolicyUrl` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "ip_address": "ip_address2",
  "terms_url": "terms_url0",
  "privacy_policy_url": "privacy_policy_url0",
  "return_refund_policy_url": "return_refund_policy_url4",
  "delivery_policy_url": "delivery_policy_url8"
}
```

