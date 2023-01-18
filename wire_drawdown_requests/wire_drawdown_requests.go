package wire_drawdown_requests

import "context"
import "increase/core"
import "fmt"
import "increase/pagination"

type WireDrawdownRequestService struct {
	Requester core.Requester
	get       func(context.Context, string, *core.CoreRequest, interface{}) error
	post      func(context.Context, string, *core.CoreRequest, interface{}) error
	patch     func(context.Context, string, *core.CoreRequest, interface{}) error
	put       func(context.Context, string, *core.CoreRequest, interface{}) error
	delete    func(context.Context, string, *core.CoreRequest, interface{}) error
}

func NewWireDrawdownRequestService(requester core.Requester) (r *WireDrawdownRequestService) {
	r = &WireDrawdownRequestService{}
	r.Requester = requester
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

//
type WireDrawdownRequest struct {
	// A constant representing the object's type. For this resource it will always be
	// `wire_drawdown_request`.
	Type WireDrawdownRequestType `json:"type"`
	// The Wire drawdown request identifier.
	ID string `json:"id"`
	// The Account Number to which the recipient of this request is being requested to
	// send funds.
	AccountNumberID string `json:"account_number_id"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber string `json:"recipient_account_number"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber string `json:"recipient_routing_number"`
	// The amount being requested in cents.
	Amount int `json:"amount"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code for the amount being
	// requested. Will always be "USD".
	Currency string `json:"currency"`
	// The message the recipient will see as part of the drawdown request.
	MessageToRecipient string `json:"message_to_recipient"`
	// The drawdown request's recipient's name.
	RecipientName *string `json:"recipient_name"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 *string `json:"recipient_address_line1"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 *string `json:"recipient_address_line2"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 *string `json:"recipient_address_line3"`
	// After the drawdown request is submitted to Fedwire, this will contain
	// supplemental details.
	Submission *WireDrawdownRequestSubmission `json:"submission"`
	// If the recipient fulfills the drawdown request by sending funds, then this will
	// be the identifier of the corresponding Transaction.
	FulfillmentTransactionID *string `json:"fulfillment_transaction_id"`
	// The lifecycle status of the drawdown request.
	Status WireDrawdownRequestStatus `json:"status"`
}

// The drawdown request's recipient's name.
func (r *WireDrawdownRequest) GetRecipientName() (RecipientName string) {
	if r != nil && r.RecipientName != nil {
		RecipientName = *r.RecipientName
	}
	return
}

// Line 1 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine1() (RecipientAddressLine1 string) {
	if r != nil && r.RecipientAddressLine1 != nil {
		RecipientAddressLine1 = *r.RecipientAddressLine1
	}
	return
}

// Line 2 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine2() (RecipientAddressLine2 string) {
	if r != nil && r.RecipientAddressLine2 != nil {
		RecipientAddressLine2 = *r.RecipientAddressLine2
	}
	return
}

// Line 3 of the drawdown request's recipient's address.
func (r *WireDrawdownRequest) GetRecipientAddressLine3() (RecipientAddressLine3 string) {
	if r != nil && r.RecipientAddressLine3 != nil {
		RecipientAddressLine3 = *r.RecipientAddressLine3
	}
	return
}

// After the drawdown request is submitted to Fedwire, this will contain
// supplemental details.
func (r *WireDrawdownRequest) GetSubmission() (Submission WireDrawdownRequestSubmission) {
	if r != nil && r.Submission != nil {
		Submission = *r.Submission
	}
	return
}

// If the recipient fulfills the drawdown request by sending funds, then this will
// be the identifier of the corresponding Transaction.
func (r *WireDrawdownRequest) GetFulfillmentTransactionID() (FulfillmentTransactionID string) {
	if r != nil && r.FulfillmentTransactionID != nil {
		FulfillmentTransactionID = *r.FulfillmentTransactionID
	}
	return
}

type WireDrawdownRequestType string

const (
	WireDrawdownRequestTypeWireDrawdownRequest WireDrawdownRequestType = "wire_drawdown_request"
)

