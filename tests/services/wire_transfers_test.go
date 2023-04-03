package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestWireTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.WireTransfers.NewInbound(context.TODO(), &requests.SimulateAWireTransferToYourAccountParameters{AccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: increase.F(int64(1000)), BeneficiaryAddressLine1: increase.F("x"), BeneficiaryAddressLine2: increase.F("x"), BeneficiaryAddressLine3: increase.F("x"), BeneficiaryName: increase.F("x"), BeneficiaryReference: increase.F("x"), OriginatorAddressLine1: increase.F("x"), OriginatorAddressLine2: increase.F("x"), OriginatorAddressLine3: increase.F("x"), OriginatorName: increase.F("x"), OriginatorToBeneficiaryInformationLine1: increase.F("x"), OriginatorToBeneficiaryInformationLine2: increase.F("x"), OriginatorToBeneficiaryInformationLine3: increase.F("x"), OriginatorToBeneficiaryInformationLine4: increase.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
