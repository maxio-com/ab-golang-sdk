// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "errors"
    "fmt"
)

// AllVaults is a string enum.
// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
type AllVaults string

// MarshalJSON implements the json.Marshaler interface for AllVaults.
// It customizes the JSON marshaling process for AllVaults objects.
func (e AllVaults) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for AllVaults")
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllVaults.
// It customizes the JSON unmarshaling process for AllVaults objects.
func (e *AllVaults) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = AllVaults(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to AllVaults")
    }
    return nil
}

// Checks whether the value is actually a member of AllVaults.
func (e AllVaults) isValid() bool {
    switch e {
    case AllVaults_ADYEN,
        AllVaults_AUTHORIZENET,
        AllVaults_BEANSTREAM,
        AllVaults_BLUESNAP,
        AllVaults_BOGUS,
        AllVaults_BRAINTREE1,
        AllVaults_BRAINTREEBLUE,
        AllVaults_CHECKOUT,
        AllVaults_CYBERSOURCE,
        AllVaults_ELAVON,
        AllVaults_EWAY,
        AllVaults_EWAYRAPID,
        AllVaults_EWAYRAPIDSTD,
        AllVaults_FIRSTDATA,
        AllVaults_FORTE,
        AllVaults_GOCARDLESS,
        AllVaults_LITLE,
        AllVaults_MAXIOPAYMENTS,
        AllVaults_MAXP,
        AllVaults_MODUSLINK,
        AllVaults_MONERIS,
        AllVaults_NMI,
        AllVaults_ORBITAL,
        AllVaults_PAYMENTEXPRESS,
        AllVaults_PAYMILL,
        AllVaults_PAYPAL,
        AllVaults_PAYPALCOMPLETE,
        AllVaults_PIN,
        AllVaults_SQUARE,
        AllVaults_STRIPE,
        AllVaults_STRIPECONNECT,
        AllVaults_TRUSTCOMMERCE,
        AllVaults_UNIPAAS,
        AllVaults_WIRECARD:
        return true
    }
    return false
}

const (
    AllVaults_ADYEN          AllVaults = "adyen"
    AllVaults_AUTHORIZENET   AllVaults = "authorizenet"
    AllVaults_BEANSTREAM     AllVaults = "beanstream"
    AllVaults_BLUESNAP       AllVaults = "blue_snap"
    AllVaults_BOGUS          AllVaults = "bogus"
    AllVaults_BRAINTREE1     AllVaults = "braintree1"
    AllVaults_BRAINTREEBLUE  AllVaults = "braintree_blue"
    AllVaults_CHECKOUT       AllVaults = "checkout"
    AllVaults_CYBERSOURCE    AllVaults = "cybersource"
    AllVaults_ELAVON         AllVaults = "elavon"
    AllVaults_EWAY           AllVaults = "eway"
    AllVaults_EWAYRAPID      AllVaults = "eway_rapid"
    AllVaults_EWAYRAPIDSTD   AllVaults = "eway_rapid_std"
    AllVaults_FIRSTDATA      AllVaults = "firstdata"
    AllVaults_FORTE          AllVaults = "forte"
    AllVaults_GOCARDLESS     AllVaults = "gocardless"
    AllVaults_LITLE          AllVaults = "litle"
    AllVaults_MAXIOPAYMENTS  AllVaults = "maxio_payments"
    AllVaults_MAXP           AllVaults = "maxp"
    AllVaults_MODUSLINK      AllVaults = "moduslink"
    AllVaults_MONERIS        AllVaults = "moneris"
    AllVaults_NMI            AllVaults = "nmi"
    AllVaults_ORBITAL        AllVaults = "orbital"
    AllVaults_PAYMENTEXPRESS AllVaults = "payment_express"
    AllVaults_PAYMILL        AllVaults = "paymill"
    AllVaults_PAYPAL         AllVaults = "paypal"
    AllVaults_PAYPALCOMPLETE AllVaults = "paypal_complete"
    AllVaults_PIN            AllVaults = "pin"
    AllVaults_SQUARE         AllVaults = "square"
    AllVaults_STRIPE         AllVaults = "stripe"
    AllVaults_STRIPECONNECT  AllVaults = "stripe_connect"
    AllVaults_TRUSTCOMMERCE  AllVaults = "trust_commerce"
    AllVaults_UNIPAAS        AllVaults = "unipaas"
    AllVaults_WIRECARD       AllVaults = "wirecard"
)

// AllocationPreviewDirection is a string enum.
type AllocationPreviewDirection string

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewDirection.
// It customizes the JSON marshaling process for AllocationPreviewDirection objects.
func (e AllocationPreviewDirection) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for AllocationPreviewDirection")
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewDirection.
// It customizes the JSON unmarshaling process for AllocationPreviewDirection objects.
func (e *AllocationPreviewDirection) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = AllocationPreviewDirection(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to AllocationPreviewDirection")
    }
    return nil
}

// Checks whether the value is actually a member of AllocationPreviewDirection.
func (e AllocationPreviewDirection) isValid() bool {
    switch e {
    case AllocationPreviewDirection_UPGRADE,
        AllocationPreviewDirection_DOWNGRADE:
        return true
    }
    return false
}

const (
    AllocationPreviewDirection_UPGRADE   AllocationPreviewDirection = "upgrade"
    AllocationPreviewDirection_DOWNGRADE AllocationPreviewDirection = "downgrade"
)

// AllocationPreviewLineItemKind is a string enum.
// A handle for the line item kind for allocation preview
type AllocationPreviewLineItemKind string

// MarshalJSON implements the json.Marshaler interface for AllocationPreviewLineItemKind.
// It customizes the JSON marshaling process for AllocationPreviewLineItemKind objects.
func (e AllocationPreviewLineItemKind) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for AllocationPreviewLineItemKind")
}

// UnmarshalJSON implements the json.Unmarshaler interface for AllocationPreviewLineItemKind.
// It customizes the JSON unmarshaling process for AllocationPreviewLineItemKind objects.
func (e *AllocationPreviewLineItemKind) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = AllocationPreviewLineItemKind(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to AllocationPreviewLineItemKind")
    }
    return nil
}

// Checks whether the value is actually a member of AllocationPreviewLineItemKind.
func (e AllocationPreviewLineItemKind) isValid() bool {
    switch e {
    case AllocationPreviewLineItemKind_QUANTITYBASEDCOMPONENT,
        AllocationPreviewLineItemKind_ONOFFCOMPONENT,
        AllocationPreviewLineItemKind_COUPON,
        AllocationPreviewLineItemKind_TAX:
        return true
    }
    return false
}

const (
    AllocationPreviewLineItemKind_QUANTITYBASEDCOMPONENT AllocationPreviewLineItemKind = "quantity_based_component"
    AllocationPreviewLineItemKind_ONOFFCOMPONENT         AllocationPreviewLineItemKind = "on_off_component"
    AllocationPreviewLineItemKind_COUPON                 AllocationPreviewLineItemKind = "coupon"
    AllocationPreviewLineItemKind_TAX                    AllocationPreviewLineItemKind = "tax"
)

// ApplePayVault is a string enum.
// The vault that stores the payment profile with the provided vault_token.
type ApplePayVault string

// MarshalJSON implements the json.Marshaler interface for ApplePayVault.
// It customizes the JSON marshaling process for ApplePayVault objects.
func (e ApplePayVault) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ApplePayVault")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApplePayVault.
// It customizes the JSON unmarshaling process for ApplePayVault objects.
func (e *ApplePayVault) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ApplePayVault(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ApplePayVault")
    }
    return nil
}

// Checks whether the value is actually a member of ApplePayVault.
func (e ApplePayVault) isValid() bool {
    switch e {
    case ApplePayVault_BRAINTREEBLUE:
        return true
    }
    return false
}

const (
    ApplePayVault_BRAINTREEBLUE ApplePayVault = "braintree_blue"
)

// AutoInvite is a int enum.
type AutoInvite int

// MarshalJSON implements the json.Marshaler interface for AutoInvite.
// It customizes the JSON marshaling process for AutoInvite objects.
func (e AutoInvite) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("%v", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for AutoInvite")
}

// UnmarshalJSON implements the json.Unmarshaler interface for AutoInvite.
// It customizes the JSON unmarshaling process for AutoInvite objects.
func (e *AutoInvite) UnmarshalJSON(input []byte) error {
    var enumValue int
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = AutoInvite(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to AutoInvite")
    }
    return nil
}

// Checks whether the value is actually a member of AutoInvite.
func (e AutoInvite) isValid() bool {
    switch e {
    case AutoInvite_NO,
        AutoInvite_YES:
        return true
    }
    return false
}

const (
    AutoInvite_NO  AutoInvite = 0
    AutoInvite_YES AutoInvite = 1
)

// BankAccountHolderType is a string enum.
// Defaults to personal
type BankAccountHolderType string

// MarshalJSON implements the json.Marshaler interface for BankAccountHolderType.
// It customizes the JSON marshaling process for BankAccountHolderType objects.
func (e BankAccountHolderType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for BankAccountHolderType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountHolderType.
// It customizes the JSON unmarshaling process for BankAccountHolderType objects.
func (e *BankAccountHolderType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = BankAccountHolderType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to BankAccountHolderType")
    }
    return nil
}

// Checks whether the value is actually a member of BankAccountHolderType.
func (e BankAccountHolderType) isValid() bool {
    switch e {
    case BankAccountHolderType_PERSONAL,
        BankAccountHolderType_BUSINESS:
        return true
    }
    return false
}

const (
    BankAccountHolderType_PERSONAL BankAccountHolderType = "personal"
    BankAccountHolderType_BUSINESS BankAccountHolderType = "business"
)

// BankAccountType is a string enum.
// Defaults to checking
type BankAccountType string

// MarshalJSON implements the json.Marshaler interface for BankAccountType.
// It customizes the JSON marshaling process for BankAccountType objects.
func (e BankAccountType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for BankAccountType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountType.
// It customizes the JSON unmarshaling process for BankAccountType objects.
func (e *BankAccountType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = BankAccountType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to BankAccountType")
    }
    return nil
}

// Checks whether the value is actually a member of BankAccountType.
func (e BankAccountType) isValid() bool {
    switch e {
    case BankAccountType_CHECKING,
        BankAccountType_SAVINGS:
        return true
    }
    return false
}

const (
    BankAccountType_CHECKING BankAccountType = "checking"
    BankAccountType_SAVINGS  BankAccountType = "savings"
)

// BankAccountVault is a string enum.
// The vault that stores the payment profile with the provided vault_token. Use `bogus` for testing.
type BankAccountVault string

// MarshalJSON implements the json.Marshaler interface for BankAccountVault.
// It customizes the JSON marshaling process for BankAccountVault objects.
func (e BankAccountVault) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for BankAccountVault")
}

// UnmarshalJSON implements the json.Unmarshaler interface for BankAccountVault.
// It customizes the JSON unmarshaling process for BankAccountVault objects.
func (e *BankAccountVault) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = BankAccountVault(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to BankAccountVault")
    }
    return nil
}

// Checks whether the value is actually a member of BankAccountVault.
func (e BankAccountVault) isValid() bool {
    switch e {
    case BankAccountVault_AUTHORIZENET,
        BankAccountVault_BLUESNAP,
        BankAccountVault_BOGUS,
        BankAccountVault_FORTE,
        BankAccountVault_GOCARDLESS,
        BankAccountVault_MAXIOPAYMENTS,
        BankAccountVault_MAXP,
        BankAccountVault_STRIPECONNECT:
        return true
    }
    return false
}

const (
    BankAccountVault_AUTHORIZENET  BankAccountVault = "authorizenet"
    BankAccountVault_BLUESNAP      BankAccountVault = "blue_snap"
    BankAccountVault_BOGUS         BankAccountVault = "bogus"
    BankAccountVault_FORTE         BankAccountVault = "forte"
    BankAccountVault_GOCARDLESS    BankAccountVault = "gocardless"
    BankAccountVault_MAXIOPAYMENTS BankAccountVault = "maxio_payments"
    BankAccountVault_MAXP          BankAccountVault = "maxp"
    BankAccountVault_STRIPECONNECT BankAccountVault = "stripe_connect"
)

// BasicDateField is a string enum.
// Allows to filter by `created_at` or `updated_at`.
type BasicDateField string

// MarshalJSON implements the json.Marshaler interface for BasicDateField.
// It customizes the JSON marshaling process for BasicDateField objects.
func (e BasicDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for BasicDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for BasicDateField.
// It customizes the JSON unmarshaling process for BasicDateField objects.
func (e *BasicDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = BasicDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to BasicDateField")
    }
    return nil
}

// Checks whether the value is actually a member of BasicDateField.
func (e BasicDateField) isValid() bool {
    switch e {
    case BasicDateField_UPDATEDAT,
        BasicDateField_CREATEDAT:
        return true
    }
    return false
}

const (
    BasicDateField_UPDATEDAT BasicDateField = "updated_at"
    BasicDateField_CREATEDAT BasicDateField = "created_at"
)

// BillingManifestLineItemKind is a string enum.
// A handle for the billing manifest line item kind
type BillingManifestLineItemKind string

// MarshalJSON implements the json.Marshaler interface for BillingManifestLineItemKind.
// It customizes the JSON marshaling process for BillingManifestLineItemKind objects.
func (e BillingManifestLineItemKind) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for BillingManifestLineItemKind")
}

// UnmarshalJSON implements the json.Unmarshaler interface for BillingManifestLineItemKind.
// It customizes the JSON unmarshaling process for BillingManifestLineItemKind objects.
func (e *BillingManifestLineItemKind) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = BillingManifestLineItemKind(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to BillingManifestLineItemKind")
    }
    return nil
}

