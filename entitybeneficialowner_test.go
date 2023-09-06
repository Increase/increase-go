// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestEntityBeneficialOwnerNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Entities.BeneficialOwners.New(context.TODO(), increase.EntityBeneficialOwnerNewParams{
		BeneficialOwner: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwner{
			Individual: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual{
				Name:        increase.F("x"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress{
					Line1: increase.F("x"),
					Line2: increase.F("x"),
					City:  increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification{
					Method: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					Passport: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport{
						FileID:         increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						Country:        increase.F("x"),
					}),
					DriversLicense: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense{
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
					}),
				}),
			}),
			CompanyTitle: increase.F("x"),
			Prongs:       increase.F([]increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProng{increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProngOwnership, increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl}),
		}),
		EntityID: increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerArchive(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Entities.BeneficialOwners.Archive(context.TODO(), increase.EntityBeneficialOwnerArchiveParams{
		BeneficialOwnerID: increase.F("string"),
		EntityID:          increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityBeneficialOwnerUpdateAddressWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	_, err := client.Entities.BeneficialOwners.UpdateAddress(context.TODO(), increase.EntityBeneficialOwnerUpdateAddressParams{
		Address: increase.F(increase.EntityBeneficialOwnerUpdateAddressParamsAddress{
			Line1: increase.F("x"),
			Line2: increase.F("x"),
			City:  increase.F("x"),
			State: increase.F("x"),
			Zip:   increase.F("x"),
		}),
		BeneficialOwnerID: increase.F("string"),
		EntityID:          increase.F("string"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
