// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package codebuildiface provides an interface to enable mocking the AWS CodeBuild service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codebuildiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

// CodeBuildAPI provides an interface to enable mocking the
// codebuild.CodeBuild service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CodeBuild.
//    func myFunc(svc codebuildiface.CodeBuildAPI) bool {
//        // Make svc.BatchGetBuilds request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := codebuild.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeBuildClient struct {
//        codebuildiface.CodeBuildAPI
//    }
//    func (m *mockCodeBuildClient) BatchGetBuilds(input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeBuildClient{}
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
type CodeBuildAPI interface {
	BatchGetBuilds(*codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsWithContext(aws.Context, *codebuild.BatchGetBuildsInput, ...request.Option) (*codebuild.BatchGetBuildsOutput, error)
	BatchGetBuildsRequest(*codebuild.BatchGetBuildsInput) (*request.Request, *codebuild.BatchGetBuildsOutput)

	BatchGetProjects(*codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsWithContext(aws.Context, *codebuild.BatchGetProjectsInput, ...request.Option) (*codebuild.BatchGetProjectsOutput, error)
	BatchGetProjectsRequest(*codebuild.BatchGetProjectsInput) (*request.Request, *codebuild.BatchGetProjectsOutput)

	CreateProject(*codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error)
	CreateProjectWithContext(aws.Context, *codebuild.CreateProjectInput, ...request.Option) (*codebuild.CreateProjectOutput, error)
	CreateProjectRequest(*codebuild.CreateProjectInput) (*request.Request, *codebuild.CreateProjectOutput)

	DeleteProject(*codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectWithContext(aws.Context, *codebuild.DeleteProjectInput, ...request.Option) (*codebuild.DeleteProjectOutput, error)
	DeleteProjectRequest(*codebuild.DeleteProjectInput) (*request.Request, *codebuild.DeleteProjectOutput)

	ListBuilds(*codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error)
	ListBuildsWithContext(aws.Context, *codebuild.ListBuildsInput, ...request.Option) (*codebuild.ListBuildsOutput, error)
	ListBuildsRequest(*codebuild.ListBuildsInput) (*request.Request, *codebuild.ListBuildsOutput)

	ListBuildsForProject(*codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectWithContext(aws.Context, *codebuild.ListBuildsForProjectInput, ...request.Option) (*codebuild.ListBuildsForProjectOutput, error)
	ListBuildsForProjectRequest(*codebuild.ListBuildsForProjectInput) (*request.Request, *codebuild.ListBuildsForProjectOutput)

	ListCuratedEnvironmentImages(*codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesWithContext(aws.Context, *codebuild.ListCuratedEnvironmentImagesInput, ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error)
	ListCuratedEnvironmentImagesRequest(*codebuild.ListCuratedEnvironmentImagesInput) (*request.Request, *codebuild.ListCuratedEnvironmentImagesOutput)

	ListProjects(*codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error)
	ListProjectsWithContext(aws.Context, *codebuild.ListProjectsInput, ...request.Option) (*codebuild.ListProjectsOutput, error)
	ListProjectsRequest(*codebuild.ListProjectsInput) (*request.Request, *codebuild.ListProjectsOutput)

	StartBuild(*codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error)
	StartBuildWithContext(aws.Context, *codebuild.StartBuildInput, ...request.Option) (*codebuild.StartBuildOutput, error)
	StartBuildRequest(*codebuild.StartBuildInput) (*request.Request, *codebuild.StartBuildOutput)

	StopBuild(*codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error)
	StopBuildWithContext(aws.Context, *codebuild.StopBuildInput, ...request.Option) (*codebuild.StopBuildOutput, error)
	StopBuildRequest(*codebuild.StopBuildInput) (*request.Request, *codebuild.StopBuildOutput)

	UpdateProject(*codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectWithContext(aws.Context, *codebuild.UpdateProjectInput, ...request.Option) (*codebuild.UpdateProjectOutput, error)
	UpdateProjectRequest(*codebuild.UpdateProjectInput) (*request.Request, *codebuild.UpdateProjectOutput)
}

var _ CodeBuildAPI = (*codebuild.CodeBuild)(nil)
