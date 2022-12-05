package digital_wallet_tokens

import "increase/core"
import "fmt"

type DigitalWalletTokens struct {
	Requester core.Requester
	get       func(string, *core.CoreRequest, interface{}) error
	post      func(string, *core.CoreRequest, interface{}) error
	patch     func(string, *core.CoreRequest, interface{}) error
	put       func(string, *core.CoreRequest, interface{}) error
	delete    func(string, *core.CoreRequest, interface{}) error
}

func NewDigitalWalletTokens(requster core.Requester) (r *DigitalWalletTokens) {
	r = &DigitalWalletTokens{}
	r.Requester = requster
	r.get = r.Requester.Get
	r.post = r.Requester.Post
	r.patch = r.Requester.Patch
	r.put = r.Requester.Put
	r.delete = r.Requester.Delete
	return
}

func (r *DigitalWalletTokens) Retrieve(digital_wallet_token_id string, opts ...*core.RequestOpts) (res DigitalWalletToken, err error) {
	err = r.get(
		fmt.Sprintf("/digital_wallet_tokens/%s", digital_wallet_token_id),
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
		},
		&res,
	)
	return
}

func (r *DigitalWalletTokens) List(query *DigitalWalletTokensListQuery, opts ...*core.RequestOpts) (res DigitalWalletTokensListResponse, err error) {
	err = r.get(
		"/digital_wallet_tokens",
		&core.CoreRequest{
			Params: core.MergeRequestOpts(opts...),
			Query:  query,
		},
		&res,
	)
	return
}
