package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type ExternalAccountNewParams struct {
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber field.Field[string] `json:"routing_number,required"`
	// The account number for the destination account.
	AccountNumber field.Field[string] `json:"account_number,required"`
	// The type of the destination account. Defaults to `checking`.
	Funding field.Field[ExternalAccountNewParamsFunding] `json:"funding"`
	// The name you choose for the Account.
	Description field.Field[string] `json:"description,required"`
}

// MarshalJSON serializes ExternalAccountNewParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountNewParamsFunding string

const (
	ExternalAccountNewParamsFundingChecking ExternalAccountNewParamsFunding = "checking"
	ExternalAccountNewParamsFundingSavings  ExternalAccountNewParamsFunding = "savings"
	ExternalAccountNewParamsFundingOther    ExternalAccountNewParamsFunding = "other"
)

type ExternalAccountUpdateParams struct {
	// The description you choose to give the external account.
	Description field.Field[string] `json:"description"`
	// The status of the External Account.
	Status field.Field[ExternalAccountUpdateParamsStatus] `json:"status"`
}

// MarshalJSON serializes ExternalAccountUpdateParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r ExternalAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExternalAccountUpdateParamsStatus string

const (
	ExternalAccountUpdateParamsStatusActive   ExternalAccountUpdateParamsStatus = "active"
	ExternalAccountUpdateParamsStatusArchived ExternalAccountUpdateParamsStatus = "archived"
)

type ExternalAccountListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  field.Field[int64]                           `query:"limit"`
	Status field.Field[ExternalAccountListParamsStatus] `query:"status"`
}

// URLQuery serializes ExternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r ExternalAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ExternalAccountListParamsStatus struct {
	// Filter External Accounts for those with the specified status or statuses. For
	// GET requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]ExternalAccountListParamsStatusIn] `query:"in"`
}

// URLQuery serializes ExternalAccountListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r ExternalAccountListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type ExternalAccountListParamsStatusIn string

const (
	ExternalAccountListParamsStatusInActive   ExternalAccountListParamsStatusIn = "active"
	ExternalAccountListParamsStatusInArchived ExternalAccountListParamsStatusIn = "archived"
)
