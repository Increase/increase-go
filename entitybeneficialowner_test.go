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
			CompanyTitle: increase.F("CEO"),
			Individual: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividual{
				Address: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualAddress{
					City:  increase.F("New York"),
					Line1: increase.F("33 Liberty Street"),
					Line2: increase.F("x"),
					State: increase.F("NY"),
					Zip:   increase.F("10045"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				DateOfBirth:        increase.F(time.Now()),
				Identification: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentification{
					DriversLicense: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationDriversLicense{
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						State:          increase.F("x"),
					}),
					Method: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
					Number: increase.F("078051120"),
					Other: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationOther{
						BackFileID:     increase.F("string"),
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
					}),
					Passport: increase.F(increase.EntityBeneficialOwnerNewParamsBeneficialOwnerIndividualIdentificationPassport{
						Country:        increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
					}),
				}),
				Name: increase.F("Ian Crease"),
			}),
			Prongs: increase.F([]increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProng{increase.EntityBeneficialOwnerNewParamsBeneficialOwnerProngControl}),
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
			City:  increase.F("New York"),
			Line1: increase.F("33 Liberty Street"),
			Line2: increase.F("Unit 2"),
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
