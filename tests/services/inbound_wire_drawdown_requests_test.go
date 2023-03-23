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

func TestInboundWireDrawdownRequestsNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(options.WithAPIKey("APIKey"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.InboundWireDrawdownRequests.New(context.TODO(), &requests.SimulateAnInboundWireDrawdownRequestBeingCreatedParameters{RecipientAccountNumberID: fields.F("account_number_v18nkfqm6afpsrvy82b2"), OriginatorAccountNumber: fields.F("987654321"), OriginatorRoutingNumber: fields.F("101050001"), BeneficiaryAccountNumber: fields.F("987654321"), BeneficiaryRoutingNumber: fields.F("101050001"), Amount: fields.F(int64(10000)), Currency: fields.F("USD"), MessageToRecipient: fields.F("Invoice 29582"), OriginatorToBeneficiaryInformationLine1: fields.F("x"), OriginatorToBeneficiaryInformationLine2: fields.F("x"), OriginatorToBeneficiaryInformationLine3: fields.F("x"), OriginatorToBeneficiaryInformationLine4: fields.F("x"), OriginatorName: fields.F("Ian Crease"), OriginatorAddressLine1: fields.F("33 Liberty Street"), OriginatorAddressLine2: fields.F("New York, NY, 10045"), OriginatorAddressLine3: fields.F("x"), BeneficiaryName: fields.F("Ian Crease"), BeneficiaryAddressLine1: fields.F("33 Liberty Street"), BeneficiaryAddressLine2: fields.F("New York, NY, 10045"), BeneficiaryAddressLine3: fields.F("x")})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
