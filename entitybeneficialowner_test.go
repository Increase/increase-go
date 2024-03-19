// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/internal/testutil"
	"github.com/increase/increase-go/option"
)

func TestEntityBeneficialOwnerNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Entities.BeneficialOwners.New(context.TODO(), increase.EntityBeneficialOwnerNewParams{
		BeneficialOwner: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwner{
			Individual: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual{
				Name:        increase.F("Ian Crease"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress{
					Line1: increase.F("33 Liberty Street"),
					Line2: increase.F("x"),
					City:  increase.F("New York"),
					State: increase.F("NY"),
					Zip:   increase.F("10045"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification{
					Method: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
					Number: increase.F("078051120"),
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
			CompanyTitle: increase.F("CEO"),
			Prongs:       increase.F([]increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProng{increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl}),
		}),
		EntityID: increase.F("entity_n8y8tnk2p9339ti393yi"),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Entities.BeneficialOwners.Archive(context.TODO(), increase.EntityBeneficialOwnerArchiveParams{
		BeneficialOwnerID: increase.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          increase.F("entity_n8y8tnk2p9339ti393yi"),
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
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Entities.BeneficialOwners.UpdateAddress(context.TODO(), increase.EntityBeneficialOwnerUpdateAddressParams{
		Address: increase.F(increase.EntityBeneficialOwnerUpdateAddressParamsAddress{
			Line1: increase.F("33 Liberty Street"),
			Line2: increase.F("Unit 2"),
			City:  increase.F("New York"),
			State: increase.F("NY"),
			Zip:   increase.F("10045"),
		}),
		BeneficialOwnerID: increase.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		EntityID:          increase.F("entity_n8y8tnk2p9339ti393yi"),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