// Checks whether the value is actually a member of BillingManifestLineItemKind.
func (e BillingManifestLineItemKind) isValid() bool {
    switch e {
    case BillingManifestLineItemKind_BASELINE,
        BillingManifestLineItemKind_INITIAL,
        BillingManifestLineItemKind_TRIAL,
        BillingManifestLineItemKind_COUPON,
        BillingManifestLineItemKind_COMPONENT,
        BillingManifestLineItemKind_TAX:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for CancellationMethod.
// It customizes the JSON marshaling process for CancellationMethod objects.
func (e CancellationMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CancellationMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CancellationMethod.
// It customizes the JSON unmarshaling process for CancellationMethod objects.
func (e *CancellationMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CancellationMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CancellationMethod")
    }
    return nil
}

// Checks whether the value is actually a member of CancellationMethod.
func (e CancellationMethod) isValid() bool {
    switch e {
    case CancellationMethod_MERCHANTUI,
        CancellationMethod_MERCHANTAPI,
        CancellationMethod_DUNNING,
        CancellationMethod_BILLINGPORTAL,
        CancellationMethod_UNKNOWN,
        CancellationMethod_IMPORTED:
        return true
    }
    return false
}

const (
    CancellationMethod_MERCHANTUI    CancellationMethod = "merchant_ui"
    CancellationMethod_MERCHANTAPI   CancellationMethod = "merchant_api"
    CancellationMethod_DUNNING       CancellationMethod = "dunning"
    CancellationMethod_BILLINGPORTAL CancellationMethod = "billing_portal"
    CancellationMethod_UNKNOWN       CancellationMethod = "unknown"
    CancellationMethod_IMPORTED      CancellationMethod = "imported"
)

// CardType is a string enum.
// The type of card used.
type CardType string

// MarshalJSON implements the json.Marshaler interface for CardType.
// It customizes the JSON marshaling process for CardType objects.
func (e CardType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CardType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CardType.
// It customizes the JSON unmarshaling process for CardType objects.
func (e *CardType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CardType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CardType")
    }
    return nil
}

// Checks whether the value is actually a member of CardType.
func (e CardType) isValid() bool {
    switch e {
    case CardType_VISA,
        CardType_MASTER,
        CardType_ELO,
        CardType_CABAL,
        CardType_ALELO,
        CardType_DISCOVER,
        CardType_AMERICANEXPRESS,
        CardType_NARANJA,
        CardType_DINERSCLUB,
        CardType_JCB,
        CardType_DANKORT,
        CardType_MAESTRO,
        CardType_MAESTRONOLUHN,
        CardType_FORBRUGSFORENINGEN,
        CardType_SODEXO,
        CardType_ALIA,
        CardType_VR,
        CardType_UNIONPAY,
        CardType_CARNET,
        CardType_CARTESBANCAIRES,
        CardType_OLIMPICA,
        CardType_CREDITEL,
        CardType_CONFIABLE,
        CardType_SYNCHRONY,
        CardType_ROUTEX,
        CardType_MADA,
        CardType_BPPLUS,
        CardType_PASSCARD,
        CardType_EDENRED,
        CardType_ANDA,
        CardType_TARJETAD,
        CardType_HIPERCARD,
        CardType_BOGUS,
        CardType_ENUMSWITCH,
        CardType_SOLO,
        CardType_LASER:
        return true
    }
    return false
}

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

// ChargebackStatus is a string enum.
// The current chargeback status.
type ChargebackStatus string

// MarshalJSON implements the json.Marshaler interface for ChargebackStatus.
// It customizes the JSON marshaling process for ChargebackStatus objects.
func (e ChargebackStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ChargebackStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChargebackStatus.
// It customizes the JSON unmarshaling process for ChargebackStatus objects.
func (e *ChargebackStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ChargebackStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ChargebackStatus")
    }
    return nil
}

// Checks whether the value is actually a member of ChargebackStatus.
func (e ChargebackStatus) isValid() bool {
    switch e {
    case ChargebackStatus_OPEN,
        ChargebackStatus_LOST,
        ChargebackStatus_WON,
        ChargebackStatus_CLOSED:
        return true
    }
    return false
}

const (
    ChargebackStatus_OPEN   ChargebackStatus = "open"
    ChargebackStatus_LOST   ChargebackStatus = "lost"
    ChargebackStatus_WON    ChargebackStatus = "won"
    ChargebackStatus_CLOSED ChargebackStatus = "closed"
)

// CleanupScope is a string enum.
// all: Will clear all products, customers, and related subscriptions from the site. customers: Will clear only customers and related subscriptions (leaving the products untouched) for the site. Revenue will also be reset to 0.
type CleanupScope string

// MarshalJSON implements the json.Marshaler interface for CleanupScope.
// It customizes the JSON marshaling process for CleanupScope objects.
func (e CleanupScope) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CleanupScope")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CleanupScope.
// It customizes the JSON unmarshaling process for CleanupScope objects.
func (e *CleanupScope) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CleanupScope(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CleanupScope")
    }
    return nil
}

// Checks whether the value is actually a member of CleanupScope.
func (e CleanupScope) isValid() bool {
    switch e {
    case CleanupScope_ALL,
        CleanupScope_CUSTOMERS:
        return true
    }
    return false
}

const (
    CleanupScope_ALL       CleanupScope = "all"
    CleanupScope_CUSTOMERS CleanupScope = "customers"
)

// CollectionMethod is a string enum.
// The type of payment collection to be used in the subscription. For legacy Statements Architecture valid options are - `invoice`, `automatic`. For current Relationship Invoicing Architecture valid options are - `remittance`, `automatic`, `prepaid`.
type CollectionMethod string

// MarshalJSON implements the json.Marshaler interface for CollectionMethod.
// It customizes the JSON marshaling process for CollectionMethod objects.
func (e CollectionMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CollectionMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CollectionMethod.
// It customizes the JSON unmarshaling process for CollectionMethod objects.
func (e *CollectionMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CollectionMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CollectionMethod")
    }
    return nil
}

// Checks whether the value is actually a member of CollectionMethod.
func (e CollectionMethod) isValid() bool {
    switch e {
    case CollectionMethod_AUTOMATIC,
        CollectionMethod_REMITTANCE,
        CollectionMethod_PREPAID,
        CollectionMethod_INVOICE:
        return true
    }
    return false
}

const (
    CollectionMethod_AUTOMATIC  CollectionMethod = "automatic"
    CollectionMethod_REMITTANCE CollectionMethod = "remittance"
    CollectionMethod_PREPAID    CollectionMethod = "prepaid"
    CollectionMethod_INVOICE    CollectionMethod = "invoice"
)

// ComponentKind is a string enum.
// A handle for the component type
type ComponentKind string

// MarshalJSON implements the json.Marshaler interface for ComponentKind.
// It customizes the JSON marshaling process for ComponentKind objects.
func (e ComponentKind) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ComponentKind")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ComponentKind.
// It customizes the JSON unmarshaling process for ComponentKind objects.
func (e *ComponentKind) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ComponentKind(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ComponentKind")
    }
    return nil
}

// Checks whether the value is actually a member of ComponentKind.
func (e ComponentKind) isValid() bool {
    switch e {
    case ComponentKind_METEREDCOMPONENT,
        ComponentKind_QUANTITYBASEDCOMPONENT,
        ComponentKind_ONOFFCOMPONENT,
        ComponentKind_PREPAIDUSAGECOMPONENT,
        ComponentKind_EVENTBASEDCOMPONENT:
        return true
    }
    return false
}

const (
    ComponentKind_METEREDCOMPONENT       ComponentKind = "metered_component"
    ComponentKind_QUANTITYBASEDCOMPONENT ComponentKind = "quantity_based_component"
    ComponentKind_ONOFFCOMPONENT         ComponentKind = "on_off_component"
    ComponentKind_PREPAIDUSAGECOMPONENT  ComponentKind = "prepaid_usage_component"
    ComponentKind_EVENTBASEDCOMPONENT    ComponentKind = "event_based_component"
)

// CompoundingStrategy is a string enum.
// Applicable only to stackable coupons. For `compound`, Percentage-based discounts will be calculated against the remaining price, after prior discounts have been calculated. For `full-price`, Percentage-based discounts will always be calculated against the original item price, before other discounts are applied.
type CompoundingStrategy string

// MarshalJSON implements the json.Marshaler interface for CompoundingStrategy.
// It customizes the JSON marshaling process for CompoundingStrategy objects.
func (e CompoundingStrategy) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CompoundingStrategy")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CompoundingStrategy.
// It customizes the JSON unmarshaling process for CompoundingStrategy objects.
func (e *CompoundingStrategy) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CompoundingStrategy(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CompoundingStrategy")
    }
    return nil
}

// Checks whether the value is actually a member of CompoundingStrategy.
func (e CompoundingStrategy) isValid() bool {
    switch e {
    case CompoundingStrategy_COMPOUND,
        CompoundingStrategy_FULLPRICE:
        return true
    }
    return false
}

const (
    CompoundingStrategy_COMPOUND  CompoundingStrategy = "compound"
    CompoundingStrategy_FULLPRICE CompoundingStrategy = "full-price"
)

// CreateInvoiceStatus is a string enum.
type CreateInvoiceStatus string

// MarshalJSON implements the json.Marshaler interface for CreateInvoiceStatus.
// It customizes the JSON marshaling process for CreateInvoiceStatus objects.
func (e CreateInvoiceStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreateInvoiceStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateInvoiceStatus.
// It customizes the JSON unmarshaling process for CreateInvoiceStatus objects.
func (e *CreateInvoiceStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreateInvoiceStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreateInvoiceStatus")
    }
    return nil
}

// Checks whether the value is actually a member of CreateInvoiceStatus.
func (e CreateInvoiceStatus) isValid() bool {
    switch e {
    case CreateInvoiceStatus_DRAFT,
        CreateInvoiceStatus_OPEN:
        return true
    }
    return false
}

const (
    CreateInvoiceStatus_DRAFT CreateInvoiceStatus = "draft"
    CreateInvoiceStatus_OPEN  CreateInvoiceStatus = "open"
)

// CreatePrepaymentMethod is a string enum.
// :- When the `method` specified is `"credit_card_on_file"`, the prepayment amount will be collected using the default credit card payment profile and applied to the prepayment account balance. This is especially useful for manual replenishment of prepaid subscriptions.
type CreatePrepaymentMethod string

// MarshalJSON implements the json.Marshaler interface for CreatePrepaymentMethod.
// It customizes the JSON marshaling process for CreatePrepaymentMethod objects.
func (e CreatePrepaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreatePrepaymentMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreatePrepaymentMethod.
// It customizes the JSON unmarshaling process for CreatePrepaymentMethod objects.
func (e *CreatePrepaymentMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreatePrepaymentMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreatePrepaymentMethod")
    }
    return nil
}

// Checks whether the value is actually a member of CreatePrepaymentMethod.
func (e CreatePrepaymentMethod) isValid() bool {
    switch e {
    case CreatePrepaymentMethod_CHECK,
        CreatePrepaymentMethod_CASH,
        CreatePrepaymentMethod_MONEYORDER,
        CreatePrepaymentMethod_ACH,
        CreatePrepaymentMethod_PAYPALACCOUNT,
        CreatePrepaymentMethod_CREDITCARD,
        CreatePrepaymentMethod_CREDITCARDONFILE,
        CreatePrepaymentMethod_OTHER:
        return true
    }
    return false
}

const (
    CreatePrepaymentMethod_CHECK            CreatePrepaymentMethod = "check"
    CreatePrepaymentMethod_CASH             CreatePrepaymentMethod = "cash"
    CreatePrepaymentMethod_MONEYORDER       CreatePrepaymentMethod = "money_order"
    CreatePrepaymentMethod_ACH              CreatePrepaymentMethod = "ach"
    CreatePrepaymentMethod_PAYPALACCOUNT    CreatePrepaymentMethod = "paypal_account"
    CreatePrepaymentMethod_CREDITCARD       CreatePrepaymentMethod = "credit_card"
    CreatePrepaymentMethod_CREDITCARDONFILE CreatePrepaymentMethod = "credit_card_on_file"
    CreatePrepaymentMethod_OTHER            CreatePrepaymentMethod = "other"
)

// CreateSignupProformaPreviewInclude is a string enum.
type CreateSignupProformaPreviewInclude string

// MarshalJSON implements the json.Marshaler interface for CreateSignupProformaPreviewInclude.
// It customizes the JSON marshaling process for CreateSignupProformaPreviewInclude objects.
func (e CreateSignupProformaPreviewInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreateSignupProformaPreviewInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreateSignupProformaPreviewInclude.
// It customizes the JSON unmarshaling process for CreateSignupProformaPreviewInclude objects.
func (e *CreateSignupProformaPreviewInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreateSignupProformaPreviewInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreateSignupProformaPreviewInclude")
    }
    return nil
}

// Checks whether the value is actually a member of CreateSignupProformaPreviewInclude.
func (e CreateSignupProformaPreviewInclude) isValid() bool {
    switch e {
    case CreateSignupProformaPreviewInclude_NEXTPROFORMAINVOICE:
        return true
    }
    return false
}

const (
    CreateSignupProformaPreviewInclude_NEXTPROFORMAINVOICE CreateSignupProformaPreviewInclude = "next_proforma_invoice"
)

// CreditCardVault is a string enum.
// The vault that stores the payment profile with the provided `vault_token`. Use `bogus` for testing.
type CreditCardVault string

// MarshalJSON implements the json.Marshaler interface for CreditCardVault.
// It customizes the JSON marshaling process for CreditCardVault objects.
func (e CreditCardVault) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreditCardVault")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditCardVault.
// It customizes the JSON unmarshaling process for CreditCardVault objects.
func (e *CreditCardVault) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreditCardVault(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreditCardVault")
    }
    return nil
}

// Checks whether the value is actually a member of CreditCardVault.
func (e CreditCardVault) isValid() bool {
    switch e {
    case CreditCardVault_ADYEN,
        CreditCardVault_AUTHORIZENET,
        CreditCardVault_BEANSTREAM,
        CreditCardVault_BLUESNAP,
        CreditCardVault_BOGUS,
        CreditCardVault_BRAINTREE1,
        CreditCardVault_BRAINTREEBLUE,
        CreditCardVault_CHECKOUT,
        CreditCardVault_CYBERSOURCE,
        CreditCardVault_ELAVON,
        CreditCardVault_EWAY,
        CreditCardVault_EWAYRAPID,
        CreditCardVault_EWAYRAPIDSTD,
        CreditCardVault_FIRSTDATA,
        CreditCardVault_FORTE,
        CreditCardVault_LITLE,
        CreditCardVault_MAXIOPAYMENTS,
        CreditCardVault_MAXP,
        CreditCardVault_MODUSLINK,
        CreditCardVault_MONERIS,
        CreditCardVault_NMI,
        CreditCardVault_ORBITAL,
        CreditCardVault_PAYMENTEXPRESS,
        CreditCardVault_PAYMILL,
        CreditCardVault_PAYPAL,
        CreditCardVault_PAYPALCOMPLETE,
        CreditCardVault_PIN,
        CreditCardVault_SQUARE,
        CreditCardVault_STRIPE,
        CreditCardVault_STRIPECONNECT,
        CreditCardVault_TRUSTCOMMERCE,
        CreditCardVault_UNIPAAS,
        CreditCardVault_WIRECARD:
        return true
    }
    return false
}

const (
    CreditCardVault_ADYEN          CreditCardVault = "adyen"
    CreditCardVault_AUTHORIZENET   CreditCardVault = "authorizenet"
    CreditCardVault_BEANSTREAM     CreditCardVault = "beanstream"
    CreditCardVault_BLUESNAP       CreditCardVault = "blue_snap"
    CreditCardVault_BOGUS          CreditCardVault = "bogus"
    CreditCardVault_BRAINTREE1     CreditCardVault = "braintree1"
    CreditCardVault_BRAINTREEBLUE  CreditCardVault = "braintree_blue"
    CreditCardVault_CHECKOUT       CreditCardVault = "checkout"
    CreditCardVault_CYBERSOURCE    CreditCardVault = "cybersource"
    CreditCardVault_ELAVON         CreditCardVault = "elavon"
    CreditCardVault_EWAY           CreditCardVault = "eway"
    CreditCardVault_EWAYRAPID      CreditCardVault = "eway_rapid"
    CreditCardVault_EWAYRAPIDSTD   CreditCardVault = "eway_rapid_std"
    CreditCardVault_FIRSTDATA      CreditCardVault = "firstdata"
    CreditCardVault_FORTE          CreditCardVault = "forte"
    CreditCardVault_LITLE          CreditCardVault = "litle"
    CreditCardVault_MAXIOPAYMENTS  CreditCardVault = "maxio_payments"
    CreditCardVault_MAXP           CreditCardVault = "maxp"
    CreditCardVault_MODUSLINK      CreditCardVault = "moduslink"
    CreditCardVault_MONERIS        CreditCardVault = "moneris"
    CreditCardVault_NMI            CreditCardVault = "nmi"
    CreditCardVault_ORBITAL        CreditCardVault = "orbital"
    CreditCardVault_PAYMENTEXPRESS CreditCardVault = "payment_express"
    CreditCardVault_PAYMILL        CreditCardVault = "paymill"
    CreditCardVault_PAYPAL         CreditCardVault = "paypal"
    CreditCardVault_PAYPALCOMPLETE CreditCardVault = "paypal_complete"
    CreditCardVault_PIN            CreditCardVault = "pin"
    CreditCardVault_SQUARE         CreditCardVault = "square"
    CreditCardVault_STRIPE         CreditCardVault = "stripe"
    CreditCardVault_STRIPECONNECT  CreditCardVault = "stripe_connect"
    CreditCardVault_TRUSTCOMMERCE  CreditCardVault = "trust_commerce"
    CreditCardVault_UNIPAAS        CreditCardVault = "unipaas"
    CreditCardVault_WIRECARD       CreditCardVault = "wirecard"
)

// CreditNoteStatus is a string enum.
// Current status of the credit note.
type CreditNoteStatus string

// MarshalJSON implements the json.Marshaler interface for CreditNoteStatus.
// It customizes the JSON marshaling process for CreditNoteStatus objects.
func (e CreditNoteStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreditNoteStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditNoteStatus.
// It customizes the JSON unmarshaling process for CreditNoteStatus objects.
func (e *CreditNoteStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreditNoteStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreditNoteStatus")
    }
    return nil
}

