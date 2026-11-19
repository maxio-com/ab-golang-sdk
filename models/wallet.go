// Package advancedbilling
// This file was automatically generated for Maxio by APIMATIC v3.0 ( https://www.apimatic.io ).
package models

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Wallet represents a Wallet struct.
// A balance, denominated in the unit of a Wallet Product, owned by a subscription. Grants and recurring funding schedules are embedded so a single read shows the full balance breakdown.
type Wallet struct {
    // The Advanced Billing id of the wallet.
    Id                         *int                     `json:"id,omitempty"`
    // The polymorphic owner type. `subscription` is the only value currently supported.
    OwnerType                  *OwnerType               `json:"owner_type,omitempty"`
    // The id of the owning subscription.
    OwnerId                    *int                     `json:"owner_id,omitempty"`
    // (Optional) A merchant-facing label for the wallet, set on create and changeable with [Update Wallet]($e/Wallets/updateWallet). Surrounding spaces are stripped; a blank value is stored as `null`.
    Name                       Optional[string]         `json:"name"`
    // The id of the [Wallet Product]($e/Wallet%20Products/createWalletProduct) that denominates the wallet. Immutable for the life of the wallet.
    WalletProductId            *int                     `json:"wallet_product_id,omitempty"`
    // The name of the Wallet Product that denominates the wallet.
    WalletProductName          *string                  `json:"wallet_product_name,omitempty"`
    // The billing product's unit name, immutable for the life of the wallet.
    Unit                       *string                  `json:"unit,omitempty"`
    // The price point that priced this wallet's [purchases]($e/Wallet%20Purchases/createWalletPurchase), set on create and immutable afterward. Defaults to the Wallet Product's default price point; `null` if the Wallet Product has none and none was given on create.
    PricePointId               Optional[int]            `json:"price_point_id"`
    // What the wallet can spend now, encoded as a decimal string to avoid floating-point precision loss. This is the number to enforce against. For a metered wallet, it includes usage not yet posted to the ledger, through `metered_through_at`, but only in the response from [Read Wallet]($e/Wallets/readWallet); every other response (list, create, update, archive) reports the balance as of the last ledger write, equal to `posted_balance`, even for a metered wallet. The grants' `remaining` follow this figure, so they add up to it whenever it is zero or more.
    Balance                    *string                  `json:"balance,omitempty"`
    // The balance from written ledger entries alone: grants funded, minus debits written to the ledger, as a decimal string. Always present, also for a wallet with no meter. It equals `balance` wherever no live metered balance is computed: list, create, update, and archive responses, an unlinked or archived wallet, and a read where `metered_stale` is `true`. Only an active, metered wallet read through [Read Wallet]($e/Wallets/readWallet), with `metered_stale` not `true`, can have a `balance` below its `posted_balance`. It stays exact when `metered_stale` is `true`. Reconcile the transactions list (`balance_after`) against this figure, not against `balance`.
    PostedBalance              *string                  `json:"posted_balance,omitempty"`
    // The metering meter this wallet is paired with, or `null` for a wallet with no meter. Set on create and immutable afterward.
    Meter                      Optional[Meter]          `json:"meter"`
    // The point up to which live meter usage has been factored into `balance`. Populated only in the response from [Read Wallet]($e/Wallets/readWallet), for an active, metered wallet on a site with metered consumption enabled; `null` everywhere else, including list, create, update, and archive responses for the same wallet.
    MeteredThroughAt           Optional[time.Time]      `json:"metered_through_at"`
    // How many seconds to wait before re-reading `balance` to see usage past `metered_through_at`. Same availability as `metered_through_at`.
    MeteredRecheckAfterSeconds Optional[int]            `json:"metered_recheck_after_seconds"`
    // `true` when live usage could not be read fresh and `balance` reflects an earlier snapshot. Same availability as `metered_through_at`.
    MeteredStale               Optional[bool]           `json:"metered_stale"`
    Status                     *WalletStatus            `json:"status,omitempty"`
    // What happens when a debit would take the wallet below zero. Known values: `hard_stop` (the only value you can set today). More values are expected over time, so treat this as an open string and don't fail on one you don't recognize. At the zero boundary, `hard_stop` behaves as set by `hard_stop_overflow`.
    DepletionPolicy            *string                  `json:"depletion_policy,omitempty"`
    // Behavior at the hard-stop boundary. `discard` drains the wallet to zero and rejects the excess (the default). `allow_negative` deducts the full amount, letting the balance go negative. The excess past the wallet's grants is booked as an overdraft ledger row with `grant_id` set to `null`.
    HardStopOverflow           *WalletHardStopOverflow  `json:"hard_stop_overflow,omitempty"`
    // How the wallet came to exist. `manual` for a wallet created directly with [Create Wallet]($e/Wallets/createWallet). `feature` for a wallet auto-provisioned because a subscription resolved a usage-limit entitlement whose Feature has `wallet_behavior` configured.
    Origin                     *WalletOrigin            `json:"origin,omitempty"`
    // The key of the Feature this wallet was provisioned for, when `origin` is `feature`. `null` for a `manual` wallet.
    FeatureKey                 Optional[string]         `json:"feature_key"`
    // Every grant on the wallet, including ones already spent or expired, in drain order (the order debits would consume the still-open ones in): the soonest-expiring grant first, evergreen grants last.
    Grants                     []WalletGrant            `json:"grants,omitempty"`
    // The wallet's automatic funding schedules, if any. Empty for a `manual` wallet with no funding entitlement. A wallet funds automatically when it's provisioned from a Feature with `wallet_behavior` configured; see [the Wallets concept guide](page:introduction/basic-concepts/wallets#provisioning-a-wallet-from-a-feature).
    RecurringFunding           []WalletRecurringFunding `json:"recurring_funding,omitempty"`
    CreatedAt                  *time.Time               `json:"created_at,omitempty"`
    UpdatedAt                  *time.Time               `json:"updated_at,omitempty"`
    AdditionalProperties       map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for Wallet,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w Wallet) String() string {
    return fmt.Sprintf(
    	"Wallet[Id=%v, OwnerType=%v, OwnerId=%v, Name=%v, WalletProductId=%v, WalletProductName=%v, Unit=%v, PricePointId=%v, Balance=%v, PostedBalance=%v, Meter=%v, MeteredThroughAt=%v, MeteredRecheckAfterSeconds=%v, MeteredStale=%v, Status=%v, DepletionPolicy=%v, HardStopOverflow=%v, Origin=%v, FeatureKey=%v, Grants=%v, RecurringFunding=%v, CreatedAt=%v, UpdatedAt=%v, AdditionalProperties=%v]",
    	w.Id, w.OwnerType, w.OwnerId, w.Name, w.WalletProductId, w.WalletProductName, w.Unit, w.PricePointId, w.Balance, w.PostedBalance, w.Meter, w.MeteredThroughAt, w.MeteredRecheckAfterSeconds, w.MeteredStale, w.Status, w.DepletionPolicy, w.HardStopOverflow, w.Origin, w.FeatureKey, w.Grants, w.RecurringFunding, w.CreatedAt, w.UpdatedAt, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Wallet.
// It customizes the JSON marshaling process for Wallet objects.
func (w Wallet) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "id", "owner_type", "owner_id", "name", "wallet_product_id", "wallet_product_name", "unit", "price_point_id", "balance", "posted_balance", "meter", "metered_through_at", "metered_recheck_after_seconds", "metered_stale", "status", "depletion_policy", "hard_stop_overflow", "origin", "feature_key", "grants", "recurring_funding", "created_at", "updated_at"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the Wallet object to a map representation for JSON marshaling.
func (w Wallet) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.OwnerType != nil {
        structMap["owner_type"] = w.OwnerType
    }
    if w.OwnerId != nil {
        structMap["owner_id"] = w.OwnerId
    }
    if w.Name.IsValueSet() {
        if w.Name.Value() != nil {
            structMap["name"] = w.Name.Value()
        } else {
            structMap["name"] = nil
        }
    }
    if w.WalletProductId != nil {
        structMap["wallet_product_id"] = w.WalletProductId
    }
    if w.WalletProductName != nil {
        structMap["wallet_product_name"] = w.WalletProductName
    }
    if w.Unit != nil {
        structMap["unit"] = w.Unit
    }
    if w.PricePointId.IsValueSet() {
        if w.PricePointId.Value() != nil {
            structMap["price_point_id"] = w.PricePointId.Value()
        } else {
            structMap["price_point_id"] = nil
        }
    }
    if w.Balance != nil {
        structMap["balance"] = w.Balance
    }
    if w.PostedBalance != nil {
        structMap["posted_balance"] = w.PostedBalance
    }
    if w.Meter.IsValueSet() {
        if w.Meter.Value() != nil {
            structMap["meter"] = w.Meter.Value().toMap()
        } else {
            structMap["meter"] = nil
        }
    }
    if w.MeteredThroughAt.IsValueSet() {
        var MeteredThroughAtVal *string = nil
        if w.MeteredThroughAt.Value() != nil {
            val := w.MeteredThroughAt.Value().Format(time.RFC3339)
            MeteredThroughAtVal = &val
        }
        if w.MeteredThroughAt.Value() != nil {
            structMap["metered_through_at"] = MeteredThroughAtVal
        } else {
            structMap["metered_through_at"] = nil
        }
    }
    if w.MeteredRecheckAfterSeconds.IsValueSet() {
        if w.MeteredRecheckAfterSeconds.Value() != nil {
            structMap["metered_recheck_after_seconds"] = w.MeteredRecheckAfterSeconds.Value()
        } else {
            structMap["metered_recheck_after_seconds"] = nil
        }
    }
    if w.MeteredStale.IsValueSet() {
        if w.MeteredStale.Value() != nil {
            structMap["metered_stale"] = w.MeteredStale.Value()
        } else {
            structMap["metered_stale"] = nil
        }
    }
    if w.Status != nil {
        structMap["status"] = w.Status
    }
    if w.DepletionPolicy != nil {
        structMap["depletion_policy"] = w.DepletionPolicy
    }
    if w.HardStopOverflow != nil {
        structMap["hard_stop_overflow"] = w.HardStopOverflow
    }
    if w.Origin != nil {
        structMap["origin"] = w.Origin
    }
    if w.FeatureKey.IsValueSet() {
        if w.FeatureKey.Value() != nil {
            structMap["feature_key"] = w.FeatureKey.Value()
        } else {
            structMap["feature_key"] = nil
        }
    }
    if w.Grants != nil {
        structMap["grants"] = w.Grants
    }
    if w.RecurringFunding != nil {
        structMap["recurring_funding"] = w.RecurringFunding
    }
    if w.CreatedAt != nil {
        structMap["created_at"] = w.CreatedAt.Format(time.RFC3339)
    }
    if w.UpdatedAt != nil {
        structMap["updated_at"] = w.UpdatedAt.Format(time.RFC3339)
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Wallet.
// It customizes the JSON unmarshaling process for Wallet objects.
func (w *Wallet) UnmarshalJSON(input []byte) error {
    var temp tempWallet
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "owner_type", "owner_id", "name", "wallet_product_id", "wallet_product_name", "unit", "price_point_id", "balance", "posted_balance", "meter", "metered_through_at", "metered_recheck_after_seconds", "metered_stale", "status", "depletion_policy", "hard_stop_overflow", "origin", "feature_key", "grants", "recurring_funding", "created_at", "updated_at")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Id = temp.Id
    w.OwnerType = temp.OwnerType
    w.OwnerId = temp.OwnerId
    w.Name = temp.Name
    w.WalletProductId = temp.WalletProductId
    w.WalletProductName = temp.WalletProductName
    w.Unit = temp.Unit
    w.PricePointId = temp.PricePointId
    w.Balance = temp.Balance
    w.PostedBalance = temp.PostedBalance
    w.Meter = temp.Meter
    w.MeteredThroughAt.ShouldSetValue(temp.MeteredThroughAt.IsValueSet())
    if temp.MeteredThroughAt.Value() != nil {
        MeteredThroughAtVal, err := time.Parse(time.RFC3339, (*temp.MeteredThroughAt.Value()))
        if err != nil {
            log.Fatalf("Cannot Parse metered_through_at as % s format.", time.RFC3339)
        }
        w.MeteredThroughAt.SetValue(&MeteredThroughAtVal)
    }
    w.MeteredRecheckAfterSeconds = temp.MeteredRecheckAfterSeconds
    w.MeteredStale = temp.MeteredStale
    w.Status = temp.Status
    w.DepletionPolicy = temp.DepletionPolicy
    w.HardStopOverflow = temp.HardStopOverflow
    w.Origin = temp.Origin
    w.FeatureKey = temp.FeatureKey
    w.Grants = temp.Grants
    w.RecurringFunding = temp.RecurringFunding
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        w.CreatedAt = &CreatedAtVal
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        w.UpdatedAt = &UpdatedAtVal
    }
    return nil
}

// tempWallet is a temporary struct used for validating the fields of Wallet.
type tempWallet  struct {
    Id                         *int                     `json:"id,omitempty"`
    OwnerType                  *OwnerType               `json:"owner_type,omitempty"`
    OwnerId                    *int                     `json:"owner_id,omitempty"`
    Name                       Optional[string]         `json:"name"`
    WalletProductId            *int                     `json:"wallet_product_id,omitempty"`
    WalletProductName          *string                  `json:"wallet_product_name,omitempty"`
    Unit                       *string                  `json:"unit,omitempty"`
    PricePointId               Optional[int]            `json:"price_point_id"`
    Balance                    *string                  `json:"balance,omitempty"`
    PostedBalance              *string                  `json:"posted_balance,omitempty"`
    Meter                      Optional[Meter]          `json:"meter"`
    MeteredThroughAt           Optional[string]         `json:"metered_through_at"`
    MeteredRecheckAfterSeconds Optional[int]            `json:"metered_recheck_after_seconds"`
    MeteredStale               Optional[bool]           `json:"metered_stale"`
    Status                     *WalletStatus            `json:"status,omitempty"`
    DepletionPolicy            *string                  `json:"depletion_policy,omitempty"`
    HardStopOverflow           *WalletHardStopOverflow  `json:"hard_stop_overflow,omitempty"`
    Origin                     *WalletOrigin            `json:"origin,omitempty"`
    FeatureKey                 Optional[string]         `json:"feature_key"`
    Grants                     []WalletGrant            `json:"grants,omitempty"`
    RecurringFunding           []WalletRecurringFunding `json:"recurring_funding,omitempty"`
    CreatedAt                  *string                  `json:"created_at,omitempty"`
    UpdatedAt                  *string                  `json:"updated_at,omitempty"`
}
