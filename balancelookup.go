package increase

import (
	"context"
	"net/http"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

// BalanceLookupService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewBalanceLookupService] method
// instead.
type BalanceLookupService struct {
	Options []option.RequestOption
}

// NewBalanceLookupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBalanceLookupService(opts ...option.RequestOption) (r *BalanceLookupService) {
	r = &BalanceLookupService{}
	r.Options = opts
	return
}

// Look up the balance for an Account
func (r *BalanceLookupService) Lookup(ctx context.Context, body BalanceLookupLookupParams, opts ...option.RequestOption) (res *BalanceLookupLookupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "balance_lookups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Represents a request to lookup the balance of an Account at a given point in
// time.
type BalanceLookupLookupResponse struct {
	// The identifier for the account for which the balance was queried.
	AccountID string `json:"account_id,required"`
	// The Account's current balance, representing the sum of all posted Transactions
	// on the Account.
	CurrentBalance int64 `json:"current_balance,required"`
	// The Account's available balance, representing the current balance less any open
	// Pending Transactions on the Account.
	AvailableBalance int64 `json:"available_balance,required"`
	// A constant representing the object's type. For this resource it will always be
	// `balance_lookup`.
	Type BalanceLookupLookupResponseType `json:"type,required"`
	JSON balanceLookupLookupResponseJSON
}

// balanceLookupLookupResponseJSON contains the JSON metadata for the struct
// [BalanceLookupLookupResponse]
type balanceLookupLookupResponseJSON struct {
	AccountID        apijson.Field
	CurrentBalance   apijson.Field
	AvailableBalance apijson.Field
	Type             apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *BalanceLookupLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceLookupLookupResponseType string

const (
	BalanceLookupLookupResponseTypeBalanceLookup BalanceLookupLookupResponseType = "balance_lookup"
)

type BalanceLookupLookupParams struct {
	// The Account to query the balance for.
	AccountID param.Field[string] `json:"account_id,required"`
	// The moment to query the balance at. If not set, returns the current balances.
	AtTime param.Field[time.Time] `json:"at_time" format:"date-time"`
}

func (r BalanceLookupLookupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
