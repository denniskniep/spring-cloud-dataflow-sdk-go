package jobs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ExecutionsItemStepsItemProgressRequestBuilder builds and executes requests for operations under \jobs\executions\{jobExecutionId}\steps\{stepExecutionId}\progress
type ExecutionsItemStepsItemProgressRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ExecutionsItemStepsItemProgressRequestBuilderGetQueryParameters 
type ExecutionsItemStepsItemProgressRequestBuilderGetQueryParameters struct {
    // 
    SchemaTarget *string `uriparametername:"schemaTarget"`
}
// ExecutionsItemStepsItemProgressRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExecutionsItemStepsItemProgressRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExecutionsItemStepsItemProgressRequestBuilderGetQueryParameters
}
// NewExecutionsItemStepsItemProgressRequestBuilderInternal instantiates a new ProgressRequestBuilder and sets the default values.
func NewExecutionsItemStepsItemProgressRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsItemProgressRequestBuilder) {
    m := &ExecutionsItemStepsItemProgressRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/jobs/executions/{jobExecutionId}/steps/{stepExecutionId}/progress{?schemaTarget*}", pathParameters),
    }
    return m
}
// NewExecutionsItemStepsItemProgressRequestBuilder instantiates a new ProgressRequestBuilder and sets the default values.
func NewExecutionsItemStepsItemProgressRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExecutionsItemStepsItemProgressRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExecutionsItemStepsItemProgressRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ExecutionsItemStepsItemProgressRequestBuilder) Get(ctx context.Context, requestConfiguration *ExecutionsItemStepsItemProgressRequestBuilderGetRequestConfiguration)([]byte, error) {
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
func (m *ExecutionsItemStepsItemProgressRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ExecutionsItemStepsItemProgressRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ExecutionsItemStepsItemProgressRequestBuilder) WithUrl(rawUrl string)(*ExecutionsItemStepsItemProgressRequestBuilder) {
    return NewExecutionsItemStepsItemProgressRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