//
type WireDrawdownRequestSubmission struct {
	// The input message accountability data (IMAD) uniquely identifying the submission
	// with Fedwire.
	InputMessageAccountabilityData string `json:"input_message_accountability_data"`
}

type WireDrawdownRequestStatus string

const (
	WireDrawdownRequestStatusPendingSubmission WireDrawdownRequestStatus = "pending_submission"
	WireDrawdownRequestStatusPendingResponse   WireDrawdownRequestStatus = "pending_response"
	WireDrawdownRequestStatusFulfilled         WireDrawdownRequestStatus = "fulfilled"
	WireDrawdownRequestStatusRefused           WireDrawdownRequestStatus = "refused"
)

type CreateAWireDrawdownRequestParameters struct {
	// The Account Number to which the recipient should send funds.
	AccountNumberID string `json:"account_number_id"`
	// The amount requested from the recipient, in cents.
	Amount int `json:"amount"`
	// A message the recipient will see as part of the request.
	MessageToRecipient string `json:"message_to_recipient"`
	// The drawdown request's recipient's account number.
	RecipientAccountNumber string `json:"recipient_account_number"`
	// The drawdown request's recipient's routing number.
	RecipientRoutingNumber string `json:"recipient_routing_number"`
	// The drawdown request's recipient's name.
	RecipientName string `json:"recipient_name,omitempty"`
	// Line 1 of the drawdown request's recipient's address.
	RecipientAddressLine1 string `json:"recipient_address_line1,omitempty"`
	// Line 2 of the drawdown request's recipient's address.
	RecipientAddressLine2 string `json:"recipient_address_line2,omitempty"`
	// Line 3 of the drawdown request's recipient's address.
	RecipientAddressLine3 string `json:"recipient_address_line3,omitempty"`
}

type ListWireDrawdownRequestsQuery struct {
	// Return the page of entries after this one.
	Cursor string `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit int `query:"limit"`
}

//
type WireDrawdownRequestList struct {
	// The contents of the list.
	Data []WireDrawdownRequest `json:"data"`
	// A pointer to a place in the list.
	NextCursor *string `json:"next_cursor"`
}

// A pointer to a place in the list.
func (r *WireDrawdownRequestList) GetNextCursor() (NextCursor string) {
	if r != nil && r.NextCursor != nil {
		NextCursor = *r.NextCursor
	}
	return
}

func (r *WireDrawdownRequestService) Create(ctx context.Context, body *CreateAWireDrawdownRequestParameters, opts ...*core.RequestOpts) (res *WireDrawdownRequest, err error) {
	err = r.post(
		ctx,
		"/wire_drawdown_requests",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Body:   body,
		},
		&res,
	)

	return
}

func (r *WireDrawdownRequestService) Retrieve(ctx context.Context, wire_drawdown_request_id string, opts ...*core.RequestOpts) (res *WireDrawdownRequest, err error) {
	err = r.get(
		ctx,
		fmt.Sprintf("/wire_drawdown_requests/%s", wire_drawdown_request_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)

	return
}

type WireDrawdownRequestsPage struct {
	*pagination.Page[WireDrawdownRequest]
}

func (r *WireDrawdownRequestsPage) WireDrawdownRequest() *WireDrawdownRequest {
	return r.Current()
}

func (r *WireDrawdownRequestsPage) GetNextPage() (*WireDrawdownRequestsPage, error) {
	if page, err := r.Page.GetNextPage(); err != nil {
		return nil, err
	} else {
		return &WireDrawdownRequestsPage{page}, nil
	}
}

func (r *WireDrawdownRequestService) List(ctx context.Context, query *ListWireDrawdownRequestsQuery, opts ...*core.RequestOpts) (res *WireDrawdownRequestsPage, err error) {
	page := &WireDrawdownRequestsPage{
		Page: &pagination.Page[WireDrawdownRequest]{
			Options: pagination.PageOptions{
				RequestParams: query,
				Path:          "/wire_drawdown_requests",
			},
			Requester: r.Requester,
			Context:   ctx,
		},
	}
	res, err = page.GetNextPage()
	return
}
