// File generated from our OpenAPI spec by Stainless.

package shared

type PointOfServiceEntryMode string

const (
	PointOfServiceEntryModeManual                     PointOfServiceEntryMode = "manual"
	PointOfServiceEntryModeMagneticStripeNoCvv        PointOfServiceEntryMode = "magnetic_stripe_no_cvv"
	PointOfServiceEntryModeOpticalCode                PointOfServiceEntryMode = "optical_code"
	PointOfServiceEntryModeIntegratedCircuitCard      PointOfServiceEntryMode = "integrated_circuit_card"
	PointOfServiceEntryModeContactless                PointOfServiceEntryMode = "contactless"
	PointOfServiceEntryModeCredentialOnFile           PointOfServiceEntryMode = "credential_on_file"
	PointOfServiceEntryModeMagneticStripe             PointOfServiceEntryMode = "magnetic_stripe"
	PointOfServiceEntryModeContactlessMagneticStripe  PointOfServiceEntryMode = "contactless_magnetic_stripe"
	PointOfServiceEntryModeIntegratedCircuitCardNoCvv PointOfServiceEntryMode = "integrated_circuit_card_no_cvv"
)
