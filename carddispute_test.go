// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package increase_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Increase/increase-go"
	"github.com/Increase/increase-go/internal/testutil"
	"github.com/Increase/increase-go/option"
)

func TestCardDisputeNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardDisputes.New(context.TODO(), increase.CardDisputeNewParams{
		DisputedTransactionID: increase.F("transaction_uyrp7fld2ium70oa7oi"),
		Network:               increase.F(increase.CardDisputeNewParamsNetworkVisa),
		Amount:                increase.F(int64(100)),
		AttachmentFiles: increase.F([]increase.CardDisputeNewParamsAttachmentFile{{
			FileID: increase.F("file_id"),
		}}),
		Explanation: increase.F("x"),
		Visa: increase.F(increase.CardDisputeNewParamsVisa{
			Category: increase.F(increase.CardDisputeNewParamsVisaCategoryFraud),
			Authorization: increase.F(increase.CardDisputeNewParamsVisaAuthorization{
				AccountStatus: increase.F(increase.CardDisputeNewParamsVisaAuthorizationAccountStatusAccountClosed),
			}),
			ConsumerCanceledMerchandise: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandise{
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted),
				PurchaseExplanation:         increase.F("x"),
				ReceivedOrExpectedAt:        increase.F(time.Now()),
				ReturnOutcome:               increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnOutcomeNotReturned),
				CardholderCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellation{
					CanceledAt:                 increase.F(time.Now()),
					CanceledPriorToShipDate:    increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate),
					CancellationPolicyProvided: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided),
					Reason:                     increase.F("x"),
				}),
				NotReturned: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseNotReturned{}),
				ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttempted{
					AttemptExplanation:     increase.F("x"),
					AttemptReason:          increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
					AttemptedAt:            increase.F(time.Now()),
					MerchandiseDisposition: increase.F("x"),
				}),
				Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturned{
					ReturnMethod:             increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledMerchandiseReturnedReturnMethodDhl),
					ReturnedAt:               increase.F(time.Now()),
					MerchantReceivedReturnAt: increase.F(time.Now()),
					OtherExplanation:         increase.F("x"),
					TrackingNumber:           increase.F("x"),
				}),
			}),
			ConsumerCanceledRecurringTransaction: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledRecurringTransaction{
				CancellationTarget: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionCancellationTargetAccount),
				MerchantContactMethods: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledRecurringTransactionMerchantContactMethods{
					ApplicationName:       increase.F("x"),
					CallCenterPhoneNumber: increase.F("x"),
					EmailAddress:          increase.F("x"),
					InPersonAddress:       increase.F("x"),
					MailingAddress:        increase.F("x"),
					TextPhoneNumber:       increase.F("x"),
				}),
				TransactionOrAccountCanceledAt: increase.F(time.Now()),
				OtherFormOfPaymentExplanation:  increase.F("x"),
			}),
			ConsumerCanceledServices: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServices{
				CardholderCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellation{
					CanceledAt:                 increase.F(time.Now()),
					CancellationPolicyProvided: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided),
					Reason:                     increase.F("x"),
				}),
				ContractedAt:                increase.F(time.Now()),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesMerchantResolutionAttemptedAttempted),
				PurchaseExplanation:         increase.F("x"),
				ServiceType:                 increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesServiceTypeGuaranteedReservation),
				GuaranteedReservation: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservation{
					Explanation: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService),
				}),
				Other:     increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesOther{}),
				Timeshare: increase.F(increase.CardDisputeNewParamsVisaConsumerCanceledServicesTimeshare{}),
			}),
			ConsumerCounterfeitMerchandise: increase.F(increase.CardDisputeNewParamsVisaConsumerCounterfeitMerchandise{
				CounterfeitExplanation: increase.F("x"),
				DispositionExplanation: increase.F("x"),
				OrderExplanation:       increase.F("x"),
				ReceivedAt:             increase.F(time.Now()),
			}),
			ConsumerCreditNotProcessed: increase.F(increase.CardDisputeNewParamsVisaConsumerCreditNotProcessed{
				CanceledOrReturnedAt: increase.F(time.Now()),
				CreditExpectedAt:     increase.F(time.Now()),
			}),
			ConsumerDamagedOrDefectiveMerchandise: increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandise{
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted),
				OrderAndIssueExplanation:    increase.F("x"),
				ReceivedAt:                  increase.F(time.Now()),
				ReturnOutcome:               increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned),
				NotReturned:                 increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseNotReturned{}),
				ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttempted{
					AttemptExplanation:     increase.F("x"),
					AttemptReason:          increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
					AttemptedAt:            increase.F(time.Now()),
					MerchandiseDisposition: increase.F("x"),
				}),
				Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturned{
					ReturnMethod:             increase.F(increase.CardDisputeNewParamsVisaConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl),
					ReturnedAt:               increase.F(time.Now()),
					MerchantReceivedReturnAt: increase.F(time.Now()),
					OtherExplanation:         increase.F("x"),
					TrackingNumber:           increase.F("x"),
				}),
			}),
			ConsumerMerchandiseMisrepresentation: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentation{
				MerchantResolutionAttempted:  increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted),
				MisrepresentationExplanation: increase.F("x"),
				PurchaseExplanation:          increase.F("x"),
				ReceivedAt:                   increase.F(time.Now()),
				ReturnOutcome:                increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned),
				NotReturned:                  increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationNotReturned{}),
				ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttempted{
					AttemptExplanation:     increase.F("x"),
					AttemptReason:          increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding),
					AttemptedAt:            increase.F(time.Now()),
					MerchandiseDisposition: increase.F("x"),
				}),
				Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturned{
					ReturnMethod:             increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl),
					ReturnedAt:               increase.F(time.Now()),
					MerchantReceivedReturnAt: increase.F(time.Now()),
					OtherExplanation:         increase.F("x"),
					TrackingNumber:           increase.F("x"),
				}),
			}),
			ConsumerMerchandiseNotAsDescribed: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribed{
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted),
				ReceivedAt:                  increase.F(time.Now()),
				ReturnOutcome:               increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnOutcomeReturned),
				ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttempted{
					AttemptExplanation:     increase.F("x"),
					AttemptReason:          increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding),
					AttemptedAt:            increase.F(time.Now()),
					MerchandiseDisposition: increase.F("x"),
				}),
				Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturned{
					ReturnMethod:             increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl),
					ReturnedAt:               increase.F(time.Now()),
					MerchantReceivedReturnAt: increase.F(time.Now()),
					OtherExplanation:         increase.F("x"),
					TrackingNumber:           increase.F("x"),
				}),
			}),
			ConsumerMerchandiseNotReceived: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceived{
				CancellationOutcome:         increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt),
				DeliveryIssue:               increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveryIssueDelayed),
				LastExpectedReceiptAt:       increase.F(time.Now()),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted),
				PurchaseInfoAndExplanation:  increase.F("x"),
				CardholderCancellationPriorToExpectedReceipt: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt{
					CanceledAt: increase.F(time.Now()),
					Reason:     increase.F("x"),
				}),
				Delayed: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayed{
					Explanation:   increase.F("x"),
					ReturnOutcome: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned),
					NotReturned:   increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedNotReturned{}),
					ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturnAttempted{
						AttemptedAt: increase.F(time.Now()),
					}),
					Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDelayedReturned{
						MerchantReceivedReturnAt: increase.F(time.Now()),
						ReturnedAt:               increase.F(time.Now()),
					}),
				}),
				DeliveredToWrongLocation: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedDeliveredToWrongLocation{
					AgreedLocation: increase.F("x"),
				}),
				MerchantCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedMerchantCancellation{
					CanceledAt: increase.F(time.Now()),
				}),
				NoCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerMerchandiseNotReceivedNoCancellation{}),
			}),
			ConsumerNonReceiptOfCash: increase.F(increase.CardDisputeNewParamsVisaConsumerNonReceiptOfCash{}),
			ConsumerOriginalCreditTransactionNotAccepted: increase.F(increase.CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAccepted{
				Explanation: increase.F("x"),
				Reason:      increase.F(increase.CardDisputeNewParamsVisaConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation),
			}),
			ConsumerQualityMerchandise: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandise{
				ExpectedAt:                  increase.F(time.Now()),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted),
				PurchaseInfoAndQualityIssue: increase.F("x"),
				ReceivedAt:                  increase.F(time.Now()),
				ReturnOutcome:               increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnOutcomeNotReturned),
				NotReturned:                 increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseNotReturned{}),
				OngoingNegotiations: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseOngoingNegotiations{
					Explanation:           increase.F("x"),
					IssuerFirstNotifiedAt: increase.F(time.Now()),
					StartedAt:             increase.F(time.Now()),
				}),
				ReturnAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttempted{
					AttemptExplanation:     increase.F("x"),
					AttemptReason:          increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
					AttemptedAt:            increase.F(time.Now()),
					MerchandiseDisposition: increase.F("x"),
				}),
				Returned: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseReturned{
					ReturnMethod:             increase.F(increase.CardDisputeNewParamsVisaConsumerQualityMerchandiseReturnedReturnMethodDhl),
					ReturnedAt:               increase.F(time.Now()),
					MerchantReceivedReturnAt: increase.F(time.Now()),
					OtherExplanation:         increase.F("x"),
					TrackingNumber:           increase.F("x"),
				}),
			}),
			ConsumerQualityServices: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServices{
				CardholderCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellation{
					AcceptedByMerchant: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted),
					CanceledAt:         increase.F(time.Now()),
					Reason:             increase.F("x"),
				}),
				NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated),
				PurchaseInfoAndQualityIssue:                                       increase.F("x"),
				ServicesReceivedAt:                                                increase.F(time.Now()),
				CardholderPaidToHaveWorkRedone:                                    increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone),
				OngoingNegotiations: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesOngoingNegotiations{
					Explanation:           increase.F("x"),
					IssuerFirstNotifiedAt: increase.F(time.Now()),
					StartedAt:             increase.F(time.Now()),
				}),
				RestaurantFoodRelated: increase.F(increase.CardDisputeNewParamsVisaConsumerQualityServicesRestaurantFoodRelatedNotRelated),
			}),
			ConsumerServicesMisrepresentation: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesMisrepresentation{
				CardholderCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellation{
					AcceptedByMerchant: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted),
					CanceledAt:         increase.F(time.Now()),
					Reason:             increase.F("x"),
				}),
				MerchantResolutionAttempted:  increase.F(increase.CardDisputeNewParamsVisaConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted),
				MisrepresentationExplanation: increase.F("x"),
				PurchaseExplanation:          increase.F("x"),
				ReceivedAt:                   increase.F(time.Now()),
			}),
			ConsumerServicesNotAsDescribed: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotAsDescribed{
				CardholderCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellation{
					AcceptedByMerchant: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted),
					CanceledAt:         increase.F(time.Now()),
					Reason:             increase.F("x"),
				}),
				Explanation:                 increase.F("x"),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted),
				ReceivedAt:                  increase.F(time.Now()),
			}),
			ConsumerServicesNotReceived: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceived{
				CancellationOutcome:         increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt),
				LastExpectedReceiptAt:       increase.F(time.Now()),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted),
				PurchaseInfoAndExplanation:  increase.F("x"),
				CardholderCancellationPriorToExpectedReceipt: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt{
					CanceledAt: increase.F(time.Now()),
					Reason:     increase.F("x"),
				}),
				MerchantCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceivedMerchantCancellation{
					CanceledAt: increase.F(time.Now()),
				}),
				NoCancellation: increase.F(increase.CardDisputeNewParamsVisaConsumerServicesNotReceivedNoCancellation{}),
			}),
			Fraud: increase.F(increase.CardDisputeNewParamsVisaFraud{
				FraudType: increase.F(increase.CardDisputeNewParamsVisaFraudFraudTypeAccountOrCredentialsTakeover),
			}),
			ProcessingError: increase.F(increase.CardDisputeNewParamsVisaProcessingError{
				ErrorReason:                 increase.F(increase.CardDisputeNewParamsVisaProcessingErrorErrorReasonDuplicateTransaction),
				MerchantResolutionAttempted: increase.F(increase.CardDisputeNewParamsVisaProcessingErrorMerchantResolutionAttemptedAttempted),
				DuplicateTransaction: increase.F(increase.CardDisputeNewParamsVisaProcessingErrorDuplicateTransaction{
					OtherTransactionID: increase.F("x"),
				}),
				IncorrectAmount: increase.F(increase.CardDisputeNewParamsVisaProcessingErrorIncorrectAmount{
					ExpectedAmount: increase.F(int64(0)),
				}),
				PaidByOtherMeans: increase.F(increase.CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeans{
					OtherFormOfPaymentEvidence: increase.F(increase.CardDisputeNewParamsVisaProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck),
					OtherTransactionID:         increase.F("x"),
				}),
			}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardDisputes.Get(context.TODO(), "card_dispute_h9sc95nbl1cgltpp7men")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardDisputes.List(context.TODO(), increase.CardDisputeListParams{
		CreatedAt: increase.F(increase.CardDisputeListParamsCreatedAt{
			After:      increase.F(time.Now()),
			Before:     increase.F(time.Now()),
			OnOrAfter:  increase.F(time.Now()),
			OnOrBefore: increase.F(time.Now()),
		}),
		Cursor:         increase.F("cursor"),
		IdempotencyKey: increase.F("x"),
		Limit:          increase.F(int64(1)),
		Status: increase.F(increase.CardDisputeListParamsStatus{
			In: increase.F([]increase.CardDisputeListParamsStatusIn{increase.CardDisputeListParamsStatusInUserSubmissionRequired}),
		}),
	})
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeSubmitUserSubmissionWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardDisputes.SubmitUserSubmission(
		context.TODO(),
		"card_dispute_h9sc95nbl1cgltpp7men",
		increase.CardDisputeSubmitUserSubmissionParams{
			Network: increase.F(increase.CardDisputeSubmitUserSubmissionParamsNetworkVisa),
			Amount:  increase.F(int64(1)),
			AttachmentFiles: increase.F([]increase.CardDisputeSubmitUserSubmissionParamsAttachmentFile{{
				FileID: increase.F("file_id"),
			}}),
			Explanation: increase.F("x"),
			Visa: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisa{
				Category: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaCategoryMerchantPrearbitrationDecline),
				Chargeback: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargeback{
					Category: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackCategoryAuthorization),
					Authorization: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorization{
						AccountStatus: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackAuthorizationAccountStatusAccountClosed),
					}),
					ConsumerCanceledMerchandise: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandise{
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseMerchantResolutionAttemptedAttempted),
						PurchaseExplanation:         increase.F("x"),
						ReceivedOrExpectedAt:        increase.F(time.Now()),
						ReturnOutcome:               increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnOutcomeNotReturned),
						CardholderCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellation{
							CanceledAt:                 increase.F(time.Now()),
							CanceledPriorToShipDate:    increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCanceledPriorToShipDateCanceledPriorToShipDate),
							CancellationPolicyProvided: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseCardholderCancellationCancellationPolicyProvidedNotProvided),
							Reason:                     increase.F("x"),
						}),
						NotReturned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseNotReturned{}),
						ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttempted{
							AttemptExplanation:     increase.F("x"),
							AttemptReason:          increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
							AttemptedAt:            increase.F(time.Now()),
							MerchandiseDisposition: increase.F("x"),
						}),
						Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturned{
							ReturnMethod:             increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledMerchandiseReturnedReturnMethodDhl),
							ReturnedAt:               increase.F(time.Now()),
							MerchantReceivedReturnAt: increase.F(time.Now()),
							OtherExplanation:         increase.F("x"),
							TrackingNumber:           increase.F("x"),
						}),
					}),
					ConsumerCanceledRecurringTransaction: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransaction{
						CancellationTarget: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionCancellationTargetAccount),
						MerchantContactMethods: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledRecurringTransactionMerchantContactMethods{
							ApplicationName:       increase.F("x"),
							CallCenterPhoneNumber: increase.F("x"),
							EmailAddress:          increase.F("x"),
							InPersonAddress:       increase.F("x"),
							MailingAddress:        increase.F("x"),
							TextPhoneNumber:       increase.F("x"),
						}),
						TransactionOrAccountCanceledAt: increase.F(time.Now()),
						OtherFormOfPaymentExplanation:  increase.F("x"),
					}),
					ConsumerCanceledServices: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServices{
						CardholderCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellation{
							CanceledAt:                 increase.F(time.Now()),
							CancellationPolicyProvided: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesCardholderCancellationCancellationPolicyProvidedNotProvided),
							Reason:                     increase.F("x"),
						}),
						ContractedAt:                increase.F(time.Now()),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesMerchantResolutionAttemptedAttempted),
						PurchaseExplanation:         increase.F("x"),
						ServiceType:                 increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesServiceTypeGuaranteedReservation),
						GuaranteedReservation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservation{
							Explanation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesGuaranteedReservationExplanationCardholderCanceledPriorToService),
						}),
						Other:     increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesOther{}),
						Timeshare: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCanceledServicesTimeshare{}),
					}),
					ConsumerCounterfeitMerchandise: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCounterfeitMerchandise{
						CounterfeitExplanation: increase.F("x"),
						DispositionExplanation: increase.F("x"),
						OrderExplanation:       increase.F("x"),
						ReceivedAt:             increase.F(time.Now()),
					}),
					ConsumerCreditNotProcessed: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerCreditNotProcessed{
						CanceledOrReturnedAt: increase.F(time.Now()),
						CreditExpectedAt:     increase.F(time.Now()),
					}),
					ConsumerDamagedOrDefectiveMerchandise: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandise{
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseMerchantResolutionAttemptedAttempted),
						OrderAndIssueExplanation:    increase.F("x"),
						ReceivedAt:                  increase.F(time.Now()),
						ReturnOutcome:               increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnOutcomeNotReturned),
						NotReturned:                 increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseNotReturned{}),
						ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttempted{
							AttemptExplanation:     increase.F("x"),
							AttemptReason:          increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
							AttemptedAt:            increase.F(time.Now()),
							MerchandiseDisposition: increase.F("x"),
						}),
						Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturned{
							ReturnMethod:             increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerDamagedOrDefectiveMerchandiseReturnedReturnMethodDhl),
							ReturnedAt:               increase.F(time.Now()),
							MerchantReceivedReturnAt: increase.F(time.Now()),
							OtherExplanation:         increase.F("x"),
							TrackingNumber:           increase.F("x"),
						}),
					}),
					ConsumerMerchandiseMisrepresentation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentation{
						MerchantResolutionAttempted:  increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationMerchantResolutionAttemptedAttempted),
						MisrepresentationExplanation: increase.F("x"),
						PurchaseExplanation:          increase.F("x"),
						ReceivedAt:                   increase.F(time.Now()),
						ReturnOutcome:                increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnOutcomeNotReturned),
						NotReturned:                  increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationNotReturned{}),
						ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttempted{
							AttemptExplanation:     increase.F("x"),
							AttemptReason:          increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnAttemptedAttemptReasonMerchantNotResponding),
							AttemptedAt:            increase.F(time.Now()),
							MerchandiseDisposition: increase.F("x"),
						}),
						Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturned{
							ReturnMethod:             increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseMisrepresentationReturnedReturnMethodDhl),
							ReturnedAt:               increase.F(time.Now()),
							MerchantReceivedReturnAt: increase.F(time.Now()),
							OtherExplanation:         increase.F("x"),
							TrackingNumber:           increase.F("x"),
						}),
					}),
					ConsumerMerchandiseNotAsDescribed: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribed{
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedMerchantResolutionAttemptedAttempted),
						ReceivedAt:                  increase.F(time.Now()),
						ReturnOutcome:               increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnOutcomeReturned),
						ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttempted{
							AttemptExplanation:     increase.F("x"),
							AttemptReason:          increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnAttemptedAttemptReasonMerchantNotResponding),
							AttemptedAt:            increase.F(time.Now()),
							MerchandiseDisposition: increase.F("x"),
						}),
						Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturned{
							ReturnMethod:             increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotAsDescribedReturnedReturnMethodDhl),
							ReturnedAt:               increase.F(time.Now()),
							MerchantReceivedReturnAt: increase.F(time.Now()),
							OtherExplanation:         increase.F("x"),
							TrackingNumber:           increase.F("x"),
						}),
					}),
					ConsumerMerchandiseNotReceived: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceived{
						CancellationOutcome:         increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt),
						DeliveryIssue:               increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveryIssueDelayed),
						LastExpectedReceiptAt:       increase.F(time.Now()),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantResolutionAttemptedAttempted),
						PurchaseInfoAndExplanation:  increase.F("x"),
						CardholderCancellationPriorToExpectedReceipt: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedCardholderCancellationPriorToExpectedReceipt{
							CanceledAt: increase.F(time.Now()),
							Reason:     increase.F("x"),
						}),
						Delayed: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayed{
							Explanation:   increase.F("x"),
							ReturnOutcome: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnOutcomeNotReturned),
							NotReturned:   increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedNotReturned{}),
							ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturnAttempted{
								AttemptedAt: increase.F(time.Now()),
							}),
							Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDelayedReturned{
								MerchantReceivedReturnAt: increase.F(time.Now()),
								ReturnedAt:               increase.F(time.Now()),
							}),
						}),
						DeliveredToWrongLocation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedDeliveredToWrongLocation{
							AgreedLocation: increase.F("x"),
						}),
						MerchantCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedMerchantCancellation{
							CanceledAt: increase.F(time.Now()),
						}),
						NoCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerMerchandiseNotReceivedNoCancellation{}),
					}),
					ConsumerNonReceiptOfCash: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerNonReceiptOfCash{}),
					ConsumerOriginalCreditTransactionNotAccepted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAccepted{
						Explanation: increase.F("x"),
						Reason:      increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerOriginalCreditTransactionNotAcceptedReasonProhibitedByLocalLawsOrRegulation),
					}),
					ConsumerQualityMerchandise: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandise{
						ExpectedAt:                  increase.F(time.Now()),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseMerchantResolutionAttemptedAttempted),
						PurchaseInfoAndQualityIssue: increase.F("x"),
						ReceivedAt:                  increase.F(time.Now()),
						ReturnOutcome:               increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnOutcomeNotReturned),
						NotReturned:                 increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseNotReturned{}),
						OngoingNegotiations: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseOngoingNegotiations{
							Explanation:           increase.F("x"),
							IssuerFirstNotifiedAt: increase.F(time.Now()),
							StartedAt:             increase.F(time.Now()),
						}),
						ReturnAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttempted{
							AttemptExplanation:     increase.F("x"),
							AttemptReason:          increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnAttemptedAttemptReasonMerchantNotResponding),
							AttemptedAt:            increase.F(time.Now()),
							MerchandiseDisposition: increase.F("x"),
						}),
						Returned: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturned{
							ReturnMethod:             increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityMerchandiseReturnedReturnMethodDhl),
							ReturnedAt:               increase.F(time.Now()),
							MerchantReceivedReturnAt: increase.F(time.Now()),
							OtherExplanation:         increase.F("x"),
							TrackingNumber:           increase.F("x"),
						}),
					}),
					ConsumerQualityServices: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServices{
						CardholderCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellation{
							AcceptedByMerchant: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderCancellationAcceptedByMerchantAccepted),
							CanceledAt:         increase.F(time.Now()),
							Reason:             increase.F("x"),
						}),
						NonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescription: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesNonFiatCurrencyOrNonFungibleTokenRelatedAndNotMatchingDescriptionNotRelated),
						PurchaseInfoAndQualityIssue:                                       increase.F("x"),
						ServicesReceivedAt:                                                increase.F(time.Now()),
						CardholderPaidToHaveWorkRedone:                                    increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesCardholderPaidToHaveWorkRedoneDidNotPayToHaveWorkRedone),
						OngoingNegotiations: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesOngoingNegotiations{
							Explanation:           increase.F("x"),
							IssuerFirstNotifiedAt: increase.F(time.Now()),
							StartedAt:             increase.F(time.Now()),
						}),
						RestaurantFoodRelated: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerQualityServicesRestaurantFoodRelatedNotRelated),
					}),
					ConsumerServicesMisrepresentation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentation{
						CardholderCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellation{
							AcceptedByMerchant: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationCardholderCancellationAcceptedByMerchantAccepted),
							CanceledAt:         increase.F(time.Now()),
							Reason:             increase.F("x"),
						}),
						MerchantResolutionAttempted:  increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesMisrepresentationMerchantResolutionAttemptedAttempted),
						MisrepresentationExplanation: increase.F("x"),
						PurchaseExplanation:          increase.F("x"),
						ReceivedAt:                   increase.F(time.Now()),
					}),
					ConsumerServicesNotAsDescribed: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribed{
						CardholderCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellation{
							AcceptedByMerchant: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedCardholderCancellationAcceptedByMerchantAccepted),
							CanceledAt:         increase.F(time.Now()),
							Reason:             increase.F("x"),
						}),
						Explanation:                 increase.F("x"),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotAsDescribedMerchantResolutionAttemptedAttempted),
						ReceivedAt:                  increase.F(time.Now()),
					}),
					ConsumerServicesNotReceived: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceived{
						CancellationOutcome:         increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCancellationOutcomeCardholderCancellationPriorToExpectedReceipt),
						LastExpectedReceiptAt:       increase.F(time.Now()),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantResolutionAttemptedAttempted),
						PurchaseInfoAndExplanation:  increase.F("x"),
						CardholderCancellationPriorToExpectedReceipt: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedCardholderCancellationPriorToExpectedReceipt{
							CanceledAt: increase.F(time.Now()),
							Reason:     increase.F("x"),
						}),
						MerchantCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedMerchantCancellation{
							CanceledAt: increase.F(time.Now()),
						}),
						NoCancellation: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackConsumerServicesNotReceivedNoCancellation{}),
					}),
					Fraud: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackFraud{
						FraudType: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackFraudFraudTypeAccountOrCredentialsTakeover),
					}),
					ProcessingError: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingError{
						ErrorReason:                 increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorErrorReasonDuplicateTransaction),
						MerchantResolutionAttempted: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorMerchantResolutionAttemptedAttempted),
						DuplicateTransaction: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorDuplicateTransaction{
							OtherTransactionID: increase.F("x"),
						}),
						IncorrectAmount: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorIncorrectAmount{
							ExpectedAmount: increase.F(int64(0)),
						}),
						PaidByOtherMeans: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeans{
							OtherFormOfPaymentEvidence: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaChargebackProcessingErrorPaidByOtherMeansOtherFormOfPaymentEvidenceCanceledCheck),
							OtherTransactionID:         increase.F("x"),
						}),
					}),
				}),
				MerchantPrearbitrationDecline: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaMerchantPrearbitrationDecline{
					Reason: increase.F("The pre-arbitration received from the merchantdoes not explain how they obtained permission to charge the card."),
				}),
				UserPrearbitration: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitration{
					Reason: increase.F("x"),
					CategoryChange: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChange{
						Category: increase.F(increase.CardDisputeSubmitUserSubmissionParamsVisaUserPrearbitrationCategoryChangeCategoryAuthorization),
						Reason:   increase.F("x"),
					}),
				}),
			}),
		},
	)
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCardDisputeWithdraw(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := increase.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.CardDisputes.Withdraw(context.TODO(), "card_dispute_h9sc95nbl1cgltpp7men")
	if err != nil {
		var apierr *increase.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
