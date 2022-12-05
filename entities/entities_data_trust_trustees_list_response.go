package entities

type EntitiesDataTrustTrusteesListResponse struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure *string `json:"structure,omitempty"`

	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual *EntitiesDataTrustTrusteesIndividualListResponse `json:"individual,omitempty"`
}

// The structure of the trustee. Will always be equal to `individual`.
func (r *EntitiesDataTrustTrusteesListResponse) GetStructure() (Structure string) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// The individual trustee of the trust. Will be present if the trustee's
// `structure` is equal to `individual`.
func (r *EntitiesDataTrustTrusteesListResponse) GetIndividual() (Individual EntitiesDataTrustTrusteesIndividualListResponse) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}
