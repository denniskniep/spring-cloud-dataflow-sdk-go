package tasks

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsExternalRequestBuilder builds and executes requests for operations under \tasks\executions\external
type ExecutionsExternalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ByExternalExecutionId gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.tasks.executions.external.item collection
func (m *ExecutionsExternalRequestBuilder) ByExternalExecutionId(externalExecutionId string)(*ExecutionsExternalWithExternalExecutionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if externalExecutionId != "" {
        urlTplParams["externalExecutionId"] = externalExecutionId
    }
    return NewExecutionsExternalWithExternalExecutionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExecutionsExternalRequestBuilderInternal instantiates a new ExternalRequestBuilder and sets the default values.
func NewExecutionsExternalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExternalRequestBuilder) {
    m := &ExecutionsExternalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tasks/executions/external", pathParameters),
    }
    return m
}
// NewExecutionsExternalRequestBuilder instantiates a new ExternalRequestBuilder and sets the default values.
func NewExecutionsExternalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExternalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsExternalRequestBuilderInternal(urlParams, requestAdapter)
}
