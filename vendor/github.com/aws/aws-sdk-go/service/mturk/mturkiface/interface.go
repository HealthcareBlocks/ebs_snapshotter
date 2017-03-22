// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mturkiface provides an interface to enable mocking the Amazon Mechanical Turk service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mturkiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mturk"
)

// MTurkAPI provides an interface to enable mocking the
// mturk.MTurk service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Mechanical Turk.
//    func myFunc(svc mturkiface.MTurkAPI) bool {
//        // Make svc.AcceptQualificationRequest request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mturk.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMTurkClient struct {
//        mturkiface.MTurkAPI
//    }
//    func (m *mockMTurkClient) AcceptQualificationRequest(input *mturk.AcceptQualificationRequestInput) (*mturk.AcceptQualificationRequestOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMTurkClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MTurkAPI interface {
	AcceptQualificationRequest(*mturk.AcceptQualificationRequestInput) (*mturk.AcceptQualificationRequestOutput, error)
	AcceptQualificationRequestWithContext(aws.Context, *mturk.AcceptQualificationRequestInput, ...request.Option) (*mturk.AcceptQualificationRequestOutput, error)
	AcceptQualificationRequestRequest(*mturk.AcceptQualificationRequestInput) (*request.Request, *mturk.AcceptQualificationRequestOutput)

	ApproveAssignment(*mturk.ApproveAssignmentInput) (*mturk.ApproveAssignmentOutput, error)
	ApproveAssignmentWithContext(aws.Context, *mturk.ApproveAssignmentInput, ...request.Option) (*mturk.ApproveAssignmentOutput, error)
	ApproveAssignmentRequest(*mturk.ApproveAssignmentInput) (*request.Request, *mturk.ApproveAssignmentOutput)

	AssociateQualificationWithWorker(*mturk.AssociateQualificationWithWorkerInput) (*mturk.AssociateQualificationWithWorkerOutput, error)
	AssociateQualificationWithWorkerWithContext(aws.Context, *mturk.AssociateQualificationWithWorkerInput, ...request.Option) (*mturk.AssociateQualificationWithWorkerOutput, error)
	AssociateQualificationWithWorkerRequest(*mturk.AssociateQualificationWithWorkerInput) (*request.Request, *mturk.AssociateQualificationWithWorkerOutput)

	CreateAdditionalAssignmentsForHIT(*mturk.CreateAdditionalAssignmentsForHITInput) (*mturk.CreateAdditionalAssignmentsForHITOutput, error)
	CreateAdditionalAssignmentsForHITWithContext(aws.Context, *mturk.CreateAdditionalAssignmentsForHITInput, ...request.Option) (*mturk.CreateAdditionalAssignmentsForHITOutput, error)
	CreateAdditionalAssignmentsForHITRequest(*mturk.CreateAdditionalAssignmentsForHITInput) (*request.Request, *mturk.CreateAdditionalAssignmentsForHITOutput)

	CreateHIT(*mturk.CreateHITInput) (*mturk.CreateHITOutput, error)
	CreateHITWithContext(aws.Context, *mturk.CreateHITInput, ...request.Option) (*mturk.CreateHITOutput, error)
	CreateHITRequest(*mturk.CreateHITInput) (*request.Request, *mturk.CreateHITOutput)

	CreateHITType(*mturk.CreateHITTypeInput) (*mturk.CreateHITTypeOutput, error)
	CreateHITTypeWithContext(aws.Context, *mturk.CreateHITTypeInput, ...request.Option) (*mturk.CreateHITTypeOutput, error)
	CreateHITTypeRequest(*mturk.CreateHITTypeInput) (*request.Request, *mturk.CreateHITTypeOutput)

	CreateHITWithHITType(*mturk.CreateHITWithHITTypeInput) (*mturk.CreateHITWithHITTypeOutput, error)
	CreateHITWithHITTypeWithContext(aws.Context, *mturk.CreateHITWithHITTypeInput, ...request.Option) (*mturk.CreateHITWithHITTypeOutput, error)
	CreateHITWithHITTypeRequest(*mturk.CreateHITWithHITTypeInput) (*request.Request, *mturk.CreateHITWithHITTypeOutput)

	CreateQualificationType(*mturk.CreateQualificationTypeInput) (*mturk.CreateQualificationTypeOutput, error)
	CreateQualificationTypeWithContext(aws.Context, *mturk.CreateQualificationTypeInput, ...request.Option) (*mturk.CreateQualificationTypeOutput, error)
	CreateQualificationTypeRequest(*mturk.CreateQualificationTypeInput) (*request.Request, *mturk.CreateQualificationTypeOutput)

	CreateWorkerBlock(*mturk.CreateWorkerBlockInput) (*mturk.CreateWorkerBlockOutput, error)
	CreateWorkerBlockWithContext(aws.Context, *mturk.CreateWorkerBlockInput, ...request.Option) (*mturk.CreateWorkerBlockOutput, error)
	CreateWorkerBlockRequest(*mturk.CreateWorkerBlockInput) (*request.Request, *mturk.CreateWorkerBlockOutput)

	DeleteHIT(*mturk.DeleteHITInput) (*mturk.DeleteHITOutput, error)
	DeleteHITWithContext(aws.Context, *mturk.DeleteHITInput, ...request.Option) (*mturk.DeleteHITOutput, error)
	DeleteHITRequest(*mturk.DeleteHITInput) (*request.Request, *mturk.DeleteHITOutput)

	DeleteQualificationType(*mturk.DeleteQualificationTypeInput) (*mturk.DeleteQualificationTypeOutput, error)
	DeleteQualificationTypeWithContext(aws.Context, *mturk.DeleteQualificationTypeInput, ...request.Option) (*mturk.DeleteQualificationTypeOutput, error)
	DeleteQualificationTypeRequest(*mturk.DeleteQualificationTypeInput) (*request.Request, *mturk.DeleteQualificationTypeOutput)

	DeleteWorkerBlock(*mturk.DeleteWorkerBlockInput) (*mturk.DeleteWorkerBlockOutput, error)
	DeleteWorkerBlockWithContext(aws.Context, *mturk.DeleteWorkerBlockInput, ...request.Option) (*mturk.DeleteWorkerBlockOutput, error)
	DeleteWorkerBlockRequest(*mturk.DeleteWorkerBlockInput) (*request.Request, *mturk.DeleteWorkerBlockOutput)

	DisassociateQualificationFromWorker(*mturk.DisassociateQualificationFromWorkerInput) (*mturk.DisassociateQualificationFromWorkerOutput, error)
	DisassociateQualificationFromWorkerWithContext(aws.Context, *mturk.DisassociateQualificationFromWorkerInput, ...request.Option) (*mturk.DisassociateQualificationFromWorkerOutput, error)
	DisassociateQualificationFromWorkerRequest(*mturk.DisassociateQualificationFromWorkerInput) (*request.Request, *mturk.DisassociateQualificationFromWorkerOutput)

	GetAccountBalance(*mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error)
	GetAccountBalanceWithContext(aws.Context, *mturk.GetAccountBalanceInput, ...request.Option) (*mturk.GetAccountBalanceOutput, error)
	GetAccountBalanceRequest(*mturk.GetAccountBalanceInput) (*request.Request, *mturk.GetAccountBalanceOutput)

	GetAssignment(*mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error)
	GetAssignmentWithContext(aws.Context, *mturk.GetAssignmentInput, ...request.Option) (*mturk.GetAssignmentOutput, error)
	GetAssignmentRequest(*mturk.GetAssignmentInput) (*request.Request, *mturk.GetAssignmentOutput)

	GetFileUploadURL(*mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error)
	GetFileUploadURLWithContext(aws.Context, *mturk.GetFileUploadURLInput, ...request.Option) (*mturk.GetFileUploadURLOutput, error)
	GetFileUploadURLRequest(*mturk.GetFileUploadURLInput) (*request.Request, *mturk.GetFileUploadURLOutput)

	GetHIT(*mturk.GetHITInput) (*mturk.GetHITOutput, error)
	GetHITWithContext(aws.Context, *mturk.GetHITInput, ...request.Option) (*mturk.GetHITOutput, error)
	GetHITRequest(*mturk.GetHITInput) (*request.Request, *mturk.GetHITOutput)

	GetQualificationScore(*mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error)
	GetQualificationScoreWithContext(aws.Context, *mturk.GetQualificationScoreInput, ...request.Option) (*mturk.GetQualificationScoreOutput, error)
	GetQualificationScoreRequest(*mturk.GetQualificationScoreInput) (*request.Request, *mturk.GetQualificationScoreOutput)

	GetQualificationType(*mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error)
	GetQualificationTypeWithContext(aws.Context, *mturk.GetQualificationTypeInput, ...request.Option) (*mturk.GetQualificationTypeOutput, error)
	GetQualificationTypeRequest(*mturk.GetQualificationTypeInput) (*request.Request, *mturk.GetQualificationTypeOutput)

	ListAssignmentsForHIT(*mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error)
	ListAssignmentsForHITWithContext(aws.Context, *mturk.ListAssignmentsForHITInput, ...request.Option) (*mturk.ListAssignmentsForHITOutput, error)
	ListAssignmentsForHITRequest(*mturk.ListAssignmentsForHITInput) (*request.Request, *mturk.ListAssignmentsForHITOutput)

	ListAssignmentsForHITPages(*mturk.ListAssignmentsForHITInput, func(*mturk.ListAssignmentsForHITOutput, bool) bool) error
	ListAssignmentsForHITPagesWithContext(aws.Context, *mturk.ListAssignmentsForHITInput, func(*mturk.ListAssignmentsForHITOutput, bool) bool, ...request.Option) error

	ListBonusPayments(*mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error)
	ListBonusPaymentsWithContext(aws.Context, *mturk.ListBonusPaymentsInput, ...request.Option) (*mturk.ListBonusPaymentsOutput, error)
	ListBonusPaymentsRequest(*mturk.ListBonusPaymentsInput) (*request.Request, *mturk.ListBonusPaymentsOutput)

	ListBonusPaymentsPages(*mturk.ListBonusPaymentsInput, func(*mturk.ListBonusPaymentsOutput, bool) bool) error
	ListBonusPaymentsPagesWithContext(aws.Context, *mturk.ListBonusPaymentsInput, func(*mturk.ListBonusPaymentsOutput, bool) bool, ...request.Option) error

	ListHITs(*mturk.ListHITsInput) (*mturk.ListHITsOutput, error)
	ListHITsWithContext(aws.Context, *mturk.ListHITsInput, ...request.Option) (*mturk.ListHITsOutput, error)
	ListHITsRequest(*mturk.ListHITsInput) (*request.Request, *mturk.ListHITsOutput)

	ListHITsPages(*mturk.ListHITsInput, func(*mturk.ListHITsOutput, bool) bool) error
	ListHITsPagesWithContext(aws.Context, *mturk.ListHITsInput, func(*mturk.ListHITsOutput, bool) bool, ...request.Option) error

	ListHITsForQualificationType(*mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error)
	ListHITsForQualificationTypeWithContext(aws.Context, *mturk.ListHITsForQualificationTypeInput, ...request.Option) (*mturk.ListHITsForQualificationTypeOutput, error)
	ListHITsForQualificationTypeRequest(*mturk.ListHITsForQualificationTypeInput) (*request.Request, *mturk.ListHITsForQualificationTypeOutput)

	ListHITsForQualificationTypePages(*mturk.ListHITsForQualificationTypeInput, func(*mturk.ListHITsForQualificationTypeOutput, bool) bool) error
	ListHITsForQualificationTypePagesWithContext(aws.Context, *mturk.ListHITsForQualificationTypeInput, func(*mturk.ListHITsForQualificationTypeOutput, bool) bool, ...request.Option) error

	ListQualificationRequests(*mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error)
	ListQualificationRequestsWithContext(aws.Context, *mturk.ListQualificationRequestsInput, ...request.Option) (*mturk.ListQualificationRequestsOutput, error)
	ListQualificationRequestsRequest(*mturk.ListQualificationRequestsInput) (*request.Request, *mturk.ListQualificationRequestsOutput)

	ListQualificationRequestsPages(*mturk.ListQualificationRequestsInput, func(*mturk.ListQualificationRequestsOutput, bool) bool) error
	ListQualificationRequestsPagesWithContext(aws.Context, *mturk.ListQualificationRequestsInput, func(*mturk.ListQualificationRequestsOutput, bool) bool, ...request.Option) error

	ListQualificationTypes(*mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error)
	ListQualificationTypesWithContext(aws.Context, *mturk.ListQualificationTypesInput, ...request.Option) (*mturk.ListQualificationTypesOutput, error)
	ListQualificationTypesRequest(*mturk.ListQualificationTypesInput) (*request.Request, *mturk.ListQualificationTypesOutput)

	ListQualificationTypesPages(*mturk.ListQualificationTypesInput, func(*mturk.ListQualificationTypesOutput, bool) bool) error
	ListQualificationTypesPagesWithContext(aws.Context, *mturk.ListQualificationTypesInput, func(*mturk.ListQualificationTypesOutput, bool) bool, ...request.Option) error

	ListReviewPolicyResultsForHIT(*mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error)
	ListReviewPolicyResultsForHITWithContext(aws.Context, *mturk.ListReviewPolicyResultsForHITInput, ...request.Option) (*mturk.ListReviewPolicyResultsForHITOutput, error)
	ListReviewPolicyResultsForHITRequest(*mturk.ListReviewPolicyResultsForHITInput) (*request.Request, *mturk.ListReviewPolicyResultsForHITOutput)

	ListReviewPolicyResultsForHITPages(*mturk.ListReviewPolicyResultsForHITInput, func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool) error
	ListReviewPolicyResultsForHITPagesWithContext(aws.Context, *mturk.ListReviewPolicyResultsForHITInput, func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool, ...request.Option) error

	ListReviewableHITs(*mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error)
	ListReviewableHITsWithContext(aws.Context, *mturk.ListReviewableHITsInput, ...request.Option) (*mturk.ListReviewableHITsOutput, error)
	ListReviewableHITsRequest(*mturk.ListReviewableHITsInput) (*request.Request, *mturk.ListReviewableHITsOutput)

	ListReviewableHITsPages(*mturk.ListReviewableHITsInput, func(*mturk.ListReviewableHITsOutput, bool) bool) error
	ListReviewableHITsPagesWithContext(aws.Context, *mturk.ListReviewableHITsInput, func(*mturk.ListReviewableHITsOutput, bool) bool, ...request.Option) error

	ListWorkerBlocks(*mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error)
	ListWorkerBlocksWithContext(aws.Context, *mturk.ListWorkerBlocksInput, ...request.Option) (*mturk.ListWorkerBlocksOutput, error)
	ListWorkerBlocksRequest(*mturk.ListWorkerBlocksInput) (*request.Request, *mturk.ListWorkerBlocksOutput)

	ListWorkerBlocksPages(*mturk.ListWorkerBlocksInput, func(*mturk.ListWorkerBlocksOutput, bool) bool) error
	ListWorkerBlocksPagesWithContext(aws.Context, *mturk.ListWorkerBlocksInput, func(*mturk.ListWorkerBlocksOutput, bool) bool, ...request.Option) error

	ListWorkersWithQualificationType(*mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error)
	ListWorkersWithQualificationTypeWithContext(aws.Context, *mturk.ListWorkersWithQualificationTypeInput, ...request.Option) (*mturk.ListWorkersWithQualificationTypeOutput, error)
	ListWorkersWithQualificationTypeRequest(*mturk.ListWorkersWithQualificationTypeInput) (*request.Request, *mturk.ListWorkersWithQualificationTypeOutput)

	ListWorkersWithQualificationTypePages(*mturk.ListWorkersWithQualificationTypeInput, func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool) error
	ListWorkersWithQualificationTypePagesWithContext(aws.Context, *mturk.ListWorkersWithQualificationTypeInput, func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool, ...request.Option) error

	NotifyWorkers(*mturk.NotifyWorkersInput) (*mturk.NotifyWorkersOutput, error)
	NotifyWorkersWithContext(aws.Context, *mturk.NotifyWorkersInput, ...request.Option) (*mturk.NotifyWorkersOutput, error)
	NotifyWorkersRequest(*mturk.NotifyWorkersInput) (*request.Request, *mturk.NotifyWorkersOutput)

	RejectAssignment(*mturk.RejectAssignmentInput) (*mturk.RejectAssignmentOutput, error)
	RejectAssignmentWithContext(aws.Context, *mturk.RejectAssignmentInput, ...request.Option) (*mturk.RejectAssignmentOutput, error)
	RejectAssignmentRequest(*mturk.RejectAssignmentInput) (*request.Request, *mturk.RejectAssignmentOutput)

	RejectQualificationRequest(*mturk.RejectQualificationRequestInput) (*mturk.RejectQualificationRequestOutput, error)
	RejectQualificationRequestWithContext(aws.Context, *mturk.RejectQualificationRequestInput, ...request.Option) (*mturk.RejectQualificationRequestOutput, error)
	RejectQualificationRequestRequest(*mturk.RejectQualificationRequestInput) (*request.Request, *mturk.RejectQualificationRequestOutput)

	SendBonus(*mturk.SendBonusInput) (*mturk.SendBonusOutput, error)
	SendBonusWithContext(aws.Context, *mturk.SendBonusInput, ...request.Option) (*mturk.SendBonusOutput, error)
	SendBonusRequest(*mturk.SendBonusInput) (*request.Request, *mturk.SendBonusOutput)

	SendTestEventNotification(*mturk.SendTestEventNotificationInput) (*mturk.SendTestEventNotificationOutput, error)
	SendTestEventNotificationWithContext(aws.Context, *mturk.SendTestEventNotificationInput, ...request.Option) (*mturk.SendTestEventNotificationOutput, error)
	SendTestEventNotificationRequest(*mturk.SendTestEventNotificationInput) (*request.Request, *mturk.SendTestEventNotificationOutput)

	UpdateExpirationForHIT(*mturk.UpdateExpirationForHITInput) (*mturk.UpdateExpirationForHITOutput, error)
	UpdateExpirationForHITWithContext(aws.Context, *mturk.UpdateExpirationForHITInput, ...request.Option) (*mturk.UpdateExpirationForHITOutput, error)
	UpdateExpirationForHITRequest(*mturk.UpdateExpirationForHITInput) (*request.Request, *mturk.UpdateExpirationForHITOutput)

	UpdateHITReviewStatus(*mturk.UpdateHITReviewStatusInput) (*mturk.UpdateHITReviewStatusOutput, error)
	UpdateHITReviewStatusWithContext(aws.Context, *mturk.UpdateHITReviewStatusInput, ...request.Option) (*mturk.UpdateHITReviewStatusOutput, error)
	UpdateHITReviewStatusRequest(*mturk.UpdateHITReviewStatusInput) (*request.Request, *mturk.UpdateHITReviewStatusOutput)

	UpdateHITTypeOfHIT(*mturk.UpdateHITTypeOfHITInput) (*mturk.UpdateHITTypeOfHITOutput, error)
	UpdateHITTypeOfHITWithContext(aws.Context, *mturk.UpdateHITTypeOfHITInput, ...request.Option) (*mturk.UpdateHITTypeOfHITOutput, error)
	UpdateHITTypeOfHITRequest(*mturk.UpdateHITTypeOfHITInput) (*request.Request, *mturk.UpdateHITTypeOfHITOutput)

	UpdateNotificationSettings(*mturk.UpdateNotificationSettingsInput) (*mturk.UpdateNotificationSettingsOutput, error)
	UpdateNotificationSettingsWithContext(aws.Context, *mturk.UpdateNotificationSettingsInput, ...request.Option) (*mturk.UpdateNotificationSettingsOutput, error)
	UpdateNotificationSettingsRequest(*mturk.UpdateNotificationSettingsInput) (*request.Request, *mturk.UpdateNotificationSettingsOutput)

	UpdateQualificationType(*mturk.UpdateQualificationTypeInput) (*mturk.UpdateQualificationTypeOutput, error)
	UpdateQualificationTypeWithContext(aws.Context, *mturk.UpdateQualificationTypeInput, ...request.Option) (*mturk.UpdateQualificationTypeOutput, error)
	UpdateQualificationTypeRequest(*mturk.UpdateQualificationTypeInput) (*request.Request, *mturk.UpdateQualificationTypeOutput)
}

var _ MTurkAPI = (*mturk.MTurk)(nil)
