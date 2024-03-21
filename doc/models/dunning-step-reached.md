
# Dunning Step Reached

## Structure

`DunningStepReached`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dunner` | [`models.DunnerData`](../../doc/models/dunner-data.md) | Required | - |
| `CurrentStep` | [`models.DunningStepData`](../../doc/models/dunning-step-data.md) | Required | - |
| `NextStep` | [`models.DunningStepData`](../../doc/models/dunning-step-data.md) | Required | - |

## Example (as JSON)

```json
{
  "dunner": {
    "state": "state8",
    "subscription_id": 194,
    "revenue_at_risk_in_cents": 98,
    "created_at": "2016-03-13T12:52:32.123Z",
    "attempts": 42,
    "last_attempted_at": "2016-03-13T12:52:32.123Z"
  },
  "current_step": {
    "day_threshold": 198,
    "action": "action4",
    "email_body": "email_body4",
    "email_subject": "email_subject6",
    "send_email": false,
    "send_bcc_email": false,
    "send_sms": false,
    "sms_body": "sms_body0"
  },
  "next_step": {
    "day_threshold": 30,
    "action": "action4",
    "email_body": "email_body4",
    "email_subject": "email_subject4",
    "send_email": false,
    "send_bcc_email": false,
    "send_sms": false,
    "sms_body": "sms_body0"
  }
}
```

