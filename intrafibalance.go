// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Increase/increase-go/internal/apijson"
	"github.com/Increase/increase-go/internal/requestconfig"
	"github.com/Increase/increase-go/option"
)

// IntrafiBalanceService contains methods and other services that help with
// interacting with the increase API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntrafiBalanceService] method instead.
type IntrafiBalanceService struct {
	Options []option.RequestOption
}

// NewIntrafiBalanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntrafiBalanceService(opts ...option.RequestOption) (r *IntrafiBalanceService) {
	r = &IntrafiBalanceService{}
	r.Options = opts
	return
}

// Returns the IntraFi balance for the given account. IntraFi may sweep funds to
// multiple banks. This endpoint will include both the total balance and the amount
// swept to each institution.
func (r *IntrafiBalanceService) IntrafiBalance(ctx context.Context, accountID string, opts ...option.RequestOption) (res *IntrafiBalance, err error) {
	opts = append(r.Options[:], opts...)
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intrafi_balance", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// When using IntraFi, each account's balance over the standard FDIC insurance
// amount is swept to various other institutions. Funds are rebalanced across banks
// as needed once per business day.
type IntrafiBalance struct {
	// The identifier of this balance.
	ID string `json:"id,required"`
	// Each entry represents a balance held at a different bank. IntraFi separates the
	// total balance across many participating banks in the network.
	Balances []IntrafiBalanceBalance `json:"balances,required"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the account
	// currency.
	Currency IntrafiBalanceCurrency `json:"currency,required"`
	// The date this balance reflects.
	EffectiveDate time.Time `json:"effective_date,required" format:"date"`
	// The total balance, in minor units of `currency`. Increase reports this balance
	// to IntraFi daily.
	TotalBalance int64 `json:"total_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `intrafi_balance`.
	Type IntrafiBalanceType `json:"type,required"`
	JSON intrafiBalanceJSON `json:"-"`
}

// intrafiBalanceJSON contains the JSON metadata for the struct [IntrafiBalance]
type intrafiBalanceJSON struct {
	ID            apijson.Field
	Balances      apijson.Field
	Currency      apijson.Field
	EffectiveDate apijson.Field
	TotalBalance  apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IntrafiBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intrafiBalanceJSON) RawJSON() string {
	return r.raw
}

type IntrafiBalanceBalance struct {
	// The identifier of this balance.
	ID string `json:"id,required"`
	// The balance, in minor units of `currency`, held with this bank.
	Balance int64 `json:"balance,required"`
	// The name of the bank holding these funds.
	Bank string `json:"bank,required"`
	// The primary location of the bank.
	BankLocation IntrafiBalanceBalancesBankLocation `json:"bank_location,required,nullable"`
	// The Federal Deposit Insurance Corporation (FDIC) certificate number of the bank.
	// Because many banks have the same or similar names, this can be used to uniquely
	// identify the institution.
	FdicCertificateNumber string                    `json:"fdic_certificate_number,required"`
	JSON                  intrafiBalanceBalanceJSON `json:"-"`
}

// intrafiBalanceBalanceJSON contains the JSON metadata for the struct
// [IntrafiBalanceBalance]
type intrafiBalanceBalanceJSON struct {
	ID                    apijson.Field
	Balance               apijson.Field
	Bank                  apijson.Field
	BankLocation          apijson.Field
	FdicCertificateNumber apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *IntrafiBalanceBalance) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intrafiBalanceBalanceJSON) RawJSON() string {
	return r.raw
}

// The primary location of the bank.
type IntrafiBalanceBalancesBankLocation struct {
	// The bank's city.
	City string `json:"city,required"`
	// The bank's state.
	State string                                 `json:"state,required"`
	JSON  intrafiBalanceBalancesBankLocationJSON `json:"-"`
}

// intrafiBalanceBalancesBankLocationJSON contains the JSON metadata for the struct
// [IntrafiBalanceBalancesBankLocation]
type intrafiBalanceBalancesBankLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntrafiBalanceBalancesBankLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intrafiBalanceBalancesBankLocationJSON) RawJSON() string {
	return r.raw
}

// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the account
// currency.
type IntrafiBalanceCurrency string

const (
	IntrafiBalanceCurrencyCad IntrafiBalanceCurrency = "CAD"
	IntrafiBalanceCurrencyChf IntrafiBalanceCurrency = "CHF"
	IntrafiBalanceCurrencyEur IntrafiBalanceCurrency = "EUR"
	IntrafiBalanceCurrencyGbp IntrafiBalanceCurrency = "GBP"
	IntrafiBalanceCurrencyJpy IntrafiBalanceCurrency = "JPY"
	IntrafiBalanceCurrencyUsd IntrafiBalanceCurrency = "USD"
)

func (r IntrafiBalanceCurrency) IsKnown() bool {
	switch r {
	case IntrafiBalanceCurrencyCad, IntrafiBalanceCurrencyChf, IntrafiBalanceCurrencyEur, IntrafiBalanceCurrencyGbp, IntrafiBalanceCurrencyJpy, IntrafiBalanceCurrencyUsd:
		return true
	}
	return false
}

// A constant representing the object's type. For this resource it will always be
// `intrafi_balance`.
type IntrafiBalanceType string

const (
	IntrafiBalanceTypeIntrafiBalance IntrafiBalanceType = "intrafi_balance"
)

func (r IntrafiBalanceType) IsKnown() bool {
	switch r {
	case IntrafiBalanceTypeIntrafiBalance:
		return true
	}
	return false
}
