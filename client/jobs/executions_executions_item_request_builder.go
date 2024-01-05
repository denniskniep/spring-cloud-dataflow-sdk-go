package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f "dataflow/client/models"
)

// ExecutionsExecutionsItemRequestBuilder builds and executes requests for operations under \jobs\executions\{executions-id}
type ExecutionsExecutionsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsExecutionsItemRequestBuilderGetQueryParameters 
type ExecutionsExecutionsItemRequestBuilderGetQueryParameters struct {
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// ExecutionsExecutionsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsExecutionsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsExecutionsItemRequestBuilderGetQueryParameters
}
// ExecutionsExecutionsItemRequestBuilderPutQueryParameters 
type ExecutionsExecutionsItemRequestBuilderPutQueryParameters struct {
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// ExecutionsExecutionsItemRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsExecutionsItemRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsExecutionsItemRequestBuilderPutQueryParameters
}
// NewExecutionsExecutionsItemRequestBuilderInternal instantiates a new ExecutionsItemRequestBuilder and sets the default values.
func NewExecutionsExecutionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExecutionsItemRequestBuilder) {
    m := &ExecutionsExecutionsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/executions/{executions%2Did}{?schemaTarget*}", pathParameters),
    }
    return m
}
// NewExecutionsExecutionsItemRequestBuilder instantiates a new ExecutionsItemRequestBuilder and sets the default values.
func NewExecutionsExecutionsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsExecutionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsExecutionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsExecutionsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsExecutionsItemRequestBuilderGetRequestConfiguration)(i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f.JobExecutionResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f.CreateJobExecutionResourceFromDiscriminatorValue, nil)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i320978ecdc392613ad9c6b55a6f9f5639a92533af8a2d757ef66f9a633865c3f.JobExecutionResourceable), nil
}
func (m *ExecutionsExecutionsItemRequestBuilder) Put(ctx context.Context, requestConfiguration *ExecutionsExecutionsItemRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, requestConfiguration);
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
// Steps the steps property
func (m *ExecutionsExecutionsItemRequestBuilder) Steps()(*ExecutionsItemStepsRequestBuilder) {
    return NewExecutionsItemStepsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
func (m *ExecutionsExecutionsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsExecutionsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExecutionsExecutionsItemRequestBuilder) ToPutRequestInformation(ctx context.Context, requestConfiguration *ExecutionsExecutionsItemRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ExecutionsExecutionsItemRequestBuilder) WithUrl(rawUrl string)(*ExecutionsExecutionsItemRequestBuilder) {
    return NewExecutionsExecutionsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
