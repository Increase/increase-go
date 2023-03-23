package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/pagination"
)

type ExternalAccount struct {
	// The External Account's identifier.
	ID string `json:"id,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the External Account was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The External Account's description for display purposes.
	Description string `json:"description,required"`
	// The External Account's status.
	Status ExternalAccountStatus `json:"status,required"`
	// The American Bankers' Association (ABA) Routing Transit Number (RTN).
	RoutingNumber string `json:"routing_number,required"`
	// The destination account number.
	AccountNumber string `json:"account_number,required"`
	// The type of the account to which the transfer will be sent.
	Funding ExternalAccountFunding `json:"funding,required"`
	// If you have verified ownership of the External Account.
	VerificationStatus ExternalAccountVerificationStatus `json:"verification_status,required"`
	// A constant representing the object's type. For this resource it will always be
	// `external_account`.
	Type ExternalAccountType `json:"type,required"`
	JSON ExternalAccountJSON
}

type ExternalAccountJSON struct {
	ID                 pjson.Metadata
	CreatedAt          pjson.Metadata
	Description        pjson.Metadata
	Status             pjson.Metadata
	RoutingNumber      pjson.Metadata
	AccountNumber      pjson.Metadata
	Funding            pjson.Metadata
	VerificationStatus pjson.Metadata
	Type               pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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

type ExternalAccountList struct {
	// The contents of the list.
	Data []ExternalAccount `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ExternalAccountListJSON
}

type ExternalAccountListJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountList using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccountList) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
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
