package account_statements

type AccountStatementsListQuery struct {
	// Return the page of entries after this one.
	Cursor *string `json:"cursor,omitempty"`

	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit *int64 `json:"limit,omitempty"`

	// Filter Account Statements to those belonging to the specified Account.
	AccountId            *string                                         `json:"account_id,omitempty"`
	StatementPeriodStart *AccountStatementsStatementPeriodStartListQuery `json:"statement_period_start,omitempty"`
}

// Return the page of entries after this one.
func (r *AccountStatementsListQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *AccountStatementsListQuery) GetLimit() (Limit int64) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

// Filter Account Statements to those belonging to the specified Account.
func (r *AccountStatementsListQuery) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

func (r *AccountStatementsListQuery) GetStatementPeriodStart() (StatementPeriodStart AccountStatementsStatementPeriodStartListQuery) {
	if r != nil && r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}
