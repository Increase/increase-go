package responses

import (
	pjson "github.com/increase/increase-go/core/json"
)

type BookkeepingAccount struct {
	// The account identifier.
	ID string `json:"id,required"`
	// The compliance category of the account.
	ComplianceCategory BookkeepingAccountComplianceCategory `json:"compliance_category,required,nullable"`
	// The API Account associated with this bookkeeping account.
	AccountID string `json:"account_id,required,nullable"`
	// The Entity associated with this bookkeeping account.
	EntityID string `json:"entity_id,required,nullable"`
	// The name you choose for the account.
	Name string `json:"name,required"`
	// A constant representing the object's type. For this resource it will always be
	// `bookkeeping_account`.
	Type BookkeepingAccountType `json:"type,required"`
	JSON BookkeepingAccountJSON
}

type BookkeepingAccountJSON struct {
	ID                 pjson.Metadata
	ComplianceCategory pjson.Metadata
	AccountID          pjson.Metadata
	EntityID           pjson.Metadata
	Name               pjson.Metadata
	Type               pjson.Metadata
	Raw                []byte
	Extras             map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into BookkeepingAccount using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *BookkeepingAccount) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type BookkeepingAccountComplianceCategory string

const (
	BookkeepingAccountComplianceCategoryCommingledCash  BookkeepingAccountComplianceCategory = "commingled_cash"
	BookkeepingAccountComplianceCategoryCustomerBalance BookkeepingAccountComplianceCategory = "customer_balance"
)

type BookkeepingAccountType string

const (
	BookkeepingAccountTypeBookkeepingAccount BookkeepingAccountType = "bookkeeping_account"
)

type BookkeepingAccountListResponse struct {
	// The contents of the list.
	Data []BookkeepingAccount `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       BookkeepingAccountListResponseJSON
}

type BookkeepingAccountListResponseJSON struct {
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into
// BookkeepingAccountListResponse using the internal pjson library. Unrecognized
// fields are stored in the `jsonFields` property.
func (r *BookkeepingAccountListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
