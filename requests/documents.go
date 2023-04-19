package requests

import (
	"net/url"
	"time"

	"github.com/increase/increase-go/core/field"
	"github.com/increase/increase-go/core/query"
)

type DocumentListParams struct {
	// Return the page of entries after this one.
	Cursor field.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit field.Field[int64] `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID  field.Field[string]                      `query:"entity_id"`
	Category  field.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt field.Field[DocumentListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes DocumentListParams into a url.Values of the query parameters
// associated with this value
func (r DocumentListParams) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DocumentListParamsCategory struct {
	// Filter Documents for those with the specified category or categories. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In field.Field[[]DocumentListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes DocumentListParamsCategory into a url.Values of the query
// parameters associated with this value
func (r DocumentListParamsCategory) URLQuery() (v url.Values) {
	return query.Marshal(r)
}

type DocumentListParamsCategoryIn string

const (
	DocumentListParamsCategoryInAntiMoneyLaunderingPolicy        DocumentListParamsCategoryIn = "anti_money_laundering_policy"
	DocumentListParamsCategoryInAntiMoneyLaunderingProcedures    DocumentListParamsCategoryIn = "anti_money_laundering_procedures"
	DocumentListParamsCategoryInAuditReport                      DocumentListParamsCategoryIn = "audit_report"
	DocumentListParamsCategoryInBackgroundChecks                 DocumentListParamsCategoryIn = "background_checks"
	DocumentListParamsCategoryInBusinessContinuityPlan           DocumentListParamsCategoryIn = "business_continuity_plan"
	DocumentListParamsCategoryInCollectionsPolicy                DocumentListParamsCategoryIn = "collections_policy"
	DocumentListParamsCategoryInComplaintsPolicy                 DocumentListParamsCategoryIn = "complaints_policy"
	DocumentListParamsCategoryInComplaintReport                  DocumentListParamsCategoryIn = "complaint_report"
	DocumentListParamsCategoryInComplianceReport                 DocumentListParamsCategoryIn = "compliance_report"
	DocumentListParamsCategoryInComplianceManagementSystemPolicy DocumentListParamsCategoryIn = "compliance_management_system_policy"
	DocumentListParamsCategoryInConsumerProtectionPolicy         DocumentListParamsCategoryIn = "consumer_protection_policy"
	DocumentListParamsCategoryInCorporateFormationDocument       DocumentListParamsCategoryIn = "corporate_formation_document"
	DocumentListParamsCategoryInCreditMonitoringReport           DocumentListParamsCategoryIn = "credit_monitoring_report"
	DocumentListParamsCategoryInCustomerInformationProgramPolicy DocumentListParamsCategoryIn = "customer_information_program_policy"
	DocumentListParamsCategoryInEmployeeOverview                 DocumentListParamsCategoryIn = "employee_overview"
	DocumentListParamsCategoryInEndUserTermsOfService            DocumentListParamsCategoryIn = "end_user_terms_of_service"
	DocumentListParamsCategoryInFinancialStatement               DocumentListParamsCategoryIn = "financial_statement"
	DocumentListParamsCategoryInForm_1099Int                     DocumentListParamsCategoryIn = "form_1099_int"
	DocumentListParamsCategoryInFraudPreventionPolicy            DocumentListParamsCategoryIn = "fraud_prevention_policy"
	DocumentListParamsCategoryInFundsFlowDiagram                 DocumentListParamsCategoryIn = "funds_flow_diagram"
	DocumentListParamsCategoryInInformationSecurityPolicy        DocumentListParamsCategoryIn = "information_security_policy"
	DocumentListParamsCategoryInInsurancePolicy                  DocumentListParamsCategoryIn = "insurance_policy"
	DocumentListParamsCategoryInInvestorPresentation             DocumentListParamsCategoryIn = "investor_presentation"
	DocumentListParamsCategoryInLoanApplicationProcessingPolicy  DocumentListParamsCategoryIn = "loan_application_processing_policy"
	DocumentListParamsCategoryInManagementBiography              DocumentListParamsCategoryIn = "management_biography"
	DocumentListParamsCategoryInMarketingAndAdvertisingPolicy    DocumentListParamsCategoryIn = "marketing_and_advertising_policy"
	DocumentListParamsCategoryInNetworkSecurityDiagram           DocumentListParamsCategoryIn = "network_security_diagram"
	DocumentListParamsCategoryInOnboardingQuestionnaire          DocumentListParamsCategoryIn = "onboarding_questionnaire"
	DocumentListParamsCategoryInPenetrationTestReport            DocumentListParamsCategoryIn = "penetration_test_report"
	DocumentListParamsCategoryInProgramRiskAssessment            DocumentListParamsCategoryIn = "program_risk_assessment"
	DocumentListParamsCategoryInSecurityAuditReport              DocumentListParamsCategoryIn = "security_audit_report"
	DocumentListParamsCategoryInServicingPolicy                  DocumentListParamsCategoryIn = "servicing_policy"
	DocumentListParamsCategoryInTransactionMonitoringReport      DocumentListParamsCategoryIn = "transaction_monitoring_report"
	DocumentListParamsCategoryInUnderwritingPolicy               DocumentListParamsCategoryIn = "underwriting_policy"
	DocumentListParamsCategoryInVendorList                       DocumentListParamsCategoryIn = "vendor_list"
	DocumentListParamsCategoryInVendorManagementPolicy           DocumentListParamsCategoryIn = "vendor_management_policy"
	DocumentListParamsCategoryInVendorRiskManagementReport       DocumentListParamsCategoryIn = "vendor_risk_management_report"
	DocumentListParamsCategoryInVolumeForecast                   DocumentListParamsCategoryIn = "volume_forecast"
)

type DocumentListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After field.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before field.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter field.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore field.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes DocumentListParamsCreatedAt into a url.Values of the query
// parameters associated with this value
func (r DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return query.Marshal(r)
}
