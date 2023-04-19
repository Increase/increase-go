package responses

import (
	"time"

	pjson "github.com/increase/increase-go/core/json"
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
	ID        pjson.Metadata
	Category  pjson.Metadata
	CreatedAt pjson.Metadata
	EntityID  pjson.Metadata
	FileID    pjson.Metadata
	Type      pjson.Metadata
	Raw       []byte
	Extras    map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Document using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DocumentCategory string

const (
	DocumentCategoryAntiMoneyLaunderingPolicy        DocumentCategory = "anti_money_laundering_policy"
	DocumentCategoryAntiMoneyLaunderingProcedures    DocumentCategory = "anti_money_laundering_procedures"
	DocumentCategoryAuditReport                      DocumentCategory = "audit_report"
	DocumentCategoryBackgroundChecks                 DocumentCategory = "background_checks"
	DocumentCategoryBusinessContinuityPlan           DocumentCategory = "business_continuity_plan"
	DocumentCategoryCollectionsPolicy                DocumentCategory = "collections_policy"
	DocumentCategoryComplaintsPolicy                 DocumentCategory = "complaints_policy"
	DocumentCategoryComplaintReport                  DocumentCategory = "complaint_report"
	DocumentCategoryComplianceReport                 DocumentCategory = "compliance_report"
	DocumentCategoryComplianceManagementSystemPolicy DocumentCategory = "compliance_management_system_policy"
	DocumentCategoryConsumerProtectionPolicy         DocumentCategory = "consumer_protection_policy"
	DocumentCategoryCorporateFormationDocument       DocumentCategory = "corporate_formation_document"
	DocumentCategoryCreditMonitoringReport           DocumentCategory = "credit_monitoring_report"
	DocumentCategoryCustomerInformationProgramPolicy DocumentCategory = "customer_information_program_policy"
	DocumentCategoryEmployeeOverview                 DocumentCategory = "employee_overview"
	DocumentCategoryEndUserTermsOfService            DocumentCategory = "end_user_terms_of_service"
	DocumentCategoryFinancialStatement               DocumentCategory = "financial_statement"
	DocumentCategoryForm_1099Int                     DocumentCategory = "form_1099_int"
	DocumentCategoryFraudPreventionPolicy            DocumentCategory = "fraud_prevention_policy"
	DocumentCategoryFundsFlowDiagram                 DocumentCategory = "funds_flow_diagram"
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
	Data       pjson.Metadata
	NextCursor pjson.Metadata
	Raw        []byte
	Extras     map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DocumentListResponse using
// the internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}
