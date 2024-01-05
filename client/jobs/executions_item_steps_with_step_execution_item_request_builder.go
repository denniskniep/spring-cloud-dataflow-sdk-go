package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsItemStepsWithStepExecutionItemRequestBuilder builds and executes requests for operations under \jobs\executions\{jobExecutionId}\steps\{stepExecutionId}
type ExecutionsItemStepsWithStepExecutionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetQueryParameters 
type ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetQueryParameters struct {
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetQueryParameters
}
// NewExecutionsItemStepsWithStepExecutionItemRequestBuilderInternal instantiates a new WithStepExecutionItemRequestBuilder and sets the default values.
func NewExecutionsItemStepsWithStepExecutionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsWithStepExecutionItemRequestBuilder) {
    m := &ExecutionsItemStepsWithStepExecutionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/executions/{jobExecutionId}/steps/{stepExecutionId}{?schemaTarget*}", pathParameters),
    }
    return m
}
// NewExecutionsItemStepsWithStepExecutionItemRequestBuilder instantiates a new WithStepExecutionItemRequestBuilder and sets the default values.
func NewExecutionsItemStepsWithStepExecutionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsWithStepExecutionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsItemStepsWithStepExecutionItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsItemStepsWithStepExecutionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Progress the progress property
func (m *ExecutionsItemStepsWithStepExecutionItemRequestBuilder) Progress()(*ExecutionsItemStepsItemProgressRequestBuilder) {
    return NewExecutionsItemStepsItemProgressRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *ExecutionsItemStepsWithStepExecutionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsItemStepsWithStepExecutionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "*/*")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ExecutionsItemStepsWithStepExecutionItemRequestBuilder) WithUrl(rawUrl string)(*ExecutionsItemStepsWithStepExecutionItemRequestBuilder) {
    return NewExecutionsItemStepsWithStepExecutionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
