// File generated from our OpenAPI spec by Stainless.

package shared

// The method used to enter the cardholder's primary account number and card
// expiration date
type PointOfServiceEntryMode string

const (
	// Unknown
	PointOfServiceEntryModeUnknown PointOfServiceEntryMode = "unknown"
	// Manual key entry
	PointOfServiceEntryModeManual PointOfServiceEntryMode = "manual"
	// Magnetic stripe read, without card verification value
	PointOfServiceEntryModeMagneticStripeNoCvv PointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	// Optical code
	PointOfServiceEntryModeOpticalCode PointOfServiceEntryMode = "optical_code"
	// Contact chip card
	PointOfServiceEntryModeIntegratedCircuitCard PointOfServiceEntryMode = "integrated_circuit_card"
	// Contactless read of chip card
	PointOfServiceEntryModeContactless PointOfServiceEntryMode = "contactless"
	// Transaction iniated using a credential that has previously been stored on file
	PointOfServiceEntryModeCredentialOnFile PointOfServiceEntryMode = "credential_on_file"
	// Magnetic stripe read
	PointOfServiceEntryModeMagneticStripe PointOfServiceEntryMode = "magnetic_stripe"
	// Contactless read of magnetic stripe data
	PointOfServiceEntryModeContactlessMagneticStripe PointOfServiceEntryMode = "contactless_magnetic_stripe"
	// Contact chip card, without card verification value
	PointOfServiceEntryModeIntegratedCircuitCardNoCvv PointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)
