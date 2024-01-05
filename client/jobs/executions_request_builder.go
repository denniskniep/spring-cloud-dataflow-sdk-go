package jobs

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8 "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/models"
    i220bf2fb75b02bbb806ce37092dbb6f24024a2ee1d18f0d9eae68355662eb76b "github.com/denniskniep/spring-cloud-dataflow-sdk-go/client/jobs/executions"
)

// ExecutionsRequestBuilder builds and executes requests for operations under \jobs\executions
type ExecutionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsRequestBuilderGetQueryParameters 
type ExecutionsRequestBuilderGetQueryParameters struct {
    // 
    Assembler *string `uriparametername:"assembler"`
    // 
    Name *string `uriparametername:"name"`
    // 
    Pageable *string `uriparametername:"pageable"`
    // 
    // Deprecated: This property is deprecated, use statusAsGetStatusQueryParameterType instead
    Status *string `uriparametername:"status"`
    // 
    StatusAsGetStatusQueryParameterType *i220bf2fb75b02bbb806ce37092dbb6f24024a2ee1d18f0d9eae68355662eb76b.GetStatusQueryParameterType `uriparametername:"status"`
}
// ExecutionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsRequestBuilderGetQueryParameters
}
// ByExecutionsId gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.jobs.executions.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ExecutionsRequestBuilder) ByExecutionsId(executionsId string)(*ExecutionsExecutionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if executionsId != "" {
        urlTplParams["executions%2Did"] = executionsId
    }
    return NewExecutionsExecutionsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByExecutionsIdInt64 gets an item from the github.com/denniskniep/spring-cloud-dataflow-sdk-go/client.jobs.executions.item collection
func (m *ExecutionsRequestBuilder) ByExecutionsIdInt64(executionsId int64)(*ExecutionsExecutionsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["executions%2Did"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(executionsId, 10)
    return NewExecutionsExecutionsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewExecutionsRequestBuilderInternal instantiates a new ExecutionsRequestBuilder and sets the default values.
func NewExecutionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsRequestBuilder) {
    m := &ExecutionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/executions{?name*,status*,pageable*,assembler*}", pathParameters),
    }
    return m
}
// NewExecutionsRequestBuilder instantiates a new ExecutionsRequestBuilder and sets the default values.
func NewExecutionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsRequestBuilderGetRequestConfiguration)(iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.PagedModelJobExecutionResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.CreatePagedModelJobExecutionResourceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iad7d7e09ef436a0a5bba1e03c342f1f06ef5b0695af5e104ba3e19ae9f63d8d8.PagedModelJobExecutionResourceable), nil
}
func (m *ExecutionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ExecutionsRequestBuilder) WithUrl(rawUrl string)(*ExecutionsRequestBuilder) {
    return NewExecutionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
