// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// DigitalWalletTokenService contains methods and other services that help with
// interacting with the increase API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDigitalWalletTokenService] method
// instead.
type DigitalWalletTokenService struct {
	Options []option.RequestOption
}

// NewDigitalWalletTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDigitalWalletTokenService(opts ...option.RequestOption) (r *DigitalWalletTokenService) {
	r = &DigitalWalletTokenService{}
	r.Options = opts
	return
}

// Retrieve a Digital Wallet Token
func (r *DigitalWalletTokenService) Get(ctx context.Context, digitalWalletTokenID string, opts ...option.RequestOption) (res *DigitalWalletToken, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("digital_wallet_tokens/%s", digitalWalletTokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) List(ctx context.Context, query DigitalWalletTokenListParams, opts ...option.RequestOption) (res *shared.Page[DigitalWalletToken], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "digital_wallet_tokens"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Digital Wallet Tokens
func (r *DigitalWalletTokenService) ListAutoPaging(ctx context.Context, query DigitalWalletTokenListParams, opts ...option.RequestOption) *shared.PageAutoPager[DigitalWalletToken] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// A Digital Wallet Token is created when a user adds a Card to their Apple Pay or
// Google Pay app. The Digital Wallet Token can be used for purchases just like a
// Card.
type DigitalWalletToken struct {
	// The Digital Wallet Token identifier.
	ID string `json:"id,required"`
	// The identifier for the Card this Digital Wallet Token belongs to.
	CardID string `json:"card_id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// This indicates if payments can be made with the Digital Wallet Token.
	Status DigitalWalletTokenStatus `json:"status,required"`
	// The digital wallet app being used.
	TokenRequestor DigitalWalletTokenTokenRequestor `json:"token_requestor,required"`
	// A constant representing the object's type. For this resource it will always be
	// `digital_wallet_token`.
	Type DigitalWalletTokenType `json:"type,required"`
	JSON digitalWalletTokenJSON `json:"-"`
}

// digitalWalletTokenJSON contains the JSON metadata for the struct
// [DigitalWalletToken]
type digitalWalletTokenJSON struct {
	ID             apijson.Field
	CardID         apijson.Field
	CreatedAt      apijson.Field
	Status         apijson.Field
	TokenRequestor apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DigitalWalletToken) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// This indicates if payments can be made with the Digital Wallet Token.
type DigitalWalletTokenStatus string

const (
	// The digital wallet token is active.
	DigitalWalletTokenStatusActive DigitalWalletTokenStatus = "active"
	// The digital wallet token has been created but not successfully activated via
	// two-factor authentication yet.
	DigitalWalletTokenStatusInactive DigitalWalletTokenStatus = "inactive"
	// The digital wallet token has been temporarily paused.
	DigitalWalletTokenStatusSuspended DigitalWalletTokenStatus = "suspended"
	// The digital wallet token has been permanently canceled.
	DigitalWalletTokenStatusDeactivated DigitalWalletTokenStatus = "deactivated"
)

// The digital wallet app being used.
type DigitalWalletTokenTokenRequestor string

const (
	// Apple Pay
	DigitalWalletTokenTokenRequestorApplePay DigitalWalletTokenTokenRequestor = "apple_pay"
	// Google Pay
	DigitalWalletTokenTokenRequestorGooglePay DigitalWalletTokenTokenRequestor = "google_pay"
	// Unknown
	DigitalWalletTokenTokenRequestorUnknown DigitalWalletTokenTokenRequestor = "unknown"
)

// A constant representing the object's type. For this resource it will always be
// `digital_wallet_token`.
type DigitalWalletTokenType string

const (
	DigitalWalletTokenTypeDigitalWalletToken DigitalWalletTokenType = "digital_wallet_token"
)

type DigitalWalletTokenListParams struct {
	// Filter Digital Wallet Tokens to ones belonging to the specified Card.
	CardID    param.Field[string]                                `query:"card_id"`
	CreatedAt param.Field[DigitalWalletTokenListParamsCreatedAt] `query:"created_at"`
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [DigitalWalletTokenListParams]'s query parameters as
// `url.Values`.
func (r DigitalWalletTokenListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DigitalWalletTokenListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [DigitalWalletTokenListParamsCreatedAt]'s query parameters
// as `url.Values`.
func (r DigitalWalletTokenListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
