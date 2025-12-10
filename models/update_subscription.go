// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// UpdateSubscription represents a UpdateSubscription struct.
type UpdateSubscription struct {
    CreditCardAttributes              *CreditCardAttributes               `json:"credit_card_attributes,omitempty"`
    // Set to the handle of a different product to change the subscription's product
    ProductHandle                     *string                             `json:"product_handle,omitempty"`
    // Set to the id of a different product to change the subscription's product
    ProductId                         *int                                `json:"product_id,omitempty"`
    ProductChangeDelayed              *bool                               `json:"product_change_delayed,omitempty"`
    // Set to an empty string to cancel a delayed product change.
    NextProductId                     *string                             `json:"next_product_id,omitempty"`
    NextProductPricePointId           *string                             `json:"next_product_price_point_id,omitempty"`
    // Use for subscriptions with product eligible for calendar billing only. Value can be 1-28 or 'end'.
    SnapDay                           Optional[UpdateSubscriptionSnapDay] `json:"snap_day"`
    // (Optional) Set this attribute to a future date/time to update a subscription in the Awaiting Signup Date state, to Awaiting Signup. In the Awaiting Signup state, a subscription behaves like any other. It can be canceled, allocated to, or have its billing date changed. etc. When the `initial_billing_at` date hits, the subscription will transition to the expected state. If the product has a trial, the subscription will enter a trial, otherwise it will go active. Setup fees will be respected either before or after the trial, as configured on the price point. If the payment is due at the initial_billing_at and it fails the subscription will be immediately canceled. You can omit the initial_billing_at date to activate the subscription immediately. See the [subscription import](https://maxio.zendesk.com/hc/en-us/articles/24251489107213-Advanced-Billing-Subscription-Imports#date-format) documentation for more information about Date/Time formats.
    InitialBillingAt                  *time.Time                          `json:"initial_billing_at,omitempty"`
    // (Optional) Set this attribute to true to move the subscription from Awaiting Signup, to Awaiting Signup Date. Use this when you want to update a subscription that has an unknown initial billing date. When the first billing date is known, update a subscription to set the `initial_billing_at` date. The subscription moves to the awaiting signup with a scheduled initial billing date. You can omit the initial_billing_at date to activate the subscription immediately. See [Subscription States](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404222005773-Subscription-States) for more information.
    DeferSignup                       *bool                               `json:"defer_signup,omitempty"`
    NextBillingAt                     *time.Time                          `json:"next_billing_at,omitempty"`
    // Timestamp giving the expiration date of this subscription (if any). You may manually change the expiration date at any point during a subscription period.
    ExpiresAt                         *time.Time                          `json:"expires_at,omitempty"`
    PaymentCollectionMethod           *string                             `json:"payment_collection_method,omitempty"`
    ReceivesInvoiceEmails             *bool                               `json:"receives_invoice_emails,omitempty"`
    NetTerms                          *UpdateSubscriptionNetTerms         `json:"net_terms,omitempty"`
    StoredCredentialTransactionId     *int                                `json:"stored_credential_transaction_id,omitempty"`
    Reference                         *string                             `json:"reference,omitempty"`
    // (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
    CustomPrice                       *SubscriptionCustomPrice            `json:"custom_price,omitempty"`
    // (Optional) An array of component ids and custom prices to be added to the subscription.
    Components                        []UpdateSubscriptionComponent       `json:"components,omitempty"`
    // Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute.
    DunningCommunicationDelayEnabled  *bool                               `json:"dunning_communication_delay_enabled,omitempty"`
    // Time zone for the Dunning Communication Delay feature.
    DunningCommunicationDelayTimeZone Optional[string]                    `json:"dunning_communication_delay_time_zone"`
    // Set to change the current product's price point.
    ProductPricePointId               *int                                `json:"product_price_point_id,omitempty"`
    // Set to change the current product's price point.
    ProductPricePointHandle           *string                             `json:"product_price_point_handle,omitempty"`
    AdditionalProperties              map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for UpdateSubscription,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UpdateSubscription) String() string {
    return fmt.Sprintf(
    	"UpdateSubscription[CreditCardAttributes=%v, ProductHandle=%v, ProductId=%v, ProductChangeDelayed=%v, NextProductId=%v, NextProductPricePointId=%v, SnapDay=%v, InitialBillingAt=%v, DeferSignup=%v, NextBillingAt=%v, ExpiresAt=%v, PaymentCollectionMethod=%v, ReceivesInvoiceEmails=%v, NetTerms=%v, StoredCredentialTransactionId=%v, Reference=%v, CustomPrice=%v, Components=%v, DunningCommunicationDelayEnabled=%v, DunningCommunicationDelayTimeZone=%v, ProductPricePointId=%v, ProductPricePointHandle=%v, AdditionalProperties=%v]",
    	u.CreditCardAttributes, u.ProductHandle, u.ProductId, u.ProductChangeDelayed, u.NextProductId, u.NextProductPricePointId, u.SnapDay, u.InitialBillingAt, u.DeferSignup, u.NextBillingAt, u.ExpiresAt, u.PaymentCollectionMethod, u.ReceivesInvoiceEmails, u.NetTerms, u.StoredCredentialTransactionId, u.Reference, u.CustomPrice, u.Components, u.DunningCommunicationDelayEnabled, u.DunningCommunicationDelayTimeZone, u.ProductPricePointId, u.ProductPricePointHandle, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscription.
// It customizes the JSON marshaling process for UpdateSubscription objects.
func (u UpdateSubscription) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "credit_card_attributes", "product_handle", "product_id", "product_change_delayed", "next_product_id", "next_product_price_point_id", "snap_day", "initial_billing_at", "defer_signup", "next_billing_at", "expires_at", "payment_collection_method", "receives_invoice_emails", "net_terms", "stored_credential_transaction_id", "reference", "custom_price", "components", "dunning_communication_delay_enabled", "dunning_communication_delay_time_zone", "product_price_point_id", "product_price_point_handle"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscription object to a map representation for JSON marshaling.
func (u UpdateSubscription) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.CreditCardAttributes != nil {
        structMap["credit_card_attributes"] = u.CreditCardAttributes.toMap()
    }
    if u.ProductHandle != nil {
        structMap["product_handle"] = u.ProductHandle
    }
    if u.ProductId != nil {
        structMap["product_id"] = u.ProductId
    }
    if u.ProductChangeDelayed != nil {
        structMap["product_change_delayed"] = u.ProductChangeDelayed
    }
    if u.NextProductId != nil {
        structMap["next_product_id"] = u.NextProductId
    }
    if u.NextProductPricePointId != nil {
        structMap["next_product_price_point_id"] = u.NextProductPricePointId
    }
    if u.SnapDay.IsValueSet() {
        if u.SnapDay.Value() != nil {
            structMap["snap_day"] = u.SnapDay.Value().toMap()
        } else {
            structMap["snap_day"] = nil
        }
    }
    if u.InitialBillingAt != nil {
        structMap["initial_billing_at"] = u.InitialBillingAt.Format(time.RFC3339)
    }
    if u.DeferSignup != nil {
        structMap["defer_signup"] = u.DeferSignup
    }
    if u.NextBillingAt != nil {
        structMap["next_billing_at"] = u.NextBillingAt.Format(time.RFC3339)
    }
    if u.ExpiresAt != nil {
        structMap["expires_at"] = u.ExpiresAt.Format(time.RFC3339)
    }
    if u.PaymentCollectionMethod != nil {
        structMap["payment_collection_method"] = u.PaymentCollectionMethod
    }
    if u.ReceivesInvoiceEmails != nil {
        structMap["receives_invoice_emails"] = u.ReceivesInvoiceEmails
    }
    if u.NetTerms != nil {
        structMap["net_terms"] = u.NetTerms.toMap()
    }
    if u.StoredCredentialTransactionId != nil {
        structMap["stored_credential_transaction_id"] = u.StoredCredentialTransactionId
    }
    if u.Reference != nil {
        structMap["reference"] = u.Reference
    }
    if u.CustomPrice != nil {
        structMap["custom_price"] = u.CustomPrice.toMap()
    }
    if u.Components != nil {
        structMap["components"] = u.Components
    }
    if u.DunningCommunicationDelayEnabled != nil {
        structMap["dunning_communication_delay_enabled"] = u.DunningCommunicationDelayEnabled
    }
    if u.DunningCommunicationDelayTimeZone.IsValueSet() {
        if u.DunningCommunicationDelayTimeZone.Value() != nil {
            structMap["dunning_communication_delay_time_zone"] = u.DunningCommunicationDelayTimeZone.Value()
        } else {
            structMap["dunning_communication_delay_time_zone"] = nil
        }
    }
    if u.ProductPricePointId != nil {
        structMap["product_price_point_id"] = u.ProductPricePointId
    }
    if u.ProductPricePointHandle != nil {
        structMap["product_price_point_handle"] = u.ProductPricePointHandle
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscription.
// It customizes the JSON unmarshaling process for UpdateSubscription objects.
func (u *UpdateSubscription) UnmarshalJSON(input []byte) error {
    var temp tempUpdateSubscription
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "credit_card_attributes", "product_handle", "product_id", "product_change_delayed", "next_product_id", "next_product_price_point_id", "snap_day", "initial_billing_at", "defer_signup", "next_billing_at", "expires_at", "payment_collection_method", "receives_invoice_emails", "net_terms", "stored_credential_transaction_id", "reference", "custom_price", "components", "dunning_communication_delay_enabled", "dunning_communication_delay_time_zone", "product_price_point_id", "product_price_point_handle")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.CreditCardAttributes = temp.CreditCardAttributes
    u.ProductHandle = temp.ProductHandle
    u.ProductId = temp.ProductId
    u.ProductChangeDelayed = temp.ProductChangeDelayed
    u.NextProductId = temp.NextProductId
    u.NextProductPricePointId = temp.NextProductPricePointId
    u.SnapDay = temp.SnapDay
    if temp.InitialBillingAt != nil {
        InitialBillingAtVal, err := time.Parse(time.RFC3339, *temp.InitialBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse initial_billing_at as % s format.", time.RFC3339)
        }
        u.InitialBillingAt = &InitialBillingAtVal
    }
    u.DeferSignup = temp.DeferSignup
    if temp.NextBillingAt != nil {
        NextBillingAtVal, err := time.Parse(time.RFC3339, *temp.NextBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
        }
        u.NextBillingAt = &NextBillingAtVal
    }
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        u.ExpiresAt = &ExpiresAtVal
    }
    u.PaymentCollectionMethod = temp.PaymentCollectionMethod
    u.ReceivesInvoiceEmails = temp.ReceivesInvoiceEmails
    u.NetTerms = temp.NetTerms
    u.StoredCredentialTransactionId = temp.StoredCredentialTransactionId
    u.Reference = temp.Reference
    u.CustomPrice = temp.CustomPrice
    u.Components = temp.Components
    u.DunningCommunicationDelayEnabled = temp.DunningCommunicationDelayEnabled
    u.DunningCommunicationDelayTimeZone = temp.DunningCommunicationDelayTimeZone
    u.ProductPricePointId = temp.ProductPricePointId
    u.ProductPricePointHandle = temp.ProductPricePointHandle
    return nil
}

// tempUpdateSubscription is a temporary struct used for validating the fields of UpdateSubscription.
type tempUpdateSubscription  struct {
    CreditCardAttributes              *CreditCardAttributes               `json:"credit_card_attributes,omitempty"`
    ProductHandle                     *string                             `json:"product_handle,omitempty"`
    ProductId                         *int                                `json:"product_id,omitempty"`
    ProductChangeDelayed              *bool                               `json:"product_change_delayed,omitempty"`
    NextProductId                     *string                             `json:"next_product_id,omitempty"`
    NextProductPricePointId           *string                             `json:"next_product_price_point_id,omitempty"`
    SnapDay                           Optional[UpdateSubscriptionSnapDay] `json:"snap_day"`
    InitialBillingAt                  *string                             `json:"initial_billing_at,omitempty"`
    DeferSignup                       *bool                               `json:"defer_signup,omitempty"`
    NextBillingAt                     *string                             `json:"next_billing_at,omitempty"`
    ExpiresAt                         *string                             `json:"expires_at,omitempty"`
    PaymentCollectionMethod           *string                             `json:"payment_collection_method,omitempty"`
    ReceivesInvoiceEmails             *bool                               `json:"receives_invoice_emails,omitempty"`
    NetTerms                          *UpdateSubscriptionNetTerms         `json:"net_terms,omitempty"`
    StoredCredentialTransactionId     *int                                `json:"stored_credential_transaction_id,omitempty"`
    Reference                         *string                             `json:"reference,omitempty"`
    CustomPrice                       *SubscriptionCustomPrice            `json:"custom_price,omitempty"`
    Components                        []UpdateSubscriptionComponent       `json:"components,omitempty"`
    DunningCommunicationDelayEnabled  *bool                               `json:"dunning_communication_delay_enabled,omitempty"`
    DunningCommunicationDelayTimeZone Optional[string]                    `json:"dunning_communication_delay_time_zone"`
    ProductPricePointId               *int                                `json:"product_price_point_id,omitempty"`
    ProductPricePointHandle           *string                             `json:"product_price_point_handle,omitempty"`
}
