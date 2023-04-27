package services

import (
	"context"
	"errors"
	"testing"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestSimulationInboundWireDrawdownRequestNewWithOptionalParams(t *testing.T) {
	c := increase.NewClient(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.InboundWireDrawdownRequests.New(context.TODO(), requests.SimulationInboundWireDrawdownRequestNewParams{RecipientAccountNumberID: increase.F("account_number_v18nkfqm6afpsrvy82b2"), OriginatorAccountNumber: increase.F("987654321"), OriginatorRoutingNumber: increase.F("101050001"), BeneficiaryAccountNumber: increase.F("987654321"), BeneficiaryRoutingNumber: increase.F("101050001"), Amount: increase.F(int64(10000)), Currency: increase.F("USD"), MessageToRecipient: increase.F("Invoice 29582"), OriginatorToBeneficiaryInformationLine1: increase.F("x"), OriginatorToBeneficiaryInformationLine2: increase.F("x"), OriginatorToBeneficiaryInformationLine3: increase.F("x"), OriginatorToBeneficiaryInformationLine4: increase.F("x"), OriginatorName: increase.F("Ian Crease"), OriginatorAddressLine1: increase.F("33 Liberty Street"), OriginatorAddressLine2: increase.F("New York, NY, 10045"), OriginatorAddressLine3: increase.F("x"), BeneficiaryName: increase.F("Ian Crease"), BeneficiaryAddressLine1: increase.F("33 Liberty Street"), BeneficiaryAddressLine2: increase.F("New York, NY, 10045"), BeneficiaryAddressLine3: increase.F("x")})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
