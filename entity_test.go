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

func TestEntityNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.Entities.New(context.TODO(), increase.EntityNewParams{
		Relationship: increase.F(increase.EntityNewParamsRelationshipAffiliated),
		Structure:    increase.F(increase.EntityNewParamsStructureCorporation),
		Corporation: increase.F(increase.EntityNewParamsCorporation{
			Name:               increase.F("x"),
			Website:            increase.F("string"),
			TaxIdentifier:      increase.F("x"),
			IncorporationState: increase.F("x"),
			Address: increase.F(increase.EntityNewParamsCorporationAddress{
				Line1: increase.F("x"),
				Line2: increase.F("x"),
				City:  increase.F("x"),
				State: increase.F("x"),
				Zip:   increase.F("x"),
			}),
			BeneficialOwners: increase.F([]increase.EntityNewParamsCorporationBeneficialOwner{{
				Individual: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
						}),
					}),
				}),
				CompanyTitle: increase.F("x"),
				Prong:        increase.F(increase.EntityNewParamsCorporationBeneficialOwnersProngOwnership),
			}, {
				Individual: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
						}),
					}),
				}),
				CompanyTitle: increase.F("x"),
				Prong:        increase.F(increase.EntityNewParamsCorporationBeneficialOwnersProngOwnership),
			}, {
				Individual: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividual{
					Name:        increase.F("x"),
					DateOfBirth: increase.F(time.Now()),
					Address: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{
						Line1: increase.F("x"),
						Line2: increase.F("x"),
						City:  increase.F("x"),
						State: increase.F("x"),
						Zip:   increase.F("x"),
					}),
					ConfirmedNoUsTaxID: increase.F(true),
					Identification: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{
						Method: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber),
						Number: increase.F("xxxx"),
						Passport: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							Country:        increase.F("x"),
						}),
						DriversLicense: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{
							FileID:         increase.F("string"),
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
						}),
					}),
				}),
				CompanyTitle: increase.F("x"),
				Prong:        increase.F(increase.EntityNewParamsCorporationBeneficialOwnersProngOwnership),
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
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
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
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
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
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsJointIndividualsIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
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
					ExpirationDate: increase.F(time.Now()),
					State:          increase.F("x"),
				}),
				Other: increase.F(increase.EntityNewParamsNaturalPersonIdentificationOther{
					Country:        increase.F("x"),
					Description:    increase.F("x"),
					ExpirationDate: increase.F(time.Now()),
					FileID:         increase.F("string"),
				}),
			}),
		}),
		SupplementalDocuments: increase.F([]increase.EntityNewParamsSupplementalDocument{{
			FileID: increase.F("string"),
		}, {
			FileID: increase.F("string"),
		}, {
			FileID: increase.F("string"),
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
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
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
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
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
							ExpirationDate: increase.F(time.Now()),
							State:          increase.F("x"),
						}),
						Other: increase.F(increase.EntityNewParamsTrustTrusteesIndividualIdentificationOther{
							Country:        increase.F("x"),
							Description:    increase.F("x"),
							ExpirationDate: increase.F(time.Now()),
							FileID:         increase.F("string"),
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
						ExpirationDate: increase.F(time.Now()),
						State:          increase.F("x"),
					}),
					Other: increase.F(increase.EntityNewParamsTrustGrantorIdentificationOther{
						Country:        increase.F("x"),
						Description:    increase.F("x"),
						ExpirationDate: increase.F(time.Now()),
						FileID:         increase.F("string"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
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
	if !testutil.CheckTestServer(t) {
		return
	}
	client := increase.NewClient(
		option.WithAPIKey("APIKey"),
		option.WithBaseURL("http://127.0.0.1:4010"),
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
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
