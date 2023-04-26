package requests

import (
	"net/url"

	"github.com/increase/increase-go/core/field"
	apijson "github.com/increase/increase-go/core/json"
	"github.com/increase/increase-go/core/query"
)

type BookkeepingAccountNewParams struct {
	// The account compliance category.
	ComplianceCategory field.Field[BookkeepingAccountNewParamsComplianceCategory] `json:"compliance_category"`
	// The entity, if `compliance_category` is `customer_balance`.
	EntityID field.Field[string] `json:"entity_id"`
	// The entity, if `compliance_category` is `commingled_cash`.
	AccountID field.Field[string] `json:"account_id"`
	// The name you choose for the account.
	Name field.Field[string] `json:"name,required"`
}

// MarshalJSON serializes BookkeepingAccountNewParams into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r BookkeepingAccountNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BookkeepingAccountNewParamsComplianceCategory string

const (
	BookkeepingAccountNewParamsComplianceCategoryCommingledCash  BookkeepingAccountNewParamsComplianceCategory = "commingled_cash"
	BookkeepingAccountNewParamsComplianceCategoryCustomerBalance BookkeepingAccountNewParamsComplianceCategory = "customer_balance"
)

type BookkeepingAccountListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
}

// URLQuery serializes BookkeepingAccountListParams into a url.Values of the query
// parameters associated with this value
func (r BookkeepingAccountListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
