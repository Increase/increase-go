package requests

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core/fields"
	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type CreateAnAccountParameters struct {
	// The identifier for the Entity that will own the Account.
	EntityID fields.Field[string] `json:"entity_id"`
	// The identifier for the Program that this Account falls under.
	ProgramID fields.Field[string] `json:"program_id"`
	// The identifier of an Entity that, while not owning the Account, is associated
	// with its activity. Its relationship to your group must be `informational`.
	InformationalEntityID fields.Field[string] `json:"informational_entity_id"`
	// The name you choose for the Account.
	Name fields.Field[string] `json:"name,required"`
}

// MarshalJSON serializes CreateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r CreateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountParameters{EntityID:%s ProgramID:%s InformationalEntityID:%s Name:%s}", r.EntityID, r.ProgramID, r.InformationalEntityID, r.Name)
}

type UpdateAnAccountParameters struct {
	// The new name of the Account.
	Name fields.Field[string] `json:"name"`
}

// MarshalJSON serializes UpdateAnAccountParameters into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *UpdateAnAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r UpdateAnAccountParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnAccountParameters{Name:%s}", r.Name)
}

type AccountListParams struct {
	// Return the page of entries after this one.
	Cursor fields.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit fields.Field[int64] `query:"limit"`
	// Filter Accounts for those belonging to the specified Entity.
	EntityID fields.Field[string] `query:"entity_id"`
	// Filter Accounts for those with the specified status.
	Status fields.Field[AccountListParamsStatus] `query:"status"`
}

// URLQuery serializes AccountListParams into a url.Values of the query parameters
// associated with this value
func (r *AccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

func (r AccountListParams) String() (result string) {
	return fmt.Sprintf("&AccountListParams{Cursor:%s Limit:%s EntityID:%s Status:%s}", r.Cursor, r.Limit, r.EntityID, r.Status)
}

type AccountListParamsStatus string

const (
	AccountListParamsStatusOpen   AccountListParamsStatus = "open"
	AccountListParamsStatusClosed AccountListParamsStatus = "closed"
)
