package increase

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/increase/increase-go/internal/apijson"
	"github.com/increase/increase-go/internal/apiquery"
	"github.com/increase/increase-go/internal/param"
	"github.com/increase/increase-go/internal/requestconfig"
	"github.com/increase/increase-go/internal/shared"
	"github.com/increase/increase-go/option"
)

// DocumentService contains methods and other services that help with interacting
// with the increase API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDocumentService] method instead.
type DocumentService struct {
	Options []option.RequestOption
}

// NewDocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDocumentService(opts ...option.RequestOption) (r *DocumentService) {
	r = &DocumentService{}
	r.Options = opts
	return
}

// Retrieve a Document
func (r *DocumentService) Get(ctx context.Context, document_id string, opts ...option.RequestOption) (res *Document, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("documents/%s", document_id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List Documents
func (r *DocumentService) List(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) (res *shared.Page[Document], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "documents"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List Documents
func (r *DocumentService) ListAutoPaging(ctx context.Context, query DocumentListParams, opts ...option.RequestOption) *shared.PageAutoPager[Document] {
	return shared.NewPageAutoPager(r.List(ctx, query, opts...))
}

// Increase generates certain documents / forms automatically for your application;
// they can be listed here. Currently the only supported document type is IRS Form
// 1099-INT.
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
	JSON documentJSON
}

// documentJSON contains the JSON metadata for the struct [Document]
type documentJSON struct {
	ID          apijson.Field
	Category    apijson.Field
	CreatedAt   apijson.Field
	EntityID    apijson.Field
	FileID      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Document) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentCategory string

const (
	DocumentCategoryAccountOpeningDisclosures                   DocumentCategory = "account_opening_disclosures"
	DocumentCategoryAntiMoneyLaunderingPolicy                   DocumentCategory = "anti_money_laundering_policy"
	DocumentCategoryAntiMoneyLaunderingProcedures               DocumentCategory = "anti_money_laundering_procedures"
	DocumentCategoryAuditReport                                 DocumentCategory = "audit_report"
	DocumentCategoryBackgroundChecks                            DocumentCategory = "background_checks"
	DocumentCategoryBusinessContinuityPlan                      DocumentCategory = "business_continuity_plan"
	DocumentCategoryCollectionsPolicy                           DocumentCategory = "collections_policy"
	DocumentCategoryComplaintsPolicy                            DocumentCategory = "complaints_policy"
	DocumentCategoryComplaintReport                             DocumentCategory = "complaint_report"
	DocumentCategoryComplianceReport                            DocumentCategory = "compliance_report"
	DocumentCategoryComplianceStaffingPlan                      DocumentCategory = "compliance_staffing_plan"
	DocumentCategoryComplianceManagementSystemPolicy            DocumentCategory = "compliance_management_system_policy"
	DocumentCategoryConsumerPrivacyNotice                       DocumentCategory = "consumer_privacy_notice"
	DocumentCategoryConsumerProtectionPolicy                    DocumentCategory = "consumer_protection_policy"
	DocumentCategoryCorporateFormationDocument                  DocumentCategory = "corporate_formation_document"
	DocumentCategoryCreditMonitoringReport                      DocumentCategory = "credit_monitoring_report"
	DocumentCategoryCustomerInformationProgramPolicy            DocumentCategory = "customer_information_program_policy"
	DocumentCategoryElectronicFundsTranferActPolicy             DocumentCategory = "electronic_funds_tranfer_act_policy"
	DocumentCategoryEmployeeOverview                            DocumentCategory = "employee_overview"
	DocumentCategoryEndUserTermsOfService                       DocumentCategory = "end_user_terms_of_service"
	DocumentCategoryESignPolicy                                 DocumentCategory = "e_sign_policy"
	DocumentCategoryFinancialStatement                          DocumentCategory = "financial_statement"
	DocumentCategoryForm1099Int                                 DocumentCategory = "form_1099_int"
	DocumentCategoryFraudPreventionPolicy                       DocumentCategory = "fraud_prevention_policy"
	DocumentCategoryFundsAvailabilityPolicy                     DocumentCategory = "funds_availability_policy"
	DocumentCategoryFundsAvailabilityDisclosure                 DocumentCategory = "funds_availability_disclosure"
	DocumentCategoryFundsFlowDiagram                            DocumentCategory = "funds_flow_diagram"
	DocumentCategoryGrammLeachBlileyActPolicy                   DocumentCategory = "gramm_leach_bliley_act_policy"
	DocumentCategoryInformationSecurityPolicy                   DocumentCategory = "information_security_policy"
	DocumentCategoryInsurancePolicy                             DocumentCategory = "insurance_policy"
	DocumentCategoryInvestorPresentation                        DocumentCategory = "investor_presentation"
	DocumentCategoryLoanApplicationProcessingPolicy             DocumentCategory = "loan_application_processing_policy"
	DocumentCategoryManagementBiography                         DocumentCategory = "management_biography"
	DocumentCategoryMarketingAndAdvertisingPolicy               DocumentCategory = "marketing_and_advertising_policy"
	DocumentCategoryNetworkSecurityDiagram                      DocumentCategory = "network_security_diagram"
	DocumentCategoryOnboardingQuestionnaire                     DocumentCategory = "onboarding_questionnaire"
	DocumentCategoryPenetrationTestReport                       DocumentCategory = "penetration_test_report"
	DocumentCategoryPlatformComplianceMetricsSubmission         DocumentCategory = "platform_compliance_metrics_submission"
	DocumentCategoryConsumerPlatformComplianceMetricsSubmission DocumentCategory = "consumer_platform_compliance_metrics_submission"
	DocumentCategoryProgramRiskAssessment                       DocumentCategory = "program_risk_assessment"
	DocumentCategorySecurityAuditReport                         DocumentCategory = "security_audit_report"
	DocumentCategoryServicingPolicy                             DocumentCategory = "servicing_policy"
	DocumentCategoryTransactionMonitoringReport                 DocumentCategory = "transaction_monitoring_report"
	DocumentCategoryTruthInSavingsActPolicy                     DocumentCategory = "truth_in_savings_act_policy"
	DocumentCategoryUnderwritingPolicy                          DocumentCategory = "underwriting_policy"
	DocumentCategoryVendorList                                  DocumentCategory = "vendor_list"
	DocumentCategoryVendorManagementPolicy                      DocumentCategory = "vendor_management_policy"
	DocumentCategoryVendorRiskManagementReport                  DocumentCategory = "vendor_risk_management_report"
	DocumentCategoryVolumeForecast                              DocumentCategory = "volume_forecast"
)

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type DocumentListParams struct {
	// Return the page of entries after this one.
	Cursor param.Field[string] `query:"cursor"`
	// Limit the size of the list that is returned. The default (and maximum) is 100
	// objects.
	Limit param.Field[int64] `query:"limit"`
	// Filter Documents to ones belonging to the specified Entity.
	EntityID  param.Field[string]                      `query:"entity_id"`
	Category  param.Field[DocumentListParamsCategory]  `query:"category"`
	CreatedAt param.Field[DocumentListParamsCreatedAt] `query:"created_at"`
}

// URLQuery serializes [DocumentListParams]'s query parameters as `url.Values`.
func (r DocumentListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DocumentListParamsCategory struct {
	// Filter Documents for those with the specified category or categories. For GET
	// requests, this should be encoded as a comma-delimited string, such as
	// `?in=one,two,three`.
	In param.Field[[]DocumentListParamsCategoryIn] `query:"in"`
}

// URLQuery serializes [DocumentListParamsCategory]'s query parameters as
// `url.Values`.
func (r DocumentListParamsCategory) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DocumentListParamsCategoryIn string

const (
	DocumentListParamsCategoryInAccountOpeningDisclosures                   DocumentListParamsCategoryIn = "account_opening_disclosures"
	DocumentListParamsCategoryInAntiMoneyLaunderingPolicy                   DocumentListParamsCategoryIn = "anti_money_laundering_policy"
	DocumentListParamsCategoryInAntiMoneyLaunderingProcedures               DocumentListParamsCategoryIn = "anti_money_laundering_procedures"
	DocumentListParamsCategoryInAuditReport                                 DocumentListParamsCategoryIn = "audit_report"
	DocumentListParamsCategoryInBackgroundChecks                            DocumentListParamsCategoryIn = "background_checks"
	DocumentListParamsCategoryInBusinessContinuityPlan                      DocumentListParamsCategoryIn = "business_continuity_plan"
	DocumentListParamsCategoryInCollectionsPolicy                           DocumentListParamsCategoryIn = "collections_policy"
	DocumentListParamsCategoryInComplaintsPolicy                            DocumentListParamsCategoryIn = "complaints_policy"
	DocumentListParamsCategoryInComplaintReport                             DocumentListParamsCategoryIn = "complaint_report"
	DocumentListParamsCategoryInComplianceReport                            DocumentListParamsCategoryIn = "compliance_report"
	DocumentListParamsCategoryInComplianceStaffingPlan                      DocumentListParamsCategoryIn = "compliance_staffing_plan"
	DocumentListParamsCategoryInComplianceManagementSystemPolicy            DocumentListParamsCategoryIn = "compliance_management_system_policy"
	DocumentListParamsCategoryInConsumerPrivacyNotice                       DocumentListParamsCategoryIn = "consumer_privacy_notice"
	DocumentListParamsCategoryInConsumerProtectionPolicy                    DocumentListParamsCategoryIn = "consumer_protection_policy"
	DocumentListParamsCategoryInCorporateFormationDocument                  DocumentListParamsCategoryIn = "corporate_formation_document"
	DocumentListParamsCategoryInCreditMonitoringReport                      DocumentListParamsCategoryIn = "credit_monitoring_report"
	DocumentListParamsCategoryInCustomerInformationProgramPolicy            DocumentListParamsCategoryIn = "customer_information_program_policy"
	DocumentListParamsCategoryInElectronicFundsTranferActPolicy             DocumentListParamsCategoryIn = "electronic_funds_tranfer_act_policy"
	DocumentListParamsCategoryInEmployeeOverview                            DocumentListParamsCategoryIn = "employee_overview"
	DocumentListParamsCategoryInEndUserTermsOfService                       DocumentListParamsCategoryIn = "end_user_terms_of_service"
	DocumentListParamsCategoryInESignPolicy                                 DocumentListParamsCategoryIn = "e_sign_policy"
	DocumentListParamsCategoryInFinancialStatement                          DocumentListParamsCategoryIn = "financial_statement"
	DocumentListParamsCategoryInForm1099Int                                 DocumentListParamsCategoryIn = "form_1099_int"
	DocumentListParamsCategoryInFraudPreventionPolicy                       DocumentListParamsCategoryIn = "fraud_prevention_policy"
	DocumentListParamsCategoryInFundsAvailabilityPolicy                     DocumentListParamsCategoryIn = "funds_availability_policy"
	DocumentListParamsCategoryInFundsAvailabilityDisclosure                 DocumentListParamsCategoryIn = "funds_availability_disclosure"
	DocumentListParamsCategoryInFundsFlowDiagram                            DocumentListParamsCategoryIn = "funds_flow_diagram"
	DocumentListParamsCategoryInGrammLeachBlileyActPolicy                   DocumentListParamsCategoryIn = "gramm_leach_bliley_act_policy"
	DocumentListParamsCategoryInInformationSecurityPolicy                   DocumentListParamsCategoryIn = "information_security_policy"
	DocumentListParamsCategoryInInsurancePolicy                             DocumentListParamsCategoryIn = "insurance_policy"
	DocumentListParamsCategoryInInvestorPresentation                        DocumentListParamsCategoryIn = "investor_presentation"
	DocumentListParamsCategoryInLoanApplicationProcessingPolicy             DocumentListParamsCategoryIn = "loan_application_processing_policy"
	DocumentListParamsCategoryInManagementBiography                         DocumentListParamsCategoryIn = "management_biography"
	DocumentListParamsCategoryInMarketingAndAdvertisingPolicy               DocumentListParamsCategoryIn = "marketing_and_advertising_policy"
	DocumentListParamsCategoryInNetworkSecurityDiagram                      DocumentListParamsCategoryIn = "network_security_diagram"
	DocumentListParamsCategoryInOnboardingQuestionnaire                     DocumentListParamsCategoryIn = "onboarding_questionnaire"
	DocumentListParamsCategoryInPenetrationTestReport                       DocumentListParamsCategoryIn = "penetration_test_report"
	DocumentListParamsCategoryInPlatformComplianceMetricsSubmission         DocumentListParamsCategoryIn = "platform_compliance_metrics_submission"
	DocumentListParamsCategoryInConsumerPlatformComplianceMetricsSubmission DocumentListParamsCategoryIn = "consumer_platform_compliance_metrics_submission"
	DocumentListParamsCategoryInProgramRiskAssessment                       DocumentListParamsCategoryIn = "program_risk_assessment"
	DocumentListParamsCategoryInSecurityAuditReport                         DocumentListParamsCategoryIn = "security_audit_report"
	DocumentListParamsCategoryInServicingPolicy                             DocumentListParamsCategoryIn = "servicing_policy"
	DocumentListParamsCategoryInTransactionMonitoringReport                 DocumentListParamsCategoryIn = "transaction_monitoring_report"
	DocumentListParamsCategoryInTruthInSavingsActPolicy                     DocumentListParamsCategoryIn = "truth_in_savings_act_policy"
	DocumentListParamsCategoryInUnderwritingPolicy                          DocumentListParamsCategoryIn = "underwriting_policy"
	DocumentListParamsCategoryInVendorList                                  DocumentListParamsCategoryIn = "vendor_list"
	DocumentListParamsCategoryInVendorManagementPolicy                      DocumentListParamsCategoryIn = "vendor_management_policy"
	DocumentListParamsCategoryInVendorRiskManagementReport                  DocumentListParamsCategoryIn = "vendor_risk_management_report"
	DocumentListParamsCategoryInVolumeForecast                              DocumentListParamsCategoryIn = "volume_forecast"
)

type DocumentListParamsCreatedAt struct {
	// Return results after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	After param.Field[time.Time] `query:"after" format:"date-time"`
	// Return results before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601)
	// timestamp.
	Before param.Field[time.Time] `query:"before" format:"date-time"`
	// Return results on or after this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrAfter param.Field[time.Time] `query:"on_or_after" format:"date-time"`
	// Return results on or before this
	// [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) timestamp.
	OnOrBefore param.Field[time.Time] `query:"on_or_before" format:"date-time"`
}

// URLQuery serializes [DocumentListParamsCreatedAt]'s query parameters as
// `url.Values`.
func (r DocumentListParamsCreatedAt) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// A list of Document objects
type DocumentListResponse struct {
	// The contents of the list.
	Data []Document `json:"data,required"`
	// A pointer to a place in the list.
	NextCursor string `json:"next_cursor,required,nullable"`
	JSON       documentListResponseJSON
}

// documentListResponseJSON contains the JSON metadata for the struct
// [DocumentListResponse]
type documentListResponseJSON struct {
	Data        apijson.Field
	NextCursor  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DocumentListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
