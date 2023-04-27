package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
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
	ID                 apijson.Metadata
	CreatedAt          apijson.Metadata
	Description        apijson.Metadata
	Status             apijson.Metadata
	RoutingNumber      apijson.Metadata
	AccountNumber      apijson.Metadata
	Funding            apijson.Metadata
	VerificationStatus apijson.Metadata
	Type               apijson.Metadata
	raw                string
	Extras             map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccount using the
// internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *ExternalAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type ExternalAccountListResponse struct {
	// The contents of the list.
	Data []ExternalAccount `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       ExternalAccountListResponseJSON
}

type ExternalAccountListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into ExternalAccountListResponse
// using the internal json library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *ExternalAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
