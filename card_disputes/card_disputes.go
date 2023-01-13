package card_disputes

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type CardDisputeService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewCardDisputeService(requester core.Requester) (r *CardDisputeService) {
	r = &CardDisputeService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

type PreloadedCardDisputeService struct {
	CardDisputes *CardDisputeService
}

func (r *PreloadedCardDisputeService) Init(service *CardDisputeService) {
	r.CardDisputes = service
}

func NewPreloadedCardDisputeService(service *CardDisputeService) (r *PreloadedCardDisputeService) {
	r = &PreloadedCardDisputeService{}
	r.Init(service)
	return
}

//
type CardDispute struct {
	// The Card Dispute identifier.
	Id *string `json:"id"`
	// Why you disputed the Transaction in question.
	Explanation *string `json:"explanation"`
	// The results of the Dispute investigation.
	Status *CardDisputeStatus `json:"status"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was created.
	CreatedAt *string `json:"created_at"`
	// The identifier of the Transaction that was disputed.
	DisputedTransactionId *string `json:"disputed_transaction_id"`
	// If the Card Dispute's status is `accepted`, this will contain details of the
	// successful dispute.
	Acceptance *CardDisputeAcceptance `json:"acceptance"`
	// If the Card Dispute's status is `rejected`, this will contain details of the
	// unsuccessful dispute.
	Rejection *CardDisputeRejection `json:"rejection"`
	// A constant representing the object's type. For this resource it will always be
	// `card_dispute`.
	Type *CardDisputeType `json:"type"`
}

// The Card Dispute identifier.
func (r *CardDispute) GetId() (Id string) {
	if r != nil && r.Id != nil {
		Id = *r.Id
	}
	return
}

// Why you disputed the Transaction in question.
func (r *CardDispute) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The results of the Dispute investigation.
func (r *CardDispute) GetStatus() (Status CardDisputeStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was created.
func (r *CardDispute) GetCreatedAt() (CreatedAt string) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

// The identifier of the Transaction that was disputed.
func (r *CardDispute) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

// If the Card Dispute's status is `accepted`, this will contain details of the
// successful dispute.
func (r *CardDispute) GetAcceptance() (Acceptance CardDisputeAcceptance) {
	if r != nil && r.Acceptance != nil {
		Acceptance = *r.Acceptance
	}
	return
}

// If the Card Dispute's status is `rejected`, this will contain details of the
// unsuccessful dispute.
func (r *CardDispute) GetRejection() (Rejection CardDisputeRejection) {
	if r != nil && r.Rejection != nil {
		Rejection = *r.Rejection
	}
	return
}

// A constant representing the object's type. For this resource it will always be
// `card_dispute`.
func (r *CardDispute) GetType() (Type CardDisputeType) {
	if r != nil && r.Type != nil {
		Type = *r.Type
	}
	return
}

type CardDisputeStatus string

const (
	CardDisputeStatusPendingReviewing CardDisputeStatus = "pending_reviewing"
	CardDisputeStatusAccepted         CardDisputeStatus = "accepted"
	CardDisputeStatusRejected         CardDisputeStatus = "rejected"
)

//
type CardDisputeAcceptance struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was accepted.
	AcceptedAt *string `json:"accepted_at"`
	// The identifier of the Card Dispute that was accepted.
	CardDisputeId *string `json:"card_dispute_id"`
	// The identifier of the Transaction that was created to return the disputed funds
	// to your account.
	TransactionId *string `json:"transaction_id"`
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was accepted.
func (r *CardDisputeAcceptance) GetAcceptedAt() (AcceptedAt string) {
	if r != nil && r.AcceptedAt != nil {
		AcceptedAt = *r.AcceptedAt
	}
	return
}

// The identifier of the Card Dispute that was accepted.
func (r *CardDisputeAcceptance) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

// The identifier of the Transaction that was created to return the disputed funds
// to your account.
func (r *CardDisputeAcceptance) GetTransactionId() (TransactionId string) {
	if r != nil && r.TransactionId != nil {
		TransactionId = *r.TransactionId
	}
	return
}

//
type CardDisputeRejection struct {
	// Why the Card Dispute was rejected.
	Explanation *string `json:"explanation"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
	// the Card Dispute was rejected.
	RejectedAt *string `json:"rejected_at"`
	// The identifier of the Card Dispute that was rejected.
	CardDisputeId *string `json:"card_dispute_id"`
}

// Why the Card Dispute was rejected.
func (r *CardDisputeRejection) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date and time at which
// the Card Dispute was rejected.
func (r *CardDisputeRejection) GetRejectedAt() (RejectedAt string) {
	if r != nil && r.RejectedAt != nil {
		RejectedAt = *r.RejectedAt
	}
	return
}

// The identifier of the Card Dispute that was rejected.
func (r *CardDisputeRejection) GetCardDisputeId() (CardDisputeId string) {
	if r != nil && r.CardDisputeId != nil {
		CardDisputeId = *r.CardDisputeId
	}
	return
}

type CardDisputeType string

const (
	CardDisputeTypeCardDispute CardDisputeType = "card_dispute"
)

type CreateACardDisputeParameters struct {
	// The Transaction you wish to dispute. This Transaction must have a `source_type`
	// of `card_settlement`.
	DisputedTransactionId *string `json:"disputed_transaction_id"`
	// Why you are disputing this Transaction.
	Explanation *string `json:"explanation"`
}

// The Transaction you wish to dispute. This Transaction must have a `source_type`
// of `card_settlement`.
func (r *CreateACardDisputeParameters) GetDisputedTransactionId() (DisputedTransactionId string) {
	if r != nil && r.DisputedTransactionId != nil {
		DisputedTransactionId = *r.DisputedTransactionId
	}
	return
}

// Why you are disputing this Transaction.
func (r *CreateACardDisputeParameters) GetExplanation() (Explanation string) {
	if r != nil && r.Explanation != nil {
		Explanation = *r.Explanation
	}
	return
}

type ListCardDisputesQuery struct {
	// Return the page of entries after this one.
	Cursor *string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit     *int                            `query:"limit"`
	CreatedAt *ListCardDisputesQueryCreatedAt `query:"created_at"`
	Status    *ListCardDisputesQueryStatus    `query:"status"`
}

// Return the page of entries after this one.
func (r *ListCardDisputesQuery) GetCursor() (Cursor string) {
	if r != nil && r.Cursor != nil {
		Cursor = *r.Cursor
	}
	return
}

// Limit the size of the list that is returned. The default (and maximum) is 100
// objects.
func (r *ListCardDisputesQuery) GetLimit() (Limit int) {
	if r != nil && r.Limit != nil {
		Limit = *r.Limit
	}
	return
}

func (r *ListCardDisputesQuery) GetCreatedAt() (CreatedAt ListCardDisputesQueryCreatedAt) {
	if r != nil && r.CreatedAt != nil {
		CreatedAt = *r.CreatedAt
	}
	return
}

func (r *ListCardDisputesQuery) GetStatus() (Status ListCardDisputesQueryStatus) {
	if r != nil && r.Status != nil {
		Status = *r.Status
	}
	return
}

type ListCardDisputesQueryCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After *string `json:"after,omitempty"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before *string `json:"before,omitempty"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter *string `json:"on_or_after,omitempty"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore *string `json:"on_or_before,omitempty"`
}

// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCardDisputesQueryCreatedAt) GetAfter() (After string) {
	if r != nil && r.After != nil {
		After = *r.After
	}
	return
}

// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
// timestamp.
func (r *ListCardDisputesQueryCreatedAt) GetBefore() (Before string) {
	if r != nil && r.Before != nil {
		Before = *r.Before
	}
	return
}

// Return results on or after this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCardDisputesQueryCreatedAt) GetOnOrAfter() (OnOrAfter string) {
	if r != nil && r.OnOrAfter != nil {
		OnOrAfter = *r.OnOrAfter
	}
	return
}

// Return results on or before this
// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
func (r *ListCardDisputesQueryCreatedAt) GetOnOrBefore() (OnOrBefore string) {
	if r != nil && r.OnOrBefore != nil {
		OnOrBefore = *r.OnOrBefore
	}
	return
}

type ListCardDisputesQueryStatus struct {
	// Return results whose value is in the provided list. For GET requests, this
	// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
	In *[]ListCardDisputesQueryStatusIn `json:"in,omitempty"`
}

// Return results whose value is in the provided list. For GET requests, this
// should be encoded as a comma-delimited string, such as `?in=one,two,three`.
func (r *ListCardDisputesQueryStatus) GetIn() (In []ListCardDisputesQueryStatusIn) {
	if r != nil && r.In != nil {
		In = *r.In
	}
	return
}

type ListCardDisputesQueryStatusIn string

const (
	ListCardDisputesQueryStatusInPendingReviewing ListCardDisputesQueryStatusIn = "pending_reviewing"
	ListCardDisputesQueryStatusInAccepted         ListCardDisputesQueryStatusIn = "accepted"
	ListCardDisputesQueryStatusInRejected         ListCardDisputesQueryStatusIn = "rejected"
)

//
type CardDisputeList struct {
	// The contents of the list.
	Data *[]CardDispute `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// The contents of the list.
func (r *CardDisputeList) GetData() (Data []CardDispute) {
	if r != nil && r.Data != nil {
		Data = *r.Data
	}
	return
}

// A pointer to a place in the list.
func (r *CardDisputeList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *CardDisputeService) Create(ctx context.Context, body *CreateACardDisputeParameters, opts ...*core.RequestOpts) (res *CardDispute, err error) {
	err = r.post(
		ctx,
		"/card_disputes",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *PreloadedCardDisputeService) Create(ctx context.Context, body *CreateACardDisputeParameters, opts ...*core.RequestOpts) (res *CardDispute, err error) {
	err = r.CardDisputes.post(
		ctx,
		"/card_disputes",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)
	return
}

func (r *CardDisputeService) Retrieve(ctx context.Context, card_dispute_id string, opts ...*core.RequestOpts) (res *CardDispute, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/card_disputes/%s", card_dispute_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *PreloadedCardDisputeService) Retrieve(ctx context.Context, card_dispute_id string, opts ...*core.RequestOpts) (res *CardDispute, err error) {
	err = r.CardDisputes.get(
		ctx,
		fmt.Sprintf("/card_disputes/%s", card_dispute_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

type CardDisputesPage struct {
	*pagination.Page[CardDispute]
}

func (r *CardDisputesPage) CardDispute() *CardDispute {
	return r.Current()
}

func (r *CardDisputesPage) GetNextPage() (*CardDisputesPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &CardDisputesPage{page}, nil
	}
}

func (r *CardDisputeService) List(ctx context.Context, query *ListCardDisputesQuery, opts ...*core.RequestOpts) (res *CardDisputesPage, err error) {
	page := &CardDisputesPage{
		Page: &pagination.Page[CardDispute]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_disputes",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}

func (r *PreloadedCardDisputeService) List(ctx context.Context, query *ListCardDisputesQuery, opts ...*core.RequestOpts) (res *CardDisputesPage, err error) {
	page := &CardDisputesPage{
		Page: &pagination.Page[CardDispute]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/card_disputes",
			},
			Requester: r.CardDisputes.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
