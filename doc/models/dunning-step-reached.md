
# Dunning Step Reached

## Structure

`DunningStepReached`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dunner` | [`models.DunnerData`](../../doc/models/dunner-data.md) | Required | - |
| `CurrentStep` | [`models.DunningStepData`](../../doc/models/dunning-step-data.md) | Required | - |
| `NextStep` | [`models.DunningStepData`](../../doc/models/dunning-step-data.md) | Required | - |

## Example

```go
package main

import (
    "log"
    "time"
    "github.com/maxio-com/ab-golang-sdk/models"
)

func main() {
    parseTime := func(layout, value string, errCallback func(error)) time.Time {
        dateTime, err := time.Parse(layout, value)
        if err != nil {
            errCallback(err) 
       }
        return dateTime
    }
    dunningStepReached := models.DunningStepReached{
        Dunner:               models.DunnerData{
            State:                "state8",
            SubscriptionId:       194,
            RevenueAtRiskInCents: int64(98),
            CreatedAt:            parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
            Attempts:             42,
            LastAttemptedAt:      parseTime(time.RFC3339, "2016-03-13T12:52:32.123Z", func(err error) { log.Fatalln(err) }),
        },
        CurrentStep:          models.DunningStepData{
            DayThreshold:         198,
            Action:               "action4",
            EmailBody:            models.NewOptional(models.ToPointer("email_body4")),
            EmailSubject:         models.NewOptional(models.ToPointer("email_subject6")),
            SendEmail:            false,
            SendBccEmail:         false,
            SendSms:              false,
            SmsBody:              models.NewOptional(models.ToPointer("sms_body0")),
        },
        NextStep:             models.DunningStepData{
            DayThreshold:         30,
            Action:               "action4",
            EmailBody:            models.NewOptional(models.ToPointer("email_body4")),
            EmailSubject:         models.NewOptional(models.ToPointer("email_subject4")),
            SendEmail:            false,
            SendBccEmail:         false,
            SendSms:              false,
            SmsBody:              models.NewOptional(models.ToPointer("sms_body0")),
        },
    }

}
```

