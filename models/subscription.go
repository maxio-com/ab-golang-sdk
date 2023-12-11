package models

import (
	"encoding/json"
	"log"
	"time"
)

// Subscription represents a Subscription struct.
type Subscription struct {
	// The subscription unique id within Chargify.
	Id *int `json:"id,omitempty"`
	// The state of a subscription.
	// * **Live States**
	//     * `active` - A normal, active subscription. It is not in a trial and is paid and up to date.
	//     * `assessing` - An internal (transient) state that indicates a subscription is in the middle of periodic assessment. Do not base any access decisions in your app on this state, as it may not always be exposed.
	//     * `pending` - An internal (transient) state that indicates a subscription is in the creation process. Do not base any access decisions in your app on this state, as it may not always be exposed.
	//     * `trialing` - A subscription in trialing state has a valid trial subscription. This type of subscription may transition to active once payment is received when the trial has ended. Otherwise, it may go to a Problem or End of Life state.
	//     * `paused` - An internal state that indicates that your account with Advanced Billing is in arrears.
	// * **Problem States**
	//     * `past_due` - Indicates that the most recent payment has failed, and payment is past due for this subscription. If you have enabled our automated dunning, this subscription will be in the dunning process (additional status and callbacks from the dunning process will be available in the future). If you are handling dunning and payment updates yourself, you will want to use this state to initiate a payment update from your customers.
	//     * `soft_failure` - Indicates that normal assessment/processing of the subscription has failed for a reason that cannot be fixed by the Customer. For example, a Soft Fail may result from a timeout at the gateway or incorrect credentials on your part. The subscriptions should be retried automatically. An interface is being built for you to review problems resulting from these events to take manual action when needed.
	//     * `unpaid` - Indicates an unpaid subscription. A subscription is marked unpaid if the retry period expires and you have configured your [Dunning](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405505141005) settings to have a Final Action of `mark the subscription unpaid`.
	// * **End of Life States**
	//     * `canceled` - Indicates a canceled subscription. This may happen at your request (via the API or the web interface) or due to the expiration of the [Dunning](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405505141005) process without payment. See the [Reactivation](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404559291021) documentation for info on how to restart a canceled subscription.
	//     While a subscription is canceled, its period will not advance, it will not accrue any new charges, and Advanced Billing will not attempt to collect the overdue balance.
	//     * `expired` - Indicates a subscription that has expired due to running its normal life cycle. Some products may be configured to have an expiration period. An expired subscription then is one that stayed active until it fulfilled its full period.
	//     * `failed_to_create` - Indicates that signup has failed. (You may see this state in a signup_failure webhook.)
	//     * `on_hold` - Indicates that a subscription’s billing has been temporarily stopped. While it is expected that the subscription will resume and return to active status, this is still treated as an “End of Life” state because the customer is not paying for services during this time.
	//     * `suspended` - Indicates that a prepaid subscription has used up all their prepayment balance. If a prepayment is applied, it will return to an active state.
	//     * `trial_ended` - A subscription in a trial_ended state is a subscription that completed a no-obligation trial and did not have a card on file at the expiration of the trial period. See [Product Pricing – No Obligation Trials](https://maxio-chargify.zendesk.com/hc/en-us/articles/5405246782221) for more details.
	// See [Subscription States](https://maxio-chargify.zendesk.com/hc/en-us/articles/5404222005773) for more info about subscription states and state transitions.
	State *SubscriptionStateEnum `json:"state,omitempty"`
	// Gives the current outstanding subscription balance in the number of cents.
	BalanceInCents *int64 `json:"balance_in_cents,omitempty"`
	// Gives the total revenue from the subscription in the number of cents.
	TotalRevenueInCents *int64 `json:"total_revenue_in_cents,omitempty"`
	// (Added Nov 5 2013) The recurring amount of the product (and version),currently subscribed. NOTE: this may differ from the current price of,the product, if you’ve changed the price of the product but haven’t,moved this subscription to a newer version.
	ProductPriceInCents *int64 `json:"product_price_in_cents,omitempty"`
	// The version of the product for the subscription. Note that this is a deprecated field kept for backwards-compatibility.
	ProductVersionNumber *int `json:"product_version_number,omitempty"`
	// Timestamp relating to the end of the current (recurring) period (i.e.,when the next regularly scheduled attempted charge will occur)
	CurrentPeriodEndsAt *time.Time `json:"current_period_ends_at,omitempty"`
	// Timestamp that indicates when capture of payment will be tried or,retried. This value will usually track the current_period_ends_at, but,will diverge if a renewal payment fails and must be retried. In that,case, the current_period_ends_at will advance to the end of the next,period (time doesn’t stop because a payment was missed) but the,next_assessment_at will be scheduled for the auto-retry time (i.e. 24,hours in the future, in some cases)
	NextAssessmentAt *time.Time `json:"next_assessment_at,omitempty"`
	// Timestamp for when the trial period (if any) began
	TrialStartedAt Optional[time.Time] `json:"trial_started_at"`
	// Timestamp for when the trial period (if any) ended
	TrialEndedAt Optional[time.Time] `json:"trial_ended_at"`
	// Timestamp for when the subscription began (i.e. when it came out of trial, or when it began in the case of no trial)
	ActivatedAt *time.Time `json:"activated_at,omitempty"`
	// Timestamp giving the expiration date of this subscription (if any)
	ExpiresAt Optional[time.Time] `json:"expires_at"`
	// The creation date for this subscription
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date of last update for this subscription
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// Seller-provided reason for, or note about, the cancellation.
	CancellationMessage Optional[string] `json:"cancellation_message"`
	// The process used to cancel the subscription, if the subscription has been canceled. It is nil if the subscription's state is not canceled.
	CancellationMethod Optional[CancellationMethodEnum] `json:"cancellation_method"`
	// Whether or not the subscription will (or has) canceled at the end of the period.
	CancelAtEndOfPeriod Optional[bool] `json:"cancel_at_end_of_period"`
	// The timestamp of the most recent cancellation
	CanceledAt Optional[time.Time] `json:"canceled_at"`
	// Timestamp relating to the start of the current (recurring) period
	CurrentPeriodStartedAt *time.Time `json:"current_period_started_at,omitempty"`
	// Only valid for webhook payloads The previous state for webhooks that have indicated a change in state. For normal API calls, this will always be the same as the state (current state)
	PreviousState *SubscriptionStateEnum `json:"previous_state,omitempty"`
	// The ID of the transaction that generated the revenue
	SignupPaymentId *int `json:"signup_payment_id,omitempty"`
	// The revenue, formatted as a string of decimal separated dollars and,cents, from the subscription signup ($50.00 would be formatted as,50.00)
	SignupRevenue *string `json:"signup_revenue,omitempty"`
	// Timestamp for when the subscription is currently set to cancel.
	DelayedCancelAt Optional[time.Time] `json:"delayed_cancel_at"`
	// (deprecated) The coupon code of the single coupon currently applied to the subscription. See coupon_codes instead as subscriptions can now have more than one coupon.
	CouponCode Optional[string] `json:"coupon_code"` // Deprecated
	// The day of the month that the subscription will charge according to calendar billing rules, if used.
	SnapDay Optional[string] `json:"snap_day"`
	// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
	PaymentCollectionMethod *PaymentCollectionMethodEnum      `json:"payment_collection_method,omitempty"`
	Customer                *Customer                         `json:"customer,omitempty"`
	Product                 *Product                          `json:"product,omitempty"`
	CreditCard              *PaymentProfile                   `json:"credit_card,omitempty"`
	Group                   Optional[NestedSubscriptionGroup] `json:"group"`
	BankAccount             *SubscriptionBankAccount          `json:"bank_account,omitempty"`
	// The payment profile type for the active profile on file.
	PaymentType Optional[string] `json:"payment_type"`
	// The subscription's unique code that can be given to referrals.
	ReferralCode Optional[string] `json:"referral_code"`
	// If a delayed product change is scheduled, the ID of the product that the subscription will be changed to at the next renewal.
	NextProductId Optional[int] `json:"next_product_id"`
	// If a delayed product change is scheduled, the handle of the product that the subscription will be changed to at the next renewal.
	NextProductHandle Optional[string] `json:"next_product_handle"`
	// (deprecated) How many times the subscription's single coupon has been used. This field has no replacement for multiple coupons.
	CouponUseCount Optional[int] `json:"coupon_use_count"` // Deprecated
	// (deprecated) How many times the subscription's single coupon may be used. This field has no replacement for multiple coupons.
	CouponUsesAllowed Optional[int] `json:"coupon_uses_allowed"` // Deprecated
	// If the subscription is canceled, this is their churn code.
	ReasonCode Optional[string] `json:"reason_code"`
	// The date the subscription is scheduled to automatically resume from the on_hold state.
	AutomaticallyResumeAt Optional[time.Time] `json:"automatically_resume_at"`
	// An array for all the coupons attached to the subscription.
	CouponCodes []string `json:"coupon_codes,omitempty"`
	// The ID of the offer associated with the subscription.
	OfferId Optional[int] `json:"offer_id"`
	// On Relationship Invoicing, the ID of the individual paying for the subscription. Defaults to the Customer ID unless the 'Customer Hierarchies & WhoPays' feature is enabled.
	PayerId Optional[int] `json:"payer_id"`
	// The balance in cents plus the estimated renewal amount in cents.
	CurrentBillingAmountInCents *int64 `json:"current_billing_amount_in_cents,omitempty"`
	// The product price point currently subscribed to.
	ProductPricePointId *int `json:"product_price_point_id,omitempty"`
	// Price point type. We expose the following types:
	// 1. **default**: a price point that is marked as a default price for a certain product.
	// 2. **custom**: a custom price point.
	// 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
	ProductPricePointType *PricePointTypeEnum `json:"product_price_point_type,omitempty"`
	// If a delayed product change is scheduled, the ID of the product price point that the subscription will be changed to at the next renewal.
	NextProductPricePointId Optional[int] `json:"next_product_price_point_id"`
	// On Relationship Invoicing, the number of days before a renewal invoice is due.
	NetTerms Optional[int] `json:"net_terms"`
	// For European sites subject to PSD2 and using 3D Secure, this can be used to reference a previous transaction for the customer. This will ensure the card will be charged successfully at renewal.
	StoredCredentialTransactionId Optional[int] `json:"stored_credential_transaction_id"`
	// The reference value (provided by your app) for the subscription itelf.
	Reference Optional[string] `json:"reference"`
	// The timestamp of the most recent on hold action.
	OnHoldAt Optional[time.Time] `json:"on_hold_at"`
	// Boolean representing whether the subscription is prepaid and currently in dunning. Only returned for Relationship Invoicing sites with the feature enabled
	PrepaidDunning *bool `json:"prepaid_dunning,omitempty"`
	// Additional coupon data. To use this data you also have to include the following param in the request`include[]=coupons`.
	// Only in Read Subscription Endpoint.
	Coupons []SubscriptionIncludedCoupon `json:"coupons,omitempty"`
	// Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute.
	DunningCommunicationDelayEnabled *bool `json:"dunning_communication_delay_enabled,omitempty"`
	// Time zone for the Dunning Communication Delay feature.
	DunningCommunicationDelayTimeZone Optional[string]      `json:"dunning_communication_delay_time_zone"`
	ReceivesInvoiceEmails             Optional[bool]        `json:"receives_invoice_emails"`
	Locale                            Optional[string]      `json:"locale"`
	Currency                          *string               `json:"currency,omitempty"`
	ScheduledCancellationAt           Optional[time.Time]   `json:"scheduled_cancellation_at"`
	CreditBalanceInCents              *int64                `json:"credit_balance_in_cents,omitempty"`
	PrepaymentBalanceInCents          *int64                `json:"prepayment_balance_in_cents,omitempty"`
	PrepaidConfiguration              *PrepaidConfiguration `json:"prepaid_configuration,omitempty"`
}

