/*
Package advancedbilling

This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package models

// AllocationPreviewDirection is a string enum.
type AllocationPreviewDirection string

const (
    AllocationPreviewDirection_UPGRADE   AllocationPreviewDirection = "upgrade"
    AllocationPreviewDirection_DOWNGRADE AllocationPreviewDirection = "downgrade"
)

// AllocationPreviewLineItemKind is a string enum.
// A handle for the line item kind for allocation preview
type AllocationPreviewLineItemKind string

const (
    AllocationPreviewLineItemKind_QUANTITYBASEDCOMPONENT AllocationPreviewLineItemKind = "quantity_based_component"
    AllocationPreviewLineItemKind_ONOFFCOMPONENT         AllocationPreviewLineItemKind = "on_off_component"
    AllocationPreviewLineItemKind_COUPON                 AllocationPreviewLineItemKind = "coupon"
    AllocationPreviewLineItemKind_TAX                    AllocationPreviewLineItemKind = "tax"
)

// AutoInvite is a int enum.
type AutoInvite int

const (
    AutoInvite_NO  AutoInvite = 0
    AutoInvite_YES AutoInvite = 1
)

// BankAccountHolderType is a string enum.
// Defaults to personal
type BankAccountHolderType string

const (
    BankAccountHolderType_PERSONAL BankAccountHolderType = "personal"
    BankAccountHolderType_BUSINESS BankAccountHolderType = "business"
)

// BankAccountType is a string enum.
// Defaults to checking
type BankAccountType string

const (
    BankAccountType_CHECKING BankAccountType = "checking"
    BankAccountType_SAVINGS  BankAccountType = "savings"
)

// BankAccountVault is a string enum.
// The vault that stores the payment profile with the provided vault_token.
type BankAccountVault string

const (
    BankAccountVault_BOGUS         BankAccountVault = "bogus"
    BankAccountVault_AUTHORIZENET  BankAccountVault = "authorizenet"
    BankAccountVault_STRIPECONNECT BankAccountVault = "stripe_connect"
    BankAccountVault_BRAINTREEBLUE BankAccountVault = "braintree_blue"
    BankAccountVault_GOCARDLESS    BankAccountVault = "gocardless"
)

// BasicDateField is a string enum.
// Allows to filter by `created_at` or `updated_at`.
type BasicDateField string

const (
    BasicDateField_UPDATEDAT BasicDateField = "updated_at"
    BasicDateField_CREATEDAT BasicDateField = "created_at"
)

// BillingManifestLineItemKind is a string enum.
// A handle for the billing manifest line item kind
type BillingManifestLineItemKind string

const (
    BillingManifestLineItemKind_BASELINE  BillingManifestLineItemKind = "baseline"
    BillingManifestLineItemKind_INITIAL   BillingManifestLineItemKind = "initial"
    BillingManifestLineItemKind_TRIAL     BillingManifestLineItemKind = "trial"
    BillingManifestLineItemKind_COUPON    BillingManifestLineItemKind = "coupon"
    BillingManifestLineItemKind_COMPONENT BillingManifestLineItemKind = "component"
    BillingManifestLineItemKind_TAX       BillingManifestLineItemKind = "tax"
)

// CancellationMethod is a string enum.
// The process used to cancel the subscription, if the subscription has been canceled. It is nil if the subscription's state is not canceled.
type CancellationMethod string

const (
    CancellationMethod_MERCHANTUI    CancellationMethod = "merchant_ui"
    CancellationMethod_MERCHANTAPI   CancellationMethod = "merchant_api"
    CancellationMethod_DUNNING       CancellationMethod = "dunning"
    CancellationMethod_BILLINGPORTAL CancellationMethod = "billing_portal"
    CancellationMethod_UNKNOWN       CancellationMethod = "unknown"
)

// CardType is a string enum.
// The type of card used.
type CardType string

const (
    CardType_VISA               CardType = "visa"
    CardType_MASTER             CardType = "master"
    CardType_ELO                CardType = "elo"
    CardType_CABAL              CardType = "cabal"
    CardType_ALELO              CardType = "alelo"
    CardType_DISCOVER           CardType = "discover"
    CardType_AMERICANEXPRESS    CardType = "american_express"
    CardType_NARANJA            CardType = "naranja"
    CardType_DINERSCLUB         CardType = "diners_club"
    CardType_JCB                CardType = "jcb"
    CardType_DANKORT            CardType = "dankort"
    CardType_MAESTRO            CardType = "maestro"
    CardType_MAESTRONOLUHN      CardType = "maestro_no_luhn"
    CardType_FORBRUGSFORENINGEN CardType = "forbrugsforeningen"
    CardType_SODEXO             CardType = "sodexo"
    CardType_ALIA               CardType = "alia"
    CardType_VR                 CardType = "vr"
    CardType_UNIONPAY           CardType = "unionpay"
    CardType_CARNET             CardType = "carnet"
    CardType_CARTESBANCAIRES    CardType = "cartes_bancaires"
    CardType_OLIMPICA           CardType = "olimpica"
    CardType_CREDITEL           CardType = "creditel"
    CardType_CONFIABLE          CardType = "confiable"
    CardType_SYNCHRONY          CardType = "synchrony"
    CardType_ROUTEX             CardType = "routex"
    CardType_MADA               CardType = "mada"
    CardType_BPPLUS             CardType = "bp_plus"
    CardType_PASSCARD           CardType = "passcard"
    CardType_EDENRED            CardType = "edenred"
    CardType_ANDA               CardType = "anda"
    CardType_TARJETAD           CardType = "tarjeta-d"
    CardType_HIPERCARD          CardType = "hipercard"
    CardType_BOGUS              CardType = "bogus"
    CardType_ENUMSWITCH         CardType = "switch"
    CardType_SOLO               CardType = "solo"
    CardType_LASER              CardType = "laser"
)

// CleanupScope is a string enum.
// all: Will clear all products, customers, and related subscriptions from the site. customers: Will clear only customers and related subscriptions (leaving the products untouched) for the site. Revenue will also be reset to 0.
type CleanupScope string

const (
    CleanupScope_ALL       CleanupScope = "all"
    CleanupScope_CUSTOMERS CleanupScope = "customers"
)

// CollectionMethod is a string enum.
// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
type CollectionMethod string

const (
    CollectionMethod_AUTOMATIC  CollectionMethod = "automatic"
    CollectionMethod_REMITTANCE CollectionMethod = "remittance"
    CollectionMethod_PREPAID    CollectionMethod = "prepaid"
    CollectionMethod_INVOICE    CollectionMethod = "invoice"
)

// ComponentKind is a string enum.
// A handle for the component type
type ComponentKind string

const (
    ComponentKind_METEREDCOMPONENT       ComponentKind = "metered_component"
    ComponentKind_QUANTITYBASEDCOMPONENT ComponentKind = "quantity_based_component"
    ComponentKind_ONOFFCOMPONENT         ComponentKind = "on_off_component"
    ComponentKind_PREPAIDUSAGECOMPONENT  ComponentKind = "prepaid_usage_component"
    ComponentKind_EVENTBASEDCOMPONENT    ComponentKind = "event_based_component"
)

// CompoundingStrategy is a string enum.
type CompoundingStrategy string

const (
    CompoundingStrategy_COMPOUND  CompoundingStrategy = "compound"
    CompoundingStrategy_FULLPRICE CompoundingStrategy = "full-price"
)

// CreateInvoiceStatus is a string enum.
type CreateInvoiceStatus string

const (
    CreateInvoiceStatus_DRAFT CreateInvoiceStatus = "draft"
    CreateInvoiceStatus_OPEN  CreateInvoiceStatus = "open"
)

// CreditNoteStatus is a string enum.
// Current status of the credit note.
type CreditNoteStatus string

const (
    CreditNoteStatus_OPEN    CreditNoteStatus = "open"
    CreditNoteStatus_APPLIED CreditNoteStatus = "applied"
)

// CreditScheme is a string enum.
type CreditScheme string

const (
    CreditScheme_NONE   CreditScheme = "none"
    CreditScheme_CREDIT CreditScheme = "credit"
    CreditScheme_REFUND CreditScheme = "refund"
)

// CreditType is a string enum.
// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
// Available values: `full`, `prorated`, `none`.
type CreditType string

const (
    CreditType_FULL     CreditType = "full"
    CreditType_PRORATED CreditType = "prorated"
    CreditType_NONE     CreditType = "none"
)

// CurrencyPriceRole is a string enum.
// Role for the price.
type CurrencyPriceRole string

const (
    CurrencyPriceRole_BASELINE CurrencyPriceRole = "baseline"
    CurrencyPriceRole_TRIAL    CurrencyPriceRole = "trial"
    CurrencyPriceRole_INITIAL  CurrencyPriceRole = "initial"
)

// CurrentVault is a string enum.
// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
type CurrentVault string

const (
    CurrentVault_ADYEN          CurrentVault = "adyen"
    CurrentVault_AUTHORIZENET   CurrentVault = "authorizenet"
    CurrentVault_AVALARA        CurrentVault = "avalara"
    CurrentVault_BEANSTREAM     CurrentVault = "beanstream"
    CurrentVault_BLUESNAP       CurrentVault = "blue_snap"
    CurrentVault_BOGUS          CurrentVault = "bogus"
    CurrentVault_BRAINTREEBLUE  CurrentVault = "braintree_blue"
    CurrentVault_CHECKOUT       CurrentVault = "checkout"
    CurrentVault_CYBERSOURCE    CurrentVault = "cybersource"
    CurrentVault_ELAVON         CurrentVault = "elavon"
    CurrentVault_EWAY           CurrentVault = "eway"
    CurrentVault_EWAYRAPIDSTD   CurrentVault = "eway_rapid_std"
    CurrentVault_FIRSTDATA      CurrentVault = "firstdata"
    CurrentVault_FORTE          CurrentVault = "forte"
    CurrentVault_GOCARDLESS     CurrentVault = "gocardless"
    CurrentVault_LITLE          CurrentVault = "litle"
    CurrentVault_MAXIOPAYMENTS  CurrentVault = "maxio_payments"
    CurrentVault_MODUSLINK      CurrentVault = "moduslink"
    CurrentVault_MONERIS        CurrentVault = "moneris"
    CurrentVault_NMI            CurrentVault = "nmi"
    CurrentVault_ORBITAL        CurrentVault = "orbital"
    CurrentVault_PAYMENTEXPRESS CurrentVault = "payment_express"
    CurrentVault_PIN            CurrentVault = "pin"
    CurrentVault_SQUARE         CurrentVault = "square"
    CurrentVault_STRIPECONNECT  CurrentVault = "stripe_connect"
    CurrentVault_TRUSTCOMMERCE  CurrentVault = "trust_commerce"
    CurrentVault_UNIPAAS        CurrentVault = "unipaas"
)

// CustomFieldOwner is a string enum.
type CustomFieldOwner string

const (
    CustomFieldOwner_CUSTOMER     CustomFieldOwner = "Customer"
    CustomFieldOwner_SUBSCRIPTION CustomFieldOwner = "Subscription"
)

// Direction is a string enum.
type Direction string

const (
    Direction_ASC  Direction = "asc"
    Direction_DESC Direction = "desc"
)

// DiscountType is a string enum.
type DiscountType string

const (
    DiscountType_AMOUNT  DiscountType = "amount"
    DiscountType_PERCENT DiscountType = "percent"
)

// EventType is a string enum.
type EventType string

const (
    EventType_ACCOUNTTRANSACTIONCHANGED                      EventType = "account_transaction_changed"
    EventType_BILLINGDATECHANGE                              EventType = "billing_date_change"
    EventType_COMPONENTALLOCATIONCHANGE                      EventType = "component_allocation_change"
    EventType_CUSTOMERUPDATE                                 EventType = "customer_update"
    EventType_CUSTOMERCREATE                                 EventType = "customer_create"
    EventType_DUNNINGSTEPREACHED                             EventType = "dunning_step_reached"
    EventType_EXPIRATIONDATECHANGE                           EventType = "expiration_date_change"
    EventType_EXPIRINGCARD                                   EventType = "expiring_card"
    EventType_METEREDUSAGE                                   EventType = "metered_usage"
    EventType_PAYMENTSUCCESS                                 EventType = "payment_success"
    EventType_PAYMENTSUCCESSRECREATED                        EventType = "payment_success_recreated"
    EventType_PAYMENTFAILURE                                 EventType = "payment_failure"
    EventType_PAYMENTFAILURERECREATED                        EventType = "payment_failure_recreated"
    EventType_REFUNDFAILURE                                  EventType = "refund_failure"
    EventType_REFUNDSUCCESS                                  EventType = "refund_success"
    EventType_RENEWALSUCCESS                                 EventType = "renewal_success"
    EventType_RENEWALSUCCESSRECREATED                        EventType = "renewal_success_recreated"
    EventType_RENEWALFAILURE                                 EventType = "renewal_failure"
    EventType_SIGNUPSUCCESS                                  EventType = "signup_success"
    EventType_SIGNUPFAILURE                                  EventType = "signup_failure"
    EventType_STATEMENTCLOSED                                EventType = "statement_closed"
    EventType_STATEMENTSETTLED                               EventType = "statement_settled"
    EventType_SUBSCRIPTIONBANKACCOUNTUPDATE                  EventType = "subscription_bank_account_update"
    EventType_SUBSCRIPTIONDELETION                           EventType = "subscription_deletion"
    EventType_SUBSCRIPTIONPAYPALACCOUNTUPDATE                EventType = "subscription_paypal_account_update"
    EventType_SUBSCRIPTIONPRODUCTCHANGE                      EventType = "subscription_product_change"
    EventType_SUBSCRIPTIONSTATECHANGE                        EventType = "subscription_state_change"
    EventType_TRIALENDNOTICE                                 EventType = "trial_end_notice"
    EventType_UPGRADEDOWNGRADESUCCESS                        EventType = "upgrade_downgrade_success"
    EventType_UPGRADEDOWNGRADEFAILURE                        EventType = "upgrade_downgrade_failure"
    EventType_UPCOMINGRENEWALNOTICE                          EventType = "upcoming_renewal_notice"
    EventType_CUSTOMFIELDVALUECHANGE                         EventType = "custom_field_value_change"
    EventType_SUBSCRIPTIONPREPAYMENTACCOUNTBALANCECHANGED    EventType = "subscription_prepayment_account_balance_changed"
    EventType_SUBSCRIPTIONSERVICECREDITACCOUNTBALANCECHANGED EventType = "subscription_service_credit_account_balance_changed"
)

// ExtendedIntervalUnit is a string enum.
type ExtendedIntervalUnit string

const (
    ExtendedIntervalUnit_DAY   ExtendedIntervalUnit = "day"
    ExtendedIntervalUnit_MONTH ExtendedIntervalUnit = "month"
    ExtendedIntervalUnit_NEVER ExtendedIntervalUnit = "never"
)

// FailedPaymentAction is a string enum.
// Action taken when payment for an invoice fails:
// - `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
// - `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
// - `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).
type FailedPaymentAction string

const (
    FailedPaymentAction_LEAVEOPENINVOICE  FailedPaymentAction = "leave_open_invoice"
    FailedPaymentAction_ROLLBACKTOPENDING FailedPaymentAction = "rollback_to_pending"
    FailedPaymentAction_INITIATEDUNNING   FailedPaymentAction = "initiate_dunning"
)

// FirstChargeType is a string enum.
type FirstChargeType string

const (
    FirstChargeType_PRORATED  FirstChargeType = "prorated"
    FirstChargeType_IMMEDIATE FirstChargeType = "immediate"
    FirstChargeType_DELAYED   FirstChargeType = "delayed"
)

// GroupTargetType is a string enum.
// The type of object indicated by the id attribute.
type GroupTargetType string

const (
    GroupTargetType_CUSTOMER     GroupTargetType = "customer"
    GroupTargetType_SUBSCRIPTION GroupTargetType = "subscription"
    GroupTargetType_SELF         GroupTargetType = "self"
    GroupTargetType_PARENT       GroupTargetType = "parent"
    GroupTargetType_ELDEST       GroupTargetType = "eldest"
)

// IncludeNotNull is a string enum.
// Passed as a parameter to list methods to return only non null values.
type IncludeNotNull string

const (
    IncludeNotNull_NOTNULL IncludeNotNull = "not_null"
)

// IncludeOption is a string enum.
type IncludeOption string

const (
    IncludeOption_EXCLUDE IncludeOption = "0"
    IncludeOption_INCLUDE IncludeOption = "1"
)

// IntervalUnit is a string enum.
type IntervalUnit string

const (
    IntervalUnit_DAY   IntervalUnit = "day"
    IntervalUnit_MONTH IntervalUnit = "month"
)

// InvoiceConsolidationLevel is a string enum.
// Consolidation level of the invoice, which is applicable to invoice consolidation.  It will hold one of the following values:
// * "none": A normal invoice with no consolidation.
// * "child": An invoice segment which has been combined into a consolidated invoice.
// * "parent": A consolidated invoice, whose contents are composed of invoice segments.
// "Parent" invoices do not have lines of their own, but they have subtotals and totals which aggregate the member invoice segments.
// See also the [invoice consolidation documentation](https://chargify.zendesk.com/hc/en-us/articles/4407746391835).
type InvoiceConsolidationLevel string

const (
    InvoiceConsolidationLevel_NONE   InvoiceConsolidationLevel = "none"
    InvoiceConsolidationLevel_CHILD  InvoiceConsolidationLevel = "child"
    InvoiceConsolidationLevel_PARENT InvoiceConsolidationLevel = "parent"
)

// InvoiceDateField is a string enum.
type InvoiceDateField string

const (
    InvoiceDateField_CREATEDAT InvoiceDateField = "created_at"
    InvoiceDateField_DUEDATE   InvoiceDateField = "due_date"
    InvoiceDateField_ISSUEDATE InvoiceDateField = "issue_date"
    InvoiceDateField_UPDATEDAT InvoiceDateField = "updated_at"
    InvoiceDateField_PAIDDATE  InvoiceDateField = "paid_date"
)

// InvoiceEventPaymentMethod is a string enum.
type InvoiceEventPaymentMethod string

const (
    InvoiceEventPaymentMethod_APPLEPAY      InvoiceEventPaymentMethod = "apple_pay"
    InvoiceEventPaymentMethod_BANKACCOUNT   InvoiceEventPaymentMethod = "bank_account"
    InvoiceEventPaymentMethod_CREDITCARD    InvoiceEventPaymentMethod = "credit_card"
    InvoiceEventPaymentMethod_EXTERNAL      InvoiceEventPaymentMethod = "external"
    InvoiceEventPaymentMethod_PAYPALACCOUNT InvoiceEventPaymentMethod = "paypal_account"
)

// InvoiceEventType is a string enum.
// Invoice Event Type
type InvoiceEventType string

const (
    InvoiceEventType_ISSUEINVOICE                  InvoiceEventType = "issue_invoice"
    InvoiceEventType_APPLYCREDITNOTE               InvoiceEventType = "apply_credit_note"
    InvoiceEventType_CREATECREDITNOTE              InvoiceEventType = "create_credit_note"
    InvoiceEventType_APPLYPAYMENT                  InvoiceEventType = "apply_payment"
    InvoiceEventType_APPLYDEBITNOTE                InvoiceEventType = "apply_debit_note"
    InvoiceEventType_REFUNDINVOICE                 InvoiceEventType = "refund_invoice"
    InvoiceEventType_VOIDINVOICE                   InvoiceEventType = "void_invoice"
    InvoiceEventType_VOIDREMAINDER                 InvoiceEventType = "void_remainder"
    InvoiceEventType_BACKPORTINVOICE               InvoiceEventType = "backport_invoice"
    InvoiceEventType_CHANGEINVOICESTATUS           InvoiceEventType = "change_invoice_status"
    InvoiceEventType_CHANGEINVOICECOLLECTIONMETHOD InvoiceEventType = "change_invoice_collection_method"
    InvoiceEventType_REMOVEPAYMENT                 InvoiceEventType = "remove_payment"
    InvoiceEventType_FAILEDPAYMENT                 InvoiceEventType = "failed_payment"
    InvoiceEventType_CHANGECHARGEBACKSTATUS        InvoiceEventType = "change_chargeback_status"
)

// InvoicePaymentMethodType is a string enum.
// The type of payment method used. Defaults to other.
type InvoicePaymentMethodType string

const (
    InvoicePaymentMethodType_CREDITCARD InvoicePaymentMethodType = "credit_card"
    InvoicePaymentMethodType_CHECK      InvoicePaymentMethodType = "check"
    InvoicePaymentMethodType_CASH       InvoicePaymentMethodType = "cash"
    InvoicePaymentMethodType_MONEYORDER InvoicePaymentMethodType = "money_order"
    InvoicePaymentMethodType_ACH        InvoicePaymentMethodType = "ach"
    InvoicePaymentMethodType_OTHER      InvoicePaymentMethodType = "other"
)

// InvoicePaymentType is a string enum.
// The type of payment to be applied to an Invoice. Defaults to external.
type InvoicePaymentType string

const (
    InvoicePaymentType_EXTERNAL      InvoicePaymentType = "external"
    InvoicePaymentType_PREPAYMENT    InvoicePaymentType = "prepayment"
    InvoicePaymentType_SERVICECREDIT InvoicePaymentType = "service_credit"
    InvoicePaymentType_PAYMENT       InvoicePaymentType = "payment"
)

// InvoiceRole is a string enum.
type InvoiceRole string

const (
    InvoiceRole_UNSET                         InvoiceRole = "unset"
    InvoiceRole_SIGNUP                        InvoiceRole = "signup"
    InvoiceRole_RENEWAL                       InvoiceRole = "renewal"
    InvoiceRole_USAGE                         InvoiceRole = "usage"
    InvoiceRole_REACTIVATION                  InvoiceRole = "reactivation"
    InvoiceRole_PRORATION                     InvoiceRole = "proration"
    InvoiceRole_MIGRATION                     InvoiceRole = "migration"
    InvoiceRole_ADHOC                         InvoiceRole = "adhoc"
    InvoiceRole_BACKPORT                      InvoiceRole = "backport"
    InvoiceRole_BACKPORTBALANCERECONCILIATION InvoiceRole = "backport-balance-reconciliation"
)

// InvoiceSortField is a string enum.
type InvoiceSortField string

const (
    InvoiceSortField_STATUS      InvoiceSortField = "status"
    InvoiceSortField_TOTALAMOUNT InvoiceSortField = "total_amount"
    InvoiceSortField_DUEAMOUNT   InvoiceSortField = "due_amount"
    InvoiceSortField_CREATEDAT   InvoiceSortField = "created_at"
    InvoiceSortField_UPDATEDAT   InvoiceSortField = "updated_at"
    InvoiceSortField_ISSUEDATE   InvoiceSortField = "issue_date"
    InvoiceSortField_DUEDATE     InvoiceSortField = "due_date"
    InvoiceSortField_NUMBER      InvoiceSortField = "number"
)

// InvoiceStatus is a string enum.
// The current status of the invoice. See [Invoice Statuses](https://chargify.zendesk.com/hc/en-us/articles/4407737494171#line-item-breakdowns) for more.
type InvoiceStatus string

const (
    InvoiceStatus_DRAFT    InvoiceStatus = "draft"
    InvoiceStatus_OPEN     InvoiceStatus = "open"
    InvoiceStatus_PAID     InvoiceStatus = "paid"
    InvoiceStatus_PENDING  InvoiceStatus = "pending"
    InvoiceStatus_VOIDED   InvoiceStatus = "voided"
    InvoiceStatus_CANCELED InvoiceStatus = "canceled"
)

// ItemCategory is a string enum.
// One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
type ItemCategory string

const (
    ItemCategory_ENUMBUSINESSSOFTWARE ItemCategory = "Business Software"
    ItemCategory_ENUMCONSUMERSOFTWARE ItemCategory = "Consumer Software"
    ItemCategory_ENUMDIGITALSERVICES  ItemCategory = "Digital Services"
    ItemCategory_ENUMPHYSICALGOODS    ItemCategory = "Physical Goods"
    ItemCategory_OTHER                ItemCategory = "Other"
)

// LineItemKind is a string enum.
// A handle for the line item kind
type LineItemKind string

const (
    LineItemKind_BASELINE               LineItemKind = "baseline"
    LineItemKind_INITIAL                LineItemKind = "initial"
    LineItemKind_TRIAL                  LineItemKind = "trial"
    LineItemKind_QUANTITYBASEDCOMPONENT LineItemKind = "quantity_based_component"
    LineItemKind_PREPAIDUSAGECOMPONENT  LineItemKind = "prepaid_usage_component"
    LineItemKind_ONOFFCOMPONENT         LineItemKind = "on_off_component"
    LineItemKind_METEREDCOMPONENT       LineItemKind = "metered_component"
    LineItemKind_EVENTBASEDCOMPONENT    LineItemKind = "event_based_component"
    LineItemKind_COUPON                 LineItemKind = "coupon"
    LineItemKind_TAX                    LineItemKind = "tax"
)

// LineItemTransactionType is a string enum.
// A handle for the line item transaction type
type LineItemTransactionType string

const (
    LineItemTransactionType_CHARGE               LineItemTransactionType = "charge"
    LineItemTransactionType_CREDIT               LineItemTransactionType = "credit"
    LineItemTransactionType_ADJUSTMENT           LineItemTransactionType = "adjustment"
    LineItemTransactionType_PAYMENT              LineItemTransactionType = "payment"
    LineItemTransactionType_REFUND               LineItemTransactionType = "refund"
    LineItemTransactionType_INFOTRANSACTION      LineItemTransactionType = "info_transaction"
    LineItemTransactionType_PAYMENTAUTHORIZATION LineItemTransactionType = "payment_authorization"
)

// ListComponentsPricePointsInclude is a string enum.
type ListComponentsPricePointsInclude string

const (
    ListComponentsPricePointsInclude_CURRENCYPRICES ListComponentsPricePointsInclude = "currency_prices"
)

// ListEventsDateField is a string enum.
type ListEventsDateField string

const (
    ListEventsDateField_CREATEDAT ListEventsDateField = "created_at"
)

// ListProductsInclude is a string enum.
type ListProductsInclude string

const (
    ListProductsInclude_PREPAIDPRODUCTPRICEPOINT ListProductsInclude = "prepaid_product_price_point"
)

// ListProductsPricePointsInclude is a string enum.
type ListProductsPricePointsInclude string

const (
    ListProductsPricePointsInclude_CURRENCYPRICES ListProductsPricePointsInclude = "currency_prices"
)

// ListSubscriptionComponentsInclude is a string enum.
type ListSubscriptionComponentsInclude string

const (
    ListSubscriptionComponentsInclude_SUBSCRIPTION ListSubscriptionComponentsInclude = "subscription"
)

// ListSubscriptionComponentsSort is a string enum.
type ListSubscriptionComponentsSort string

const (
    ListSubscriptionComponentsSort_ID        ListSubscriptionComponentsSort = "id"
    ListSubscriptionComponentsSort_UPDATEDAT ListSubscriptionComponentsSort = "updated_at"
)

// ListSubscriptionGroupPrepaymentDateField is a string enum.
type ListSubscriptionGroupPrepaymentDateField string

const (
    ListSubscriptionGroupPrepaymentDateField_CREATEDAT     ListSubscriptionGroupPrepaymentDateField = "created_at"
    ListSubscriptionGroupPrepaymentDateField_APPLICATIONAT ListSubscriptionGroupPrepaymentDateField = "application_at"
)

// MetafieldInput is a string enum.
// Indicates how data should be added to the metafield. For example, a text type is just a string, so a given metafield of this type can have any value attached. On the other hand, dropdown and radio have a set of allowed values that can be input, and appear differently on a Public Signup Page. Defaults to 'text'
type MetafieldInput string

const (
    MetafieldInput_BALANCETRACKER MetafieldInput = "balance_tracker"
    MetafieldInput_TEXT           MetafieldInput = "text"
    MetafieldInput_RADIO          MetafieldInput = "radio"
    MetafieldInput_DROPDOWN       MetafieldInput = "dropdown"
)

// PaymentType is a string enum.
type PaymentType string

const (
    PaymentType_CREDITCARD    PaymentType = "credit_card"
    PaymentType_BANKACCOUNT   PaymentType = "bank_account"
    PaymentType_PAYPALACCOUNT PaymentType = "paypal_account"
)

// PrepaymentMethod is a string enum.
// :- When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions.
type PrepaymentMethod string

const (
    PrepaymentMethod_CHECK            PrepaymentMethod = "check"
    PrepaymentMethod_CASH             PrepaymentMethod = "cash"
    PrepaymentMethod_MONEYORDER       PrepaymentMethod = "money_order"
    PrepaymentMethod_ACH              PrepaymentMethod = "ach"
    PrepaymentMethod_PAYPALACCOUNT    PrepaymentMethod = "paypal_account"
    PrepaymentMethod_CREDITCARDONFILE PrepaymentMethod = "credit_card_on_file"
    PrepaymentMethod_OTHER            PrepaymentMethod = "other"
)

// PricePointType is a string enum.
// Price point type. We expose the following types:
// 1. **default**: a price point that is marked as a default price for a certain product.
// 2. **custom**: a custom price point.
// 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
type PricePointType string

const (
    PricePointType_CATALOG     PricePointType = "catalog"
    PricePointType_ENUMDEFAULT PricePointType = "default"
    PricePointType_CUSTOM      PricePointType = "custom"
)

// PricingScheme is a string enum.
// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
type PricingScheme string

const (
    PricingScheme_STAIRSTEP PricingScheme = "stairstep"
    PricingScheme_VOLUME    PricingScheme = "volume"
    PricingScheme_PERUNIT   PricingScheme = "per_unit"
    PricingScheme_TIERED    PricingScheme = "tiered"
)

// ReactivationCharge is a string enum.
// You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted for to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal
type ReactivationCharge string

const (
    ReactivationCharge_PRORATED  ReactivationCharge = "prorated"
    ReactivationCharge_IMMEDIATE ReactivationCharge = "immediate"
    ReactivationCharge_DELAYED   ReactivationCharge = "delayed"
)

// RecurringScheme is a string enum.
type RecurringScheme string

const (
    RecurringScheme_DONOTRECUR        RecurringScheme = "do_not_recur"
    RecurringScheme_RECURINDEFINITELY RecurringScheme = "recur_indefinitely"
    RecurringScheme_RECURWITHDURATION RecurringScheme = "recur_with_duration"
)

// ResourceType is a string enum.
type ResourceType string

const (
    ResourceType_SUBSCRIPTIONS ResourceType = "subscriptions"
    ResourceType_CUSTOMERS     ResourceType = "customers"
)

// RestrictionType is a string enum.
type RestrictionType string

const (
    RestrictionType_COMPONENT RestrictionType = "Component"
    RestrictionType_PRODUCT   RestrictionType = "Product"
)

// ResumptionCharge is a string enum.
// (For calendar billing subscriptions only) The way that the resumed subscription's charge should be handled
type ResumptionCharge string

const (
    ResumptionCharge_PRORATED  ResumptionCharge = "prorated"
    ResumptionCharge_IMMEDIATE ResumptionCharge = "immediate"
    ResumptionCharge_DELAYED   ResumptionCharge = "delayed"
)

// ServiceCreditType is a string enum.
// The type of entry
type ServiceCreditType string

const (
    ServiceCreditType_CREDIT ServiceCreditType = "Credit"
    ServiceCreditType_DEBIT  ServiceCreditType = "Debit"
)

// SortingDirection is a string enum.
// Used for sorting results.
type SortingDirection string

const (
    SortingDirection_ASC  SortingDirection = "asc"
    SortingDirection_DESC SortingDirection = "desc"
)

// SubscriptionDateField is a string enum.
type SubscriptionDateField string

const (
    SubscriptionDateField_CURRENTPERIODENDSAT   SubscriptionDateField = "current_period_ends_at"
    SubscriptionDateField_CURRENTPERIODSTARTSAT SubscriptionDateField = "current_period_starts_at"
    SubscriptionDateField_CREATEDAT             SubscriptionDateField = "created_at"
    SubscriptionDateField_ACTIVATEDAT           SubscriptionDateField = "activated_at"
    SubscriptionDateField_CANCELEDAT            SubscriptionDateField = "canceled_at"
    SubscriptionDateField_EXPIRESAT             SubscriptionDateField = "expires_at"
    SubscriptionDateField_TRIALSTARTEDAT        SubscriptionDateField = "trial_started_at"
    SubscriptionDateField_TRIALENDEDAT          SubscriptionDateField = "trial_ended_at"
    SubscriptionDateField_UPDATEDAT             SubscriptionDateField = "updated_at"
)

// SubscriptionGroupPrepaymentMethod is a string enum.
type SubscriptionGroupPrepaymentMethod string

const (
    SubscriptionGroupPrepaymentMethod_CHECK         SubscriptionGroupPrepaymentMethod = "check"
    SubscriptionGroupPrepaymentMethod_CASH          SubscriptionGroupPrepaymentMethod = "cash"
    SubscriptionGroupPrepaymentMethod_MONEYORDER    SubscriptionGroupPrepaymentMethod = "money_order"
    SubscriptionGroupPrepaymentMethod_ACH           SubscriptionGroupPrepaymentMethod = "ach"
    SubscriptionGroupPrepaymentMethod_PAYPALACCOUNT SubscriptionGroupPrepaymentMethod = "paypal_account"
    SubscriptionGroupPrepaymentMethod_OTHER         SubscriptionGroupPrepaymentMethod = "other"
)

// SubscriptionInclude is a string enum.
type SubscriptionInclude string

const (
    SubscriptionInclude_COUPONS              SubscriptionInclude = "coupons"
    SubscriptionInclude_SELFSERVICEPAGETOKEN SubscriptionInclude = "self_service_page_token"
)

// SubscriptionListDateField is a string enum.
type SubscriptionListDateField string

const (
    SubscriptionListDateField_UPDATEDAT SubscriptionListDateField = "updated_at"
)

// SubscriptionListInclude is a string enum.
type SubscriptionListInclude string

const (
    SubscriptionListInclude_SELFSERVICEPAGETOKEN SubscriptionListInclude = "self_service_page_token"
)

// SubscriptionPurgeType is a string enum.
type SubscriptionPurgeType string

const (
    SubscriptionPurgeType_CUSTOMER       SubscriptionPurgeType = "customer"
    SubscriptionPurgeType_PAYMENTPROFILE SubscriptionPurgeType = "payment_profile"
)

// SubscriptionSort is a string enum.
type SubscriptionSort string

const (
    SubscriptionSort_SIGNUPDATE     SubscriptionSort = "signup_date"
    SubscriptionSort_PERIODSTART    SubscriptionSort = "period_start"
    SubscriptionSort_PERIODEND      SubscriptionSort = "period_end"
    SubscriptionSort_NEXTASSESSMENT SubscriptionSort = "next_assessment"
    SubscriptionSort_UPDATEDAT      SubscriptionSort = "updated_at"
    SubscriptionSort_CREATEDAT      SubscriptionSort = "created_at"
)

// SubscriptionState is a string enum.
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
type SubscriptionState string

const (
    SubscriptionState_PENDING        SubscriptionState = "pending"
    SubscriptionState_FAILEDTOCREATE SubscriptionState = "failed_to_create"
    SubscriptionState_TRIALING       SubscriptionState = "trialing"
    SubscriptionState_ASSESSING      SubscriptionState = "assessing"
    SubscriptionState_ACTIVE         SubscriptionState = "active"
    SubscriptionState_SOFTFAILURE    SubscriptionState = "soft_failure"
    SubscriptionState_PASTDUE        SubscriptionState = "past_due"
    SubscriptionState_SUSPENDED      SubscriptionState = "suspended"
    SubscriptionState_CANCELED       SubscriptionState = "canceled"
    SubscriptionState_EXPIRED        SubscriptionState = "expired"
    SubscriptionState_PAUSED         SubscriptionState = "paused"
    SubscriptionState_UNPAID         SubscriptionState = "unpaid"
    SubscriptionState_TRIALENDED     SubscriptionState = "trial_ended"
    SubscriptionState_ONHOLD         SubscriptionState = "on_hold"
    SubscriptionState_AWAITINGSIGNUP SubscriptionState = "awaiting_signup"
)

// SubscriptionStateFilter is a string enum.
// Allowed values for filtering by the current state of the subscription.
type SubscriptionStateFilter string

const (
    SubscriptionStateFilter_ACTIVE              SubscriptionStateFilter = "active"
    SubscriptionStateFilter_CANCELED            SubscriptionStateFilter = "canceled"
    SubscriptionStateFilter_EXPIRED             SubscriptionStateFilter = "expired"
    SubscriptionStateFilter_EXPIREDCARDS        SubscriptionStateFilter = "expired_cards"
    SubscriptionStateFilter_ONHOLD              SubscriptionStateFilter = "on_hold"
    SubscriptionStateFilter_PASTDUE             SubscriptionStateFilter = "past_due"
    SubscriptionStateFilter_PENDINGCANCELLATION SubscriptionStateFilter = "pending_cancellation"
    SubscriptionStateFilter_PENDINGRENEWAL      SubscriptionStateFilter = "pending_renewal"
    SubscriptionStateFilter_SUSPENDED           SubscriptionStateFilter = "suspended"
    SubscriptionStateFilter_TRIALENDED          SubscriptionStateFilter = "trial_ended"
    SubscriptionStateFilter_TRIALING            SubscriptionStateFilter = "trialing"
    SubscriptionStateFilter_UNPAID              SubscriptionStateFilter = "unpaid"
)

// TaxConfigurationKind is a string enum.
type TaxConfigurationKind string

const (
    TaxConfigurationKind_CUSTOM             TaxConfigurationKind = "custom"
    TaxConfigurationKind_ENUMMANAGEDAVALARA TaxConfigurationKind = "managed avalara"
    TaxConfigurationKind_ENUMLINKEDAVALARA  TaxConfigurationKind = "linked avalara"
    TaxConfigurationKind_ENUMDIGITALRIVER   TaxConfigurationKind = "digital river"
)

// TaxDestinationAddress is a string enum.
type TaxDestinationAddress string

const (
    TaxDestinationAddress_SHIPPINGTHENBILLING TaxDestinationAddress = "shipping_then_billing"
    TaxDestinationAddress_BILLINGTHENSHIPPING TaxDestinationAddress = "billing_then_shipping"
    TaxDestinationAddress_SHIPPINGONLY        TaxDestinationAddress = "shipping_only"
    TaxDestinationAddress_BILLINGONLY         TaxDestinationAddress = "billing_only"
)

// WebhookOrder is a string enum.
type WebhookOrder string

const (
    WebhookOrder_NEWESTFIRST WebhookOrder = "newest_first"
    WebhookOrder_OLDESTFIRST WebhookOrder = "oldest_first"
)

// WebhookStatus is a string enum.
type WebhookStatus string

const (
    WebhookStatus_SUCCESSFUL WebhookStatus = "successful"
    WebhookStatus_FAILED     WebhookStatus = "failed"
    WebhookStatus_PENDING    WebhookStatus = "pending"
    WebhookStatus_PAUSED     WebhookStatus = "paused"
)

// WebhookSubscription is a string enum.
type WebhookSubscription string

const (
    WebhookSubscription_BILLINGDATECHANGE                              WebhookSubscription = "billing_date_change"
    WebhookSubscription_COMPONENTALLOCATIONCHANGE                      WebhookSubscription = "component_allocation_change"
    WebhookSubscription_CUSTOMERCREATE                                 WebhookSubscription = "customer_create"
    WebhookSubscription_CUSTOMERUPDATE                                 WebhookSubscription = "customer_update"
    WebhookSubscription_DUNNINGSTEPREACHED                             WebhookSubscription = "dunning_step_reached"
    WebhookSubscription_EXPIRINGCARD                                   WebhookSubscription = "expiring_card"
    WebhookSubscription_EXPIRATIONDATECHANGE                           WebhookSubscription = "expiration_date_change"
    WebhookSubscription_INVOICEISSUED                                  WebhookSubscription = "invoice_issued"
    WebhookSubscription_METEREDUSAGE                                   WebhookSubscription = "metered_usage"
    WebhookSubscription_PAYMENTFAILURE                                 WebhookSubscription = "payment_failure"
    WebhookSubscription_PAYMENTSUCCESS                                 WebhookSubscription = "payment_success"
    WebhookSubscription_DIRECTDEBITPAYMENTPENDING                      WebhookSubscription = "direct_debit_payment_pending"
    WebhookSubscription_DIRECTDEBITPAYMENTPAIDOUT                      WebhookSubscription = "direct_debit_payment_paid_out"
    WebhookSubscription_DIRECTDEBITPAYMENTREJECTED                     WebhookSubscription = "direct_debit_payment_rejected"
    WebhookSubscription_PREPAIDSUBSCRIPTIONBALANCECHANGED              WebhookSubscription = "prepaid_subscription_balance_changed"
    WebhookSubscription_PREPAIDUSAGE                                   WebhookSubscription = "prepaid_usage"
    WebhookSubscription_REFUNDFAILURE                                  WebhookSubscription = "refund_failure"
    WebhookSubscription_REFUNDSUCCESS                                  WebhookSubscription = "refund_success"
    WebhookSubscription_RENEWALFAILURE                                 WebhookSubscription = "renewal_failure"
    WebhookSubscription_RENEWALSUCCESS                                 WebhookSubscription = "renewal_success"
    WebhookSubscription_SIGNUPFAILURE                                  WebhookSubscription = "signup_failure"
    WebhookSubscription_SIGNUPSUCCESS                                  WebhookSubscription = "signup_success"
    WebhookSubscription_STATEMENTCLOSED                                WebhookSubscription = "statement_closed"
    WebhookSubscription_STATEMENTSETTLED                               WebhookSubscription = "statement_settled"
    WebhookSubscription_SUBSCRIPTIONCARDUPDATE                         WebhookSubscription = "subscription_card_update"
    WebhookSubscription_SUBSCRIPTIONGROUPCARDUPDATE                    WebhookSubscription = "subscription_group_card_update"
    WebhookSubscription_SUBSCRIPTIONPRODUCTCHANGE                      WebhookSubscription = "subscription_product_change"
    WebhookSubscription_SUBSCRIPTIONSTATECHANGE                        WebhookSubscription = "subscription_state_change"
    WebhookSubscription_TRIALENDNOTICE                                 WebhookSubscription = "trial_end_notice"
    WebhookSubscription_UPCOMINGRENEWALNOTICE                          WebhookSubscription = "upcoming_renewal_notice"
    WebhookSubscription_UPGRADEDOWNGRADEFAILURE                        WebhookSubscription = "upgrade_downgrade_failure"
    WebhookSubscription_UPGRADEDOWNGRADESUCCESS                        WebhookSubscription = "upgrade_downgrade_success"
    WebhookSubscription_PENDINGCANCELLATIONCHANGE                      WebhookSubscription = "pending_cancellation_change"
    WebhookSubscription_SUBSCRIPTIONPREPAYMENTACCOUNTBALANCECHANGED    WebhookSubscription = "subscription_prepayment_account_balance_changed"
    WebhookSubscription_SUBSCRIPTIONSERVICECREDITACCOUNTBALANCECHANGED WebhookSubscription = "subscription_service_credit_account_balance_changed"
)
