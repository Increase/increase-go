package types

import (
	"fmt"
	"net/url"
	"time"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type ExternalAccount struct {
	// The External Account's identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt *time.Time `pjson:"created_at" format:"2006-01-02T15:04:05Z07:00"`
	// The External Account's description for display purposes.
	Description *string `pjson:"description"`
	// The External Account's status.
	Status *ExternalAccountStatus `pjson:"status"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// The destination account number.
	AccountNumber *string `pjson:"account_number"`
	// The type of the account to which the transfer will be sent.
	Funding *ExternalAccountFunding `pjson:"funding"`
	// If you have verified ownership of the External Account.
	VerificationStatus *ExternalAccountVerificationStatus `pjson:"verification_status"`
	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type       *ExternalAccountType   `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ExternalAccount into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *ExternalAccount) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The External Account's identifier.
func (r ExternalAccount) GetID() (ID string) {
	if r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the External Account was created.
func (r ExternalAccount) GetCreatedAt() (CreatedAt time.Time) {
	if r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The External Account's description for display purposes.
func (r ExternalAccount) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The External Account's status.
func (r ExternalAccount) GetStatus() (Status ExternalAccountStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r ExternalAccount) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The destination account number.
func (r ExternalAccount) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The type of the account to which the transfer will be sent.
func (r ExternalAccount) GetFunding() (Funding ExternalAccountFunding) {
	if r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// If you have verified ownership of the External Account.
func (r ExternalAccount) GetVerificationStatus() (VerificationStatus ExternalAccountVerificationStatus) {
	if r.VerificationStatus != nil {
		VerificationStatus = *r.VerificationStatus
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `external_account`.
func (r ExternalAccount) GetType() (Type ExternalAccountType) {
	if r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r ExternalAccount) String() (result string) {
	return fmt.Sprintf("&ExternalAccount{ID:%s CreatedAt:%s Description:%s Status:%s RoutingNumber:%s AccountNumber:%s Funding:%s VerificationStatus:%s Type:%s}", core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.Description), core.FmtP(r.Status), core.FmtP(r.RoutingNumber), core.FmtP(r.AccountNumber), core.FmtP(r.Funding), core.FmtP(r.VerificationStatus), core.FmtP(r.Type))
}

type ExternalAccountStatus string

const (
	ExternalAccountStatusActive   ExternalAccountStatus = "active"
	ExternalAccountStatusArchived ExternalAccountStatus = "archived"
)

type ExternalAccountFunding string

const (
	ExternalAccountFundingChecking ExternalAccountFunding = "checking"
	ExternalAccountFundingSavings  ExternalAccountFunding = "savings"
	ExternalAccountFundingOther    ExternalAccountFunding = "other"
)

type ExternalAccountVerificationStatus string

const (
	ExternalAccountVerificationStatusUnverified ExternalAccountVerificationStatus = "unverified"
	ExternalAccountVerificationStatusPending    ExternalAccountVerificationStatus = "pending"
	ExternalAccountVerificationStatusVerified   ExternalAccountVerificationStatus = "verified"
)

type ExternalAccountType string

const (
	ExternalAccountTypeExternalAccount ExternalAccountType = "external_account"
)

type CreateAnExternalAccountParameters struct {
	// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
	// destination account.
	RoutingNumber *string `pjson:"routing_number"`
	// The account number for the destination account.
	AccountNumber *string `pjson:"account_number"`
	// The type of the destination account. Defaults to `checking`.
	Funding *CreateAnExternalAccountParametersFunding `pjson:"funding"`
	// The name you choose for the Account.
	Description *string                `pjson:"description"`
	jsonFields  map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnExternalAccountParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CreateAnExternalAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnExternalAccountParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnExternalAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN) for the
// destination account.
func (r CreateAnExternalAccountParameters) GetRoutingNumber() (RoutingNumber string) {
	if r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// The account number for the destination account.
func (r CreateAnExternalAccountParameters) GetAccountNumber() (AccountNumber string) {
	if r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The type of the destination account. Defaults to `checking`.
func (r CreateAnExternalAccountParameters) GetFunding() (Funding CreateAnExternalAccountParametersFunding) {
	if r.Funding != nil {
		Funding = *r.Funding
	}
	return
}

// The name you choose for the Account.
func (r CreateAnExternalAccountParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

func (r CreateAnExternalAccountParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnExternalAccountParameters{RoutingNumber:%s AccountNumber:%s Funding:%s Description:%s}", core.FmtP(r.RoutingNumber), core.FmtP(r.AccountNumber), core.FmtP(r.Funding), core.FmtP(r.Description))
}

type CreateAnExternalAccountParametersFunding string

const (
	CreateAnExternalAccountParametersFundingChecking CreateAnExternalAccountParametersFunding = "checking"
	CreateAnExternalAccountParametersFundingSavings  CreateAnExternalAccountParametersFunding = "savings"
	CreateAnExternalAccountParametersFundingOther    CreateAnExternalAccountParametersFunding = "other"
)

type UpdateAnExternalAccountParameters struct {
	// The description you choose to give the external account.
	Description *string `pjson:"description"`
	// The status of the External Account.
	Status     *UpdateAnExternalAccountParametersStatus `pjson:"status"`
	jsonFields map[string]interface{}                   `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// UpdateAnExternalAccountParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *UpdateAnExternalAccountParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateAnExternalAccountParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *UpdateAnExternalAccountParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The description you choose to give the external account.
func (r UpdateAnExternalAccountParameters) GetDescription() (Description string) {
	if r.Description != nil {
		Description = *r.Description
	}
	return
}

// The status of the External Account.
func (r UpdateAnExternalAccountParameters) GetStatus() (Status UpdateAnExternalAccountParametersStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r UpdateAnExternalAccountParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnExternalAccountParameters{Description:%s Status:%s}", core.FmtP(r.Description), core.FmtP(r.Status))
}

type UpdateAnExternalAccountParametersStatus string

const (
	UpdateAnExternalAccountParametersStatusActive   UpdateAnExternalAccountParametersStatus = "active"
	UpdateAnExternalAccountParametersStatusArchived UpdateAnExternalAccountParametersStatus = "archived"
)

type ExternalAccountListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit      *int64                           `query:"limit"`
	Status     *ExternalAccountListParamsStatus `query:"status"`
	jsonFields map[string]interface{}           `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountListParams
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ExternalAccountListParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ExternalAccountListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ExternalAccountListParams into a url.Values of the query
// parameters associated with this value
func (r *ExternalAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r ExternalAccountListParams) GetCursor() (Cursor string) {
	if r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r ExternalAccountListParams) GetLimit() (Limit int64) {
	if r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r ExternalAccountListParams) GetStatus() (Status ExternalAccountListParamsStatus) {
	if r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r ExternalAccountListParams) String() (result string) {
	return fmt.Sprintf("&ExternalAccountListParams{Cursor:%s Limit:%s Status:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), r.Status)
}

type ExternalAccountListParamsStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In         *[]ExternalAccountListParamsStatusIn `pjson:"in"`
	jsonFields map[string]interface{}               `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// ExternalAccountListParamsStatus using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *ExternalAccountListParamsStatus) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ExternalAccountListParamsStatus into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *ExternalAccountListParamsStatus) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ExternalAccountListParamsStatus into a url.Values of the
// query parameters associated with this value
func (r *ExternalAccountListParamsStatus) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r ExternalAccountListParamsStatus) GetIn() (In []ExternalAccountListParamsStatusIn) {
	if r.In != nil {
		In = *r.In
	}
	return
}

func (r ExternalAccountListParamsStatus) String() (result string) {
	return fmt.Sprintf("&ExternalAccountListParamsStatus{In:%s}", core.Fmt(r.In))
}

type ExternalAccountListParamsStatusIn string

const (
	ExternalAccountListParamsStatusInActive   ExternalAccountListParamsStatusIn = "active"
	ExternalAccountListParamsStatusInArchived ExternalAccountListParamsStatusIn = "archived"
)

type ExternalAccountList struct {
	// The contents of the list.
	Data *[]ExternalAccount `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccountList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes ExternalAccountList into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *ExternalAccountList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes ExternalAccountList into a url.Values of the query
// parameters associated with this value
func (r *ExternalAccountList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r ExternalAccountList) GetData() (Data []ExternalAccount) {
	if r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r ExternalAccountList) GetNextCursor() (NextCursor string) {
	if r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r ExternalAccountList) String() (result string) {
	return fmt.Sprintf("&ExternalAccountList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type ExternalAccountsPage struct {
	*pagination.Page[ExternalAccount]
}

func (r *ExternalAccountsPage) ExternalAccount() *ExternalAccount {
	return r.Current()
}

func (r *ExternalAccountsPage) NextPage() (*ExternalAccountsPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &ExternalAccountsPage{page}, nil
	}
}
