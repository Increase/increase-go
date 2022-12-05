package transactions

type TransactionSourceEmpyrealCashDeposit struct {
	// The amount in the minor unit of the transaction's currency. For dollars, for
	// example, this is cents.
	Amount *int64 `json:"amount,omitempty"`

	BagId *string `json:"bag_id,omitempty"`

	DepositDate *string `json:"deposit_date,omitempty"`
}

// The amount in the minor unit of the transaction's currency. For dollars, for
// example, this is cents.
func (r *TransactionSourceEmpyrealCashDeposit) GetAmount() (Amount int64) {
	if r != nil && r.Amount != nil {
		Amount = *r.Amount
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetBagId() (BagId string) {
	if r != nil && r.BagId != nil {
		BagId = *r.BagId
	}
	return
}

func (r *TransactionSourceEmpyrealCashDeposit) GetDepositDate() (DepositDate string) {
	if r != nil && r.DepositDate != nil {
		DepositDate = *r.DepositDate
	}
	return
}
