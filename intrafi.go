// File generated from our OpenAPI spec by Stainless.

package increase

import (
	"github.com/increase/increase-go/option"
)

// IntrafiService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntrafiService] method instead.
type IntrafiService struct {
	Options            []option.RequestOption
	AccountEnrollments *IntrafiAccountEnrollmentService
	Balances           *IntrafiBalanceService
	Exclusions         *IntrafiExclusionService
}

// NewIntrafiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntrafiService(opts ...option.RequestOption) (r *IntrafiService) {
	r = &IntrafiService{}
	r.Options = opts
	r.AccountEnrollments = NewIntrafiAccountEnrollmentService(opts...)
	r.Balances = NewIntrafiBalanceService(opts...)
	r.Exclusions = NewIntrafiExclusionService(opts...)
	return
}
