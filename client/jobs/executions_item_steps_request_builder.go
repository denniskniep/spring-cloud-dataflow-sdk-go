package jobs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsItemStepsRequestBuilder builds and executes requests for operations under \jobs\executions\{executions-id}\steps
type ExecutionsItemStepsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsItemStepsRequestBuilderGetQueryParameters 
type ExecutionsItemStepsRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// ExecutionsItemStepsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsItemStepsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsItemStepsRequestBuilderGetQueryParameters
}
// ByStepExecutionId gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.jobs.executions.item.steps.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ExecutionsItemStepsRequestBuilder) ByStepExecutionId(stepExecutionId string)(*ExecutionsItemStepsWithStepExecutionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if stepExecutionId != "" {
        urlTplParams["stepExecutionId"] = stepExecutionId
    }
    return NewExecutionsItemStepsWithStepExecutionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByStepExecutionIdInt64 gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/v2/client.jobs.executions.item.steps.item collection
func (m *ExecutionsItemStepsRequestBuilder) ByStepExecutionIdInt64(stepExecutionId int64)(*ExecutionsItemStepsWithStepExecutionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["stepExecutionId"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(stepExecutionId, 10)
    return NewExecutionsItemStepsWithStepExecutionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExecutionsItemStepsRequestBuilderInternal instantiates a new StepsRequestBuilder and sets the default values.
func NewExecutionsItemStepsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsRequestBuilder) {
    m := &ExecutionsItemStepsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/executions/{executions%2Did}/steps{?schemaTarget*,pageable*,assembler*}", pathParameters),
    }
    return m
}
// NewExecutionsItemStepsRequestBuilder instantiates a new StepsRequestBuilder and sets the default values.
func NewExecutionsItemStepsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsItemStepsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsItemStepsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsItemStepsRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *ExecutionsItemStepsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsItemStepsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExecutionsItemStepsRequestBuilder) WithUrl(rawUrl string)(*ExecutionsItemStepsRequestBuilder) {
    return NewExecutionsItemStepsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
