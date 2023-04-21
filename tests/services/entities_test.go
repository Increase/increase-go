package services

import (
	"context"
	"errors"
	"net/http/httputil"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/core"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func TestEntityNewWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.New(context.TODO(), &requests.EntityNewParams{Structure: increase.F(requests.EntityNewParamsStructureCorporation), Corporation: increase.F(requests.EntityNewParamsCorporation{Name: increase.F("National Phonograph Company"), Website: increase.F("https://example.com"), TaxIdentifier: increase.F("602214076"), IncorporationState: increase.F("NY"), Address: increase.F(requests.EntityNewParamsCorporationAddress{Line1: increase.F("33 Liberty Street"), Line2: increase.F("x"), City: increase.F("New York"), State: increase.F("NY"), Zip: increase.F("10045")}), BeneficialOwners: increase.F([]requests.EntityNewParamsCorporationBeneficialOwners{{Individual: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{Method: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}), CompanyTitle: increase.F("x"), Prong: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersProngOwnership)}, {Individual: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{Method: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}), CompanyTitle: increase.F("x"), Prong: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersProngOwnership)}, {Individual: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentification{Method: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}), CompanyTitle: increase.F("x"), Prong: increase.F(requests.EntityNewParamsCorporationBeneficialOwnersProngOwnership)}})}), NaturalPerson: increase.F(requests.EntityNewParamsNaturalPerson{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsNaturalPersonAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsNaturalPersonIdentification{Method: increase.F(requests.EntityNewParamsNaturalPersonIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsNaturalPersonIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsNaturalPersonIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsNaturalPersonIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}), Joint: increase.F(requests.EntityNewParamsJoint{Name: increase.F("x"), Individuals: increase.F([]requests.EntityNewParamsJointIndividuals{{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsJointIndividualsAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsJointIndividualsIdentification{Method: increase.F(requests.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsJointIndividualsIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsJointIndividualsIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsJointIndividualsIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}, {Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsJointIndividualsAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsJointIndividualsIdentification{Method: increase.F(requests.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsJointIndividualsIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsJointIndividualsIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsJointIndividualsIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}, {Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsJointIndividualsAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsJointIndividualsIdentification{Method: increase.F(requests.EntityNewParamsJointIndividualsIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsJointIndividualsIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsJointIndividualsIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsJointIndividualsIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})}})}), Trust: increase.F(requests.EntityNewParamsTrust{Name: increase.F("x"), Category: increase.F(requests.EntityNewParamsTrustCategoryRevocable), TaxIdentifier: increase.F("x"), FormationState: increase.F("x"), Address: increase.F(requests.EntityNewParamsTrustAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), FormationDocumentFileID: increase.F("string"), Trustees: increase.F([]requests.EntityNewParamsTrustTrustees{{Structure: increase.F(requests.EntityNewParamsTrustTrusteesStructureIndividual), Individual: increase.F(requests.EntityNewParamsTrustTrusteesIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsTrustTrusteesIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentification{Method: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})})}, {Structure: increase.F(requests.EntityNewParamsTrustTrusteesStructureIndividual), Individual: increase.F(requests.EntityNewParamsTrustTrusteesIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsTrustTrusteesIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentification{Method: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})})}, {Structure: increase.F(requests.EntityNewParamsTrustTrusteesStructureIndividual), Individual: increase.F(requests.EntityNewParamsTrustTrusteesIndividual{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsTrustTrusteesIndividualAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentification{Method: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsTrustTrusteesIndividualIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})})}}), Grantor: increase.F(requests.EntityNewParamsTrustGrantor{Name: increase.F("x"), DateOfBirth: increase.F(time.Now()), Address: increase.F(requests.EntityNewParamsTrustGrantorAddress{Line1: increase.F("x"), Line2: increase.F("x"), City: increase.F("x"), State: increase.F("x"), Zip: increase.F("x")}), ConfirmedNoUsTaxID: increase.F(true), Identification: increase.F(requests.EntityNewParamsTrustGrantorIdentification{Method: increase.F(requests.EntityNewParamsTrustGrantorIdentificationMethodSocialSecurityNumber), Number: increase.F("xxxx"), Passport: increase.F(requests.EntityNewParamsTrustGrantorIdentificationPassport{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), Country: increase.F("x")}), DriversLicense: increase.F(requests.EntityNewParamsTrustGrantorIdentificationDriversLicense{FileID: increase.F("string"), ExpirationDate: increase.F(time.Now()), State: increase.F("x")}), Other: increase.F(requests.EntityNewParamsTrustGrantorIdentificationOther{Country: increase.F("x"), Description: increase.F("x"), ExpirationDate: increase.F(time.Now()), FileID: increase.F("string")})})})}), Description: increase.F("x"), Relationship: increase.F(requests.EntityNewParamsRelationshipAffiliated), SupplementalDocuments: increase.F([]requests.EntityNewParamsSupplementalDocuments{{FileID: increase.F("string")}, {FileID: increase.F("string")}, {FileID: increase.F("string")}})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityGet(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.Get(
		context.TODO(),
		"entity_n8y8tnk2p9339ti393yi",
	)
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestEntityListWithOptionalParams(t *testing.T) {
	c := increase.NewIncrease(option.WithAPIKey("APIKey"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Entities.List(context.TODO(), &requests.EntityListParams{Cursor: increase.F("string"), Limit: increase.F(int64(0)), CreatedAt: increase.F(requests.EntityListParamsCreatedAt{After: increase.F(time.Now()), Before: increase.F(time.Now()), OnOrAfter: increase.F(time.Now()), OnOrBefore: increase.F(time.Now())})})
	if err != nil {
		var apiError core.APIError
		if errors.As(err, &apiError) {
			body, _ := httputil.DumpRequest(apiError.Request(), true)
			println(string(body))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
