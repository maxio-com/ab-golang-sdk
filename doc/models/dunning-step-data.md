
# Dunning Step Data

## Structure

`DunningStepData`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DayThreshold` | `int` | Required | - |
| `Action` | `string` | Required | - |
| `EmailBody` | `models.Optional[string]` | Optional | - |
| `EmailSubject` | `models.Optional[string]` | Optional | - |
| `SendEmail` | `bool` | Required | - |
| `SendBccEmail` | `bool` | Required | - |
| `SendSms` | `bool` | Required | - |
| `SmsBody` | `models.Optional[string]` | Optional | - |

## Example

```go
package main

import (
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    dunningStepData := models.DunningStepData{
        DayThreshold:         206,
        Action:               "action6",
        EmailBody:            models.NewOptional(models.ToPointer("email_body6")),
        EmailSubject:         models.NewOptional(models.ToPointer("email_subject6")),
        SendEmail:            false,
        SendBccEmail:         false,
        SendSms:              false,
        SmsBody:              models.NewOptional(models.ToPointer("sms_body8")),
    }

}
```

