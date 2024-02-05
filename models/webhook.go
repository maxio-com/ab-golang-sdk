package models

import (
    "encoding/json"
)

// Webhook represents a Webhook struct.
type Webhook struct {
    // A string describing which event type produced the given webhook
    Event               *string          `json:"event,omitempty"`
    // The unique identifier for the webhooks (unique across all of Chargify). This is not changed on a retry/replay of the same webhook, so it may be used to avoid duplicate action for the same event.
    Id                  *int             `json:"id,omitempty"`
    // Timestamp indicating when the webhook was created
    CreatedAt           *string          `json:"created_at,omitempty"`
    // Text describing the status code and/or error from the last failed attempt to send the Webhook. When a webhook is retried and accepted, this field will be cleared.
    LastError           *string          `json:"last_error,omitempty"`
    // Timestamp indicating when the last non-acceptance occurred. If a webhook is later resent and accepted, this field will be cleared.
    LastErrorAt         *string          `json:"last_error_at,omitempty"`
    // Timestamp indicating when the webhook was accepted by the merchant endpoint. When a webhook is explicitly replayed by the merchant, this value will be cleared until it is accepted again.
    AcceptedAt          Optional[string] `json:"accepted_at"`
    // Timestamp indicating when the most recent attempt was made to send the webhook
    LastSentAt          *string          `json:"last_sent_at,omitempty"`
    // The url that the endpoint was last sent to.
    LastSentUrl         *string          `json:"last_sent_url,omitempty"`
    // A boolean flag describing whether the webhook was accepted by the webhook endpoint for the most recent attempt. (Acceptance is defined by receiving a “200 OK” HTTP response within a reasonable timeframe, i.e. 15 seconds)
    Successful          *bool            `json:"successful,omitempty"`
    // The data sent within the webhook post
    Body                *string          `json:"body,omitempty"`
    // The calculated webhook signature
    Signature           *string          `json:"signature,omitempty"`
    // The calculated HMAC-SHA-256 webhook signature
    SignatureHmacSha256 *string          `json:"signature_hmac_sha_256,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Webhook.
// It customizes the JSON marshaling process for Webhook objects.
func (w *Webhook) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the Webhook object to a map representation for JSON marshaling.
func (w *Webhook) toMap() map[string]any {
    structMap := make(map[string]any)
    if w.Event != nil {
        structMap["event"] = w.Event
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt
    }
    if w.LastError != nil {
        structMap["last_error"] = w.LastError
    }
    if w.LastErrorAt != nil {
        structMap["last_error_at"] = w.LastErrorAt
    }
    if w.AcceptedAt.IsValueSet() {
        structMap["accepted_at"] = w.AcceptedAt.Value()
    }
    if w.LastSentAt != nil {
        structMap["last_sent_at"] = w.LastSentAt
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
    temp := &struct {
        Event               *string          `json:"event,omitempty"`
        Id                  *int             `json:"id,omitempty"`
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
    }{}
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    
    w.Event = temp.Event
    w.Id = temp.Id
    w.CreatedAt = temp.CreatedAt
    w.LastError = temp.LastError
    w.LastErrorAt = temp.LastErrorAt
    w.AcceptedAt = temp.AcceptedAt
    w.LastSentAt = temp.LastSentAt
    w.LastSentUrl = temp.LastSentUrl
    w.Successful = temp.Successful
    w.Body = temp.Body
    w.Signature = temp.Signature
    w.SignatureHmacSha256 = temp.SignatureHmacSha256
    return nil
}
