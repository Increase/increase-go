package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateAnExternalAccountParameters struct {
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber fields.Field[string] `json:"routing_number,required"`
	// The account number for the destination account.
	AccountNumber fields.Field[string] `json:"account_number,required"`
	// The type of the destination account. Defaults to `checking`.
	Funding fields.Field[CreateAnExternalAccountParametersFunding] `json:"funding"`
	// The name you choose for the Account.
	Description fields.Field[string] `json:"description,required"`
}

// MarshalJSON serializes CreateAnExternalAccountParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnExternalAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnExternalAccountParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnExternalAccountParameters{RoutingNumber:%s AccountNumber:%s Funding:%s Description:%s}", r.RoutingNumber, r.AccountNumber, r.Funding, r.Description)
}

type CreateAnExternalAccountParametersFunding string

const (
	CreateAnExternalAccountParametersFundingChecking CreateAnExternalAccountParametersFunding = "checking"
	CreateAnExternalAccountParametersFundingSavings  CreateAnExternalAccountParametersFunding = "savings"
	CreateAnExternalAccountParametersFundingOther    CreateAnExternalAccountParametersFunding = "other"
)

type UpdateAnExternalAccountParameters struct {
	// The description you choose to give the external account.
	Description fields.Field[string] `json:"description"`
	// The status of the External Account.
	Status fields.Field[UpdateAnExternalAccountParametersStatus] `json:"status"`
}

// MarshalJSON serializes UpdateAnExternalAccountParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *UpdateAnExternalAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateAnExternalAccountParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnExternalAccountParameters{Description:%s Status:%s}", r.Description, r.Status)
}

type UpdateAnExternalAccountParametersStatus string

const (
	UpdateAnExternalAccountParametersStatusActive   UpdateAnExternalAccountParametersStatus = "active"
	UpdateAnExternalAccountParametersStatusArchived UpdateAnExternalAccountParametersStatus = "archived"
)

type ExternalAccountListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit  fields.Field[int64]                           `query:"limit"`
	Status fields.Field[ExternalAccountListParamsStatus] `query:"status"`
}

// URLQuery serializes ExternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r *ExternalAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ExternalAccountListParams) String() (result string) {
	return fmt.Sprintf("&ExternalAccountListParams{Cursor:%s Limit:%s Status:%s}", r.Cursor, r.Limit, r.Status)
}

type ExternalAccountListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In fields.Field[[]ExternalAccountListParamsStatusIn] `query:"in"`
}

// URLQuery serializes ExternalAccountListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *ExternalAccountListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r ExternalAccountListParamsStatus) String() (result string) {
	return fmt.Sprintf("&ExternalAccountListParamsStatus{In:%s}", core.Fmt(r.In))
}

type ExternalAccountListParamsStatusIn string

const (
	ExternalAccountListParamsStatusInActive   ExternalAccountListParamsStatusIn = "active"
	ExternalAccountListParamsStatusInArchived ExternalAccountListParamsStatusIn = "archived"
)
