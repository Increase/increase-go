// File generated from our OpenAPI spec by Stainless.

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
			Name:               increase.F("National Phonograph Company"),
			Website:            increase.F("https://example.com"),
			TaxIdentifier:      increase.F("602214076"),
			IncorporationState: increase.F("NY"),
			Address: increase.F(increase.EntityNewParamsCorporationAddress{
				Line1: increase.F("33 Liberty Street"),
				Line2: increase.F("x"),
				City:  increase.F("New York"),
				State: increase.F("NY"),
				Zip:   increase.F("10045"),
			}),
			BeneficialOwners: increase.F([]increase.EntityNewParamsCorporationBeneficialOwner{{
				Individual: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Name:        increase.F("Ian Crease"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						Line1: increase.F("33 Liberty Street"),
						Line2: increase.F("x"),
						City:  increase.F("New York"),
						State: increase.F("NY"),
						Zip:   increase.F("10045"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("078051120"),
						Passport: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
						}),
					}),
				}),
				CompanyTitle: increase.F("CEO"),
				Prongs:       increase.F([]increase.EntityNewParamsCorporationBeneficialOwnersProng{increase.EntityNewParamsCorporationBeneficialOwnersProngControl}),
			}}),
		}),
		Description: increase.F("x"),
		Joint: increase.F(increase.EntityNewParamsJoint{
			Name: increase.F("x"),
			Individuals: increase.F([]increase.EntityNewParamsJointIndividual{{
				Name:        increase.F("x"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityNewParamsJointIndividualsAddress{
					Line1: increase.F("x"),
					Line2: increase.F("x"),
					City:  increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityNewParamsJointIndividualsIdentification{
					Method: increase.F(increase.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					Passport: increase.F(increase.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						Country:        increase.F("x"),
					}),
					DriversLicense: increase.F(increase.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
					}),
				}),
			}, {
				Name:        increase.F("x"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityNewParamsJointIndividualsAddress{
					Line1: increase.F("x"),
					Line2: increase.F("x"),
					City:  increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityNewParamsJointIndividualsIdentification{
					Method: increase.F(increase.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					Passport: increase.F(increase.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						Country:        increase.F("x"),
					}),
					DriversLicense: increase.F(increase.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
					}),
				}),
			}, {
				Name:        increase.F("x"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityNewParamsJointIndividualsAddress{
					Line1: increase.F("x"),
					Line2: increase.F("x"),
					City:  increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityNewParamsJointIndividualsIdentification{
					Method: increase.F(increase.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					Passport: increase.F(increase.EntityNewParamsJointIndividualsIdentificationPassport{
						FileID:         increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						Country:        increase.F("x"),
					}),
					DriversLicense: increase.F(increase.EntityNewParamsJointIndividualsIdentificationDriversLicense{
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
					}),
				}),
			}}),
		}),
		NaturalPerson: increase.F(increase.EntityNewParamsNaturalPerson{
			Name:        increase.F("x"),
			DateOfBirth: increase.F(time.Now()),
			Address: increase.F(increase.EntityNewParamsNaturalPersonAddress{
				Line1: increase.F("x"),
				Line2: increase.F("x"),
				City:  increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
			}),
			ConfirmedNoUsTaxID: increase.F(true),
			Identification: increase.F(increase.EntityNewParamsNaturalPersonIdentification{
				Method: increase.F(increase.EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber),
				Number: increase.F("xxxx"),
				Passport: increase.F(increase.EntityNewParamsNaturalPersonIdentificationPassport{
					FileID:         increase.F("string"),
					ExpirationDate: increase.F(time.Now()),
					Country:        increase.F("x"),
				}),
				DriversLicense: increase.F(increase.EntityNewParamsNaturalPersonIdentificationDriversLicense{
					FileID:         increase.F("string"),
					BackFileID:     increase.F("string"),
					ExpirationDate: increase.F(time.Now()),
					State:          increase.F("x"),
				}),
				Other: increase.F(increase.EntityNewParamsNaturalPersonIdentificationOther{
					Country:        increase.F("x"),
					Description:    increase.F("x"),
					ExpirationDate: increase.F(time.Now()),
					FileID:         increase.F("string"),
					BackFileID:     increase.F("string"),
				}),
			}),
		}),
		Relationship: increase.F(increase.EntityNewParamsRelationshipAffiliated),
		SupplementalDocuments: increase.F([]increase.EntityNewParamsSupplementalDocument{{
			FileID: increase.F("file_makxrc67oh9l6sg7w9yc"),
		}}),
		Trust: increase.F(increase.EntityNewParamsTrust{
			Name:           increase.F("x"),
			Category:       increase.F(increase.EntityNewParamsTrustCategoryRevocable),
			TaxIdentifier:  increase.F("x"),
			FormationState: increase.F("x"),
			Address: increase.F(increase.EntityNewParamsTrustAddress{
				Line1: increase.F("x"),
				Line2: increase.F("x"),
				City:  increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
			}),
			FormationDocumentFileID: increase.F("string"),
			Trustees: increase.F([]increase.EntityNewParamsTrustTrustee{{
				Structure: increase.F(increase.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: increase.F(increase.EntityNewParamsTrustTrusteesIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
						}),
					}),
				}),
			}, {
				Structure: increase.F(increase.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: increase.F(increase.EntityNewParamsTrustTrusteesIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
						}),
					}),
				}),
			}, {
				Structure: increase.F(increase.EntityNewParamsTrustTrusteesStructureIndividual),
				Individual: increase.F(increase.EntityNewParamsTrustTrusteesIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsTrustTrusteesIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
							BackFileID:     increase.F("string"),
						}),
					}),
				}),
			}}),
			Grantor: increase.F(increase.EntityNewParamsTrustGrantor{
				Name:        increase.F("x"),
				DateOfBirth: increase.F(time.Now()),
				Address: increase.F(increase.EntityNewParamsTrustGrantorAddress{
					Line1: increase.F("x"),
					Line2: increase.F("x"),
					City:  increase.F("x"),
					State: increase.F("x"),
					Zip:   increase.F("x"),
				}),
				ConfirmedNoUsTaxID: increase.F(true),
				Identification: increase.F(increase.EntityNewParamsTrustGrantorIdentification{
					Method: increase.F(increase.EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber),
					Number: increase.F("xxxx"),
					Passport: increase.F(increase.EntityNewParamsTrustGrantorIdentificationPassport{
						FileID:         increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						Country:        increase.F("x"),
					}),
					DriversLicense: increase.F(increase.EntityNewParamsTrustGrantorIdentificationDriversLicense{
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsTrustGrantorIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
						BackFileID:     increase.F("string"),
					}),
				}),
			}),
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
		Cursor: increase.F("string"),
		Limit:  increase.F(int64(0)),
		Status: increase.F(increase.EntityListParamsStatus{
			In: increase.F([]increase.EntityListParamsStatusIn{increase.EntityListParamsStatusInActive, increase.EntityListParamsStatusInArchived, increase.EntityListParamsStatusInDisabled}),
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
				Line1: increase.F("33 Liberty Street"),
				Line2: increase.F("Unit 2"),
				City:  increase.F("New York"),
				State: increase.F("NY"),
				Zip:   increase.F("10045"),
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
