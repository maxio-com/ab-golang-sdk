
# Scheduled Renewal Configuration Request

## Structure

`ScheduledRenewalConfigurationRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RenewalConfiguration` | [`models.ScheduledRenewalConfigurationRequestBody`](../../doc/models/scheduled-renewal-configuration-request-body.md) | Required | - |

## Example (as JSON)

```json
{
  "renewal_configuration": {
    "starts_at": "2016-03-13T12:52:32.123Z",
    "ends_at": "2016-03-13T12:52:32.123Z",
    "lock_in_at": "2016-03-13T12:52:32.123Z",
    "contract_id": 244,
    "create_new_contract": false
  }
}
```

