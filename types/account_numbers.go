package types

import (
	"fmt"
	"net/url"

	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/core/pjson"
	"github.com/increase/increase-go/core/query"
	"github.com/increase/increase-go/pagination"
)

type AccountNumber struct {
	// The identifier for the account this Account Number belongs to.
	AccountID *string `pjson:"account_id"`
	// The account number.
	AccountNumber *string `pjson:"account_number"`
	// The Account Number identifier.
	ID *string `pjson:"id"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Number was created.
	CreatedAt *string `pjson:"created_at"`
	// The name you choose for the Account Number.
	Name *string `pjson:"name"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber *string `pjson:"routing_number"`
	// This indicates if payments can be made to the Account Number.
	Status *AccountNumberStatus `pjson:"status"`
	// A constant representing the object's type. For this resource it will always be
	// `account_number`.
	Type       *AccountNumberType     `pjson:"type"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountNumber using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumber) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountNumber into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountNumber) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The identifier for the account this Account Number belongs to.
func (r *AccountNumber) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The account number.
func (r *AccountNumber) GetAccountNumber() (AccountNumber string) {
	if r != nil && r.AccountNumber != nil {
		AccountNumber = *r.AccountNumber
	}
	return
}

// The Account Number identifier.
func (r *AccountNumber) GetID() (ID string) {
	if r != nil && r.ID != nil {
		ID = *r.ID
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Number was created.
func (r *AccountNumber) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The name you choose for the Account Number.
func (r *AccountNumber) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// The American Bankers' Association (ABA) Routing Transit Number (RTN).
func (r *AccountNumber) GetRoutingNumber() (RoutingNumber string) {
	if r != nil && r.RoutingNumber != nil {
		RoutingNumber = *r.RoutingNumber
	}
	return
}

// This indicates if payments can be made to the Account Number.
func (r *AccountNumber) GetStatus() (Status AccountNumberStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_number`.
func (r *AccountNumber) GetType() (Type AccountNumberType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

func (r AccountNumber) String() (result string) {
	return fmt.Sprintf("&AccountNumber{AccountID:%s AccountNumber:%s ID:%s CreatedAt:%s Name:%s RoutingNumber:%s Status:%s Type:%s}", core.FmtP(r.AccountID), core.FmtP(r.AccountNumber), core.FmtP(r.ID), core.FmtP(r.CreatedAt), core.FmtP(r.Name), core.FmtP(r.RoutingNumber), core.FmtP(r.Status), core.FmtP(r.Type))
}

type AccountNumberStatus string

const (
	AccountNumberStatusActive   AccountNumberStatus = "active"
	AccountNumberStatusDisabled AccountNumberStatus = "disabled"
	AccountNumberStatusCanceled AccountNumberStatus = "canceled"
)

type AccountNumberType string

const (
	AccountNumberTypeAccountNumber AccountNumberType = "account_number"
)

type CreateAnAccountNumberParameters struct {
	// The Account the Account Number should belong to.
	AccountID *string `pjson:"account_id"`
	// The name you choose for the Account Number.
	Name       *string                `pjson:"name"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// CreateAnAccountNumberParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *CreateAnAccountNumberParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes CreateAnAccountNumberParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *CreateAnAccountNumberParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The Account the Account Number should belong to.
func (r *CreateAnAccountNumberParameters) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

// The name you choose for the Account Number.
func (r *CreateAnAccountNumberParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

func (r CreateAnAccountNumberParameters) String() (result string) {
	return fmt.Sprintf("&CreateAnAccountNumberParameters{AccountID:%s Name:%s}", core.FmtP(r.AccountID), core.FmtP(r.Name))
}

type UpdateAnAccountNumberParameters struct {
	// The name you choose for the Account Number.
	Name *string `pjson:"name"`
	// This indicates if transfers can be made to the Account Number.
	Status     *UpdateAnAccountNumberParametersStatus `pjson:"status"`
	jsonFields map[string]interface{}                 `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into
// UpdateAnAccountNumberParameters using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *UpdateAnAccountNumberParameters) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes UpdateAnAccountNumberParameters into an array of bytes
// using the gjson library. Members of the `jsonFields` field are serialized into
// the top-level, and will overwrite known members of the same name.
func (r *UpdateAnAccountNumberParameters) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// The name you choose for the Account Number.
func (r *UpdateAnAccountNumberParameters) GetName() (Name string) {
	if r != nil && r.Name != nil {
		Name = *r.Name
	}
	return
}

// This indicates if transfers can be made to the Account Number.
func (r *UpdateAnAccountNumberParameters) GetStatus() (Status UpdateAnAccountNumberParametersStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

func (r UpdateAnAccountNumberParameters) String() (result string) {
	return fmt.Sprintf("&UpdateAnAccountNumberParameters{Name:%s Status:%s}", core.FmtP(r.Name), core.FmtP(r.Status))
}

type UpdateAnAccountNumberParametersStatus string

const (
	UpdateAnAccountNumberParametersStatusActive   UpdateAnAccountNumberParametersStatus = "active"
	UpdateAnAccountNumberParametersStatusDisabled UpdateAnAccountNumberParametersStatus = "disabled"
	UpdateAnAccountNumberParametersStatusCanceled UpdateAnAccountNumberParametersStatus = "canceled"
)

type AccountNumberListParams struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `query:"limit"`
	// The status to retrieve Account Numbers for.
	Status *AccountNumbersListParamsStatus `query:"status"`
	// Filter Account Numbers to those belonging to the specified Account.
	AccountID  *string                `query:"account_id"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountNumberListParams using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumberListParams) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountNumberListParams into an array of bytes using the
// gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *AccountNumberListParams) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountNumberListParams into a url.Values of the query
// parameters associated with this value
func (r *AccountNumberListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// Return the page of entries after this one.
func (r *AccountNumberListParams) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountNumberListParams) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// The status to retrieve Account Numbers for.
func (r *AccountNumberListParams) GetStatus() (Status AccountNumbersListParamsStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// Filter Account Numbers to those belonging to the specified Account.
func (r *AccountNumberListParams) GetAccountID() (AccountID string) {
	if r != nil && r.AccountID != nil {
		AccountID = *r.AccountID
	}
	return
}

func (r AccountNumberListParams) String() (result string) {
	return fmt.Sprintf("&AccountNumberListParams{Cursor:%s Limit:%s Status:%s AccountID:%s}", core.FmtP(r.Cursor), core.FmtP(r.Limit), core.FmtP(r.Status), core.FmtP(r.AccountID))
}

type AccountNumbersListParamsStatus string

const (
	AccountNumbersListParamsStatusActive   AccountNumbersListParamsStatus = "active"
	AccountNumbersListParamsStatusDisabled AccountNumbersListParamsStatus = "disabled"
	AccountNumbersListParamsStatusCanceled AccountNumbersListParamsStatus = "canceled"
)

type AccountNumberList struct {
	// The contents of the list.
	Data *[]AccountNumber `pjson:"data"`
	// A pointer to a place in the list.
	NextCursor *string                `pjson:"next_cursor"`
	jsonFields map[string]interface{} `pjson:"-,extras"`
}

// UnmarshalJSON deserializes the provided bytes into AccountNumberList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *AccountNumberList) UnmarshalJSON(data []byte) (err error) {
	return pjson.Unmarshal(data, r)
}

// MarshalJSON serializes AccountNumberList into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *AccountNumberList) MarshalJSON() (data []byte, err error) {
	return pjson.Marshal(r)
}

// URLQuery serializes AccountNumberList into a url.Values of the query parameters
// associated with this value
func (r *AccountNumberList) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

// The contents of the list.
func (r *AccountNumberList) GetData() (Data []AccountNumber) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *AccountNumberList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r AccountNumberList) String() (result string) {
	return fmt.Sprintf("&AccountNumberList{Data:%s NextCursor:%s}", core.Fmt(r.Data), core.FmtP(r.NextCursor))
}

type AccountNumbersPage struct {
	*pagination.Page[AccountNumber]
}

func (r *AccountNumbersPage) AccountNumber() *AccountNumber {
	return r.Current()
}

func (r *AccountNumbersPage) NextPage() (*AccountNumbersPage, error) {
	if page, err := r.Page.NextPage(); err != nil {
		return nil, err
	} else {
		return &AccountNumbersPage{page}, nil
	}
}
