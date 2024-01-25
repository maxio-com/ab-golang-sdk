
# Invoice Pre Payment

## Structure

`InvoicePrePayment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SubscriptionId` | `*string` | Optional | The subscription id for the prepayment account |
| `AmountInCents` | `*string` | Optional | The amount in cents of the prepayment that was created as a result of this payment. |
| `EndingBalanceInCents` | `*string` | Optional | The total balance of the prepayment account for this subscription including any prior prepayments |

## Example (as JSON)

```json
{
  "subscription_id": "subscription_id0",
  "amount_in_cents": "amount_in_cents2",
  "ending_balance_in_cents": "ending_balance_in_cents4"
}
```

