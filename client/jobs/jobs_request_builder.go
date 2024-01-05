package jobs

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// JobsRequestBuilder builds and executes requests for operations under \jobs
type JobsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewJobsRequestBuilderInternal instantiates a new JobsRequestBuilder and sets the default values.
func NewJobsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JobsRequestBuilder) {
    m := &JobsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs", pathParameters),
    }
    return m
}
// NewJobsRequestBuilder instantiates a new JobsRequestBuilder and sets the default values.
func NewJobsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*JobsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewJobsRequestBuilderInternal(urlParams, requestAdapter)
}
// Executions the executions property
func (m *JobsRequestBuilder) Executions()(*ExecutionsRequestBuilder) {
    return NewExecutionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Instances the instances property
func (m *JobsRequestBuilder) Instances()(*InstancesRequestBuilder) {
    return NewInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Thinexecutions the thinexecutions property
func (m *JobsRequestBuilder) Thinexecutions()(*ThinexecutionsRequestBuilder) {
    return NewThinexecutionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