// Checks whether the value is actually a member of CreditNoteStatus.
func (e CreditNoteStatus) isValid() bool {
    switch e {
    case CreditNoteStatus_OPEN,
        CreditNoteStatus_APPLIED:
        return true
    }
    return false
}

const (
    CreditNoteStatus_OPEN    CreditNoteStatus = "open"
    CreditNoteStatus_APPLIED CreditNoteStatus = "applied"
)

// CreditScheme is a string enum.
type CreditScheme string

// MarshalJSON implements the json.Marshaler interface for CreditScheme.
// It customizes the JSON marshaling process for CreditScheme objects.
func (e CreditScheme) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreditScheme")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditScheme.
// It customizes the JSON unmarshaling process for CreditScheme objects.
func (e *CreditScheme) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreditScheme(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreditScheme")
    }
    return nil
}

// Checks whether the value is actually a member of CreditScheme.
func (e CreditScheme) isValid() bool {
    switch e {
    case CreditScheme_NONE,
        CreditScheme_CREDIT,
        CreditScheme_REFUND:
        return true
    }
    return false
}

const (
    CreditScheme_NONE   CreditScheme = "none"
    CreditScheme_CREDIT CreditScheme = "credit"
    CreditScheme_REFUND CreditScheme = "refund"
)

// CreditType is a string enum.
// The type of credit to be created when upgrading/downgrading. Defaults to the component and then site setting if one is not provided.
// Available values: `full`, `prorated`, `none`.
type CreditType string

// MarshalJSON implements the json.Marshaler interface for CreditType.
// It customizes the JSON marshaling process for CreditType objects.
func (e CreditType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CreditType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CreditType.
// It customizes the JSON unmarshaling process for CreditType objects.
func (e *CreditType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CreditType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CreditType")
    }
    return nil
}

// Checks whether the value is actually a member of CreditType.
func (e CreditType) isValid() bool {
    switch e {
    case CreditType_FULL,
        CreditType_PRORATED,
        CreditType_NONE:
        return true
    }
    return false
}

const (
    CreditType_FULL     CreditType = "full"
    CreditType_PRORATED CreditType = "prorated"
    CreditType_NONE     CreditType = "none"
)

// CurrencyPriceRole is a string enum.
// Role for the price.
type CurrencyPriceRole string

// MarshalJSON implements the json.Marshaler interface for CurrencyPriceRole.
// It customizes the JSON marshaling process for CurrencyPriceRole objects.
func (e CurrencyPriceRole) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CurrencyPriceRole")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CurrencyPriceRole.
// It customizes the JSON unmarshaling process for CurrencyPriceRole objects.
func (e *CurrencyPriceRole) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CurrencyPriceRole(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CurrencyPriceRole")
    }
    return nil
}

// Checks whether the value is actually a member of CurrencyPriceRole.
func (e CurrencyPriceRole) isValid() bool {
    switch e {
    case CurrencyPriceRole_BASELINE,
        CurrencyPriceRole_TRIAL,
        CurrencyPriceRole_INITIAL:
        return true
    }
    return false
}

const (
    CurrencyPriceRole_BASELINE CurrencyPriceRole = "baseline"
    CurrencyPriceRole_TRIAL    CurrencyPriceRole = "trial"
    CurrencyPriceRole_INITIAL  CurrencyPriceRole = "initial"
)

// CustomFieldOwner is a string enum.
type CustomFieldOwner string

// MarshalJSON implements the json.Marshaler interface for CustomFieldOwner.
// It customizes the JSON marshaling process for CustomFieldOwner objects.
func (e CustomFieldOwner) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for CustomFieldOwner")
}

// UnmarshalJSON implements the json.Unmarshaler interface for CustomFieldOwner.
// It customizes the JSON unmarshaling process for CustomFieldOwner objects.
func (e *CustomFieldOwner) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = CustomFieldOwner(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to CustomFieldOwner")
    }
    return nil
}

// Checks whether the value is actually a member of CustomFieldOwner.
func (e CustomFieldOwner) isValid() bool {
    switch e {
    case CustomFieldOwner_CUSTOMER,
        CustomFieldOwner_SUBSCRIPTION:
        return true
    }
    return false
}

const (
    CustomFieldOwner_CUSTOMER     CustomFieldOwner = "Customer"
    CustomFieldOwner_SUBSCRIPTION CustomFieldOwner = "Subscription"
)

// DebitNoteRole is a string enum.
// The role of the debit note.
type DebitNoteRole string

// MarshalJSON implements the json.Marshaler interface for DebitNoteRole.
// It customizes the JSON marshaling process for DebitNoteRole objects.
func (e DebitNoteRole) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for DebitNoteRole")
}

// UnmarshalJSON implements the json.Unmarshaler interface for DebitNoteRole.
// It customizes the JSON unmarshaling process for DebitNoteRole objects.
func (e *DebitNoteRole) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = DebitNoteRole(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to DebitNoteRole")
    }
    return nil
}

// Checks whether the value is actually a member of DebitNoteRole.
func (e DebitNoteRole) isValid() bool {
    switch e {
    case DebitNoteRole_CHARGEBACK,
        DebitNoteRole_REFUND:
        return true
    }
    return false
}

const (
    DebitNoteRole_CHARGEBACK DebitNoteRole = "chargeback"
    DebitNoteRole_REFUND     DebitNoteRole = "refund"
)

// DebitNoteStatus is a string enum.
// Current status of the debit note.
type DebitNoteStatus string

// MarshalJSON implements the json.Marshaler interface for DebitNoteStatus.
// It customizes the JSON marshaling process for DebitNoteStatus objects.
func (e DebitNoteStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for DebitNoteStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for DebitNoteStatus.
// It customizes the JSON unmarshaling process for DebitNoteStatus objects.
func (e *DebitNoteStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = DebitNoteStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to DebitNoteStatus")
    }
    return nil
}

// Checks whether the value is actually a member of DebitNoteStatus.
func (e DebitNoteStatus) isValid() bool {
    switch e {
    case DebitNoteStatus_OPEN,
        DebitNoteStatus_APPLIED,
        DebitNoteStatus_BANISHED,
        DebitNoteStatus_PAID:
        return true
    }
    return false
}

const (
    DebitNoteStatus_OPEN     DebitNoteStatus = "open"
    DebitNoteStatus_APPLIED  DebitNoteStatus = "applied"
    DebitNoteStatus_BANISHED DebitNoteStatus = "banished"
    DebitNoteStatus_PAID     DebitNoteStatus = "paid"
)

// Direction is a string enum.
type Direction string

// MarshalJSON implements the json.Marshaler interface for Direction.
// It customizes the JSON marshaling process for Direction objects.
func (e Direction) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for Direction")
}

// UnmarshalJSON implements the json.Unmarshaler interface for Direction.
// It customizes the JSON unmarshaling process for Direction objects.
func (e *Direction) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = Direction(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to Direction")
    }
    return nil
}

// Checks whether the value is actually a member of Direction.
func (e Direction) isValid() bool {
    switch e {
    case Direction_ASC,
        Direction_DESC:
        return true
    }
    return false
}

const (
    Direction_ASC  Direction = "asc"
    Direction_DESC Direction = "desc"
)

// DiscountType is a string enum.
type DiscountType string

// MarshalJSON implements the json.Marshaler interface for DiscountType.
// It customizes the JSON marshaling process for DiscountType objects.
func (e DiscountType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for DiscountType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for DiscountType.
// It customizes the JSON unmarshaling process for DiscountType objects.
func (e *DiscountType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = DiscountType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to DiscountType")
    }
    return nil
}

// Checks whether the value is actually a member of DiscountType.
func (e DiscountType) isValid() bool {
    switch e {
    case DiscountType_AMOUNT,
        DiscountType_PERCENT:
        return true
    }
    return false
}

const (
    DiscountType_AMOUNT  DiscountType = "amount"
    DiscountType_PERCENT DiscountType = "percent"
)

// EventKey is a string enum.
type EventKey string

// MarshalJSON implements the json.Marshaler interface for EventKey.
// It customizes the JSON marshaling process for EventKey objects.
func (e EventKey) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for EventKey")
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventKey.
// It customizes the JSON unmarshaling process for EventKey objects.
func (e *EventKey) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = EventKey(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to EventKey")
    }
    return nil
}

// Checks whether the value is actually a member of EventKey.
func (e EventKey) isValid() bool {
    switch e {
    case EventKey_PAYMENTSUCCESS,
        EventKey_PAYMENTFAILURE,
        EventKey_SIGNUPSUCCESS,
        EventKey_SIGNUPFAILURE,
        EventKey_DELAYEDSIGNUPCREATIONSUCCESS,
        EventKey_DELAYEDSIGNUPCREATIONFAILURE,
        EventKey_BILLINGDATECHANGE,
        EventKey_EXPIRATIONDATECHANGE,
        EventKey_RENEWALSUCCESS,
        EventKey_RENEWALFAILURE,
        EventKey_SUBSCRIPTIONSTATECHANGE,
        EventKey_SUBSCRIPTIONPRODUCTCHANGE,
        EventKey_PENDINGCANCELLATIONCHANGE,
        EventKey_EXPIRINGCARD,
        EventKey_CUSTOMERUPDATE,
        EventKey_CUSTOMERCREATE,
        EventKey_CUSTOMERDELETE,
        EventKey_COMPONENTALLOCATIONCHANGE,
        EventKey_METEREDUSAGE,
        EventKey_PREPAIDUSAGE,
        EventKey_UPGRADEDOWNGRADESUCCESS,
        EventKey_UPGRADEDOWNGRADEFAILURE,
        EventKey_STATEMENTCLOSED,
        EventKey_STATEMENTSETTLED,
        EventKey_SUBSCRIPTIONCARDUPDATE,
        EventKey_SUBSCRIPTIONGROUPCARDUPDATE,
        EventKey_SUBSCRIPTIONBANKACCOUNTUPDATE,
        EventKey_REFUNDSUCCESS,
        EventKey_REFUNDFAILURE,
        EventKey_UPCOMINGRENEWALNOTICE,
        EventKey_TRIALENDNOTICE,
        EventKey_DUNNINGSTEPREACHED,
        EventKey_INVOICEISSUED,
        EventKey_PREPAIDSUBSCRIPTIONBALANCECHANGED,
        EventKey_SUBSCRIPTIONGROUPSIGNUPSUCCESS,
        EventKey_SUBSCRIPTIONGROUPSIGNUPFAILURE,
        EventKey_DIRECTDEBITPAYMENTPAIDOUT,
        EventKey_DIRECTDEBITPAYMENTREJECTED,
        EventKey_DIRECTDEBITPAYMENTPENDING,
        EventKey_PENDINGPAYMENTCREATED,
        EventKey_PENDINGPAYMENTFAILED,
        EventKey_PENDINGPAYMENTCOMPLETED,
        EventKey_PROFORMAINVOICEISSUED,
        EventKey_SUBSCRIPTIONPREPAYMENTACCOUNTBALANCECHANGED,
        EventKey_SUBSCRIPTIONSERVICECREDITACCOUNTBALANCECHANGED,
        EventKey_CUSTOMFIELDVALUECHANGE,
        EventKey_ITEMPRICEPOINTCHANGED,
        EventKey_RENEWALSUCCESSRECREATED,
        EventKey_RENEWALFAILURERECREATED,
        EventKey_PAYMENTSUCCESSRECREATED,
        EventKey_PAYMENTFAILURERECREATED,
        EventKey_SUBSCRIPTIONDELETION,
        EventKey_SUBSCRIPTIONGROUPBANKACCOUNTUPDATE,
        EventKey_SUBSCRIPTIONPAYPALACCOUNTUPDATE,
        EventKey_SUBSCRIPTIONGROUPPAYPALACCOUNTUPDATE,
        EventKey_SUBSCRIPTIONCUSTOMERCHANGE,
        EventKey_ACCOUNTTRANSACTIONCHANGED,
        EventKey_GOCARDLESSPAYMENTPAIDOUT,
        EventKey_GOCARDLESSPAYMENTREJECTED,
        EventKey_GOCARDLESSPAYMENTPENDING,
        EventKey_STRIPEDIRECTDEBITPAYMENTPAIDOUT,
        EventKey_STRIPEDIRECTDEBITPAYMENTREJECTED,
        EventKey_STRIPEDIRECTDEBITPAYMENTPENDING,
        EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTPAIDOUT,
        EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTREJECTED,
        EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTPENDING,
        EventKey_INVOICEINCOLLECTIONSCANCELED,
        EventKey_SUBSCRIPTIONADDEDTOGROUP,
        EventKey_SUBSCRIPTIONREMOVEDFROMGROUP,
        EventKey_CHARGEBACKOPENED,
        EventKey_CHARGEBACKLOST,
        EventKey_CHARGEBACKACCEPTED,
        EventKey_CHARGEBACKCLOSED,
        EventKey_CHARGEBACKWON,
        EventKey_PAYMENTCOLLECTIONMETHODCHANGED,
        EventKey_COMPONENTBILLINGDATECHANGED,
        EventKey_SUBSCRIPTIONTERMRENEWALSCHEDULED,
        EventKey_SUBSCRIPTIONTERMRENEWALPENDING,
        EventKey_SUBSCRIPTIONTERMRENEWALACTIVATED,
        EventKey_SUBSCRIPTIONTERMRENEWALREMOVED:
        return true
    }
    return false
}

