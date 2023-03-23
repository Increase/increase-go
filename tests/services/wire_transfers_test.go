package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/fields"
	"github.com/increase/increase-go/options"
	"github.com/increase/increase-go/requests"
)

func TestWireTransfersNewInboundWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.WireTransfers.NewInbound(context.TODO(), &requests.SimulateAWireTransferToYourAccountParameters{AccountNumberID: fields.F("account_number_v18nkfqm6afpsrvy82b2"), Amount: fields.F(int64(1000)), BeneficiaryAddressLine1: fields.F("x"), BeneficiaryAddressLine2: fields.F("x"), BeneficiaryAddressLine3: fields.F("x"), BeneficiaryName: fields.F("x"), BeneficiaryReference: fields.F("x"), OriginatorAddressLine1: fields.F("x"), OriginatorAddressLine2: fields.F("x"), OriginatorAddressLine3: fields.F("x"), OriginatorName: fields.F("x"), OriginatorToBeneficiaryInformationLine1: fields.F("x"), OriginatorToBeneficiaryInformationLine2: fields.F("x"), OriginatorToBeneficiaryInformationLine3: fields.F("x"), OriginatorToBeneficiaryInformationLine4: fields.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
