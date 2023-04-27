package responses

import (
	"time"

	apijson "github.com/increase/increase-go/internal/json"
)

type Document struct {
	// The Document identifier.
	ID string `json:"id,required"`
	// The type of document.
	Category DocumentCategory `json:"category,required"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) time at which the
	// Document was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The identifier of the Entity the document was generated for.
	EntityID string `json:"entity_id,required,nullable"`
	// The identifier of the File containing the Document's contents.
	FileID string `json:"file_id,required"`
	// A constant representing the object's type. For this resource it will always be
	// `document`.
	Type DocumentType `json:"type,required"`
	JSON DocumentJSON
}

type DocumentJSON struct {
	ID        apijson.Metadata
	Category  apijson.Metadata
	CreatedAt apijson.Metadata
	EntityID  apijson.Metadata
	FileID    apijson.Metadata
	Type      apijson.Metadata
	raw       string
	Extras    map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// json library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentCategory string

const (
	DocumentCategoryAccountOpeningDisclosures        DocumentCategory = "account_opening_disclosures"
	DocumentCategoryAntiMoneyLaunderingPolicy        DocumentCategory = "anti_money_laundering_policy"
	DocumentCategoryAntiMoneyLaunderingProcedures    DocumentCategory = "anti_money_laundering_procedures"
	DocumentCategoryAuditReport                      DocumentCategory = "audit_report"
	DocumentCategoryBackgroundChecks                 DocumentCategory = "background_checks"
	DocumentCategoryBusinessContinuityPlan           DocumentCategory = "business_continuity_plan"
	DocumentCategoryCollectionsPolicy                DocumentCategory = "collections_policy"
	DocumentCategoryComplaintsPolicy                 DocumentCategory = "complaints_policy"
	DocumentCategoryComplaintReport                  DocumentCategory = "complaint_report"
	DocumentCategoryComplianceReport                 DocumentCategory = "compliance_report"
	DocumentCategoryComplianceStaffingPlan           DocumentCategory = "compliance_staffing_plan"
	DocumentCategoryComplianceManagementSystemPolicy DocumentCategory = "compliance_management_system_policy"
	DocumentCategoryConsumerPrivacyNotice            DocumentCategory = "consumer_privacy_notice"
	DocumentCategoryConsumerProtectionPolicy         DocumentCategory = "consumer_protection_policy"
	DocumentCategoryCorporateFormationDocument       DocumentCategory = "corporate_formation_document"
	DocumentCategoryCreditMonitoringReport           DocumentCategory = "credit_monitoring_report"
	DocumentCategoryCustomerInformationProgramPolicy DocumentCategory = "customer_information_program_policy"
	DocumentCategoryElectronicFundsTranferActPolicy  DocumentCategory = "electronic_funds_tranfer_act_policy"
	DocumentCategoryEmployeeOverview                 DocumentCategory = "employee_overview"
	DocumentCategoryEndUserTermsOfService            DocumentCategory = "end_user_terms_of_service"
	DocumentCategoryESignPolicy                      DocumentCategory = "e_sign_policy"
	DocumentCategoryFinancialStatement               DocumentCategory = "financial_statement"
	DocumentCategoryForm_1099Int                     DocumentCategory = "form_1099_int"
	DocumentCategoryFraudPreventionPolicy            DocumentCategory = "fraud_prevention_policy"
	DocumentCategoryFundsAvailabilityPolicy          DocumentCategory = "funds_availability_policy"
	DocumentCategoryFundsAvailabilityDisclosure      DocumentCategory = "funds_availability_disclosure"
	DocumentCategoryFundsFlowDiagram                 DocumentCategory = "funds_flow_diagram"
	DocumentCategoryGrammLeachBlileyActPolicy        DocumentCategory = "gramm_leach_bliley_act_policy"
	DocumentCategoryInformationSecurityPolicy        DocumentCategory = "information_security_policy"
	DocumentCategoryInsurancePolicy                  DocumentCategory = "insurance_policy"
	DocumentCategoryInvestorPresentation             DocumentCategory = "investor_presentation"
	DocumentCategoryLoanApplicationProcessingPolicy  DocumentCategory = "loan_application_processing_policy"
	DocumentCategoryManagementBiography              DocumentCategory = "management_biography"
	DocumentCategoryMarketingAndAdvertisingPolicy    DocumentCategory = "marketing_and_advertising_policy"
	DocumentCategoryNetworkSecurityDiagram           DocumentCategory = "network_security_diagram"
	DocumentCategoryOnboardingQuestionnaire          DocumentCategory = "onboarding_questionnaire"
	DocumentCategoryPenetrationTestReport            DocumentCategory = "penetration_test_report"
	DocumentCategoryProgramRiskAssessment            DocumentCategory = "program_risk_assessment"
	DocumentCategorySecurityAuditReport              DocumentCategory = "security_audit_report"
	DocumentCategoryServicingPolicy                  DocumentCategory = "servicing_policy"
	DocumentCategoryTransactionMonitoringReport      DocumentCategory = "transaction_monitoring_report"
	DocumentCategoryTruthInSavingsActPolicy          DocumentCategory = "truth_in_savings_act_policy"
	DocumentCategoryUnderwritingPolicy               DocumentCategory = "underwriting_policy"
	DocumentCategoryVendorList                       DocumentCategory = "vendor_list"
	DocumentCategoryVendorManagementPolicy           DocumentCategory = "vendor_management_policy"
	DocumentCategoryVendorRiskManagementReport       DocumentCategory = "vendor_risk_management_report"
	DocumentCategoryVolumeForecast                   DocumentCategory = "volume_forecast"
)

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type DocumentListResponse struct {
	// The contents of the list.
	Data []Document `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       DocumentListResponseJSON
}

type DocumentListResponseJSON struct {
	Data       apijson.Metadata
	NextCursor apijson.Metadata
	raw        string
	Extras     map[string]apijson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DocumentListResponse using
// the internal json library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