const (
    EventKey_PAYMENTSUCCESS                                 EventKey = "payment_success"
    EventKey_PAYMENTFAILURE                                 EventKey = "payment_failure"
    EventKey_SIGNUPSUCCESS                                  EventKey = "signup_success"
    EventKey_SIGNUPFAILURE                                  EventKey = "signup_failure"
    EventKey_DELAYEDSIGNUPCREATIONSUCCESS                   EventKey = "delayed_signup_creation_success"
    EventKey_DELAYEDSIGNUPCREATIONFAILURE                   EventKey = "delayed_signup_creation_failure"
    EventKey_BILLINGDATECHANGE                              EventKey = "billing_date_change"
    EventKey_EXPIRATIONDATECHANGE                           EventKey = "expiration_date_change"
    EventKey_RENEWALSUCCESS                                 EventKey = "renewal_success"
    EventKey_RENEWALFAILURE                                 EventKey = "renewal_failure"
    EventKey_SUBSCRIPTIONSTATECHANGE                        EventKey = "subscription_state_change"
    EventKey_SUBSCRIPTIONPRODUCTCHANGE                      EventKey = "subscription_product_change"
    EventKey_PENDINGCANCELLATIONCHANGE                      EventKey = "pending_cancellation_change"
    EventKey_EXPIRINGCARD                                   EventKey = "expiring_card"
    EventKey_CUSTOMERUPDATE                                 EventKey = "customer_update"
    EventKey_CUSTOMERCREATE                                 EventKey = "customer_create"
    EventKey_CUSTOMERDELETE                                 EventKey = "customer_delete"
    EventKey_COMPONENTALLOCATIONCHANGE                      EventKey = "component_allocation_change"
    EventKey_METEREDUSAGE                                   EventKey = "metered_usage"
    EventKey_PREPAIDUSAGE                                   EventKey = "prepaid_usage"
    EventKey_UPGRADEDOWNGRADESUCCESS                        EventKey = "upgrade_downgrade_success"
    EventKey_UPGRADEDOWNGRADEFAILURE                        EventKey = "upgrade_downgrade_failure"
    EventKey_STATEMENTCLOSED                                EventKey = "statement_closed"
    EventKey_STATEMENTSETTLED                               EventKey = "statement_settled"
    EventKey_SUBSCRIPTIONCARDUPDATE                         EventKey = "subscription_card_update"
    EventKey_SUBSCRIPTIONGROUPCARDUPDATE                    EventKey = "subscription_group_card_update"
    EventKey_SUBSCRIPTIONBANKACCOUNTUPDATE                  EventKey = "subscription_bank_account_update"
    EventKey_REFUNDSUCCESS                                  EventKey = "refund_success"
    EventKey_REFUNDFAILURE                                  EventKey = "refund_failure"
    EventKey_UPCOMINGRENEWALNOTICE                          EventKey = "upcoming_renewal_notice"
    EventKey_TRIALENDNOTICE                                 EventKey = "trial_end_notice"
    EventKey_DUNNINGSTEPREACHED                             EventKey = "dunning_step_reached"
    EventKey_INVOICEISSUED                                  EventKey = "invoice_issued"
    EventKey_PREPAIDSUBSCRIPTIONBALANCECHANGED              EventKey = "prepaid_subscription_balance_changed"
    EventKey_SUBSCRIPTIONGROUPSIGNUPSUCCESS                 EventKey = "subscription_group_signup_success"
    EventKey_SUBSCRIPTIONGROUPSIGNUPFAILURE                 EventKey = "subscription_group_signup_failure"
    EventKey_DIRECTDEBITPAYMENTPAIDOUT                      EventKey = "direct_debit_payment_paid_out"
    EventKey_DIRECTDEBITPAYMENTREJECTED                     EventKey = "direct_debit_payment_rejected"
    EventKey_DIRECTDEBITPAYMENTPENDING                      EventKey = "direct_debit_payment_pending"
    EventKey_PENDINGPAYMENTCREATED                          EventKey = "pending_payment_created"
    EventKey_PENDINGPAYMENTFAILED                           EventKey = "pending_payment_failed"
    EventKey_PENDINGPAYMENTCOMPLETED                        EventKey = "pending_payment_completed"
    EventKey_PROFORMAINVOICEISSUED                          EventKey = "proforma_invoice_issued"
    EventKey_SUBSCRIPTIONPREPAYMENTACCOUNTBALANCECHANGED    EventKey = "subscription_prepayment_account_balance_changed"
    EventKey_SUBSCRIPTIONSERVICECREDITACCOUNTBALANCECHANGED EventKey = "subscription_service_credit_account_balance_changed"
    EventKey_CUSTOMFIELDVALUECHANGE                         EventKey = "custom_field_value_change"
    EventKey_ITEMPRICEPOINTCHANGED                          EventKey = "item_price_point_changed"
    EventKey_RENEWALSUCCESSRECREATED                        EventKey = "renewal_success_recreated"
    EventKey_RENEWALFAILURERECREATED                        EventKey = "renewal_failure_recreated"
    EventKey_PAYMENTSUCCESSRECREATED                        EventKey = "payment_success_recreated"
    EventKey_PAYMENTFAILURERECREATED                        EventKey = "payment_failure_recreated"
    EventKey_SUBSCRIPTIONDELETION                           EventKey = "subscription_deletion"
    EventKey_SUBSCRIPTIONGROUPBANKACCOUNTUPDATE             EventKey = "subscription_group_bank_account_update"
    EventKey_SUBSCRIPTIONPAYPALACCOUNTUPDATE                EventKey = "subscription_paypal_account_update"
    EventKey_SUBSCRIPTIONGROUPPAYPALACCOUNTUPDATE           EventKey = "subscription_group_paypal_account_update"
    EventKey_SUBSCRIPTIONCUSTOMERCHANGE                     EventKey = "subscription_customer_change"
    EventKey_ACCOUNTTRANSACTIONCHANGED                      EventKey = "account_transaction_changed"
    EventKey_GOCARDLESSPAYMENTPAIDOUT                       EventKey = "go_cardless_payment_paid_out"
    EventKey_GOCARDLESSPAYMENTREJECTED                      EventKey = "go_cardless_payment_rejected"
    EventKey_GOCARDLESSPAYMENTPENDING                       EventKey = "go_cardless_payment_pending"
    EventKey_STRIPEDIRECTDEBITPAYMENTPAIDOUT                EventKey = "stripe_direct_debit_payment_paid_out"
    EventKey_STRIPEDIRECTDEBITPAYMENTREJECTED               EventKey = "stripe_direct_debit_payment_rejected"
    EventKey_STRIPEDIRECTDEBITPAYMENTPENDING                EventKey = "stripe_direct_debit_payment_pending"
    EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTPAIDOUT         EventKey = "maxio_payments_direct_debit_payment_paid_out"
    EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTREJECTED        EventKey = "maxio_payments_direct_debit_payment_rejected"
    EventKey_MAXIOPAYMENTSDIRECTDEBITPAYMENTPENDING         EventKey = "maxio_payments_direct_debit_payment_pending"
    EventKey_INVOICEINCOLLECTIONSCANCELED                   EventKey = "invoice_in_collections_canceled"
    EventKey_SUBSCRIPTIONADDEDTOGROUP                       EventKey = "subscription_added_to_group"
    EventKey_SUBSCRIPTIONREMOVEDFROMGROUP                   EventKey = "subscription_removed_from_group"
    EventKey_CHARGEBACKOPENED                               EventKey = "chargeback_opened"
    EventKey_CHARGEBACKLOST                                 EventKey = "chargeback_lost"
    EventKey_CHARGEBACKACCEPTED                             EventKey = "chargeback_accepted"
    EventKey_CHARGEBACKCLOSED                               EventKey = "chargeback_closed"
    EventKey_CHARGEBACKWON                                  EventKey = "chargeback_won"
    EventKey_PAYMENTCOLLECTIONMETHODCHANGED                 EventKey = "payment_collection_method_changed"
    EventKey_COMPONENTBILLINGDATECHANGED                    EventKey = "component_billing_date_changed"
    EventKey_SUBSCRIPTIONTERMRENEWALSCHEDULED               EventKey = "subscription_term_renewal_scheduled"
    EventKey_SUBSCRIPTIONTERMRENEWALPENDING                 EventKey = "subscription_term_renewal_pending"
    EventKey_SUBSCRIPTIONTERMRENEWALACTIVATED               EventKey = "subscription_term_renewal_activated"
    EventKey_SUBSCRIPTIONTERMRENEWALREMOVED                 EventKey = "subscription_term_renewal_removed"
)

// ExpirationIntervalUnit is a string enum.
type ExpirationIntervalUnit string

// MarshalJSON implements the json.Marshaler interface for ExpirationIntervalUnit.
// It customizes the JSON marshaling process for ExpirationIntervalUnit objects.
func (e ExpirationIntervalUnit) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ExpirationIntervalUnit")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExpirationIntervalUnit.
// It customizes the JSON unmarshaling process for ExpirationIntervalUnit objects.
func (e *ExpirationIntervalUnit) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ExpirationIntervalUnit(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ExpirationIntervalUnit")
    }
    return nil
}

// Checks whether the value is actually a member of ExpirationIntervalUnit.
func (e ExpirationIntervalUnit) isValid() bool {
    switch e {
    case ExpirationIntervalUnit_DAY,
        ExpirationIntervalUnit_MONTH,
        ExpirationIntervalUnit_NEVER:
        return true
    }
    return false
}

const (
    ExpirationIntervalUnit_DAY   ExpirationIntervalUnit = "day"
    ExpirationIntervalUnit_MONTH ExpirationIntervalUnit = "month"
    ExpirationIntervalUnit_NEVER ExpirationIntervalUnit = "never"
)

// FailedPaymentAction is a string enum.
// Action taken when payment for an invoice fails:
// - `leave_open_invoice` - prepayments and credits applied to invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history. This is the default option.
// - `rollback_to_pending` - prepayments and credits not applied; invoice remains in "pending" status; no email sent to the customer; payment failure recorded in the invoice history.
// - `initiate_dunning` - prepayments and credits applied to the invoice; invoice status set to "open"; email sent to the customer for the issued invoice (if setting applies); payment failure recorded in the invoice history; subscription will  most likely go into "past_due" or "canceled" state (depending upon net terms and dunning settings).
type FailedPaymentAction string

// MarshalJSON implements the json.Marshaler interface for FailedPaymentAction.
// It customizes the JSON marshaling process for FailedPaymentAction objects.
func (e FailedPaymentAction) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for FailedPaymentAction")
}

// UnmarshalJSON implements the json.Unmarshaler interface for FailedPaymentAction.
// It customizes the JSON unmarshaling process for FailedPaymentAction objects.
func (e *FailedPaymentAction) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = FailedPaymentAction(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to FailedPaymentAction")
    }
    return nil
}

// Checks whether the value is actually a member of FailedPaymentAction.
func (e FailedPaymentAction) isValid() bool {
    switch e {
    case FailedPaymentAction_LEAVEOPENINVOICE,
        FailedPaymentAction_ROLLBACKTOPENDING,
        FailedPaymentAction_INITIATEDUNNING:
        return true
    }
    return false
}

const (
    FailedPaymentAction_LEAVEOPENINVOICE  FailedPaymentAction = "leave_open_invoice"
    FailedPaymentAction_ROLLBACKTOPENDING FailedPaymentAction = "rollback_to_pending"
    FailedPaymentAction_INITIATEDUNNING   FailedPaymentAction = "initiate_dunning"
)

// FirstChargeType is a string enum.
type FirstChargeType string

// MarshalJSON implements the json.Marshaler interface for FirstChargeType.
// It customizes the JSON marshaling process for FirstChargeType objects.
func (e FirstChargeType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for FirstChargeType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for FirstChargeType.
// It customizes the JSON unmarshaling process for FirstChargeType objects.
func (e *FirstChargeType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = FirstChargeType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to FirstChargeType")
    }
    return nil
}

// Checks whether the value is actually a member of FirstChargeType.
func (e FirstChargeType) isValid() bool {
    switch e {
    case FirstChargeType_PRORATED,
        FirstChargeType_IMMEDIATE,
        FirstChargeType_DELAYED:
        return true
    }
    return false
}

const (
    FirstChargeType_PRORATED  FirstChargeType = "prorated"
    FirstChargeType_IMMEDIATE FirstChargeType = "immediate"
    FirstChargeType_DELAYED   FirstChargeType = "delayed"
)

// GroupTargetType is a string enum.
// The type of object indicated by the id attribute.
type GroupTargetType string

// MarshalJSON implements the json.Marshaler interface for GroupTargetType.
// It customizes the JSON marshaling process for GroupTargetType objects.
func (e GroupTargetType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for GroupTargetType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupTargetType.
// It customizes the JSON unmarshaling process for GroupTargetType objects.
func (e *GroupTargetType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = GroupTargetType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to GroupTargetType")
    }
    return nil
}

// Checks whether the value is actually a member of GroupTargetType.
func (e GroupTargetType) isValid() bool {
    switch e {
    case GroupTargetType_CUSTOMER,
        GroupTargetType_SUBSCRIPTION,
        GroupTargetType_SELF,
        GroupTargetType_PARENT,
        GroupTargetType_ELDEST:
        return true
    }
    return false
}

const (
    GroupTargetType_CUSTOMER     GroupTargetType = "customer"
    GroupTargetType_SUBSCRIPTION GroupTargetType = "subscription"
    GroupTargetType_SELF         GroupTargetType = "self"
    GroupTargetType_PARENT       GroupTargetType = "parent"
    GroupTargetType_ELDEST       GroupTargetType = "eldest"
)

// GroupType is a string enum.
type GroupType string

// MarshalJSON implements the json.Marshaler interface for GroupType.
// It customizes the JSON marshaling process for GroupType objects.
func (e GroupType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for GroupType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for GroupType.
// It customizes the JSON unmarshaling process for GroupType objects.
func (e *GroupType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = GroupType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to GroupType")
    }
    return nil
}

// Checks whether the value is actually a member of GroupType.
func (e GroupType) isValid() bool {
    switch e {
    case GroupType_SINGLECUSTOMER,
        GroupType_MULTIPLECUSTOMERS:
        return true
    }
    return false
}

const (
    GroupType_SINGLECUSTOMER    GroupType = "single_customer"
    GroupType_MULTIPLECUSTOMERS GroupType = "multiple_customers"
)

// IncludeNotNull is a string enum.
// Passed as a parameter to list methods to return only non null values.
type IncludeNotNull string

// MarshalJSON implements the json.Marshaler interface for IncludeNotNull.
// It customizes the JSON marshaling process for IncludeNotNull objects.
func (e IncludeNotNull) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for IncludeNotNull")
}

// UnmarshalJSON implements the json.Unmarshaler interface for IncludeNotNull.
// It customizes the JSON unmarshaling process for IncludeNotNull objects.
func (e *IncludeNotNull) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = IncludeNotNull(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to IncludeNotNull")
    }
    return nil
}

// Checks whether the value is actually a member of IncludeNotNull.
func (e IncludeNotNull) isValid() bool {
    switch e {
    case IncludeNotNull_NOTNULL:
        return true
    }
    return false
}

const (
    IncludeNotNull_NOTNULL IncludeNotNull = "not_null"
)

// IncludeNullOrNotNull is a string enum.
// Allows to filter by `not_null` or `null`.
type IncludeNullOrNotNull string

// MarshalJSON implements the json.Marshaler interface for IncludeNullOrNotNull.
// It customizes the JSON marshaling process for IncludeNullOrNotNull objects.
func (e IncludeNullOrNotNull) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for IncludeNullOrNotNull")
}

// UnmarshalJSON implements the json.Unmarshaler interface for IncludeNullOrNotNull.
// It customizes the JSON unmarshaling process for IncludeNullOrNotNull objects.
func (e *IncludeNullOrNotNull) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = IncludeNullOrNotNull(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to IncludeNullOrNotNull")
    }
    return nil
}

// Checks whether the value is actually a member of IncludeNullOrNotNull.
func (e IncludeNullOrNotNull) isValid() bool {
    switch e {
    case IncludeNullOrNotNull_NOTNULL,
        IncludeNullOrNotNull_NULL:
        return true
    }
    return false
}

const (
    IncludeNullOrNotNull_NOTNULL IncludeNullOrNotNull = "not_null"
    IncludeNullOrNotNull_NULL    IncludeNullOrNotNull = "null"
)

// IncludeOption is a string enum.
type IncludeOption string

// MarshalJSON implements the json.Marshaler interface for IncludeOption.
// It customizes the JSON marshaling process for IncludeOption objects.
func (e IncludeOption) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for IncludeOption")
}

// UnmarshalJSON implements the json.Unmarshaler interface for IncludeOption.
// It customizes the JSON unmarshaling process for IncludeOption objects.
func (e *IncludeOption) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = IncludeOption(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to IncludeOption")
    }
    return nil
}

// Checks whether the value is actually a member of IncludeOption.
func (e IncludeOption) isValid() bool {
    switch e {
    case IncludeOption_EXCLUDE,
        IncludeOption_INCLUDE:
        return true
    }
    return false
}

const (
    IncludeOption_EXCLUDE IncludeOption = "0"
    IncludeOption_INCLUDE IncludeOption = "1"
)

// IntervalUnit is a string enum.
type IntervalUnit string

// MarshalJSON implements the json.Marshaler interface for IntervalUnit.
// It customizes the JSON marshaling process for IntervalUnit objects.
func (e IntervalUnit) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for IntervalUnit")
}

// UnmarshalJSON implements the json.Unmarshaler interface for IntervalUnit.
// It customizes the JSON unmarshaling process for IntervalUnit objects.
func (e *IntervalUnit) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = IntervalUnit(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to IntervalUnit")
    }
    return nil
}

// Checks whether the value is actually a member of IntervalUnit.
func (e IntervalUnit) isValid() bool {
    switch e {
    case IntervalUnit_DAY,
        IntervalUnit_MONTH:
        return true
    }
    return false
}

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
// See also the [invoice consolidation documentation](https://maxio.zendesk.com/hc/en-us/articles/24252269909389-Invoice-Consolidation).
type InvoiceConsolidationLevel string

