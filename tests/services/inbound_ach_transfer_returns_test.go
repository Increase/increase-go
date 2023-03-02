package services

import (
	"context"
	"increase"
	"increase/options"
	"increase/types"
	"testing"
)

func TestInboundACHTransferReturnsNew(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.New(context.TODO(), &types.CreateAnACHReturnParameters{TransactionID: increase.P("transaction_uyrp7fld2ium70oa7oi"), Reason: increase.P(types.CreateAnACHReturnParametersReasonAuthorizationRevokedByCustomer)})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestInboundACHTransferReturnsGet(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.Get(context.TODO(), "inbound_ach_transfer_return_fhcxk5huskwhmt7iz0gk")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestInboundACHTransferReturnsListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.InboundACHTransferReturns.List(context.TODO(), &types.InboundACHTransferReturnListParams{Cursor: increase.P("string"), Limit: increase.P(int64(0))})
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
