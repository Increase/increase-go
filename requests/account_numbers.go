package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateAnAccountNumberParameters struct {
	// The Account the Account Number should belong to.
	AccountID fields.Field[string] `json:"account_id,required"`
	// The name you choose for the Account Number.
	Name fields.Field[string] `json:"name,required"`
}

// MarshalJSON serializes CreateAnAccountNumberParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountNumberParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnAccountNumberParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountNumberParameters{AccountID:%s Name:%s}", r.AccountID, r.Name)
}

type UpdateAnAccountNumberParameters struct {
	// The name you choose for the Account Number.
	Name fields.Field[string] `json:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status fields.Field[UpdateAnAccountNumberParametersStatus] `json:"status"`
}

// MarshalJSON serializes UpdateAnAccountNumberParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *UpdateAnAccountNumberParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateAnAccountNumberParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnAccountNumberParameters{Name:%s Status:%s}", r.Name, r.Status)
}

type UpdateAnAccountNumberParametersStatus string

const (
	UpdateAnAccountNumberParametersStatusActive   UpdateAnAccountNumberParametersStatus = "active"
	UpdateAnAccountNumberParametersStatusDisabled UpdateAnAccountNumberParametersStatus = "disabled"
	UpdateAnAccountNumberParametersStatusCanceled UpdateAnAccountNumberParametersStatus = "canceled"
)

type AccountNumberListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status fields.Field[AccountNumberListParamsStatus] `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID fields.Field[string] `query:"account_id"`
}

// URLQuery serializes AccountNumberListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountNumberListParams) String() (result string) {
	return fmt.Sprintf("&AccountNumberListParams{Cursor:%s Limit:%s Status:%s AccountID:%s}", r.Cursor, r.Limit, r.Status, r.AccountID)
}

type AccountNumberListParamsStatus string

const (
	AccountNumberListParamsStatusActive   AccountNumberListParamsStatus = "active"
	AccountNumberListParamsStatusDisabled AccountNumberListParamsStatus = "disabled"
	AccountNumberListParamsStatusCanceled AccountNumberListParamsStatus = "canceled"
)
