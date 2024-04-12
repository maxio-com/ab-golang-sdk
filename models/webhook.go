package models

import (
    "encoding/json"
    "log"
    "time"
)

// Webhook represents a Webhook struct.
type Webhook struct {
    // A string describing which event type produced the given webhook
    Event                *string             `json:"event,omitempty"`
    // The unique identifier for the webhooks (unique across all of Chargify). This is not changed on a retry/replay of the same webhook, so it may be used to avoid duplicate action for the same event.
    Id                   *int64              `json:"id,omitempty"`
    // Timestamp indicating when the webhook was created
    CreatedAt            *time.Time          `json:"created_at,omitempty"`
    // Text describing the status code and/or error from the last failed attempt to send the Webhook. When a webhook is retried and accepted, this field will be cleared.
    LastError            *string             `json:"last_error,omitempty"`
    // Timestamp indicating when the last non-acceptance occurred. If a webhook is later resent and accepted, this field will be cleared.
    LastErrorAt          *time.Time          `json:"last_error_at,omitempty"`
    // Timestamp indicating when the webhook was accepted by the merchant endpoint. When a webhook is explicitly replayed by the merchant, this value will be cleared until it is accepted again.
    AcceptedAt           Optional[time.Time] `json:"accepted_at"`
    // Timestamp indicating when the most recent attempt was made to send the webhook
    LastSentAt           *time.Time          `json:"last_sent_at,omitempty"`
    // The url that the endpoint was last sent to.
    LastSentUrl          *string             `json:"last_sent_url,omitempty"`
    // A boolean flag describing whether the webhook was accepted by the webhook endpoint for the most recent attempt. (Acceptance is defined by receiving a “200 OK” HTTP response within a reasonable timeframe, i.e. 15 seconds)
    Successful           *bool               `json:"successful,omitempty"`
    // The data sent within the webhook post
    Body                 *string             `json:"body,omitempty"`
    // The calculated webhook signature
    Signature            *string             `json:"signature,omitempty"`
    // The calculated HMAC-SHA-256 webhook signature
    SignatureHmacSha256  *string             `json:"signature_hmac_sha_256,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Webhook.
// It customizes the JSON marshaling process for Webhook objects.
func (w Webhook) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the Webhook object to a map representation for JSON marshaling.
func (w Webhook) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Event != nil {
        structMap["event"] = w.Event
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt.Format(time.RFC3339)
    }
    if w.LastError != nil {
        structMap["last_error"] = w.LastError
    }
    if w.LastErrorAt != nil {
        structMap["last_error_at"] = w.LastErrorAt.Format(time.RFC3339)
    }
    if w.AcceptedAt.IsValueSet() {
        var AcceptedAtVal *string = nil
        if w.AcceptedAt.Value() != nil {
            val := w.AcceptedAt.Value().Format(time.RFC3339)
            AcceptedAtVal = &val
        }
        if w.AcceptedAt.Value() != nil {
            structMap["accepted_at"] = AcceptedAtVal
        } else {
            structMap["accepted_at"] = nil
        }
    }
    if w.LastSentAt != nil {
        structMap["last_sent_at"] = w.LastSentAt.Format(time.RFC3339)
    }
    if w.LastSentUrl != nil {
        structMap["last_sent_url"] = w.LastSentUrl
    }
    if w.Successful != nil {
        structMap["successful"] = w.Successful
    }
    if w.Body != nil {
        structMap["body"] = w.Body
    }
    if w.Signature != nil {
        structMap["signature"] = w.Signature
    }
    if w.SignatureHmacSha256 != nil {
        structMap["signature_hmac_sha_256"] = w.SignatureHmacSha256
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Webhook.
// It customizes the JSON unmarshaling process for Webhook objects.
func (w *Webhook) UnmarshalJSON(input []byte) error {
    var temp webhook
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "event", "id", "created_at", "last_error", "last_error_at", "accepted_at", "last_sent_at", "last_sent_url", "successful", "body", "signature", "signature_hmac_sha_256")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Event = temp.Event
    w.Id = temp.Id
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        w.CreatedAt = &CreatedAtVal
    }
    w.LastError = temp.LastError
    if temp.LastErrorAt != nil {
        LastErrorAtVal, err := time.Parse(time.RFC3339, *temp.LastErrorAt)
        if err != nil {
            log.Fatalf("Cannot Parse last_error_at as % s format.", time.RFC3339)
        }
        w.LastErrorAt = &LastErrorAtVal
    }
    w.AcceptedAt.ShouldSetValue(temp.AcceptedAt.IsValueSet())
    if temp.AcceptedAt.Value() != nil {
        AcceptedAtVal, err := time.Parse(time.RFC3339, (*temp.AcceptedAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse accepted_at as % s format.", time.RFC3339)
        }
        w.AcceptedAt.SetValue(&AcceptedAtVal)
    }
    if temp.LastSentAt != nil {
        LastSentAtVal, err := time.Parse(time.RFC3339, *temp.LastSentAt)
        if err != nil {
            log.Fatalf("Cannot Parse last_sent_at as % s format.", time.RFC3339)
        }
        w.LastSentAt = &LastSentAtVal
    }
    w.LastSentUrl = temp.LastSentUrl
    w.Successful = temp.Successful
    w.Body = temp.Body
    w.Signature = temp.Signature
    w.SignatureHmacSha256 = temp.SignatureHmacSha256
    return nil
}

// TODO
type webhook  struct {
    Event               *string          `json:"event,omitempty"`
    Id                  *int64           `json:"id,omitempty"`
    CreatedAt           *string          `json:"created_at,omitempty"`
    LastError           *string          `json:"last_error,omitempty"`
    LastErrorAt         *string          `json:"last_error_at,omitempty"`
    AcceptedAt          Optional[string] `json:"accepted_at"`
    LastSentAt          *string          `json:"last_sent_at,omitempty"`
    LastSentUrl         *string          `json:"last_sent_url,omitempty"`
    Successful          *bool            `json:"successful,omitempty"`
    Body                *string          `json:"body,omitempty"`
    Signature           *string          `json:"signature,omitempty"`
    SignatureHmacSha256 *string          `json:"signature_hmac_sha_256,omitempty"`
}
