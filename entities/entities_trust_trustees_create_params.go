package entities

type EntitiesTrustTrusteesCreateParams struct {
	// The structure of the trustee.
	Structure *string `json:"structure,omitempty"`

	// Details of the individual trustee. Required when the trustee `structure` is
	// equal to `individual`.
	Individual *EntitiesTrustTrusteesIndividualCreateParams `json:"individual,omitempty"`
}

// The structure of the trustee.
func (r *EntitiesTrustTrusteesCreateParams) GetStructure() (Structure string) {
	if r != nil && r.Structure != nil {
		Structure = *r.Structure
	}
	return
}

// Details of the individual trustee. Required when the trustee `structure` is
// equal to `individual`.
func (r *EntitiesTrustTrusteesCreateParams) GetIndividual() (Individual EntitiesTrustTrusteesIndividualCreateParams) {
	if r != nil && r.Individual != nil {
		Individual = *r.Individual
	}
	return
}
