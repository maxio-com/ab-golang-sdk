package models

import (
	"encoding/json"
)

// UpdateSubscription represents a UpdateSubscription struct.
type UpdateSubscription struct {
	CreditCardAttributes *CreditCardAttributes `json:"credit_card_attributes,omitempty"`
	// Set to the handle of a different product to change the subscription's product
	ProductHandle *string `json:"product_handle,omitempty"`
	// Set to the id of a different product to change the subscription's product
	ProductId            *int  `json:"product_id,omitempty"`
	ProductChangeDelayed *bool `json:"product_change_delayed,omitempty"`
	// Set to an empty string to cancel a delayed product change.
	NextProductId           *string `json:"next_product_id,omitempty"`
	NextProductPricePointId *string `json:"next_product_price_point_id,omitempty"`
	// Use for subscriptions with product eligible for calendar billing only. Value can be 1-28 or 'end'.
	SnapDay                       *interface{} `json:"snap_day,omitempty"`
	NextBillingAt                 *string      `json:"next_billing_at,omitempty"`
	PaymentCollectionMethod       *string      `json:"payment_collection_method,omitempty"`
	ReceivesInvoiceEmails         *bool        `json:"receives_invoice_emails,omitempty"`
	NetTerms                      *interface{} `json:"net_terms,omitempty"`
	StoredCredentialTransactionId *int         `json:"stored_credential_transaction_id,omitempty"`
	Reference                     *string      `json:"reference,omitempty"`
	// (Optional) Used in place of `product_price_point_id` to define a custom price point unique to the subscription
	CustomPrice *SubscriptionCustomPrice `json:"custom_price,omitempty"`
	// (Optional) An array of component ids and custom prices to be added to the subscription.
	Components []UpdateSubscriptionComponent `json:"components,omitempty"`
	// Enable Communication Delay feature, making sure no communication (email or SMS) is sent to the Customer between 9PM and 8AM in time zone set by the `dunning_communication_delay_time_zone` attribute.
	DunningCommunicationDelayEnabled Optional[bool] `json:"dunning_communication_delay_enabled"`
	// Time zone for the Dunning Communication Delay feature.
	DunningCommunicationDelayTimeZone Optional[string] `json:"dunning_communication_delay_time_zone"`
}

// MarshalJSON implements the json.Marshaler interface for UpdateSubscription.
// It customizes the JSON marshaling process for UpdateSubscription objects.
func (u *UpdateSubscription) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UpdateSubscription object to a map representation for JSON marshaling.
func (u *UpdateSubscription) toMap() map[string]any {
	structMap := make(map[string]any)
	if u.CreditCardAttributes != nil {
		structMap["credit_card_attributes"] = u.CreditCardAttributes
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
	if u.SnapDay != nil {
		structMap["snap_day"] = u.SnapDay
	}
	if u.NextBillingAt != nil {
		structMap["next_billing_at"] = u.NextBillingAt
	}
	if u.PaymentCollectionMethod != nil {
		structMap["payment_collection_method"] = u.PaymentCollectionMethod
	}
	if u.ReceivesInvoiceEmails != nil {
		structMap["receives_invoice_emails"] = u.ReceivesInvoiceEmails
	}
	if u.NetTerms != nil {
		structMap["net_terms"] = u.NetTerms
	}
	if u.StoredCredentialTransactionId != nil {
		structMap["stored_credential_transaction_id"] = u.StoredCredentialTransactionId
	}
	if u.Reference != nil {
		structMap["reference"] = u.Reference
	}
	if u.CustomPrice != nil {
		structMap["custom_price"] = u.CustomPrice
	}
	if u.Components != nil {
		structMap["components"] = u.Components
	}
	if u.DunningCommunicationDelayEnabled.IsValueSet() {
		structMap["dunning_communication_delay_enabled"] = u.DunningCommunicationDelayEnabled.Value()
	}
	if u.DunningCommunicationDelayTimeZone.IsValueSet() {
		structMap["dunning_communication_delay_time_zone"] = u.DunningCommunicationDelayTimeZone.Value()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UpdateSubscription.
// It customizes the JSON unmarshaling process for UpdateSubscription objects.
func (u *UpdateSubscription) UnmarshalJSON(input []byte) error {
	temp := &struct {
		CreditCardAttributes              *CreditCardAttributes         `json:"credit_card_attributes,omitempty"`
		ProductHandle                     *string                       `json:"product_handle,omitempty"`
		ProductId                         *int                          `json:"product_id,omitempty"`
		ProductChangeDelayed              *bool                         `json:"product_change_delayed,omitempty"`
		NextProductId                     *string                       `json:"next_product_id,omitempty"`
		NextProductPricePointId           *string                       `json:"next_product_price_point_id,omitempty"`
		SnapDay                           *interface{}                  `json:"snap_day,omitempty"`
		NextBillingAt                     *string                       `json:"next_billing_at,omitempty"`
		PaymentCollectionMethod           *string                       `json:"payment_collection_method,omitempty"`
		ReceivesInvoiceEmails             *bool                         `json:"receives_invoice_emails,omitempty"`
		NetTerms                          *interface{}                  `json:"net_terms,omitempty"`
		StoredCredentialTransactionId     *int                          `json:"stored_credential_transaction_id,omitempty"`
		Reference                         *string                       `json:"reference,omitempty"`
		CustomPrice                       *SubscriptionCustomPrice      `json:"custom_price,omitempty"`
		Components                        []UpdateSubscriptionComponent `json:"components,omitempty"`
		DunningCommunicationDelayEnabled  Optional[bool]                `json:"dunning_communication_delay_enabled"`
		DunningCommunicationDelayTimeZone Optional[string]              `json:"dunning_communication_delay_time_zone"`
	}{}
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}

	u.CreditCardAttributes = temp.CreditCardAttributes
	u.ProductHandle = temp.ProductHandle
	u.ProductId = temp.ProductId
	u.ProductChangeDelayed = temp.ProductChangeDelayed
	u.NextProductId = temp.NextProductId
	u.NextProductPricePointId = temp.NextProductPricePointId
	u.SnapDay = temp.SnapDay
	u.NextBillingAt = temp.NextBillingAt
	u.PaymentCollectionMethod = temp.PaymentCollectionMethod
	u.ReceivesInvoiceEmails = temp.ReceivesInvoiceEmails
	u.NetTerms = temp.NetTerms
	u.StoredCredentialTransactionId = temp.StoredCredentialTransactionId
	u.Reference = temp.Reference
	u.CustomPrice = temp.CustomPrice
	u.Components = temp.Components
	u.DunningCommunicationDelayEnabled = temp.DunningCommunicationDelayEnabled
	u.DunningCommunicationDelayTimeZone = temp.DunningCommunicationDelayTimeZone
	return nil
}