// MarshalJSON implements the json.Marshaler interface for Subscription.
// It customizes the JSON marshaling process for Subscription objects.
func (s *Subscription) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the Subscription object to a map representation for JSON marshaling.
func (s *Subscription) toMap() map[string]any {
	structMap := make(map[string]any)
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.State != nil {
		structMap["state"] = s.State
	}
	if s.BalanceInCents != nil {
		structMap["balance_in_cents"] = s.BalanceInCents
	}
	if s.TotalRevenueInCents != nil {
		structMap["total_revenue_in_cents"] = s.TotalRevenueInCents
	}
	if s.ProductPriceInCents != nil {
		structMap["product_price_in_cents"] = s.ProductPriceInCents
	}
	if s.ProductVersionNumber != nil {
		structMap["product_version_number"] = s.ProductVersionNumber
	}
	if s.CurrentPeriodEndsAt != nil {
		structMap["current_period_ends_at"] = s.CurrentPeriodEndsAt.Format(time.RFC3339)
	}
	if s.NextAssessmentAt != nil {
		structMap["next_assessment_at"] = s.NextAssessmentAt.Format(time.RFC3339)
	}
	if s.TrialStartedAt.IsValueSet() {
		var TrialStartedAtVal *string = nil
		if s.TrialStartedAt.Value() != nil {
			val := s.TrialStartedAt.Value().Format(time.RFC3339)
			TrialStartedAtVal = &val
		}
		structMap["trial_started_at"] = TrialStartedAtVal
	}
	if s.TrialEndedAt.IsValueSet() {
		var TrialEndedAtVal *string = nil
		if s.TrialEndedAt.Value() != nil {
			val := s.TrialEndedAt.Value().Format(time.RFC3339)
			TrialEndedAtVal = &val
		}
		structMap["trial_ended_at"] = TrialEndedAtVal
	}
	if s.ActivatedAt != nil {
		structMap["activated_at"] = s.ActivatedAt.Format(time.RFC3339)
	}
	if s.ExpiresAt.IsValueSet() {
		var ExpiresAtVal *string = nil
		if s.ExpiresAt.Value() != nil {
			val := s.ExpiresAt.Value().Format(time.RFC3339)
			ExpiresAtVal = &val
		}
		structMap["expires_at"] = ExpiresAtVal
	}
	if s.CreatedAt != nil {
		structMap["created_at"] = s.CreatedAt.Format(time.RFC3339)
	}
	if s.UpdatedAt != nil {
		structMap["updated_at"] = s.UpdatedAt.Format(time.RFC3339)
	}
	if s.CancellationMessage.IsValueSet() {
		structMap["cancellation_message"] = s.CancellationMessage.Value()
	}
	if s.CancellationMethod.IsValueSet() {
		structMap["cancellation_method"] = s.CancellationMethod.Value()
	}
	if s.CancelAtEndOfPeriod.IsValueSet() {
		structMap["cancel_at_end_of_period"] = s.CancelAtEndOfPeriod.Value()
	}
	if s.CanceledAt.IsValueSet() {
		var CanceledAtVal *string = nil
		if s.CanceledAt.Value() != nil {
			val := s.CanceledAt.Value().Format(time.RFC3339)
			CanceledAtVal = &val
		}
		structMap["canceled_at"] = CanceledAtVal
	}
	if s.CurrentPeriodStartedAt != nil {
		structMap["current_period_started_at"] = s.CurrentPeriodStartedAt.Format(time.RFC3339)
	}
	if s.PreviousState != nil {
		structMap["previous_state"] = s.PreviousState
	}
	if s.SignupPaymentId != nil {
		structMap["signup_payment_id"] = s.SignupPaymentId
	}
	if s.SignupRevenue != nil {
		structMap["signup_revenue"] = s.SignupRevenue
	}
	if s.DelayedCancelAt.IsValueSet() {
		var DelayedCancelAtVal *string = nil
		if s.DelayedCancelAt.Value() != nil {
			val := s.DelayedCancelAt.Value().Format(time.RFC3339)
			DelayedCancelAtVal = &val
		}
		structMap["delayed_cancel_at"] = DelayedCancelAtVal
	}
	if s.CouponCode.IsValueSet() {
		structMap["coupon_code"] = s.CouponCode.Value()
	}
	if s.SnapDay.IsValueSet() {
		structMap["snap_day"] = s.SnapDay.Value()
	}
	if s.PaymentCollectionMethod != nil {
		structMap["payment_collection_method"] = s.PaymentCollectionMethod
	}
	if s.Customer != nil {
		structMap["customer"] = s.Customer
	}
	if s.Product != nil {
		structMap["product"] = s.Product
	}
	if s.CreditCard != nil {
		structMap["credit_card"] = s.CreditCard
	}
	if s.Group.IsValueSet() {
		structMap["group"] = s.Group.Value()
	}
	if s.BankAccount != nil {
		structMap["bank_account"] = s.BankAccount
	}
	if s.PaymentType.IsValueSet() {
		structMap["payment_type"] = s.PaymentType.Value()
	}
	if s.ReferralCode.IsValueSet() {
		structMap["referral_code"] = s.ReferralCode.Value()
	}
	if s.NextProductId.IsValueSet() {
		structMap["next_product_id"] = s.NextProductId.Value()
	}
	if s.NextProductHandle.IsValueSet() {
		structMap["next_product_handle"] = s.NextProductHandle.Value()
	}
	if s.CouponUseCount.IsValueSet() {
		structMap["coupon_use_count"] = s.CouponUseCount.Value()
	}
	if s.CouponUsesAllowed.IsValueSet() {
		structMap["coupon_uses_allowed"] = s.CouponUsesAllowed.Value()
	}
	if s.ReasonCode.IsValueSet() {
		structMap["reason_code"] = s.ReasonCode.Value()
	}
	if s.AutomaticallyResumeAt.IsValueSet() {
		var AutomaticallyResumeAtVal *string = nil
		if s.AutomaticallyResumeAt.Value() != nil {
			val := s.AutomaticallyResumeAt.Value().Format(time.RFC3339)
			AutomaticallyResumeAtVal = &val
		}
		structMap["automatically_resume_at"] = AutomaticallyResumeAtVal
	}
	if s.CouponCodes != nil {
		structMap["coupon_codes"] = s.CouponCodes
	}
	if s.OfferId.IsValueSet() {
		structMap["offer_id"] = s.OfferId.Value()
	}
	if s.PayerId.IsValueSet() {
		structMap["payer_id"] = s.PayerId.Value()
	}
	if s.CurrentBillingAmountInCents != nil {
		structMap["current_billing_amount_in_cents"] = s.CurrentBillingAmountInCents
	}
	if s.ProductPricePointId != nil {
		structMap["product_price_point_id"] = s.ProductPricePointId
	}
	if s.ProductPricePointType != nil {
		structMap["product_price_point_type"] = s.ProductPricePointType
	}
	if s.NextProductPricePointId.IsValueSet() {
		structMap["next_product_price_point_id"] = s.NextProductPricePointId.Value()
	}
	if s.NetTerms.IsValueSet() {
		structMap["net_terms"] = s.NetTerms.Value()
	}
	if s.StoredCredentialTransactionId.IsValueSet() {
		structMap["stored_credential_transaction_id"] = s.StoredCredentialTransactionId.Value()
	}
	if s.Reference.IsValueSet() {
		structMap["reference"] = s.Reference.Value()
	}
	if s.OnHoldAt.IsValueSet() {
		var OnHoldAtVal *string = nil
		if s.OnHoldAt.Value() != nil {
			val := s.OnHoldAt.Value().Format(time.RFC3339)
			OnHoldAtVal = &val
		}
		structMap["on_hold_at"] = OnHoldAtVal
	}
	if s.PrepaidDunning != nil {
		structMap["prepaid_dunning"] = s.PrepaidDunning
	}
	if s.Coupons != nil {
		structMap["coupons"] = s.Coupons
	}
	if s.DunningCommunicationDelayEnabled != nil {
		structMap["dunning_communication_delay_enabled"] = s.DunningCommunicationDelayEnabled
	}
	if s.DunningCommunicationDelayTimeZone.IsValueSet() {
		structMap["dunning_communication_delay_time_zone"] = s.DunningCommunicationDelayTimeZone.Value()
	}
	if s.ReceivesInvoiceEmails.IsValueSet() {
		structMap["receives_invoice_emails"] = s.ReceivesInvoiceEmails.Value()
	}
	if s.Locale.IsValueSet() {
		structMap["locale"] = s.Locale.Value()
	}
	if s.Currency != nil {
		structMap["currency"] = s.Currency
	}
	if s.ScheduledCancellationAt.IsValueSet() {
		var ScheduledCancellationAtVal *string = nil
		if s.ScheduledCancellationAt.Value() != nil {
			val := s.ScheduledCancellationAt.Value().Format(time.RFC3339)
			ScheduledCancellationAtVal = &val
		}
		structMap["scheduled_cancellation_at"] = ScheduledCancellationAtVal
	}
	if s.CreditBalanceInCents != nil {
		structMap["credit_balance_in_cents"] = s.CreditBalanceInCents
	}
	if s.PrepaymentBalanceInCents != nil {
		structMap["prepayment_balance_in_cents"] = s.PrepaymentBalanceInCents
	}
	if s.PrepaidConfiguration != nil {
		structMap["prepaid_configuration"] = s.PrepaidConfiguration
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Subscription.
// It customizes the JSON unmarshaling process for Subscription objects.
func (s *Subscription) UnmarshalJSON(input []byte) error {
	temp := &struct {
		Id                                *int                              `json:"id,omitempty"`
		State                             *SubscriptionStateEnum            `json:"state,omitempty"`
		BalanceInCents                    *int64                            `json:"balance_in_cents,omitempty"`
		TotalRevenueInCents               *int64                            `json:"total_revenue_in_cents,omitempty"`
		ProductPriceInCents               *int64                            `json:"product_price_in_cents,omitempty"`
		ProductVersionNumber              *int                              `json:"product_version_number,omitempty"`
		CurrentPeriodEndsAt               *string                           `json:"current_period_ends_at,omitempty"`
		NextAssessmentAt                  *string                           `json:"next_assessment_at,omitempty"`
		TrialStartedAt                    Optional[string]                  `json:"trial_started_at"`
		TrialEndedAt                      Optional[string]                  `json:"trial_ended_at"`
		ActivatedAt                       *string                           `json:"activated_at,omitempty"`
		ExpiresAt                         Optional[string]                  `json:"expires_at"`
		CreatedAt                         *string                           `json:"created_at,omitempty"`
		UpdatedAt                         *string                           `json:"updated_at,omitempty"`
		CancellationMessage               Optional[string]                  `json:"cancellation_message"`
		CancellationMethod                Optional[CancellationMethodEnum]  `json:"cancellation_method"`
		CancelAtEndOfPeriod               Optional[bool]                    `json:"cancel_at_end_of_period"`
		CanceledAt                        Optional[string]                  `json:"canceled_at"`
		CurrentPeriodStartedAt            *string                           `json:"current_period_started_at,omitempty"`
		PreviousState                     *SubscriptionStateEnum            `json:"previous_state,omitempty"`
		SignupPaymentId                   *int                              `json:"signup_payment_id,omitempty"`
		SignupRevenue                     *string                           `json:"signup_revenue,omitempty"`
		DelayedCancelAt                   Optional[string]                  `json:"delayed_cancel_at"`
		CouponCode                        Optional[string]                  `json:"coupon_code"`
		SnapDay                           Optional[string]                  `json:"snap_day"`
		PaymentCollectionMethod           *PaymentCollectionMethodEnum      `json:"payment_collection_method,omitempty"`
		Customer                          *Customer                         `json:"customer,omitempty"`
		Product                           *Product                          `json:"product,omitempty"`
		CreditCard                        *PaymentProfile                   `json:"credit_card,omitempty"`
		Group                             Optional[NestedSubscriptionGroup] `json:"group"`
		BankAccount                       *SubscriptionBankAccount          `json:"bank_account,omitempty"`
		PaymentType                       Optional[string]                  `json:"payment_type"`
		ReferralCode                      Optional[string]                  `json:"referral_code"`
		NextProductId                     Optional[int]                     `json:"next_product_id"`
		NextProductHandle                 Optional[string]                  `json:"next_product_handle"`
		CouponUseCount                    Optional[int]                     `json:"coupon_use_count"`
		CouponUsesAllowed                 Optional[int]                     `json:"coupon_uses_allowed"`
		ReasonCode                        Optional[string]                  `json:"reason_code"`
		AutomaticallyResumeAt             Optional[string]                  `json:"automatically_resume_at"`
		CouponCodes                       []string                          `json:"coupon_codes,omitempty"`
		OfferId                           Optional[int]                     `json:"offer_id"`
		PayerId                           Optional[int]                     `json:"payer_id"`
		CurrentBillingAmountInCents       *int64                            `json:"current_billing_amount_in_cents,omitempty"`
		ProductPricePointId               *int                              `json:"product_price_point_id,omitempty"`
		ProductPricePointType             *PricePointTypeEnum               `json:"product_price_point_type,omitempty"`
		NextProductPricePointId           Optional[int]                     `json:"next_product_price_point_id"`
		NetTerms                          Optional[int]                     `json:"net_terms"`
		StoredCredentialTransactionId     Optional[int]                     `json:"stored_credential_transaction_id"`
		Reference                         Optional[string]                  `json:"reference"`
		OnHoldAt                          Optional[string]                  `json:"on_hold_at"`
		PrepaidDunning                    *bool                             `json:"prepaid_dunning,omitempty"`
		Coupons                           []SubscriptionIncludedCoupon      `json:"coupons,omitempty"`
		DunningCommunicationDelayEnabled  *bool                             `json:"dunning_communication_delay_enabled,omitempty"`
		DunningCommunicationDelayTimeZone Optional[string]                  `json:"dunning_communication_delay_time_zone"`
		ReceivesInvoiceEmails             Optional[bool]                    `json:"receives_invoice_emails"`
		Locale                            Optional[string]                  `json:"locale"`
		Currency                          *string                           `json:"currency,omitempty"`
		ScheduledCancellationAt           Optional[string]                  `json:"scheduled_cancellation_at"`
		CreditBalanceInCents              *int64                            `json:"credit_balance_in_cents,omitempty"`
		PrepaymentBalanceInCents          *int64                            `json:"prepayment_balance_in_cents,omitempty"`
		PrepaidConfiguration              *PrepaidConfiguration             `json:"prepaid_configuration,omitempty"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	s.Id = temp.Id
	s.State = temp.State
	s.BalanceInCents = temp.BalanceInCents
	s.TotalRevenueInCents = temp.TotalRevenueInCents
	s.ProductPriceInCents = temp.ProductPriceInCents
	s.ProductVersionNumber = temp.ProductVersionNumber
	if temp.CurrentPeriodEndsAt != nil {
		CurrentPeriodEndsAtVal, err := time.Parse(time.RFC3339, *temp.CurrentPeriodEndsAt)
		if err != nil {
			log.Fatalf("Cannot Parse current_period_ends_at as % s format.", time.RFC3339)
		}
		s.CurrentPeriodEndsAt = &CurrentPeriodEndsAtVal
	}
	if temp.NextAssessmentAt != nil {
		NextAssessmentAtVal, err := time.Parse(time.RFC3339, *temp.NextAssessmentAt)
		if err != nil {
			log.Fatalf("Cannot Parse next_assessment_at as % s format.", time.RFC3339)
		}
		s.NextAssessmentAt = &NextAssessmentAtVal
	}
	s.TrialStartedAt.ShouldSetValue(temp.TrialStartedAt.IsValueSet())
	if temp.TrialStartedAt.Value() != nil {
		TrialStartedAtVal, err := time.Parse(time.RFC3339, (*temp.TrialStartedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse trial_started_at as % s format.", time.RFC3339)
		}
		s.TrialStartedAt.SetValue(&TrialStartedAtVal)
	}
	s.TrialEndedAt.ShouldSetValue(temp.TrialEndedAt.IsValueSet())
	if temp.TrialEndedAt.Value() != nil {
		TrialEndedAtVal, err := time.Parse(time.RFC3339, (*temp.TrialEndedAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse trial_ended_at as % s format.", time.RFC3339)
		}
		s.TrialEndedAt.SetValue(&TrialEndedAtVal)
	}
	if temp.ActivatedAt != nil {
		ActivatedAtVal, err := time.Parse(time.RFC3339, *temp.ActivatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse activated_at as % s format.", time.RFC3339)
		}
		s.ActivatedAt = &ActivatedAtVal
	}
	s.ExpiresAt.ShouldSetValue(temp.ExpiresAt.IsValueSet())
	if temp.ExpiresAt.Value() != nil {
		ExpiresAtVal, err := time.Parse(time.RFC3339, (*temp.ExpiresAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse expires_at as % s format.", time.RFC3339)
		}
		s.ExpiresAt.SetValue(&ExpiresAtVal)
	}
	if temp.CreatedAt != nil {
		CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
		}
		s.CreatedAt = &CreatedAtVal
	}
	if temp.UpdatedAt != nil {
		UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
		if err != nil {
			log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
		}
		s.UpdatedAt = &UpdatedAtVal
	}
	s.CancellationMessage = temp.CancellationMessage
	s.CancellationMethod = temp.CancellationMethod
	s.CancelAtEndOfPeriod = temp.CancelAtEndOfPeriod
	s.CanceledAt.ShouldSetValue(temp.CanceledAt.IsValueSet())
	if temp.CanceledAt.Value() != nil {
		CanceledAtVal, err := time.Parse(time.RFC3339, (*temp.CanceledAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse canceled_at as % s format.", time.RFC3339)
		}
		s.CanceledAt.SetValue(&CanceledAtVal)
	}
	if temp.CurrentPeriodStartedAt != nil {
		CurrentPeriodStartedAtVal, err := time.Parse(time.RFC3339, *temp.CurrentPeriodStartedAt)
		if err != nil {
			log.Fatalf("Cannot Parse current_period_started_at as % s format.", time.RFC3339)
		}
		s.CurrentPeriodStartedAt = &CurrentPeriodStartedAtVal
	}
	s.PreviousState = temp.PreviousState
	s.SignupPaymentId = temp.SignupPaymentId
	s.SignupRevenue = temp.SignupRevenue
	s.DelayedCancelAt.ShouldSetValue(temp.DelayedCancelAt.IsValueSet())
	if temp.DelayedCancelAt.Value() != nil {
		DelayedCancelAtVal, err := time.Parse(time.RFC3339, (*temp.DelayedCancelAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse delayed_cancel_at as % s format.", time.RFC3339)
		}
		s.DelayedCancelAt.SetValue(&DelayedCancelAtVal)
	}
	s.CouponCode = temp.CouponCode
	s.SnapDay = temp.SnapDay
	s.PaymentCollectionMethod = temp.PaymentCollectionMethod
	s.Customer = temp.Customer
	s.Product = temp.Product
	s.CreditCard = temp.CreditCard
	s.Group = temp.Group
	s.BankAccount = temp.BankAccount
	s.PaymentType = temp.PaymentType
	s.ReferralCode = temp.ReferralCode
	s.NextProductId = temp.NextProductId
	s.NextProductHandle = temp.NextProductHandle
	s.CouponUseCount = temp.CouponUseCount
	s.CouponUsesAllowed = temp.CouponUsesAllowed
	s.ReasonCode = temp.ReasonCode
	s.AutomaticallyResumeAt.ShouldSetValue(temp.AutomaticallyResumeAt.IsValueSet())
	if temp.AutomaticallyResumeAt.Value() != nil {
		AutomaticallyResumeAtVal, err := time.Parse(time.RFC3339, (*temp.AutomaticallyResumeAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse automatically_resume_at as % s format.", time.RFC3339)
		}
		s.AutomaticallyResumeAt.SetValue(&AutomaticallyResumeAtVal)
	}
	s.CouponCodes = temp.CouponCodes
	s.OfferId = temp.OfferId
	s.PayerId = temp.PayerId
	s.CurrentBillingAmountInCents = temp.CurrentBillingAmountInCents
	s.ProductPricePointId = temp.ProductPricePointId
	s.ProductPricePointType = temp.ProductPricePointType
	s.NextProductPricePointId = temp.NextProductPricePointId
	s.NetTerms = temp.NetTerms
	s.StoredCredentialTransactionId = temp.StoredCredentialTransactionId
	s.Reference = temp.Reference
	s.OnHoldAt.ShouldSetValue(temp.OnHoldAt.IsValueSet())
	if temp.OnHoldAt.Value() != nil {
		OnHoldAtVal, err := time.Parse(time.RFC3339, (*temp.OnHoldAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse on_hold_at as % s format.", time.RFC3339)
		}
		s.OnHoldAt.SetValue(&OnHoldAtVal)
	}
	s.PrepaidDunning = temp.PrepaidDunning
	s.Coupons = temp.Coupons
	s.DunningCommunicationDelayEnabled = temp.DunningCommunicationDelayEnabled
	s.DunningCommunicationDelayTimeZone = temp.DunningCommunicationDelayTimeZone
	s.ReceivesInvoiceEmails = temp.ReceivesInvoiceEmails
	s.Locale = temp.Locale
	s.Currency = temp.Currency
	s.ScheduledCancellationAt.ShouldSetValue(temp.ScheduledCancellationAt.IsValueSet())
	if temp.ScheduledCancellationAt.Value() != nil {
		ScheduledCancellationAtVal, err := time.Parse(time.RFC3339, (*temp.ScheduledCancellationAt.Value()))
		if err != nil {
			log.Fatalf("Cannot Parse scheduled_cancellation_at as % s format.", time.RFC3339)
		}
		s.ScheduledCancellationAt.SetValue(&ScheduledCancellationAtVal)
	}
	s.CreditBalanceInCents = temp.CreditBalanceInCents
	s.PrepaymentBalanceInCents = temp.PrepaymentBalanceInCents
	s.PrepaidConfiguration = temp.PrepaidConfiguration
	return nil
}
