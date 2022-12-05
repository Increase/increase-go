package entities

type EntitiesTrustTrusteesCreateResponse struct {
	// The structure of the trustee. Will always be equal to `individual`.
	Structure *string `json:"structure,omitempty"`

	// The individual trustee of the trust. Will be present if the trustee's
	// `structure` is equal to `individual`.
	Individual *EntitiesTrustTrusteesIndividualCreateResponse `json:"individual,omitempty"`
}

// The structure of the trustee. Will always be equal to `individual`.
func (r *EntitiesTrustTrusteesCreateResponse) GetStructure() (Structure string) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// The individual trustee of the trust. Will be present if the trustee's
// `structure` is equal to `individual`.
func (r *EntitiesTrustTrusteesCreateResponse) GetIndividual() (Individual EntitiesTrustTrusteesIndividualCreateResponse) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}
