// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestEntityNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.New(context.TODO(), increase.EntityNewParams{
		Structure: increase.F(increase.EntityNewParamsStructureCorporation),
		Corporation: increase.F(increase.EntityNewParamsCorporation{
			Address: increase.F(increase.EntityNewParamsCorporationAddress{
				City:  increase.F("New York"),
				Line1: increase.F("33 Liberty Street"),
				State: increase.F("NY"),
				Zip:   increase.F("10045"),
				Line2: increase.F("x"),
			}),
			BeneficialOwners: increase.F([]increase.EntityNewParamsCorporationBeneficialOwner{{
				Individual: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Address: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						City:    increase.F("New York"),
						Country: increase.F("x"),
						Line1:   increase.F("33 Liberty Street"),
						Line2:   increase.F("x"),
						State:   increase.F("NY"),
						Zip:     increase.F("10045"),
					}),
					DateOfBirth: increase.F(time.Now()),
					Identification: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("078051120"),
						DriversLicense: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
							State:          increase.F("x"),
							BackFileID:     increase.F("back_file_id"),
						}),
						Other: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							FileID:         increase.F("file_id"),
							BackFileID:     increase.F("back_file_id"),
							ExpirationDate: increase.F(time.Now()),
						}),
						Passport: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							Country:        increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
						}),
					}),
					Name:               increase.F("Ian Crease"),
					ConfirmedNoUsTaxID: increase.F(true),
				}),
				Prongs:       increase.F([]increase.EntityNewParamsCorporationBeneficialOwnersProng{increase.EntityNewParamsCorporationBeneficialOwnersProngControl}),
				CompanyTitle: increase.F("CEO"),
			}}),
			Name:                               increase.F("National Phonograph Company"),
			TaxIdentifier:                      increase.F("602214076"),
			BeneficialOwnershipExemptionReason: increase.F(increase.EntityNewParamsCorporationBeneficialOwnershipExemptionReasonRegulatedFinancialInstitution),
			IncorporationState:                 increase.F("NY"),
			IndustryCode:                       increase.F("x"),
			Website:                            increase.F("https://example.com"),
		}),
		Description: increase.F("x"),
		GovernmentAuthority: increase.F(increase.EntityNewParamsGovernmentAuthority{
			Address: increase.F(increase.EntityNewParamsGovernmentAuthorityAddress{
				City:  increase.F("x"),
				Line1: increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
				Line2: increase.F("x"),
			}),
			AuthorizedPersons: increase.F([]increase.EntityNewParamsGovernmentAuthorityAuthorizedPerson{{
				Name: increase.F("x"),
			}}),
			Category:      increase.F(increase.EntityNewParamsGovernmentAuthorityCategoryMunicipality),
			Name:          increase.F("x"),
			TaxIdentifier: increase.F("x"),
			Website:       increase.F("website"),
		}),
		Joint: increase.F(increase.EntityNewParamsJoint{
			Individuals: increase.F([]increase.EntityNewParamsJointIndividual{{
				Address: increase.F(increase.EntityNewParamsJointIndividualsAddress{
					City:  increase.F("x"),
					Line1: increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
					Line2: increase.F("x"),
				}),
				DateOfBirth: increase.F(time.Now()),
				Identification: increase.F(increase.EntityNewParamsJointIndividualsIdentification{
					Method: increase.F(increase.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					DriversLicense: increase.F(increase.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("file_id"),
						State:          increase.F("x"),
						BackFileID:     increase.F("back_file_id"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						FileID:         increase.F("file_id"),
						BackFileID:     increase.F("back_file_id"),
						ExpirationDate: increase.F(time.Now()),
					}),
					Passport: increase.F(increase.EntityNewParamsJointIndividualsIdentificationPassport{
						Country:        increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("file_id"),
					}),
				}),
				Name:               increase.F("x"),
				ConfirmedNoUsTaxID: increase.F(true),
			}}),
		}),
		NaturalPerson: increase.F(increase.EntityNewParamsNaturalPerson{
			Address: increase.F(increase.EntityNewParamsNaturalPersonAddress{
				City:  increase.F("x"),
				Line1: increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
				Line2: increase.F("x"),
			}),
			DateOfBirth: increase.F(time.Now()),
			Identification: increase.F(increase.EntityNewParamsNaturalPersonIdentification{
				Method: increase.F(increase.EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber),
				Number: increase.F("xxxx"),
				DriversLicense: increase.F(increase.EntityNewParamsNaturalPersonIdentificationDriversLicense{
					ExpirationDate: increase.F(time.Now()),
					FileID:         increase.F("file_id"),
					State:          increase.F("x"),
					BackFileID:     increase.F("back_file_id"),
				}),
				Other: increase.F(increase.EntityNewParamsNaturalPersonIdentificationOther{
					Country:        increase.F("x"),
					Description:    increase.F("x"),
					FileID:         increase.F("file_id"),
					BackFileID:     increase.F("back_file_id"),
					ExpirationDate: increase.F(time.Now()),
				}),
				Passport: increase.F(increase.EntityNewParamsNaturalPersonIdentificationPassport{
					Country:        increase.F("x"),
					ExpirationDate: increase.F(time.Now()),
					FileID:         increase.F("file_id"),
				}),
			}),
			Name:               increase.F("x"),
			ConfirmedNoUsTaxID: increase.F(true),
		}),
		RiskRating: increase.F(increase.EntityNewParamsRiskRating{
			RatedAt: increase.F(time.Now()),
			Rating:  increase.F(increase.EntityNewParamsRiskRatingRatingLow),
		}),
		SupplementalDocuments: increase.F([]increase.EntityNewParamsSupplementalDocument{{
			FileID: increase.F("file_makxrc67oh9l6sg7w9yc"),
		}}),
		ThirdPartyVerification: increase.F(increase.EntityNewParamsThirdPartyVerification{
			Reference: increase.F("x"),
			Vendor:    increase.F(increase.EntityNewParamsThirdPartyVerificationVendorAlloy),
		}),
		Trust: increase.F(increase.EntityNewParamsTrust{
			Address: increase.F(increase.EntityNewParamsTrustAddress{
				City:  increase.F("x"),
				Line1: increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
				Line2: increase.F("x"),
			}),
			Category: increase.F(increase.EntityNewParamsTrustCategoryRevocable),
			Name:     increase.F("x"),
			Trustees: increase.F([]increase.EntityNewParamsTrustTrustee{{
				Structure: increase.F(increase.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: increase.F(increase.EntityNewParamsTrustTrusteesIndividual{
					Address: increase.F(increase.EntityNewParamsTrustTrusteesIndividualAddress{
						City:  increase.F("x"),
						Line1: increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
						Line2: increase.F("x"),
					}),
					DateOfBirth: increase.F(time.Now()),
					Identification: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						DriversLicense: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
							State:          increase.F("x"),
							BackFileID:     increase.F("back_file_id"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							FileID:         increase.F("file_id"),
							BackFileID:     increase.F("back_file_id"),
							ExpirationDate: increase.F(time.Now()),
						}),
						Passport: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							Country:        increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
						}),
					}),
					Name:               increase.F("x"),
					ConfirmedNoUsTaxID: increase.F(true),
				}),
			}}),
			FormationDocumentFileID: increase.F("formation_document_file_id"),
			FormationState:          increase.F("x"),
			Grantor: increase.F(increase.EntityNewParamsTrustGrantor{
				Address: increase.F(increase.EntityNewParamsTrustGrantorAddress{
					City:  increase.F("x"),
					Line1: increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
					Line2: increase.F("x"),
				}),
				DateOfBirth: increase.F(time.Now()),
				Identification: increase.F(increase.EntityNewParamsTrustGrantorIdentification{
					Method: increase.F(increase.EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					DriversLicense: increase.F(increase.EntityNewParamsTrustGrantorIdentificationDriversLicense{
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("file_id"),
						State:          increase.F("x"),
						BackFileID:     increase.F("back_file_id"),
					}),
					Other: increase.F(increase.EntityNewParamsTrustGrantorIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						FileID:         increase.F("file_id"),
						BackFileID:     increase.F("back_file_id"),
						ExpirationDate: increase.F(time.Now()),
					}),
					Passport: increase.F(increase.EntityNewParamsTrustGrantorIdentificationPassport{
						Country:        increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("file_id"),
					}),
				}),
				Name:               increase.F("x"),
				ConfirmedNoUsTaxID: increase.F(true),
			}),
			TaxIdentifier: increase.F("x"),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGet(t *testing.T) {
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
	_, err := client.Entities.Get(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Update(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityUpdateParams{
			RiskRating: increase.F(increase.EntityUpdateParamsRiskRating{
				RatedAt: increase.F(time.Now()),
				Rating:  increase.F(increase.EntityUpdateParamsRiskRatingRatingLow),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityListWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.List(context.TODO(), increase.EntityListParams{
		CreatedAt: increase.F(increase.EntityListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.EntityListParamsStatus{
			In: increase.F([]increase.EntityListParamsStatusIn{increase.EntityListParamsStatusInActive}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityArchive(t *testing.T) {
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
	_, err := client.Entities.Archive(context.TODO(), "entity_n8y8tnk2p9339ti393yi")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityArchiveBeneficialOwner(t *testing.T) {
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
	_, err := client.Entities.ArchiveBeneficialOwner(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityArchiveBeneficialOwnerParams{
			BeneficialOwnerID: increase.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityConfirmWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.Confirm(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityConfirmParams{
			ConfirmedAt: increase.F(time.Now()),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityNewBeneficialOwnerWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.NewBeneficialOwner(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityNewBeneficialOwnerParams{
			BeneficialOwner: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwner{
				Individual: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividual{
					Address: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualAddress{
						City:    increase.F("New York"),
						Country: increase.F("US"),
						Line1:   increase.F("33 Liberty Street"),
						Line2:   increase.F("x"),
						State:   increase.F("NY"),
						Zip:     increase.F("10045"),
					}),
					DateOfBirth: increase.F(time.Now()),
					Identification: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentification{
						Method: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("078051120"),
						DriversLicense: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationDriversLicense{
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
							State:          increase.F("x"),
							BackFileID:     increase.F("back_file_id"),
						}),
						Other: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							FileID:         increase.F("file_id"),
							BackFileID:     increase.F("back_file_id"),
							ExpirationDate: increase.F(time.Now()),
						}),
						Passport: increase.F(increase.EntityNewBeneficialOwnerParamsBeneficialOwnerIndividualIdentificationPassport{
							Country:        increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("file_id"),
						}),
					}),
					Name:               increase.F("Ian Crease"),
					ConfirmedNoUsTaxID: increase.F(true),
				}),
				Prongs:       increase.F([]increase.EntityNewBeneficialOwnerParamsBeneficialOwnerProng{increase.EntityNewBeneficialOwnerParamsBeneficialOwnerProngControl}),
				CompanyTitle: increase.F("CEO"),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityUpdateAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.UpdateAddress(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityUpdateAddressParams{
			Address: increase.F(increase.EntityUpdateAddressParamsAddress{
				City:  increase.F("New York"),
				Line1: increase.F("33 Liberty Street"),
				State: increase.F("NY"),
				Zip:   increase.F("10045"),
				Line2: increase.F("Unit 2"),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityUpdateBeneficialOwnerAddressWithOptionalParams(t *testing.T) {
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
	_, err := client.Entities.UpdateBeneficialOwnerAddress(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityUpdateBeneficialOwnerAddressParams{
			Address: increase.F(increase.EntityUpdateBeneficialOwnerAddressParamsAddress{
				City:    increase.F("New York"),
				Country: increase.F("US"),
				Line1:   increase.F("33 Liberty Street"),
				Line2:   increase.F("Unit 2"),
				State:   increase.F("NY"),
				Zip:     increase.F("10045"),
			}),
			BeneficialOwnerID: increase.F("entity_setup_beneficial_owner_submission_vgkyk7dj5eb4sfhdbkx7"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityUpdateIndustryCode(t *testing.T) {
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
	_, err := client.Entities.UpdateIndustryCode(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
		increase.EntityUpdateIndustryCodeParams{
			IndustryCode: increase.F("5132"),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