// MarshalJSON implements the json.Marshaler interface for InvoiceConsolidationLevel.
// It customizes the JSON marshaling process for InvoiceConsolidationLevel objects.
func (e InvoiceConsolidationLevel) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceConsolidationLevel")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceConsolidationLevel.
// It customizes the JSON unmarshaling process for InvoiceConsolidationLevel objects.
func (e *InvoiceConsolidationLevel) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceConsolidationLevel(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceConsolidationLevel")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceConsolidationLevel.
func (e InvoiceConsolidationLevel) isValid() bool {
    switch e {
    case InvoiceConsolidationLevel_NONE,
        InvoiceConsolidationLevel_CHILD,
        InvoiceConsolidationLevel_PARENT:
        return true
    }
    return false
}

const (
    InvoiceConsolidationLevel_NONE   InvoiceConsolidationLevel = "none"
    InvoiceConsolidationLevel_CHILD  InvoiceConsolidationLevel = "child"
    InvoiceConsolidationLevel_PARENT InvoiceConsolidationLevel = "parent"
)

// InvoiceDateField is a string enum.
type InvoiceDateField string

// MarshalJSON implements the json.Marshaler interface for InvoiceDateField.
// It customizes the JSON marshaling process for InvoiceDateField objects.
func (e InvoiceDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDateField.
// It customizes the JSON unmarshaling process for InvoiceDateField objects.
func (e *InvoiceDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceDateField")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceDateField.
func (e InvoiceDateField) isValid() bool {
    switch e {
    case InvoiceDateField_CREATEDAT,
        InvoiceDateField_DUEDATE,
        InvoiceDateField_ISSUEDATE,
        InvoiceDateField_UPDATEDAT,
        InvoiceDateField_PAIDDATE:
        return true
    }
    return false
}

const (
    InvoiceDateField_CREATEDAT InvoiceDateField = "created_at"
    InvoiceDateField_DUEDATE   InvoiceDateField = "due_date"
    InvoiceDateField_ISSUEDATE InvoiceDateField = "issue_date"
    InvoiceDateField_UPDATEDAT InvoiceDateField = "updated_at"
    InvoiceDateField_PAIDDATE  InvoiceDateField = "paid_date"
)

// InvoiceDiscountSourceType is a string enum.
type InvoiceDiscountSourceType string

// MarshalJSON implements the json.Marshaler interface for InvoiceDiscountSourceType.
// It customizes the JSON marshaling process for InvoiceDiscountSourceType objects.
func (e InvoiceDiscountSourceType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceDiscountSourceType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDiscountSourceType.
// It customizes the JSON unmarshaling process for InvoiceDiscountSourceType objects.
func (e *InvoiceDiscountSourceType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceDiscountSourceType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceDiscountSourceType")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceDiscountSourceType.
func (e InvoiceDiscountSourceType) isValid() bool {
    switch e {
    case InvoiceDiscountSourceType_COUPON,
        InvoiceDiscountSourceType_REFERRAL,
        InvoiceDiscountSourceType_ENUMADHOCCOUPON:
        return true
    }
    return false
}

const (
    InvoiceDiscountSourceType_COUPON          InvoiceDiscountSourceType = "Coupon"
    InvoiceDiscountSourceType_REFERRAL        InvoiceDiscountSourceType = "Referral"
    InvoiceDiscountSourceType_ENUMADHOCCOUPON InvoiceDiscountSourceType = "Ad Hoc Coupon"
)

// InvoiceDiscountType is a string enum.
type InvoiceDiscountType string

// MarshalJSON implements the json.Marshaler interface for InvoiceDiscountType.
// It customizes the JSON marshaling process for InvoiceDiscountType objects.
func (e InvoiceDiscountType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceDiscountType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceDiscountType.
// It customizes the JSON unmarshaling process for InvoiceDiscountType objects.
func (e *InvoiceDiscountType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceDiscountType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceDiscountType")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceDiscountType.
func (e InvoiceDiscountType) isValid() bool {
    switch e {
    case InvoiceDiscountType_PERCENTAGE,
        InvoiceDiscountType_FLATAMOUNT,
        InvoiceDiscountType_ROLLOVER:
        return true
    }
    return false
}

const (
    InvoiceDiscountType_PERCENTAGE InvoiceDiscountType = "percentage"
    InvoiceDiscountType_FLATAMOUNT InvoiceDiscountType = "flat_amount"
    InvoiceDiscountType_ROLLOVER   InvoiceDiscountType = "rollover"
)

// InvoiceEventPaymentMethod is a string enum.
type InvoiceEventPaymentMethod string

// MarshalJSON implements the json.Marshaler interface for InvoiceEventPaymentMethod.
// It customizes the JSON marshaling process for InvoiceEventPaymentMethod objects.
func (e InvoiceEventPaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceEventPaymentMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventPaymentMethod.
// It customizes the JSON unmarshaling process for InvoiceEventPaymentMethod objects.
func (e *InvoiceEventPaymentMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceEventPaymentMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceEventPaymentMethod")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceEventPaymentMethod.
func (e InvoiceEventPaymentMethod) isValid() bool {
    switch e {
    case InvoiceEventPaymentMethod_APPLEPAY,
        InvoiceEventPaymentMethod_BANKACCOUNT,
        InvoiceEventPaymentMethod_CREDITCARD,
        InvoiceEventPaymentMethod_EXTERNAL,
        InvoiceEventPaymentMethod_PAYPALACCOUNT:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for InvoiceEventType.
// It customizes the JSON marshaling process for InvoiceEventType objects.
func (e InvoiceEventType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceEventType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceEventType.
// It customizes the JSON unmarshaling process for InvoiceEventType objects.
func (e *InvoiceEventType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceEventType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceEventType")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceEventType.
func (e InvoiceEventType) isValid() bool {
    switch e {
    case InvoiceEventType_ISSUEINVOICE,
        InvoiceEventType_APPLYCREDITNOTE,
        InvoiceEventType_CREATECREDITNOTE,
        InvoiceEventType_APPLYPAYMENT,
        InvoiceEventType_APPLYDEBITNOTE,
        InvoiceEventType_CREATEDEBITNOTE,
        InvoiceEventType_REFUNDINVOICE,
        InvoiceEventType_VOIDINVOICE,
        InvoiceEventType_VOIDREMAINDER,
        InvoiceEventType_BACKPORTINVOICE,
        InvoiceEventType_CHANGEINVOICESTATUS,
        InvoiceEventType_CHANGEINVOICECOLLECTIONMETHOD,
        InvoiceEventType_REMOVEPAYMENT,
        InvoiceEventType_FAILEDPAYMENT,
        InvoiceEventType_CHANGECHARGEBACKSTATUS:
        return true
    }
    return false
}

const (
    InvoiceEventType_ISSUEINVOICE                  InvoiceEventType = "issue_invoice"
    InvoiceEventType_APPLYCREDITNOTE               InvoiceEventType = "apply_credit_note"
    InvoiceEventType_CREATECREDITNOTE              InvoiceEventType = "create_credit_note"
    InvoiceEventType_APPLYPAYMENT                  InvoiceEventType = "apply_payment"
    InvoiceEventType_APPLYDEBITNOTE                InvoiceEventType = "apply_debit_note"
    InvoiceEventType_CREATEDEBITNOTE               InvoiceEventType = "create_debit_note"
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

// MarshalJSON implements the json.Marshaler interface for InvoicePaymentMethodType.
// It customizes the JSON marshaling process for InvoicePaymentMethodType objects.
func (e InvoicePaymentMethodType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoicePaymentMethodType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePaymentMethodType.
// It customizes the JSON unmarshaling process for InvoicePaymentMethodType objects.
func (e *InvoicePaymentMethodType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoicePaymentMethodType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoicePaymentMethodType")
    }
    return nil
}

// Checks whether the value is actually a member of InvoicePaymentMethodType.
func (e InvoicePaymentMethodType) isValid() bool {
    switch e {
    case InvoicePaymentMethodType_CREDITCARD,
        InvoicePaymentMethodType_CHECK,
        InvoicePaymentMethodType_CASH,
        InvoicePaymentMethodType_MONEYORDER,
        InvoicePaymentMethodType_ACH,
        InvoicePaymentMethodType_OTHER:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for InvoicePaymentType.
// It customizes the JSON marshaling process for InvoicePaymentType objects.
func (e InvoicePaymentType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoicePaymentType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoicePaymentType.
// It customizes the JSON unmarshaling process for InvoicePaymentType objects.
func (e *InvoicePaymentType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoicePaymentType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoicePaymentType")
    }
    return nil
}

// Checks whether the value is actually a member of InvoicePaymentType.
func (e InvoicePaymentType) isValid() bool {
    switch e {
    case InvoicePaymentType_EXTERNAL,
        InvoicePaymentType_PREPAYMENT,
        InvoicePaymentType_SERVICECREDIT,
        InvoicePaymentType_PAYMENT:
        return true
    }
    return false
}

const (
    InvoicePaymentType_EXTERNAL      InvoicePaymentType = "external"
    InvoicePaymentType_PREPAYMENT    InvoicePaymentType = "prepayment"
    InvoicePaymentType_SERVICECREDIT InvoicePaymentType = "service_credit"
    InvoicePaymentType_PAYMENT       InvoicePaymentType = "payment"
)

// InvoiceRole is a string enum.
type InvoiceRole string

// MarshalJSON implements the json.Marshaler interface for InvoiceRole.
// It customizes the JSON marshaling process for InvoiceRole objects.
func (e InvoiceRole) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceRole")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceRole.
// It customizes the JSON unmarshaling process for InvoiceRole objects.
func (e *InvoiceRole) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceRole(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceRole")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceRole.
func (e InvoiceRole) isValid() bool {
    switch e {
    case InvoiceRole_UNSET,
        InvoiceRole_SIGNUP,
        InvoiceRole_RENEWAL,
        InvoiceRole_USAGE,
        InvoiceRole_REACTIVATION,
        InvoiceRole_PRORATION,
        InvoiceRole_MIGRATION,
        InvoiceRole_ADHOC,
        InvoiceRole_BACKPORT,
        InvoiceRole_BACKPORTBALANCERECONCILIATION:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for InvoiceSortField.
// It customizes the JSON marshaling process for InvoiceSortField objects.
func (e InvoiceSortField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceSortField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceSortField.
// It customizes the JSON unmarshaling process for InvoiceSortField objects.
func (e *InvoiceSortField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceSortField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceSortField")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceSortField.
func (e InvoiceSortField) isValid() bool {
    switch e {
    case InvoiceSortField_STATUS,
        InvoiceSortField_TOTALAMOUNT,
        InvoiceSortField_DUEAMOUNT,
        InvoiceSortField_CREATEDAT,
        InvoiceSortField_UPDATEDAT,
        InvoiceSortField_ISSUEDATE,
        InvoiceSortField_DUEDATE,
        InvoiceSortField_NUMBER:
        return true
    }
    return false
}

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
// The current status of the invoice. See [Invoice Statuses](https://maxio.zendesk.com/hc/en-us/articles/24252287829645-Advanced-Billing-Invoices-Overview#invoice-statuses) for more.
type InvoiceStatus string

// MarshalJSON implements the json.Marshaler interface for InvoiceStatus.
// It customizes the JSON marshaling process for InvoiceStatus objects.
func (e InvoiceStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for InvoiceStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for InvoiceStatus.
// It customizes the JSON unmarshaling process for InvoiceStatus objects.
func (e *InvoiceStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = InvoiceStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to InvoiceStatus")
    }
    return nil
}

// Checks whether the value is actually a member of InvoiceStatus.
func (e InvoiceStatus) isValid() bool {
    switch e {
    case InvoiceStatus_DRAFT,
        InvoiceStatus_OPEN,
        InvoiceStatus_PAID,
        InvoiceStatus_PENDING,
        InvoiceStatus_VOIDED,
        InvoiceStatus_CANCELED,
        InvoiceStatus_PROCESSING:
        return true
    }
    return false
}

const (
    InvoiceStatus_DRAFT      InvoiceStatus = "draft"
    InvoiceStatus_OPEN       InvoiceStatus = "open"
    InvoiceStatus_PAID       InvoiceStatus = "paid"
    InvoiceStatus_PENDING    InvoiceStatus = "pending"
    InvoiceStatus_VOIDED     InvoiceStatus = "voided"
    InvoiceStatus_CANCELED   InvoiceStatus = "canceled"
    InvoiceStatus_PROCESSING InvoiceStatus = "processing"
)

// ItemCategory is a string enum.
// One of the following: Business Software, Consumer Software, Digital Services, Physical Goods, Other
type ItemCategory string

// MarshalJSON implements the json.Marshaler interface for ItemCategory.
// It customizes the JSON marshaling process for ItemCategory objects.
func (e ItemCategory) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ItemCategory")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ItemCategory.
// It customizes the JSON unmarshaling process for ItemCategory objects.
func (e *ItemCategory) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ItemCategory(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ItemCategory")
    }
    return nil
}

// Checks whether the value is actually a member of ItemCategory.
func (e ItemCategory) isValid() bool {
    switch e {
    case ItemCategory_ENUMBUSINESSSOFTWARE,
        ItemCategory_ENUMCONSUMERSOFTWARE,
        ItemCategory_ENUMDIGITALSERVICES,
        ItemCategory_ENUMPHYSICALGOODS,
        ItemCategory_OTHER:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for LineItemKind.
// It customizes the JSON marshaling process for LineItemKind objects.
func (e LineItemKind) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for LineItemKind")
}

// UnmarshalJSON implements the json.Unmarshaler interface for LineItemKind.
// It customizes the JSON unmarshaling process for LineItemKind objects.
func (e *LineItemKind) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = LineItemKind(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to LineItemKind")
    }
    return nil
}

// Checks whether the value is actually a member of LineItemKind.
func (e LineItemKind) isValid() bool {
    switch e {
    case LineItemKind_BASELINE,
        LineItemKind_INITIAL,
        LineItemKind_TRIAL,
        LineItemKind_QUANTITYBASEDCOMPONENT,
        LineItemKind_PREPAIDUSAGECOMPONENT,
        LineItemKind_ONOFFCOMPONENT,
        LineItemKind_METEREDCOMPONENT,
        LineItemKind_EVENTBASEDCOMPONENT,
        LineItemKind_COUPON,
        LineItemKind_TAX:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for LineItemTransactionType.
// It customizes the JSON marshaling process for LineItemTransactionType objects.
func (e LineItemTransactionType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for LineItemTransactionType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for LineItemTransactionType.
// It customizes the JSON unmarshaling process for LineItemTransactionType objects.
func (e *LineItemTransactionType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = LineItemTransactionType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to LineItemTransactionType")
    }
    return nil
}

// Checks whether the value is actually a member of LineItemTransactionType.
func (e LineItemTransactionType) isValid() bool {
    switch e {
    case LineItemTransactionType_CHARGE,
        LineItemTransactionType_CREDIT,
        LineItemTransactionType_ADJUSTMENT,
        LineItemTransactionType_PAYMENT,
        LineItemTransactionType_REFUND,
        LineItemTransactionType_INFOTRANSACTION,
        LineItemTransactionType_PAYMENTAUTHORIZATION:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for ListComponentsPricePointsInclude.
// It customizes the JSON marshaling process for ListComponentsPricePointsInclude objects.
func (e ListComponentsPricePointsInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListComponentsPricePointsInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListComponentsPricePointsInclude.
// It customizes the JSON unmarshaling process for ListComponentsPricePointsInclude objects.
func (e *ListComponentsPricePointsInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListComponentsPricePointsInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListComponentsPricePointsInclude")
    }
    return nil
}

// Checks whether the value is actually a member of ListComponentsPricePointsInclude.
func (e ListComponentsPricePointsInclude) isValid() bool {
    switch e {
    case ListComponentsPricePointsInclude_CURRENCYPRICES:
        return true
    }
    return false
}

const (
    ListComponentsPricePointsInclude_CURRENCYPRICES ListComponentsPricePointsInclude = "currency_prices"
)

// ListEventsDateField is a string enum.
type ListEventsDateField string

// MarshalJSON implements the json.Marshaler interface for ListEventsDateField.
// It customizes the JSON marshaling process for ListEventsDateField objects.
func (e ListEventsDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListEventsDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListEventsDateField.
// It customizes the JSON unmarshaling process for ListEventsDateField objects.
func (e *ListEventsDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListEventsDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListEventsDateField")
    }
    return nil
}

// Checks whether the value is actually a member of ListEventsDateField.
func (e ListEventsDateField) isValid() bool {
    switch e {
    case ListEventsDateField_CREATEDAT:
        return true
    }
    return false
}

const (
    ListEventsDateField_CREATEDAT ListEventsDateField = "created_at"
)

// ListPrepaymentDateField is a string enum.
type ListPrepaymentDateField string

// MarshalJSON implements the json.Marshaler interface for ListPrepaymentDateField.
// It customizes the JSON marshaling process for ListPrepaymentDateField objects.
func (e ListPrepaymentDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListPrepaymentDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListPrepaymentDateField.
// It customizes the JSON unmarshaling process for ListPrepaymentDateField objects.
func (e *ListPrepaymentDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListPrepaymentDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListPrepaymentDateField")
    }
    return nil
}

// Checks whether the value is actually a member of ListPrepaymentDateField.
func (e ListPrepaymentDateField) isValid() bool {
    switch e {
    case ListPrepaymentDateField_CREATEDAT,
        ListPrepaymentDateField_APPLICATIONAT:
        return true
    }
    return false
}

const (
    ListPrepaymentDateField_CREATEDAT     ListPrepaymentDateField = "created_at"
    ListPrepaymentDateField_APPLICATIONAT ListPrepaymentDateField = "application_at"
)

// ListProductsInclude is a string enum.
type ListProductsInclude string

// MarshalJSON implements the json.Marshaler interface for ListProductsInclude.
// It customizes the JSON marshaling process for ListProductsInclude objects.
func (e ListProductsInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListProductsInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductsInclude.
// It customizes the JSON unmarshaling process for ListProductsInclude objects.
func (e *ListProductsInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListProductsInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListProductsInclude")
    }
    return nil
}

// Checks whether the value is actually a member of ListProductsInclude.
func (e ListProductsInclude) isValid() bool {
    switch e {
    case ListProductsInclude_PREPAIDPRODUCTPRICEPOINT:
        return true
    }
    return false
}

const (
    ListProductsInclude_PREPAIDPRODUCTPRICEPOINT ListProductsInclude = "prepaid_product_price_point"
)

// ListProductsPricePointsInclude is a string enum.
type ListProductsPricePointsInclude string

// MarshalJSON implements the json.Marshaler interface for ListProductsPricePointsInclude.
// It customizes the JSON marshaling process for ListProductsPricePointsInclude objects.
func (e ListProductsPricePointsInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListProductsPricePointsInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListProductsPricePointsInclude.
// It customizes the JSON unmarshaling process for ListProductsPricePointsInclude objects.
func (e *ListProductsPricePointsInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListProductsPricePointsInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListProductsPricePointsInclude")
    }
    return nil
}

// Checks whether the value is actually a member of ListProductsPricePointsInclude.
func (e ListProductsPricePointsInclude) isValid() bool {
    switch e {
    case ListProductsPricePointsInclude_CURRENCYPRICES:
        return true
    }
    return false
}

const (
    ListProductsPricePointsInclude_CURRENCYPRICES ListProductsPricePointsInclude = "currency_prices"
)

// ListSubscriptionComponentsInclude is a string enum.
type ListSubscriptionComponentsInclude string

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsInclude.
// It customizes the JSON marshaling process for ListSubscriptionComponentsInclude objects.
func (e ListSubscriptionComponentsInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListSubscriptionComponentsInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsInclude.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsInclude objects.
func (e *ListSubscriptionComponentsInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListSubscriptionComponentsInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListSubscriptionComponentsInclude")
    }
    return nil
}

// Checks whether the value is actually a member of ListSubscriptionComponentsInclude.
func (e ListSubscriptionComponentsInclude) isValid() bool {
    switch e {
    case ListSubscriptionComponentsInclude_SUBSCRIPTION,
        ListSubscriptionComponentsInclude_HISTORICUSAGES:
        return true
    }
    return false
}

const (
    ListSubscriptionComponentsInclude_SUBSCRIPTION   ListSubscriptionComponentsInclude = "subscription"
    ListSubscriptionComponentsInclude_HISTORICUSAGES ListSubscriptionComponentsInclude = "historic_usages"
)

// ListSubscriptionComponentsSort is a string enum.
type ListSubscriptionComponentsSort string

// MarshalJSON implements the json.Marshaler interface for ListSubscriptionComponentsSort.
// It customizes the JSON marshaling process for ListSubscriptionComponentsSort objects.
func (e ListSubscriptionComponentsSort) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ListSubscriptionComponentsSort")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ListSubscriptionComponentsSort.
// It customizes the JSON unmarshaling process for ListSubscriptionComponentsSort objects.
func (e *ListSubscriptionComponentsSort) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ListSubscriptionComponentsSort(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ListSubscriptionComponentsSort")
    }
    return nil
}

// Checks whether the value is actually a member of ListSubscriptionComponentsSort.
func (e ListSubscriptionComponentsSort) isValid() bool {
    switch e {
    case ListSubscriptionComponentsSort_ID,
        ListSubscriptionComponentsSort_UPDATEDAT:
        return true
    }
    return false
}

const (
    ListSubscriptionComponentsSort_ID        ListSubscriptionComponentsSort = "id"
    ListSubscriptionComponentsSort_UPDATEDAT ListSubscriptionComponentsSort = "updated_at"
)

// MetafieldInput is a string enum.
// Indicates the type of metafield. A text metafield allows any string value. Dropdown and radio metafields have a set of values that can be selected.  Defaults to 'text'.
type MetafieldInput string

// MarshalJSON implements the json.Marshaler interface for MetafieldInput.
// It customizes the JSON marshaling process for MetafieldInput objects.
func (e MetafieldInput) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for MetafieldInput")
}

// UnmarshalJSON implements the json.Unmarshaler interface for MetafieldInput.
// It customizes the JSON unmarshaling process for MetafieldInput objects.
func (e *MetafieldInput) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = MetafieldInput(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to MetafieldInput")
    }
    return nil
}

// Checks whether the value is actually a member of MetafieldInput.
func (e MetafieldInput) isValid() bool {
    switch e {
    case MetafieldInput_BALANCETRACKER,
        MetafieldInput_TEXT,
        MetafieldInput_RADIO,
        MetafieldInput_DROPDOWN:
        return true
    }
    return false
}

const (
    MetafieldInput_BALANCETRACKER MetafieldInput = "balance_tracker"
    MetafieldInput_TEXT           MetafieldInput = "text"
    MetafieldInput_RADIO          MetafieldInput = "radio"
    MetafieldInput_DROPDOWN       MetafieldInput = "dropdown"
)

// PaymentType is a string enum.
type PaymentType string

// MarshalJSON implements the json.Marshaler interface for PaymentType.
// It customizes the JSON marshaling process for PaymentType objects.
func (e PaymentType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for PaymentType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for PaymentType.
// It customizes the JSON unmarshaling process for PaymentType objects.
func (e *PaymentType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = PaymentType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to PaymentType")
    }
    return nil
}

// Checks whether the value is actually a member of PaymentType.
func (e PaymentType) isValid() bool {
    switch e {
    case PaymentType_CREDITCARD,
        PaymentType_BANKACCOUNT,
        PaymentType_PAYPALACCOUNT,
        PaymentType_APPLEPAY:
        return true
    }
    return false
}

const (
    PaymentType_CREDITCARD    PaymentType = "credit_card"
    PaymentType_BANKACCOUNT   PaymentType = "bank_account"
    PaymentType_PAYPALACCOUNT PaymentType = "paypal_account"
    PaymentType_APPLEPAY      PaymentType = "apple_pay"
)

// PayPalVault is a string enum.
// The vault that stores the payment profile with the provided vault_token.
type PayPalVault string

// MarshalJSON implements the json.Marshaler interface for PayPalVault.
// It customizes the JSON marshaling process for PayPalVault objects.
func (e PayPalVault) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for PayPalVault")
}

// UnmarshalJSON implements the json.Unmarshaler interface for PayPalVault.
// It customizes the JSON unmarshaling process for PayPalVault objects.
func (e *PayPalVault) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = PayPalVault(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to PayPalVault")
    }
    return nil
}

// Checks whether the value is actually a member of PayPalVault.
func (e PayPalVault) isValid() bool {
    switch e {
    case PayPalVault_BRAINTREEBLUE,
        PayPalVault_PAYPAL,
        PayPalVault_MODUSLINK,
        PayPalVault_PAYPALCOMPLETE:
        return true
    }
    return false
}

const (
    PayPalVault_BRAINTREEBLUE  PayPalVault = "braintree_blue"
    PayPalVault_PAYPAL         PayPalVault = "paypal"
    PayPalVault_MODUSLINK      PayPalVault = "moduslink"
    PayPalVault_PAYPALCOMPLETE PayPalVault = "paypal_complete"
)

// PrepaymentMethod is a string enum.
type PrepaymentMethod string

// MarshalJSON implements the json.Marshaler interface for PrepaymentMethod.
// It customizes the JSON marshaling process for PrepaymentMethod objects.
func (e PrepaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for PrepaymentMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for PrepaymentMethod.
// It customizes the JSON unmarshaling process for PrepaymentMethod objects.
func (e *PrepaymentMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = PrepaymentMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to PrepaymentMethod")
    }
    return nil
}

// Checks whether the value is actually a member of PrepaymentMethod.
func (e PrepaymentMethod) isValid() bool {
    switch e {
    case PrepaymentMethod_CHECK,
        PrepaymentMethod_CASH,
        PrepaymentMethod_MONEYORDER,
        PrepaymentMethod_ACH,
        PrepaymentMethod_PAYPALACCOUNT,
        PrepaymentMethod_CREDITCARD,
        PrepaymentMethod_OTHER:
        return true
    }
    return false
}

const (
    PrepaymentMethod_CHECK         PrepaymentMethod = "check"
    PrepaymentMethod_CASH          PrepaymentMethod = "cash"
    PrepaymentMethod_MONEYORDER    PrepaymentMethod = "money_order"
    PrepaymentMethod_ACH           PrepaymentMethod = "ach"
    PrepaymentMethod_PAYPALACCOUNT PrepaymentMethod = "paypal_account"
    PrepaymentMethod_CREDITCARD    PrepaymentMethod = "credit_card"
    PrepaymentMethod_OTHER         PrepaymentMethod = "other"
)

// PricePointType is a string enum.
// Price point type. We expose the following types:
// 1. **default**: a price point that is marked as a default price for a certain product.
// 2. **custom**: a custom price point.
// 3. **catalog**: a price point that is **not** marked as a default price for a certain product and is **not** a custom one.
type PricePointType string

// MarshalJSON implements the json.Marshaler interface for PricePointType.
// It customizes the JSON marshaling process for PricePointType objects.
func (e PricePointType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for PricePointType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for PricePointType.
// It customizes the JSON unmarshaling process for PricePointType objects.
func (e *PricePointType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = PricePointType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to PricePointType")
    }
    return nil
}

// Checks whether the value is actually a member of PricePointType.
func (e PricePointType) isValid() bool {
    switch e {
    case PricePointType_CATALOG,
        PricePointType_ENUMDEFAULT,
        PricePointType_CUSTOM:
        return true
    }
    return false
}

const (
    PricePointType_CATALOG     PricePointType = "catalog"
    PricePointType_ENUMDEFAULT PricePointType = "default"
    PricePointType_CUSTOM      PricePointType = "custom"
)

// PricingScheme is a string enum.
// The identifier for the pricing scheme. See [Product Components](https://help.chargify.com/products/product-components.html) for an overview of pricing schemes.
type PricingScheme string

// MarshalJSON implements the json.Marshaler interface for PricingScheme.
// It customizes the JSON marshaling process for PricingScheme objects.
func (e PricingScheme) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for PricingScheme")
}

// UnmarshalJSON implements the json.Unmarshaler interface for PricingScheme.
// It customizes the JSON unmarshaling process for PricingScheme objects.
func (e *PricingScheme) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = PricingScheme(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to PricingScheme")
    }
    return nil
}

// Checks whether the value is actually a member of PricingScheme.
func (e PricingScheme) isValid() bool {
    switch e {
    case PricingScheme_STAIRSTEP,
        PricingScheme_VOLUME,
        PricingScheme_PERUNIT,
        PricingScheme_TIERED:
        return true
    }
    return false
}

const (
    PricingScheme_STAIRSTEP PricingScheme = "stairstep"
    PricingScheme_VOLUME    PricingScheme = "volume"
    PricingScheme_PERUNIT   PricingScheme = "per_unit"
    PricingScheme_TIERED    PricingScheme = "tiered"
)

// ProformaInvoiceDiscountSourceType is a string enum.
type ProformaInvoiceDiscountSourceType string

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceDiscountSourceType.
// It customizes the JSON marshaling process for ProformaInvoiceDiscountSourceType objects.
func (e ProformaInvoiceDiscountSourceType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ProformaInvoiceDiscountSourceType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceDiscountSourceType.
// It customizes the JSON unmarshaling process for ProformaInvoiceDiscountSourceType objects.
func (e *ProformaInvoiceDiscountSourceType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ProformaInvoiceDiscountSourceType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ProformaInvoiceDiscountSourceType")
    }
    return nil
}

// Checks whether the value is actually a member of ProformaInvoiceDiscountSourceType.
func (e ProformaInvoiceDiscountSourceType) isValid() bool {
    switch e {
    case ProformaInvoiceDiscountSourceType_COUPON,
        ProformaInvoiceDiscountSourceType_REFERRAL:
        return true
    }
    return false
}

const (
    ProformaInvoiceDiscountSourceType_COUPON   ProformaInvoiceDiscountSourceType = "Coupon"
    ProformaInvoiceDiscountSourceType_REFERRAL ProformaInvoiceDiscountSourceType = "Referral"
)

// ProformaInvoiceRole is a string enum.
// 'proforma' value is deprecated in favor of proforma_adhoc and proforma_automatic
type ProformaInvoiceRole string

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceRole.
// It customizes the JSON marshaling process for ProformaInvoiceRole objects.
func (e ProformaInvoiceRole) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ProformaInvoiceRole")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceRole.
// It customizes the JSON unmarshaling process for ProformaInvoiceRole objects.
func (e *ProformaInvoiceRole) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ProformaInvoiceRole(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ProformaInvoiceRole")
    }
    return nil
}

// Checks whether the value is actually a member of ProformaInvoiceRole.
func (e ProformaInvoiceRole) isValid() bool {
    switch e {
    case ProformaInvoiceRole_UNSET,
        ProformaInvoiceRole_PROFORMA,
        ProformaInvoiceRole_PROFORMAADHOC,
        ProformaInvoiceRole_PROFORMAAUTOMATIC:
        return true
    }
    return false
}

const (
    ProformaInvoiceRole_UNSET             ProformaInvoiceRole = "unset"
    ProformaInvoiceRole_PROFORMA          ProformaInvoiceRole = "proforma"
    ProformaInvoiceRole_PROFORMAADHOC     ProformaInvoiceRole = "proforma_adhoc"
    ProformaInvoiceRole_PROFORMAAUTOMATIC ProformaInvoiceRole = "proforma_automatic"
)

// ProformaInvoiceStatus is a string enum.
type ProformaInvoiceStatus string

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceStatus.
// It customizes the JSON marshaling process for ProformaInvoiceStatus objects.
func (e ProformaInvoiceStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ProformaInvoiceStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceStatus.
// It customizes the JSON unmarshaling process for ProformaInvoiceStatus objects.
func (e *ProformaInvoiceStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ProformaInvoiceStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ProformaInvoiceStatus")
    }
    return nil
}

// Checks whether the value is actually a member of ProformaInvoiceStatus.
func (e ProformaInvoiceStatus) isValid() bool {
    switch e {
    case ProformaInvoiceStatus_DRAFT,
        ProformaInvoiceStatus_VOIDED,
        ProformaInvoiceStatus_ARCHIVED:
        return true
    }
    return false
}

const (
    ProformaInvoiceStatus_DRAFT    ProformaInvoiceStatus = "draft"
    ProformaInvoiceStatus_VOIDED   ProformaInvoiceStatus = "voided"
    ProformaInvoiceStatus_ARCHIVED ProformaInvoiceStatus = "archived"
)

// ProformaInvoiceTaxSourceType is a string enum.
type ProformaInvoiceTaxSourceType string

// MarshalJSON implements the json.Marshaler interface for ProformaInvoiceTaxSourceType.
// It customizes the JSON marshaling process for ProformaInvoiceTaxSourceType objects.
func (e ProformaInvoiceTaxSourceType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ProformaInvoiceTaxSourceType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ProformaInvoiceTaxSourceType.
// It customizes the JSON unmarshaling process for ProformaInvoiceTaxSourceType objects.
func (e *ProformaInvoiceTaxSourceType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ProformaInvoiceTaxSourceType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ProformaInvoiceTaxSourceType")
    }
    return nil
}

// Checks whether the value is actually a member of ProformaInvoiceTaxSourceType.
func (e ProformaInvoiceTaxSourceType) isValid() bool {
    switch e {
    case ProformaInvoiceTaxSourceType_TAX,
        ProformaInvoiceTaxSourceType_AVALARA:
        return true
    }
    return false
}

const (
    ProformaInvoiceTaxSourceType_TAX     ProformaInvoiceTaxSourceType = "Tax"
    ProformaInvoiceTaxSourceType_AVALARA ProformaInvoiceTaxSourceType = "Avalara"
)

// ReactivationCharge is a string enum.
// You may choose how to handle the reactivation charge for that subscription: 1) `prorated` A prorated charge for the product price will be attempted for to complete the period 2) `immediate` A full-price charge for the product price will be attempted immediately 3) `delayed` A full-price charge for the product price will be attempted at the next renewal
type ReactivationCharge string

// MarshalJSON implements the json.Marshaler interface for ReactivationCharge.
// It customizes the JSON marshaling process for ReactivationCharge objects.
func (e ReactivationCharge) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ReactivationCharge")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReactivationCharge.
// It customizes the JSON unmarshaling process for ReactivationCharge objects.
func (e *ReactivationCharge) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ReactivationCharge(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ReactivationCharge")
    }
    return nil
}

// Checks whether the value is actually a member of ReactivationCharge.
func (e ReactivationCharge) isValid() bool {
    switch e {
    case ReactivationCharge_PRORATED,
        ReactivationCharge_IMMEDIATE,
        ReactivationCharge_DELAYED:
        return true
    }
    return false
}

const (
    ReactivationCharge_PRORATED  ReactivationCharge = "prorated"
    ReactivationCharge_IMMEDIATE ReactivationCharge = "immediate"
    ReactivationCharge_DELAYED   ReactivationCharge = "delayed"
)

// RecurringScheme is a string enum.
type RecurringScheme string

// MarshalJSON implements the json.Marshaler interface for RecurringScheme.
// It customizes the JSON marshaling process for RecurringScheme objects.
func (e RecurringScheme) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for RecurringScheme")
}

// UnmarshalJSON implements the json.Unmarshaler interface for RecurringScheme.
// It customizes the JSON unmarshaling process for RecurringScheme objects.
func (e *RecurringScheme) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = RecurringScheme(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to RecurringScheme")
    }
    return nil
}

// Checks whether the value is actually a member of RecurringScheme.
func (e RecurringScheme) isValid() bool {
    switch e {
    case RecurringScheme_DONOTRECUR,
        RecurringScheme_RECURINDEFINITELY,
        RecurringScheme_RECURWITHDURATION:
        return true
    }
    return false
}

const (
    RecurringScheme_DONOTRECUR        RecurringScheme = "do_not_recur"
    RecurringScheme_RECURINDEFINITELY RecurringScheme = "recur_indefinitely"
    RecurringScheme_RECURWITHDURATION RecurringScheme = "recur_with_duration"
)

// ResourceType is a string enum.
type ResourceType string

// MarshalJSON implements the json.Marshaler interface for ResourceType.
// It customizes the JSON marshaling process for ResourceType objects.
func (e ResourceType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ResourceType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResourceType.
// It customizes the JSON unmarshaling process for ResourceType objects.
func (e *ResourceType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ResourceType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ResourceType")
    }
    return nil
}

// Checks whether the value is actually a member of ResourceType.
func (e ResourceType) isValid() bool {
    switch e {
    case ResourceType_SUBSCRIPTIONS,
        ResourceType_CUSTOMERS:
        return true
    }
    return false
}

const (
    ResourceType_SUBSCRIPTIONS ResourceType = "subscriptions"
    ResourceType_CUSTOMERS     ResourceType = "customers"
)

// RestrictionType is a string enum.
type RestrictionType string

// MarshalJSON implements the json.Marshaler interface for RestrictionType.
// It customizes the JSON marshaling process for RestrictionType objects.
func (e RestrictionType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for RestrictionType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for RestrictionType.
// It customizes the JSON unmarshaling process for RestrictionType objects.
func (e *RestrictionType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = RestrictionType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to RestrictionType")
    }
    return nil
}

// Checks whether the value is actually a member of RestrictionType.
func (e RestrictionType) isValid() bool {
    switch e {
    case RestrictionType_COMPONENT,
        RestrictionType_PRODUCT:
        return true
    }
    return false
}

const (
    RestrictionType_COMPONENT RestrictionType = "Component"
    RestrictionType_PRODUCT   RestrictionType = "Product"
)

// ResumptionCharge is a string enum.
// (For calendar billing subscriptions only) The way that the resumed subscription's charge should be handled
type ResumptionCharge string

// MarshalJSON implements the json.Marshaler interface for ResumptionCharge.
// It customizes the JSON marshaling process for ResumptionCharge objects.
func (e ResumptionCharge) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ResumptionCharge")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResumptionCharge.
// It customizes the JSON unmarshaling process for ResumptionCharge objects.
func (e *ResumptionCharge) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ResumptionCharge(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ResumptionCharge")
    }
    return nil
}

// Checks whether the value is actually a member of ResumptionCharge.
func (e ResumptionCharge) isValid() bool {
    switch e {
    case ResumptionCharge_PRORATED,
        ResumptionCharge_IMMEDIATE,
        ResumptionCharge_DELAYED:
        return true
    }
    return false
}

const (
    ResumptionCharge_PRORATED  ResumptionCharge = "prorated"
    ResumptionCharge_IMMEDIATE ResumptionCharge = "immediate"
    ResumptionCharge_DELAYED   ResumptionCharge = "delayed"
)

// ServiceCreditType is a string enum.
// The type of entry
type ServiceCreditType string

// MarshalJSON implements the json.Marshaler interface for ServiceCreditType.
// It customizes the JSON marshaling process for ServiceCreditType objects.
func (e ServiceCreditType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for ServiceCreditType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceCreditType.
// It customizes the JSON unmarshaling process for ServiceCreditType objects.
func (e *ServiceCreditType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = ServiceCreditType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to ServiceCreditType")
    }
    return nil
}

// Checks whether the value is actually a member of ServiceCreditType.
func (e ServiceCreditType) isValid() bool {
    switch e {
    case ServiceCreditType_CREDIT,
        ServiceCreditType_DEBIT:
        return true
    }
    return false
}

const (
    ServiceCreditType_CREDIT ServiceCreditType = "Credit"
    ServiceCreditType_DEBIT  ServiceCreditType = "Debit"
)

// SnapDay is a string enum.
type SnapDay string

// MarshalJSON implements the json.Marshaler interface for SnapDay.
// It customizes the JSON marshaling process for SnapDay objects.
func (e SnapDay) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SnapDay")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnapDay.
// It customizes the JSON unmarshaling process for SnapDay objects.
func (e *SnapDay) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SnapDay(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SnapDay")
    }
    return nil
}

// Checks whether the value is actually a member of SnapDay.
func (e SnapDay) isValid() bool {
    switch e {
    case SnapDay_END:
        return true
    }
    return false
}

const (
    SnapDay_END SnapDay = "end"
)

// SortingDirection is a string enum.
// Used for sorting results.
type SortingDirection string

// MarshalJSON implements the json.Marshaler interface for SortingDirection.
// It customizes the JSON marshaling process for SortingDirection objects.
func (e SortingDirection) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SortingDirection")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SortingDirection.
// It customizes the JSON unmarshaling process for SortingDirection objects.
func (e *SortingDirection) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SortingDirection(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SortingDirection")
    }
    return nil
}

// Checks whether the value is actually a member of SortingDirection.
func (e SortingDirection) isValid() bool {
    switch e {
    case SortingDirection_ASC,
        SortingDirection_DESC:
        return true
    }
    return false
}

const (
    SortingDirection_ASC  SortingDirection = "asc"
    SortingDirection_DESC SortingDirection = "desc"
)

// SubscriptionDateField is a string enum.
type SubscriptionDateField string

// MarshalJSON implements the json.Marshaler interface for SubscriptionDateField.
// It customizes the JSON marshaling process for SubscriptionDateField objects.
func (e SubscriptionDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionDateField.
// It customizes the JSON unmarshaling process for SubscriptionDateField objects.
func (e *SubscriptionDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionDateField")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionDateField.
func (e SubscriptionDateField) isValid() bool {
    switch e {
    case SubscriptionDateField_CURRENTPERIODENDSAT,
        SubscriptionDateField_CURRENTPERIODSTARTSAT,
        SubscriptionDateField_CREATEDAT,
        SubscriptionDateField_ACTIVATEDAT,
        SubscriptionDateField_CANCELEDAT,
        SubscriptionDateField_EXPIRESAT,
        SubscriptionDateField_TRIALSTARTEDAT,
        SubscriptionDateField_TRIALENDEDAT,
        SubscriptionDateField_UPDATEDAT:
        return true
    }
    return false
}

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

// SubscriptionGroupInclude is a string enum.
type SubscriptionGroupInclude string

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupInclude.
// It customizes the JSON marshaling process for SubscriptionGroupInclude objects.
func (e SubscriptionGroupInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionGroupInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupInclude.
// It customizes the JSON unmarshaling process for SubscriptionGroupInclude objects.
func (e *SubscriptionGroupInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionGroupInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionGroupInclude")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionGroupInclude.
func (e SubscriptionGroupInclude) isValid() bool {
    switch e {
    case SubscriptionGroupInclude_CURRENTBILLINGAMOUNTINCENTS:
        return true
    }
    return false
}

const (
    SubscriptionGroupInclude_CURRENTBILLINGAMOUNTINCENTS SubscriptionGroupInclude = "current_billing_amount_in_cents"
)

// SubscriptionGroupPrepaymentMethod is a string enum.
type SubscriptionGroupPrepaymentMethod string

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupPrepaymentMethod.
// It customizes the JSON marshaling process for SubscriptionGroupPrepaymentMethod objects.
func (e SubscriptionGroupPrepaymentMethod) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionGroupPrepaymentMethod")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupPrepaymentMethod.
// It customizes the JSON unmarshaling process for SubscriptionGroupPrepaymentMethod objects.
func (e *SubscriptionGroupPrepaymentMethod) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionGroupPrepaymentMethod(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionGroupPrepaymentMethod")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionGroupPrepaymentMethod.
func (e SubscriptionGroupPrepaymentMethod) isValid() bool {
    switch e {
    case SubscriptionGroupPrepaymentMethod_CHECK,
        SubscriptionGroupPrepaymentMethod_CASH,
        SubscriptionGroupPrepaymentMethod_MONEYORDER,
        SubscriptionGroupPrepaymentMethod_ACH,
        SubscriptionGroupPrepaymentMethod_PAYPALACCOUNT,
        SubscriptionGroupPrepaymentMethod_OTHER:
        return true
    }
    return false
}

const (
    SubscriptionGroupPrepaymentMethod_CHECK         SubscriptionGroupPrepaymentMethod = "check"
    SubscriptionGroupPrepaymentMethod_CASH          SubscriptionGroupPrepaymentMethod = "cash"
    SubscriptionGroupPrepaymentMethod_MONEYORDER    SubscriptionGroupPrepaymentMethod = "money_order"
    SubscriptionGroupPrepaymentMethod_ACH           SubscriptionGroupPrepaymentMethod = "ach"
    SubscriptionGroupPrepaymentMethod_PAYPALACCOUNT SubscriptionGroupPrepaymentMethod = "paypal_account"
    SubscriptionGroupPrepaymentMethod_OTHER         SubscriptionGroupPrepaymentMethod = "other"
)

// SubscriptionGroupsListInclude is a string enum.
type SubscriptionGroupsListInclude string

// MarshalJSON implements the json.Marshaler interface for SubscriptionGroupsListInclude.
// It customizes the JSON marshaling process for SubscriptionGroupsListInclude objects.
func (e SubscriptionGroupsListInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionGroupsListInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionGroupsListInclude.
// It customizes the JSON unmarshaling process for SubscriptionGroupsListInclude objects.
func (e *SubscriptionGroupsListInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionGroupsListInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionGroupsListInclude")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionGroupsListInclude.
func (e SubscriptionGroupsListInclude) isValid() bool {
    switch e {
    case SubscriptionGroupsListInclude_ACCOUNTBALANCES:
        return true
    }
    return false
}

const (
    SubscriptionGroupsListInclude_ACCOUNTBALANCES SubscriptionGroupsListInclude = "account_balances"
)

// SubscriptionInclude is a string enum.
type SubscriptionInclude string

// MarshalJSON implements the json.Marshaler interface for SubscriptionInclude.
// It customizes the JSON marshaling process for SubscriptionInclude objects.
func (e SubscriptionInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionInclude.
// It customizes the JSON unmarshaling process for SubscriptionInclude objects.
func (e *SubscriptionInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionInclude")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionInclude.
func (e SubscriptionInclude) isValid() bool {
    switch e {
    case SubscriptionInclude_COUPONS,
        SubscriptionInclude_SELFSERVICEPAGETOKEN:
        return true
    }
    return false
}

const (
    SubscriptionInclude_COUPONS              SubscriptionInclude = "coupons"
    SubscriptionInclude_SELFSERVICEPAGETOKEN SubscriptionInclude = "self_service_page_token"
)

// SubscriptionListDateField is a string enum.
type SubscriptionListDateField string

// MarshalJSON implements the json.Marshaler interface for SubscriptionListDateField.
// It customizes the JSON marshaling process for SubscriptionListDateField objects.
func (e SubscriptionListDateField) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionListDateField")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionListDateField.
// It customizes the JSON unmarshaling process for SubscriptionListDateField objects.
func (e *SubscriptionListDateField) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionListDateField(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionListDateField")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionListDateField.
func (e SubscriptionListDateField) isValid() bool {
    switch e {
    case SubscriptionListDateField_UPDATEDAT:
        return true
    }
    return false
}

const (
    SubscriptionListDateField_UPDATEDAT SubscriptionListDateField = "updated_at"
)

// SubscriptionListInclude is a string enum.
type SubscriptionListInclude string

// MarshalJSON implements the json.Marshaler interface for SubscriptionListInclude.
// It customizes the JSON marshaling process for SubscriptionListInclude objects.
func (e SubscriptionListInclude) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionListInclude")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionListInclude.
// It customizes the JSON unmarshaling process for SubscriptionListInclude objects.
func (e *SubscriptionListInclude) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionListInclude(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionListInclude")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionListInclude.
func (e SubscriptionListInclude) isValid() bool {
    switch e {
    case SubscriptionListInclude_SELFSERVICEPAGETOKEN:
        return true
    }
    return false
}

const (
    SubscriptionListInclude_SELFSERVICEPAGETOKEN SubscriptionListInclude = "self_service_page_token"
)

// SubscriptionPurgeType is a string enum.
type SubscriptionPurgeType string

// MarshalJSON implements the json.Marshaler interface for SubscriptionPurgeType.
// It customizes the JSON marshaling process for SubscriptionPurgeType objects.
func (e SubscriptionPurgeType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionPurgeType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionPurgeType.
// It customizes the JSON unmarshaling process for SubscriptionPurgeType objects.
func (e *SubscriptionPurgeType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionPurgeType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionPurgeType")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionPurgeType.
func (e SubscriptionPurgeType) isValid() bool {
    switch e {
    case SubscriptionPurgeType_CUSTOMER,
        SubscriptionPurgeType_PAYMENTPROFILE:
        return true
    }
    return false
}

const (
    SubscriptionPurgeType_CUSTOMER       SubscriptionPurgeType = "customer"
    SubscriptionPurgeType_PAYMENTPROFILE SubscriptionPurgeType = "payment_profile"
)

// SubscriptionSort is a string enum.
type SubscriptionSort string

// MarshalJSON implements the json.Marshaler interface for SubscriptionSort.
// It customizes the JSON marshaling process for SubscriptionSort objects.
func (e SubscriptionSort) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionSort")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionSort.
// It customizes the JSON unmarshaling process for SubscriptionSort objects.
func (e *SubscriptionSort) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionSort(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionSort")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionSort.
func (e SubscriptionSort) isValid() bool {
    switch e {
    case SubscriptionSort_SIGNUPDATE,
        SubscriptionSort_PERIODSTART,
        SubscriptionSort_PERIODEND,
        SubscriptionSort_NEXTASSESSMENT,
        SubscriptionSort_UPDATEDAT,
        SubscriptionSort_CREATEDAT:
        return true
    }
    return false
}

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
// * `active` - A normal, active subscription. It is not in a trial and is paid and up to date.
// * `assessing` - An internal (transient) state that indicates a subscription is in the middle of periodic assessment. Do not base any access decisions in your app on this state, as it may not always be exposed.
// * `pending` - An internal (transient) state that indicates a subscription is in the creation process. Do not base any access decisions in your app on this state, as it may not always be exposed.
// * `trialing` - A subscription in trialing state has a valid trial subscription. This type of subscription may transition to active once payment is received when the trial has ended. Otherwise, it may go to a Problem or End of Life state.
// * `paused` - An internal state that indicates that your account with Advanced Billing is in arrears.
// * **Problem States**
// * `past_due` - Indicates that the most recent payment has failed, and payment is past due for this subscription. If you have enabled our automated dunning, this subscription will be in the dunning process (additional status and callbacks from the dunning process will be available in the future). If you are handling dunning and payment updates yourself, you will want to use this state to initiate a payment update from your customers.
// * `soft_failure` - Indicates that normal assessment/processing of the subscription has failed for a reason that cannot be fixed by the Customer. For example, a Soft Fail may result from a timeout at the gateway or incorrect credentials on your part. The subscriptions should be retried automatically. An interface is being built for you to review problems resulting from these events to take manual action when needed.
// * `unpaid` - Indicates an unpaid subscription. A subscription is marked unpaid if the retry period expires and you have configured your [Dunning](https://maxio.zendesk.com/hc/en-us/articles/24287076583565-Dunning-Overview) settings to have a Final Action of `mark the subscription unpaid`.
// * **End of Life States**
// * `canceled` - Indicates a canceled subscription. This may happen at your request (via the API or the web interface) or due to the expiration of the [Dunning](https://maxio.zendesk.com/hc/en-us/articles/24287076583565-Dunning-Overview) process without payment. See the [Reactivation](https://maxio.zendesk.com/hc/en-us/articles/24252109503629-Reactivating-and-Resuming) documentation for info on how to restart a canceled subscription.
// While a subscription is canceled, its period will not advance, it will not accrue any new charges, and Advanced Billing will not attempt to collect the overdue balance.
// * `expired` - Indicates a subscription that has expired due to running its normal life cycle. Some products may be configured to have an expiration period. An expired subscription then is one that stayed active until it fulfilled its full period.
// * `failed_to_create` - Indicates that signup has failed. (You may see this state in a signup_failure webhook.)
// * `on_hold` - Indicates that a subscription’s billing has been temporarily stopped. While it is expected that the subscription will resume and return to active status, this is still treated as an “End of Life” state because the customer is not paying for services during this time.
// * `suspended` - Indicates that a prepaid subscription has used up all their prepayment balance. If a prepayment is applied, it will return to an active state.
// * `trial_ended` - A subscription in a trial_ended state is a subscription that completed a no-obligation trial and did not have a card on file at the expiration of the trial period. See [Product Pricing – No Obligation Trials](https://maxio.zendesk.com/hc/en-us/articles/24261076617869-Product-Editing) for more details.
// See [Subscription States](https://maxio.zendesk.com/hc/en-us/articles/24252119027853-Subscription-States) for more info about subscription states and state transitions.
type SubscriptionState string

// MarshalJSON implements the json.Marshaler interface for SubscriptionState.
// It customizes the JSON marshaling process for SubscriptionState objects.
func (e SubscriptionState) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionState")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionState.
// It customizes the JSON unmarshaling process for SubscriptionState objects.
func (e *SubscriptionState) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionState(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionState")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionState.
func (e SubscriptionState) isValid() bool {
    switch e {
    case SubscriptionState_PENDING,
        SubscriptionState_FAILEDTOCREATE,
        SubscriptionState_TRIALING,
        SubscriptionState_ASSESSING,
        SubscriptionState_ACTIVE,
        SubscriptionState_SOFTFAILURE,
        SubscriptionState_PASTDUE,
        SubscriptionState_SUSPENDED,
        SubscriptionState_CANCELED,
        SubscriptionState_EXPIRED,
        SubscriptionState_PAUSED,
        SubscriptionState_UNPAID,
        SubscriptionState_TRIALENDED,
        SubscriptionState_ONHOLD,
        SubscriptionState_AWAITINGSIGNUP:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for SubscriptionStateFilter.
// It customizes the JSON marshaling process for SubscriptionStateFilter objects.
func (e SubscriptionStateFilter) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for SubscriptionStateFilter")
}

// UnmarshalJSON implements the json.Unmarshaler interface for SubscriptionStateFilter.
// It customizes the JSON unmarshaling process for SubscriptionStateFilter objects.
func (e *SubscriptionStateFilter) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = SubscriptionStateFilter(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to SubscriptionStateFilter")
    }
    return nil
}

// Checks whether the value is actually a member of SubscriptionStateFilter.
func (e SubscriptionStateFilter) isValid() bool {
    switch e {
    case SubscriptionStateFilter_ACTIVE,
        SubscriptionStateFilter_CANCELED,
        SubscriptionStateFilter_EXPIRED,
        SubscriptionStateFilter_EXPIREDCARDS,
        SubscriptionStateFilter_ONHOLD,
        SubscriptionStateFilter_PASTDUE,
        SubscriptionStateFilter_PENDINGCANCELLATION,
        SubscriptionStateFilter_PENDINGRENEWAL,
        SubscriptionStateFilter_SUSPENDED,
        SubscriptionStateFilter_TRIALENDED,
        SubscriptionStateFilter_TRIALING,
        SubscriptionStateFilter_UNPAID:
        return true
    }
    return false
}

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

// MarshalJSON implements the json.Marshaler interface for TaxConfigurationKind.
// It customizes the JSON marshaling process for TaxConfigurationKind objects.
func (e TaxConfigurationKind) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for TaxConfigurationKind")
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaxConfigurationKind.
// It customizes the JSON unmarshaling process for TaxConfigurationKind objects.
func (e *TaxConfigurationKind) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = TaxConfigurationKind(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to TaxConfigurationKind")
    }
    return nil
}

// Checks whether the value is actually a member of TaxConfigurationKind.
func (e TaxConfigurationKind) isValid() bool {
    switch e {
    case TaxConfigurationKind_CUSTOM,
        TaxConfigurationKind_ENUMMANAGEDAVALARA,
        TaxConfigurationKind_ENUMLINKEDAVALARA,
        TaxConfigurationKind_ENUMDIGITALRIVER:
        return true
    }
    return false
}

const (
    TaxConfigurationKind_CUSTOM             TaxConfigurationKind = "custom"
    TaxConfigurationKind_ENUMMANAGEDAVALARA TaxConfigurationKind = "managed avalara"
    TaxConfigurationKind_ENUMLINKEDAVALARA  TaxConfigurationKind = "linked avalara"
    TaxConfigurationKind_ENUMDIGITALRIVER   TaxConfigurationKind = "digital river"
)

// TaxDestinationAddress is a string enum.
type TaxDestinationAddress string

// MarshalJSON implements the json.Marshaler interface for TaxDestinationAddress.
// It customizes the JSON marshaling process for TaxDestinationAddress objects.
func (e TaxDestinationAddress) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for TaxDestinationAddress")
}

// UnmarshalJSON implements the json.Unmarshaler interface for TaxDestinationAddress.
// It customizes the JSON unmarshaling process for TaxDestinationAddress objects.
func (e *TaxDestinationAddress) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = TaxDestinationAddress(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to TaxDestinationAddress")
    }
    return nil
}

// Checks whether the value is actually a member of TaxDestinationAddress.
func (e TaxDestinationAddress) isValid() bool {
    switch e {
    case TaxDestinationAddress_SHIPPINGTHENBILLING,
        TaxDestinationAddress_BILLINGTHENSHIPPING,
        TaxDestinationAddress_SHIPPINGONLY,
        TaxDestinationAddress_BILLINGONLY:
        return true
    }
    return false
}

const (
    TaxDestinationAddress_SHIPPINGTHENBILLING TaxDestinationAddress = "shipping_then_billing"
    TaxDestinationAddress_BILLINGTHENSHIPPING TaxDestinationAddress = "billing_then_shipping"
    TaxDestinationAddress_SHIPPINGONLY        TaxDestinationAddress = "shipping_only"
    TaxDestinationAddress_BILLINGONLY         TaxDestinationAddress = "billing_only"
)

// TrialType is a string enum.
// Indicates how a trial is handled when the trail period ends and there is no credit card on file. For `no_obligation`, the subscription transitions to a Trial Ended state. Maxio will not send any emails or statements. For `payment_expected`, the subscription transitions to a Past Due state. Maxio will send normal dunning emails and statements according to your other settings.
type TrialType string

// MarshalJSON implements the json.Marshaler interface for TrialType.
// It customizes the JSON marshaling process for TrialType objects.
func (e TrialType) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for TrialType")
}

// UnmarshalJSON implements the json.Unmarshaler interface for TrialType.
// It customizes the JSON unmarshaling process for TrialType objects.
func (e *TrialType) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = TrialType(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to TrialType")
    }
    return nil
}

// Checks whether the value is actually a member of TrialType.
func (e TrialType) isValid() bool {
    switch e {
    case TrialType_NOOBLIGATION,
        TrialType_PAYMENTEXPECTED:
        return true
    }
    return false
}

const (
    TrialType_NOOBLIGATION    TrialType = "no_obligation"
    TrialType_PAYMENTEXPECTED TrialType = "payment_expected"
)

// WebhookOrder is a string enum.
type WebhookOrder string

// MarshalJSON implements the json.Marshaler interface for WebhookOrder.
// It customizes the JSON marshaling process for WebhookOrder objects.
func (e WebhookOrder) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for WebhookOrder")
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookOrder.
// It customizes the JSON unmarshaling process for WebhookOrder objects.
func (e *WebhookOrder) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = WebhookOrder(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to WebhookOrder")
    }
    return nil
}

// Checks whether the value is actually a member of WebhookOrder.
func (e WebhookOrder) isValid() bool {
    switch e {
    case WebhookOrder_NEWESTFIRST,
        WebhookOrder_OLDESTFIRST:
        return true
    }
    return false
}

const (
    WebhookOrder_NEWESTFIRST WebhookOrder = "newest_first"
    WebhookOrder_OLDESTFIRST WebhookOrder = "oldest_first"
)

// WebhookStatus is a string enum.
type WebhookStatus string

// MarshalJSON implements the json.Marshaler interface for WebhookStatus.
// It customizes the JSON marshaling process for WebhookStatus objects.
func (e WebhookStatus) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for WebhookStatus")
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookStatus.
// It customizes the JSON unmarshaling process for WebhookStatus objects.
func (e *WebhookStatus) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = WebhookStatus(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to WebhookStatus")
    }
    return nil
}

// Checks whether the value is actually a member of WebhookStatus.
func (e WebhookStatus) isValid() bool {
    switch e {
    case WebhookStatus_SUCCESSFUL,
        WebhookStatus_FAILED,
        WebhookStatus_PENDING,
        WebhookStatus_PAUSED:
        return true
    }
    return false
}

const (
    WebhookStatus_SUCCESSFUL WebhookStatus = "successful"
    WebhookStatus_FAILED     WebhookStatus = "failed"
    WebhookStatus_PENDING    WebhookStatus = "pending"
    WebhookStatus_PAUSED     WebhookStatus = "paused"
)

// WebhookSubscription is a string enum.
type WebhookSubscription string

// MarshalJSON implements the json.Marshaler interface for WebhookSubscription.
// It customizes the JSON marshaling process for WebhookSubscription objects.
func (e WebhookSubscription) MarshalJSON() (
    []byte,
    error) {
    if e.isValid() {
        return []byte(fmt.Sprintf("\"%v\"", e)), nil
    }
    return nil, errors.New("the provided enum value is not allowed for WebhookSubscription")
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSubscription.
// It customizes the JSON unmarshaling process for WebhookSubscription objects.
func (e *WebhookSubscription) UnmarshalJSON(input []byte) error {
    var enumValue string
    err := json.Unmarshal(input, &enumValue)
    if err != nil {
        return err
    }
    *e = WebhookSubscription(enumValue)
    if !e.isValid() {
        return errors.New("the value " + string(input) + " cannot be unmarshalled to WebhookSubscription")
    }
    return nil
}

// Checks whether the value is actually a member of WebhookSubscription.
func (e WebhookSubscription) isValid() bool {
    switch e {
    case WebhookSubscription_BILLINGDATECHANGE,
        WebhookSubscription_COMPONENTALLOCATIONCHANGE,
        WebhookSubscription_CUSTOMERCREATE,
        WebhookSubscription_CUSTOMERUPDATE,
        WebhookSubscription_DUNNINGSTEPREACHED,
        WebhookSubscription_EXPIRINGCARD,
        WebhookSubscription_EXPIRATIONDATECHANGE,
        WebhookSubscription_INVOICEISSUED,
        WebhookSubscription_METEREDUSAGE,
        WebhookSubscription_PAYMENTFAILURE,
        WebhookSubscription_PAYMENTSUCCESS,
        WebhookSubscription_DIRECTDEBITPAYMENTPENDING,
        WebhookSubscription_DIRECTDEBITPAYMENTPAIDOUT,
        WebhookSubscription_DIRECTDEBITPAYMENTREJECTED,
        WebhookSubscription_PREPAIDSUBSCRIPTIONBALANCECHANGED,
        WebhookSubscription_PREPAIDUSAGE,
        WebhookSubscription_REFUNDFAILURE,
        WebhookSubscription_REFUNDSUCCESS,
        WebhookSubscription_RENEWALFAILURE,
        WebhookSubscription_RENEWALSUCCESS,
        WebhookSubscription_SIGNUPFAILURE,
        WebhookSubscription_SIGNUPSUCCESS,
        WebhookSubscription_STATEMENTCLOSED,
        WebhookSubscription_STATEMENTSETTLED,
        WebhookSubscription_SUBSCRIPTIONCARDUPDATE,
        WebhookSubscription_SUBSCRIPTIONGROUPCARDUPDATE,
        WebhookSubscription_SUBSCRIPTIONPRODUCTCHANGE,
        WebhookSubscription_SUBSCRIPTIONSTATECHANGE,
        WebhookSubscription_TRIALENDNOTICE,
        WebhookSubscription_UPCOMINGRENEWALNOTICE,
        WebhookSubscription_UPGRADEDOWNGRADEFAILURE,
        WebhookSubscription_UPGRADEDOWNGRADESUCCESS,
        WebhookSubscription_PENDINGCANCELLATIONCHANGE,
        WebhookSubscription_SUBSCRIPTIONPREPAYMENTACCOUNTBALANCECHANGED,
        WebhookSubscription_SUBSCRIPTIONSERVICECREDITACCOUNTBALANCECHANGED:
        return true
    }
    return false
}

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
