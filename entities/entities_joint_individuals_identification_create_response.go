package entities

type EntitiesJointIndividualsIdentificationCreateResponse struct {
	// A method that can be used to verify the individual's identity.
	Method *string `json:"method,omitempty"`

	// The last 4 digits of the identification number that can be used to verify the
	// individual's identity.
	NumberLast4 *string `json:"number_last4,omitempty"`
}

// A method that can be used to verify the individual's identity.
func (r *EntitiesJointIndividualsIdentificationCreateResponse) GetMethod() (Method string) {
	if r != nil && r.Method != nil {
		Method = *r.Method
	}
	return
}

// The last 4 digits of the identification number that can be used to verify the
// individual's identity.
func (r *EntitiesJointIndividualsIdentificationCreateResponse) GetNumberLast4() (NumberLast4 string) {
	if r != nil && r.NumberLast4 != nil {
		NumberLast4 = *r.NumberLast4
	}
	return
}
