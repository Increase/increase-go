// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase

import (
	"github.com/Increase/increase-go/internal/apierror"
)

type Error = apierror.Error
type ErrorStatus = apierror.ErrorStatus

const ErrorStatus429 = apierror.ErrorStatus429
const ErrorStatus403 = apierror.ErrorStatus403
const ErrorStatus404 = apierror.ErrorStatus404
const ErrorStatus400 = apierror.ErrorStatus400
const ErrorStatus409 = apierror.ErrorStatus409
const ErrorStatus401 = apierror.ErrorStatus401
const ErrorStatus500 = apierror.ErrorStatus500

type ErrorType = apierror.ErrorType

const ErrorTypeRateLimitedError = apierror.ErrorTypeRateLimitedError
const ErrorTypePrivateFeatureError = apierror.ErrorTypePrivateFeatureError
const ErrorTypeObjectNotFoundError = apierror.ErrorTypeObjectNotFoundError
const ErrorTypeMalformedRequestError = apierror.ErrorTypeMalformedRequestError
const ErrorTypeInvalidParametersError = apierror.ErrorTypeInvalidParametersError
const ErrorTypeInvalidOperationError = apierror.ErrorTypeInvalidOperationError
const ErrorTypeInvalidAPIKeyError = apierror.ErrorTypeInvalidAPIKeyError
const ErrorTypeInternalServerError = apierror.ErrorTypeInternalServerError
const ErrorTypeInsufficientPermissionsError = apierror.ErrorTypeInsufficientPermissionsError
const ErrorTypeIdempotencyKeyAlreadyUsedError = apierror.ErrorTypeIdempotencyKeyAlreadyUsedError
const ErrorTypeEnvironmentMismatchError = apierror.ErrorTypeEnvironmentMismatchError
const ErrorTypeAPIMethodNotFoundError = apierror.ErrorTypeAPIMethodNotFoundError
