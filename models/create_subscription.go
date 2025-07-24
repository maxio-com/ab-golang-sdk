// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// CreateSubscription represents a CreateSubscription struct.
type CreateSubscription struct {
    // The API Handle of the product for which you are creating a subscription. Required, unless a `product_id` is given instead.
    ProductHandle                     *string                       `json:"product_handle,omitempty"`
    // The Product ID of the product for which you are creating a subscription. The product ID is not currently published, so we recommend using the API Handle instead.
    ProductId                         *int                          `json:"product_id,omitempty"`
    // The user-friendly API handle of a product's particular price point.
    ProductPricePointHandle           *string                       `json:"product_price_point_handle,omitempty"`
    // The ID of the particular price point on the product.
    ProductPricePointId               *int                          `json:"product_price_point_id,omitempty"`
    // (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
    CustomPrice                       *SubscriptionCustomPrice      `json:"custom_price,omitempty"`
    // (deprecated) The coupon code of the single coupon currently applied to the subscription. See coupon_codes instead as subscriptions can now have more than one coupon.
    CouponCode                        *string                       `json:"coupon_code,omitempty"`
    // An array for all the coupons attached to the subscription.
    CouponCodes                       []string                      `json:"coupon_codes,omitempty"`
    // The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
    PaymentCollectionMethod           *CollectionMethod             `json:"payment_collection_method,omitempty"`
    // (Optional) Default: True - Whether or not this subscription is set to receive emails related to this subscription.
    ReceivesInvoiceEmails             *string                       `json:"receives_invoice_emails,omitempty"`
    // (Optional) Default: null The number of days after renewal (on invoice billing) that a subscription is due. A value between 0 (due immediately) and 180.
    NetTerms                          *string                       `json:"net_terms,omitempty"`
    // The ID of an existing customer within Chargify. Required, unless a `customer_reference` or a set of `customer_attributes` is given.
    CustomerId                        *int                          `json:"customer_id,omitempty"`
    // (Optional) Set this attribute to a future date/time to sync imported subscriptions to your existing renewal schedule. See the notes on “Date/Time Format” in our [subscription import documentation](https://maxio.zendesk.com/hc/en-us/articles/24251489107213-Advanced-Billing-Subscription-Imports#date-format). If you provide a next_billing_at timestamp that is in the future, no trial or initial charges will be applied when you create the subscription. In fact, no payment will be captured at all. The first payment will be captured, according to the prices defined by the product, near the time specified by next_billing_at. If you do not provide a value for next_billing_at, any trial and/or initial charges will be assessed and charged at the time of subscription creation. If the card cannot be successfully charged, the subscription will not be created. See further notes in the section on Importing Subscriptions.
    NextBillingAt                     *time.Time                    `json:"next_billing_at,omitempty"`
    // (Optional) Set this attribute to a future date/time to create a subscription in the Awaiting Signup state, rather than Active or Trialing. You can omit the initial_billing_at date to activate the subscription immediately. In the Awaiting Signup state, a subscription behaves like any other. It can be canceled, allocated to, or have its billing date changed. etc. When the initial_billing_at date hits, the subscription will transition to the expected state. If the product has a trial, the subscription will enter a trial, otherwise it will go active. Setup fees will be respected either before or after the trial, as configured on the price point. If the payment is due at the initial_billing_at and it fails the subscription will be immediately canceled. See the [subscription import](https://maxio.zendesk.com/hc/en-us/articles/24251489107213-Advanced-Billing-Subscription-Imports#date-format) documentation for more information about Date/Time Formats.
    InitialBillingAt                  *time.Time                    `json:"initial_billing_at,omitempty"`
    // (Optional) Set this attribute to true to create the subscription in the Awaiting Signup Date state. Use this when you want to create a subscription that has an unknown first  billing date. When the first billing date is known, update a subscription and set the `initial_billing_at` date. The subscription moves to the Awaiting Signup state with a scheduled initial billing date. You can omit the initial_billing_at date to activate the subscription immediately. See [Subscription States](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404222005773-Subscription-States) for more information.
    DeferSignup                       *bool                         `json:"defer_signup,omitempty"`
    // For European sites subject to PSD2 and using 3D Secure, this can be used to reference a previous transaction for the customer. This will ensure the card will be charged successfully at renewal.
    StoredCredentialTransactionId     *int                          `json:"stored_credential_transaction_id,omitempty"`
    SalesRepId                        *int                          `json:"sales_rep_id,omitempty"`
    // The Payment Profile ID of an existing card or bank account, which belongs to an existing customer to use for payment for this subscription. If the card, bank account, or customer does not exist already, or if you want to use a new (unstored) card or bank account for the subscription, use `payment_profile_attributes` instead to create a new payment profile along with the subscription. (This value is available on an existing subscription via the API as `credit_card` > id or `bank_account` > id)
    PaymentProfileId                  *int                          `json:"payment_profile_id,omitempty"`
    // The reference value (provided by your app) for the subscription itself.
    Reference                         *string                       `json:"reference,omitempty"`
    CustomerAttributes                *CustomerAttributes           `json:"customer_attributes,omitempty"`
    // alias to credit_card_attributes
    PaymentProfileAttributes          *PaymentProfileAttributes     `json:"payment_profile_attributes,omitempty"`
    // Credit Card data to create a new Subscription. Interchangeable with `payment_profile_attributes` property.
    CreditCardAttributes              *PaymentProfileAttributes     `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes             *BankAccountAttributes        `json:"bank_account_attributes,omitempty"`
    // (Optional) An array of component ids and quantities to be added to the subscription. See [Components](https://maxio.zendesk.com/hc/en-us/articles/24261141522189-Components-Overview) for more information.
    Components                        []CreateSubscriptionComponent `json:"components,omitempty"`
    // (Optional). Cannot be used when also specifying next_billing_at
    CalendarBilling                   *CalendarBilling              `json:"calendar_billing,omitempty"`
    // (Optional) A set of key/value pairs representing custom fields and their values. Metafields will be created “on-the-fly” in your site for a given key, if they have not been created yet.
    Metafields                        map[string]string             `json:"metafields,omitempty"`
    // The reference value (provided by your app) of an existing customer within Chargify. Required, unless a `customer_id` or a set of `customer_attributes` is given.
    CustomerReference                 *string                       `json:"customer_reference,omitempty"`
    Group                             *GroupSettings                `json:"group,omitempty"`
    // A valid referral code. (optional, see [Referrals](https://maxio.zendesk.com/hc/en-us/articles/24286981223693-Referrals-Reference#how-to-obtain-referral-codes) for more details). If supplied, must be valid, or else subscription creation will fail.
    Ref                               *string                       `json:"ref,omitempty"`
    // (Optional) Can be used when canceling a subscription (via the HTTP DELETE method) to make a note about the reason for cancellation.
    CancellationMessage               *string                       `json:"cancellation_message,omitempty"`
    // (Optional) Can be used when canceling a subscription (via the HTTP DELETE method) to make a note about how the subscription was canceled.
    CancellationMethod                *string                       `json:"cancellation_method,omitempty"`
    // (Optional) If Multi-Currency is enabled and the currency is configured in Chargify, pass it at signup to create a subscription on a non-default currency. Note that you cannot update the currency of an existing subscription.
    Currency                          *string                       `json:"currency,omitempty"`
    // Timestamp giving the expiration date of this subscription (if any). You may manually change the expiration date at any point during a subscription period.
    ExpiresAt                         *time.Time                    `json:"expires_at,omitempty"`
    // (Optional, default false) When set to true, and when next_billing_at is present, if the subscription expires, the expires_at will be shifted by the same amount of time as the difference between the old and new “next billing” dates.
    ExpirationTracksNextBillingChange *string                       `json:"expiration_tracks_next_billing_change,omitempty"`
    // (Optional) The ACH authorization agreement terms. If enabled, an email will be sent to the customer with a copy of the terms.
    AgreementTerms                    *string                       `json:"agreement_terms,omitempty"`
    // (Optional) The first name of the person authorizing the ACH agreement.
    AuthorizerFirstName               *string                       `json:"authorizer_first_name,omitempty"`
    // (Optional) The last name of the person authorizing the ACH agreement.
    AuthorizerLastName                *string                       `json:"authorizer_last_name,omitempty"`
    // (Optional) One of “prorated” (the default – the prorated product price will be charged immediately), “immediate” (the full product price will be charged immediately), or “delayed” (the full product price will be charged with the first scheduled renewal).
    CalendarBillingFirstCharge        *string                       `json:"calendar_billing_first_charge,omitempty"`
    // (Optional) Can be used when canceling a subscription (via the HTTP DELETE method) to indicate why a subscription was canceled.
    ReasonCode                        *string                       `json:"reason_code,omitempty"`
    // (Optional) used only for Delayed Product Change When set to true, indicates that a changed value for product_handle should schedule the product change to the next subscription renewal.
    ProductChangeDelayed              *bool                         `json:"product_change_delayed,omitempty"`
    // Use in place of passing product and component information to set up the subscription with an existing offer. May be either the Chargify id of the offer or its handle prefixed with `handle:`.er
    OfferId                           *CreateSubscriptionOfferId    `json:"offer_id,omitempty"`
    PrepaidConfiguration              *UpsertPrepaidConfiguration   `json:"prepaid_configuration,omitempty"`
    // Providing a previous_billing_at that is in the past will set the current_period_starts_at when the subscription is created. It will also set activated_at if not explicitly passed during the subscription import. Can only be used if next_billing_at is also passed. Using this option will allow you to set the period start for the subscription so mid period component allocations have the correct prorated amount.
    PreviousBillingAt                 *time.Time                    `json:"previous_billing_at,omitempty"`
    // Setting this attribute to true will cause the subscription's MRR to be added to your MRR analytics immediately. For this value to be honored, a next_billing_at must be present and set to a future date. This key/value will not be returned in the subscription response body.
    ImportMrr                         *bool                         `json:"import_mrr,omitempty"`
    CanceledAt                        *time.Time                    `json:"canceled_at,omitempty"`
    ActivatedAt                       *time.Time                    `json:"activated_at,omitempty"`
    // Required when creating a subscription with Maxio Payments.
    AgreementAcceptance               *AgreementAcceptance          `json:"agreement_acceptance,omitempty"`
    // (Optional) If passed, the proof of the authorized ACH agreement terms will be persisted.
    AchAgreement                      *ACHAgreement                 `json:"ach_agreement,omitempty"`
    // Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute.
    DunningCommunicationDelayEnabled  *bool                         `json:"dunning_communication_delay_enabled,omitempty"`
    // Time zone for the Dunning Communication Delay feature.
    DunningCommunicationDelayTimeZone Optional[string]              `json:"dunning_communication_delay_time_zone"`
    // Valid only for the Subscription Preview endpoint. When set to `true` it skips calculating taxes for the current and next billing manifests.
    SkipBillingManifestTaxes          *bool                         `json:"skip_billing_manifest_taxes,omitempty"`
    AdditionalProperties              map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for CreateSubscription,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CreateSubscription) String() string {
    return fmt.Sprintf(
    	"CreateSubscription[ProductHandle=%v, ProductId=%v, ProductPricePointHandle=%v, ProductPricePointId=%v, CustomPrice=%v, CouponCode=%v, CouponCodes=%v, PaymentCollectionMethod=%v, ReceivesInvoiceEmails=%v, NetTerms=%v, CustomerId=%v, NextBillingAt=%v, InitialBillingAt=%v, DeferSignup=%v, StoredCredentialTransactionId=%v, SalesRepId=%v, PaymentProfileId=%v, Reference=%v, CustomerAttributes=%v, PaymentProfileAttributes=%v, CreditCardAttributes=%v, BankAccountAttributes=%v, Components=%v, CalendarBilling=%v, Metafields=%v, CustomerReference=%v, Group=%v, Ref=%v, CancellationMessage=%v, CancellationMethod=%v, Currency=%v, ExpiresAt=%v, ExpirationTracksNextBillingChange=%v, AgreementTerms=%v, AuthorizerFirstName=%v, AuthorizerLastName=%v, CalendarBillingFirstCharge=%v, ReasonCode=%v, ProductChangeDelayed=%v, OfferId=%v, PrepaidConfiguration=%v, PreviousBillingAt=%v, ImportMrr=%v, CanceledAt=%v, ActivatedAt=%v, AgreementAcceptance=%v, AchAgreement=%v, DunningCommunicationDelayEnabled=%v, DunningCommunicationDelayTimeZone=%v, SkipBillingManifestTaxes=%v, AdditionalProperties=%v]",
    	c.ProductHandle, c.ProductId, c.ProductPricePointHandle, c.ProductPricePointId, c.CustomPrice, c.CouponCode, c.CouponCodes, c.PaymentCollectionMethod, c.ReceivesInvoiceEmails, c.NetTerms, c.CustomerId, c.NextBillingAt, c.InitialBillingAt, c.DeferSignup, c.StoredCredentialTransactionId, c.SalesRepId, c.PaymentProfileId, c.Reference, c.CustomerAttributes, c.PaymentProfileAttributes, c.CreditCardAttributes, c.BankAccountAttributes, c.Components, c.CalendarBilling, c.Metafields, c.CustomerReference, c.Group, c.Ref, c.CancellationMessage, c.CancellationMethod, c.Currency, c.ExpiresAt, c.ExpirationTracksNextBillingChange, c.AgreementTerms, c.AuthorizerFirstName, c.AuthorizerLastName, c.CalendarBillingFirstCharge, c.ReasonCode, c.ProductChangeDelayed, c.OfferId, c.PrepaidConfiguration, c.PreviousBillingAt, c.ImportMrr, c.CanceledAt, c.ActivatedAt, c.AgreementAcceptance, c.AchAgreement, c.DunningCommunicationDelayEnabled, c.DunningCommunicationDelayTimeZone, c.SkipBillingManifestTaxes, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CreateSubscription.
// It customizes the JSON marshaling process for CreateSubscription objects.
func (c CreateSubscription) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "product_handle", "product_id", "product_price_point_handle", "product_price_point_id", "custom_price", "coupon_code", "coupon_codes", "payment_collection_method", "receives_invoice_emails", "net_terms", "customer_id", "next_billing_at", "initial_billing_at", "defer_signup", "stored_credential_transaction_id", "sales_rep_id", "payment_profile_id", "reference", "customer_attributes", "payment_profile_attributes", "credit_card_attributes", "bank_account_attributes", "components", "calendar_billing", "metafields", "customer_reference", "group", "ref", "cancellation_message", "cancellation_method", "currency", "expires_at", "expiration_tracks_next_billing_change", "agreement_terms", "authorizer_first_name", "authorizer_last_name", "calendar_billing_first_charge", "reason_code", "product_change_delayed", "offer_id", "prepaid_configuration", "previous_billing_at", "import_mrr", "canceled_at", "activated_at", "agreement_acceptance", "ach_agreement", "dunning_communication_delay_enabled", "dunning_communication_delay_time_zone", "skip_billing_manifest_taxes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CreateSubscription object to a map representation for JSON marshaling.
func (c CreateSubscription) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.ProductHandle != nil {
        structMap["product_handle"] = c.ProductHandle
    }
    if c.ProductId != nil {
        structMap["product_id"] = c.ProductId
    }
    if c.ProductPricePointHandle != nil {
        structMap["product_price_point_handle"] = c.ProductPricePointHandle
    }
    if c.ProductPricePointId != nil {
        structMap["product_price_point_id"] = c.ProductPricePointId
    }
    if c.CustomPrice != nil {
        structMap["custom_price"] = c.CustomPrice.toMap()
    }
    if c.CouponCode != nil {
        structMap["coupon_code"] = c.CouponCode
    }
    if c.CouponCodes != nil {
        structMap["coupon_codes"] = c.CouponCodes
    }
    if c.PaymentCollectionMethod != nil {
        structMap["payment_collection_method"] = c.PaymentCollectionMethod
    }
    if c.ReceivesInvoiceEmails != nil {
        structMap["receives_invoice_emails"] = c.ReceivesInvoiceEmails
    }
    if c.NetTerms != nil {
        structMap["net_terms"] = c.NetTerms
    }
    if c.CustomerId != nil {
        structMap["customer_id"] = c.CustomerId
    }
    if c.NextBillingAt != nil {
        structMap["next_billing_at"] = c.NextBillingAt.Format(time.RFC3339)
    }
    if c.InitialBillingAt != nil {
        structMap["initial_billing_at"] = c.InitialBillingAt.Format(time.RFC3339)
    }
    if c.DeferSignup != nil {
        structMap["defer_signup"] = c.DeferSignup
    }
    if c.StoredCredentialTransactionId != nil {
        structMap["stored_credential_transaction_id"] = c.StoredCredentialTransactionId
    }
    if c.SalesRepId != nil {
        structMap["sales_rep_id"] = c.SalesRepId
    }
    if c.PaymentProfileId != nil {
        structMap["payment_profile_id"] = c.PaymentProfileId
    }
    if c.Reference != nil {
        structMap["reference"] = c.Reference
    }
    if c.CustomerAttributes != nil {
        structMap["customer_attributes"] = c.CustomerAttributes.toMap()
    }
    if c.PaymentProfileAttributes != nil {
        structMap["payment_profile_attributes"] = c.PaymentProfileAttributes.toMap()
    }
    if c.CreditCardAttributes != nil {
        structMap["credit_card_attributes"] = c.CreditCardAttributes.toMap()
    }
    if c.BankAccountAttributes != nil {
        structMap["bank_account_attributes"] = c.BankAccountAttributes.toMap()
    }
    if c.Components != nil {
        structMap["components"] = c.Components
    }
    if c.CalendarBilling != nil {
        structMap["calendar_billing"] = c.CalendarBilling.toMap()
    }
    if c.Metafields != nil {
        structMap["metafields"] = c.Metafields
    }
    if c.CustomerReference != nil {
        structMap["customer_reference"] = c.CustomerReference
    }
    if c.Group != nil {
        structMap["group"] = c.Group.toMap()
    }
    if c.Ref != nil {
        structMap["ref"] = c.Ref
    }
    if c.CancellationMessage != nil {
        structMap["cancellation_message"] = c.CancellationMessage
    }
    if c.CancellationMethod != nil {
        structMap["cancellation_method"] = c.CancellationMethod
    }
    if c.Currency != nil {
        structMap["currency"] = c.Currency
    }
    if c.ExpiresAt != nil {
        structMap["expires_at"] = c.ExpiresAt.Format(time.RFC3339)
    }
    if c.ExpirationTracksNextBillingChange != nil {
        structMap["expiration_tracks_next_billing_change"] = c.ExpirationTracksNextBillingChange
    }
    if c.AgreementTerms != nil {
        structMap["agreement_terms"] = c.AgreementTerms
    }
    if c.AuthorizerFirstName != nil {
        structMap["authorizer_first_name"] = c.AuthorizerFirstName
    }
    if c.AuthorizerLastName != nil {
        structMap["authorizer_last_name"] = c.AuthorizerLastName
    }
    if c.CalendarBillingFirstCharge != nil {
        structMap["calendar_billing_first_charge"] = c.CalendarBillingFirstCharge
    }
    if c.ReasonCode != nil {
        structMap["reason_code"] = c.ReasonCode
    }
    if c.ProductChangeDelayed != nil {
        structMap["product_change_delayed"] = c.ProductChangeDelayed
    }
    if c.OfferId != nil {
        structMap["offer_id"] = c.OfferId.toMap()
    }
    if c.PrepaidConfiguration != nil {
        structMap["prepaid_configuration"] = c.PrepaidConfiguration.toMap()
    }
    if c.PreviousBillingAt != nil {
        structMap["previous_billing_at"] = c.PreviousBillingAt.Format(time.RFC3339)
    }
    if c.ImportMrr != nil {
        structMap["import_mrr"] = c.ImportMrr
    }
    if c.CanceledAt != nil {
        structMap["canceled_at"] = c.CanceledAt.Format(time.RFC3339)
    }
    if c.ActivatedAt != nil {
        structMap["activated_at"] = c.ActivatedAt.Format(time.RFC3339)
    }
    if c.AgreementAcceptance != nil {
        structMap["agreement_acceptance"] = c.AgreementAcceptance.toMap()
    }
    if c.AchAgreement != nil {
        structMap["ach_agreement"] = c.AchAgreement.toMap()
    }
    if c.DunningCommunicationDelayEnabled != nil {
        structMap["dunning_communication_delay_enabled"] = c.DunningCommunicationDelayEnabled
    }
    if c.DunningCommunicationDelayTimeZone.IsValueSet() {
        if c.DunningCommunicationDelayTimeZone.Value() != nil {
            structMap["dunning_communication_delay_time_zone"] = c.DunningCommunicationDelayTimeZone.Value()
        } else {
            structMap["dunning_communication_delay_time_zone"] = nil
        }
    }
    if c.SkipBillingManifestTaxes != nil {
        structMap["skip_billing_manifest_taxes"] = c.SkipBillingManifestTaxes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSubscription.
// It customizes the JSON unmarshaling process for CreateSubscription objects.
func (c *CreateSubscription) UnmarshalJSON(input []byte) error {
    var temp tempCreateSubscription
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "product_handle", "product_id", "product_price_point_handle", "product_price_point_id", "custom_price", "coupon_code", "coupon_codes", "payment_collection_method", "receives_invoice_emails", "net_terms", "customer_id", "next_billing_at", "initial_billing_at", "defer_signup", "stored_credential_transaction_id", "sales_rep_id", "payment_profile_id", "reference", "customer_attributes", "payment_profile_attributes", "credit_card_attributes", "bank_account_attributes", "components", "calendar_billing", "metafields", "customer_reference", "group", "ref", "cancellation_message", "cancellation_method", "currency", "expires_at", "expiration_tracks_next_billing_change", "agreement_terms", "authorizer_first_name", "authorizer_last_name", "calendar_billing_first_charge", "reason_code", "product_change_delayed", "offer_id", "prepaid_configuration", "previous_billing_at", "import_mrr", "canceled_at", "activated_at", "agreement_acceptance", "ach_agreement", "dunning_communication_delay_enabled", "dunning_communication_delay_time_zone", "skip_billing_manifest_taxes")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.ProductHandle = temp.ProductHandle
    c.ProductId = temp.ProductId
    c.ProductPricePointHandle = temp.ProductPricePointHandle
    c.ProductPricePointId = temp.ProductPricePointId
    c.CustomPrice = temp.CustomPrice
    c.CouponCode = temp.CouponCode
    c.CouponCodes = temp.CouponCodes
    c.PaymentCollectionMethod = temp.PaymentCollectionMethod
    c.ReceivesInvoiceEmails = temp.ReceivesInvoiceEmails
    c.NetTerms = temp.NetTerms
    c.CustomerId = temp.CustomerId
    if temp.NextBillingAt != nil {
        NextBillingAtVal, err := time.Parse(time.RFC3339, *temp.NextBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse next_billing_at as % s format.", time.RFC3339)
        }
        c.NextBillingAt = &NextBillingAtVal
    }
    if temp.InitialBillingAt != nil {
        InitialBillingAtVal, err := time.Parse(time.RFC3339, *temp.InitialBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse initial_billing_at as % s format.", time.RFC3339)
        }
        c.InitialBillingAt = &InitialBillingAtVal
    }
    c.DeferSignup = temp.DeferSignup
    c.StoredCredentialTransactionId = temp.StoredCredentialTransactionId
    c.SalesRepId = temp.SalesRepId
    c.PaymentProfileId = temp.PaymentProfileId
    c.Reference = temp.Reference
    c.CustomerAttributes = temp.CustomerAttributes
    c.PaymentProfileAttributes = temp.PaymentProfileAttributes
    c.CreditCardAttributes = temp.CreditCardAttributes
    c.BankAccountAttributes = temp.BankAccountAttributes
    c.Components = temp.Components
    c.CalendarBilling = temp.CalendarBilling
    c.Metafields = temp.Metafields
    c.CustomerReference = temp.CustomerReference
    c.Group = temp.Group
    c.Ref = temp.Ref
    c.CancellationMessage = temp.CancellationMessage
    c.CancellationMethod = temp.CancellationMethod
    c.Currency = temp.Currency
    if temp.ExpiresAt != nil {
        ExpiresAtVal, err := time.Parse(time.RFC3339, *temp.ExpiresAt)
        if err != nil {
            log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
        }
        c.ExpiresAt = &ExpiresAtVal
    }
    c.ExpirationTracksNextBillingChange = temp.ExpirationTracksNextBillingChange
    c.AgreementTerms = temp.AgreementTerms
    c.AuthorizerFirstName = temp.AuthorizerFirstName
    c.AuthorizerLastName = temp.AuthorizerLastName
    c.CalendarBillingFirstCharge = temp.CalendarBillingFirstCharge
    c.ReasonCode = temp.ReasonCode
    c.ProductChangeDelayed = temp.ProductChangeDelayed
    c.OfferId = temp.OfferId
    c.PrepaidConfiguration = temp.PrepaidConfiguration
    if temp.PreviousBillingAt != nil {
        PreviousBillingAtVal, err := time.Parse(time.RFC3339, *temp.PreviousBillingAt)
        if err != nil {
            log.Fatalf("Cannot Parse previous_billing_at as % s format.", time.RFC3339)
        }
        c.PreviousBillingAt = &PreviousBillingAtVal
    }
    c.ImportMrr = temp.ImportMrr
    if temp.CanceledAt != nil {
        CanceledAtVal, err := time.Parse(time.RFC3339, *temp.CanceledAt)
        if err != nil {
            log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
        }
        c.CanceledAt = &CanceledAtVal
    }
    if temp.ActivatedAt != nil {
        ActivatedAtVal, err := time.Parse(time.RFC3339, *temp.ActivatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse activated_at as % s format.", time.RFC3339)
        }
        c.ActivatedAt = &ActivatedAtVal
    }
    c.AgreementAcceptance = temp.AgreementAcceptance
    c.AchAgreement = temp.AchAgreement
    c.DunningCommunicationDelayEnabled = temp.DunningCommunicationDelayEnabled
    c.DunningCommunicationDelayTimeZone = temp.DunningCommunicationDelayTimeZone
    c.SkipBillingManifestTaxes = temp.SkipBillingManifestTaxes
    return nil
}

// tempCreateSubscription is a temporary struct used for validating the fields of CreateSubscription.
type tempCreateSubscription  struct {
    ProductHandle                     *string                       `json:"product_handle,omitempty"`
    ProductId                         *int                          `json:"product_id,omitempty"`
    ProductPricePointHandle           *string                       `json:"product_price_point_handle,omitempty"`
    ProductPricePointId               *int                          `json:"product_price_point_id,omitempty"`
    CustomPrice                       *SubscriptionCustomPrice      `json:"custom_price,omitempty"`
    CouponCode                        *string                       `json:"coupon_code,omitempty"`
    CouponCodes                       []string                      `json:"coupon_codes,omitempty"`
    PaymentCollectionMethod           *CollectionMethod             `json:"payment_collection_method,omitempty"`
    ReceivesInvoiceEmails             *string                       `json:"receives_invoice_emails,omitempty"`
    NetTerms                          *string                       `json:"net_terms,omitempty"`
    CustomerId                        *int                          `json:"customer_id,omitempty"`
    NextBillingAt                     *string                       `json:"next_billing_at,omitempty"`
    InitialBillingAt                  *string                       `json:"initial_billing_at,omitempty"`
    DeferSignup                       *bool                         `json:"defer_signup,omitempty"`
    StoredCredentialTransactionId     *int                          `json:"stored_credential_transaction_id,omitempty"`
    SalesRepId                        *int                          `json:"sales_rep_id,omitempty"`
    PaymentProfileId                  *int                          `json:"payment_profile_id,omitempty"`
    Reference                         *string                       `json:"reference,omitempty"`
    CustomerAttributes                *CustomerAttributes           `json:"customer_attributes,omitempty"`
    PaymentProfileAttributes          *PaymentProfileAttributes     `json:"payment_profile_attributes,omitempty"`
    CreditCardAttributes              *PaymentProfileAttributes     `json:"credit_card_attributes,omitempty"`
    BankAccountAttributes             *BankAccountAttributes        `json:"bank_account_attributes,omitempty"`
    Components                        []CreateSubscriptionComponent `json:"components,omitempty"`
    CalendarBilling                   *CalendarBilling              `json:"calendar_billing,omitempty"`
    Metafields                        map[string]string             `json:"metafields,omitempty"`
    CustomerReference                 *string                       `json:"customer_reference,omitempty"`
    Group                             *GroupSettings                `json:"group,omitempty"`
    Ref                               *string                       `json:"ref,omitempty"`
    CancellationMessage               *string                       `json:"cancellation_message,omitempty"`
    CancellationMethod                *string                       `json:"cancellation_method,omitempty"`
    Currency                          *string                       `json:"currency,omitempty"`
    ExpiresAt                         *string                       `json:"expires_at,omitempty"`
    ExpirationTracksNextBillingChange *string                       `json:"expiration_tracks_next_billing_change,omitempty"`
    AgreementTerms                    *string                       `json:"agreement_terms,omitempty"`
    AuthorizerFirstName               *string                       `json:"authorizer_first_name,omitempty"`
    AuthorizerLastName                *string                       `json:"authorizer_last_name,omitempty"`
    CalendarBillingFirstCharge        *string                       `json:"calendar_billing_first_charge,omitempty"`
    ReasonCode                        *string                       `json:"reason_code,omitempty"`
    ProductChangeDelayed              *bool                         `json:"product_change_delayed,omitempty"`
    OfferId                           *CreateSubscriptionOfferId    `json:"offer_id,omitempty"`
    PrepaidConfiguration              *UpsertPrepaidConfiguration   `json:"prepaid_configuration,omitempty"`
    PreviousBillingAt                 *string                       `json:"previous_billing_at,omitempty"`
    ImportMrr                         *bool                         `json:"import_mrr,omitempty"`
    CanceledAt                        *string                       `json:"canceled_at,omitempty"`
    ActivatedAt                       *string                       `json:"activated_at,omitempty"`
    AgreementAcceptance               *AgreementAcceptance          `json:"agreement_acceptance,omitempty"`
    AchAgreement                      *ACHAgreement                 `json:"ach_agreement,omitempty"`
    DunningCommunicationDelayEnabled  *bool                         `json:"dunning_communication_delay_enabled,omitempty"`
    DunningCommunicationDelayTimeZone Optional[string]              `json:"dunning_communication_delay_time_zone"`
    SkipBillingManifestTaxes          *bool                         `json:"skip_billing_manifest_taxes,omitempty"`
}
