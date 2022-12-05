package account_statements

type AccountStatement struct {
	// The Account Statement identifier.
	Id *string `json:"id,omitempty"`

	// The identifier for the Account this Account Statement belongs to.
	AccountId *string `json:"account_id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
	// Statement was created.
	CreatedAt *string `json:"created_at,omitempty"`

	// The identifier of the File containing a PDF of the statement.
	FileId *string `json:"file_id,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
	// start of the period the Account Statement covers.
	StatementPeriodStart *string `json:"statement_period_start,omitempty"`

	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
	// of the period the Account Statement covers.
	StatementPeriodEnd *string `json:"statement_period_end,omitempty"`

	// The Account's balance at the start of its statement period.
	StartingBalance *int64 `json:"starting_balance,omitempty"`

	// The Account's balance at the start of its statement period.
	EndingBalance *int64 `json:"ending_balance,omitempty"`

	// A constant representing the object's type. For this resource it will always be
	// `account_statement`.
	Type *string `json:"type,omitempty"`
}

// The Account Statement identifier.
func (r *AccountStatement) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// The identifier for the Account this Account Statement belongs to.
func (r *AccountStatement) GetAccountId() (AccountId string) {
	if r != nil && r.AccountId != nil {
		AccountId = *r.AccountId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the Account
// Statement was created.
func (r *AccountStatement) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the File containing a PDF of the statement.
func (r *AccountStatement) GetFileId() (FileId string) {
	if r != nil && r.FileId != nil {
		FileId = *r.FileId
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the
// start of the period the Account Statement covers.
func (r *AccountStatement) GetStatementPeriodStart() (StatementPeriodStart string) {
	if r != nil && r.StatementPeriodStart != nil {
		StatementPeriodStart = *r.StatementPeriodStart
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time representing the end
// of the period the Account Statement covers.
func (r *AccountStatement) GetStatementPeriodEnd() (StatementPeriodEnd string) {
	if r != nil && r.StatementPeriodEnd != nil {
		StatementPeriodEnd = *r.StatementPeriodEnd
	}
	return
}

// The Account's balance at the start of its statement period.
func (r *AccountStatement) GetStartingBalance() (StartingBalance int64) {
	if r != nil && r.StartingBalance != nil {
		StartingBalance = *r.StartingBalance
	}
	return
}

// The Account's balance at the start of its statement period.
func (r *AccountStatement) GetEndingBalance() (EndingBalance int64) {
	if r != nil && r.EndingBalance != nil {
		EndingBalance = *r.EndingBalance
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `account_statement`.
func (r *AccountStatement) GetType() (Type string) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}
