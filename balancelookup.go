package increase

import (
	"context"
	"net/http"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/field"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/option"
)

type BalanceLookupService struct {
	Options []option.RequestOption
}

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
	JSON BalanceLookupLookupResponseJSON
}

type BalanceLookupLookupResponseJSON struct {
	AccountID        apijson.Metadata
	CurrentBalance   apijson.Metadata
	AvailableBalance apijson.Metadata
	Type             apijson.Metadata
	raw              string
	Extras           map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BalanceLookupLookupResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *BalanceLookupLookupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BalanceLookupLookupResponseType string

const (
	BalanceLookupLookupResponseTypeBalanceLookup BalanceLookupLookupResponseType = "balance_lookup"
)

type BalanceLookupLookupParams struct {
	// The Account to query the balance for.
	AccountID field.Field[string] `json:"account_id,required"`
}

// MarshalJSON serializes BalanceLookupLookupParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r BalanceLookupLookupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
