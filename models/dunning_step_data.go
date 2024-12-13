/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DunningStepData represents a DunningStepData struct.
type DunningStepData struct {
    DayThreshold         int                    `json:"day_threshold"`
    Action               string                 `json:"action"`
    EmailBody            Optional[string]       `json:"email_body"`
    EmailSubject         Optional[string]       `json:"email_subject"`
    SendEmail            bool                   `json:"send_email"`
    SendBccEmail         bool                   `json:"send_bcc_email"`
    SendSms              bool                   `json:"send_sms"`
    SmsBody              Optional[string]       `json:"sms_body"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DunningStepData.
// It customizes the JSON marshaling process for DunningStepData objects.
func (d DunningStepData) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "day_threshold", "action", "email_body", "email_subject", "send_email", "send_bcc_email", "send_sms", "sms_body"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DunningStepData object to a map representation for JSON marshaling.
func (d DunningStepData) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["day_threshold"] = d.DayThreshold
    structMap["action"] = d.Action
    if d.EmailBody.IsValueSet() {
        if d.EmailBody.Value() != nil {
            structMap["email_body"] = d.EmailBody.Value()
        } else {
            structMap["email_body"] = nil
        }
    }
    if d.EmailSubject.IsValueSet() {
        if d.EmailSubject.Value() != nil {
            structMap["email_subject"] = d.EmailSubject.Value()
        } else {
            structMap["email_subject"] = nil
        }
    }
    structMap["send_email"] = d.SendEmail
    structMap["send_bcc_email"] = d.SendBccEmail
    structMap["send_sms"] = d.SendSms
    if d.SmsBody.IsValueSet() {
        if d.SmsBody.Value() != nil {
            structMap["sms_body"] = d.SmsBody.Value()
        } else {
            structMap["sms_body"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DunningStepData.
// It customizes the JSON unmarshaling process for DunningStepData objects.
func (d *DunningStepData) UnmarshalJSON(input []byte) error {
    var temp tempDunningStepData
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "day_threshold", "action", "email_body", "email_subject", "send_email", "send_bcc_email", "send_sms", "sms_body")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.DayThreshold = *temp.DayThreshold
    d.Action = *temp.Action
    d.EmailBody = temp.EmailBody
    d.EmailSubject = temp.EmailSubject
    d.SendEmail = *temp.SendEmail
    d.SendBccEmail = *temp.SendBccEmail
    d.SendSms = *temp.SendSms
    d.SmsBody = temp.SmsBody
    return nil
}

// tempDunningStepData is a temporary struct used for validating the fields of DunningStepData.
type tempDunningStepData  struct {
    DayThreshold *int             `json:"day_threshold"`
    Action       *string          `json:"action"`
    EmailBody    Optional[string] `json:"email_body"`
    EmailSubject Optional[string] `json:"email_subject"`
    SendEmail    *bool            `json:"send_email"`
    SendBccEmail *bool            `json:"send_bcc_email"`
    SendSms      *bool            `json:"send_sms"`
    SmsBody      Optional[string] `json:"sms_body"`
}

func (d *tempDunningStepData) validate() error {
    var errs []string
    if d.DayThreshold == nil {
        errs = append(errs, "required field `day_threshold` is missing for type `Dunning Step Data`")
    }
    if d.Action == nil {
        errs = append(errs, "required field `action` is missing for type `Dunning Step Data`")
    }
    if d.SendEmail == nil {
        errs = append(errs, "required field `send_email` is missing for type `Dunning Step Data`")
    }
    if d.SendBccEmail == nil {
        errs = append(errs, "required field `send_bcc_email` is missing for type `Dunning Step Data`")
    }
    if d.SendSms == nil {
        errs = append(errs, "required field `send_sms` is missing for type `Dunning Step Data`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
